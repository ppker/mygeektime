package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/zkep/my-geektime/internal/config"
	"github.com/zkep/my-geektime/internal/global"
	"github.com/zkep/my-geektime/internal/initialize"
	"github.com/zkep/my-geektime/internal/model"
	"github.com/zkep/my-geektime/internal/service"
	"github.com/zkep/my-geektime/internal/types/geek"
	"github.com/zkep/my-geektime/internal/types/task"
	"github.com/zkep/my-geektime/internal/types/user"
	"github.com/zkep/my-geektime/lib/utils"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

type DataFlags struct {
	Config   string  `name:"config" description:"Path to config file"`
	Ids      []int32 `name:"id" description:"1: 体系课，4:公开课" default:"1"`
	Cookies  string  `name:"cookies" description:"geektime cookies string"`
	Download bool    `name:"download" description:"download geektime source" default:"false"`
}

func (app *App) Data(f *DataFlags) error {
	var (
		cfg         config.Config
		accessToken string
	)
	if f.Config == "" {
		fi, err := app.assets.Open("config.yml")
		if err != nil {
			return err
		}
		defer func() { _ = fi.Close() }()
		if err = yaml.NewDecoder(fi).Decode(&cfg); err != nil {
			return err
		}
	} else {
		fi, err := os.Open(f.Config)
		if err != nil {
			return err
		}
		defer func() { _ = fi.Close() }()
		if err = yaml.NewDecoder(fi).Decode(&cfg); err != nil {
			return err
		}
	}
	global.CONF = &cfg
	global.CONF.Site.Download = f.Download
	if err := initialize.Gorm(app.ctx); err != nil {
		return err
	}
	if err := initialize.Logger(app.ctx); err != nil {
		return err
	}
	if err := initialize.Storage(app.ctx); err != nil {
		return err
	}
	if len(f.Cookies) > 0 {
		accessToken = f.Cookies
		if cookies := os.Getenv("cookies"); len(cookies) > 0 {
			accessToken = cookies
		}
	} else {
		var u model.User
		if err := global.DB.
			Where(&model.User{RoleId: user.AdminRoleId}).
			First(&u).Error; err != nil {
			return err
		}
		accessToken = u.AccessToken
	}
	if accessToken == "" {
		return errors.New("no access token")
	}
	var tags []Tag
	if err := json.Unmarshal([]byte(TagJSON), &tags); err != nil {
		return err
	}
	for _, id := range f.Ids {
		typ, ok := ProductTypes[id]
		if !ok {
			fmt.Printf("not found product id [%d]", id)
			continue
		}
		for _, form := range ProductForms {
			for _, tag := range tags {
				for _, opt := range tag.Options {
					if err := app.iterators(typ, form, opt, tag, id, accessToken); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (app *App) iterators(typ, form, opt Option, tag Tag, id int32, accessToken string) error {
	prev, psize, hasNext, total := 1, 20, true, 0
	fmt.Printf(
		"download start [%s/%s/%s/%s] \n",
		typ.Label, form.Label, tag.Label, opt.Label,
	)
	for hasNext {
		req := geek.PvipProductRequest{
			TagIds:       []int32{opt.Value},
			ProductType:  id,
			ProductForm:  form.Value,
			Sort:         8,
			Size:         psize,
			Prev:         prev,
			WithArticles: true,
		}
		resp, err := service.GetPvipProduct(app.ctx, "", accessToken, req)
		if err != nil {
			return err
		}
		total += len(resp.Data.Products)
		if len(resp.Data.Products) < psize {
			fmt.Printf(
				"download end [%s/%s/%s/%s] total: %d \n",
				typ.Label, form.Label, tag.Label, opt.Label, total,
			)
			hasNext = false
		}
		prev++
		for _, product := range resp.Data.Products {
			articles, err1 := service.GetArticles(app.ctx, "", accessToken,
				geek.ArticlesListRequest{
					Cid:   fmt.Sprintf("%d", product.ID),
					Order: "earliest",
					Prev:  1,
					Size:  1000,
				})
			if err1 != nil {
				return err1
			}
			jobId := utils.HalfUUID()
			itemRaw, _ := json.Marshal(product)
			job := &model.Task{
				TaskId:     jobId,
				TaskName:   product.Title,
				TaskType:   service.TASK_TYPE_PRODUCT,
				OtherId:    fmt.Sprintf("%d", product.ID),
				Cover:      product.Cover.Square,
				Raw:        itemRaw,
				OtherType:  typ.Value,
				OtherForm:  form.Value,
				OtherGroup: tag.Value,
				OtherTag:   opt.Value,
				Status:     service.TASK_STATUS_PENDING,
			}
			tasks := make([]*model.Task, 0, len(articles.Data.List))
			for _, article := range articles.Data.List {
				info, er := service.GetArticleInfo(app.ctx, "",
					accessToken, geek.ArticlesInfoRequest{Id: article.ID})
				if er != nil {
					return er
				}
				var m geek.ArticleInfoRaw
				if err = json.Unmarshal(info.Raw, &m); err != nil {
					return err
				}
				raw := m.Data
				otherId := fmt.Sprintf("%d", info.Data.Info.ID)
				taskName := info.Data.Info.Title
				cover := info.Data.Info.Cover.Default
				item := model.Task{
					TaskPid:    jobId,
					TaskId:     utils.HalfUUID(),
					OtherId:    otherId,
					TaskName:   taskName,
					TaskType:   service.TASK_TYPE_ARTICLE,
					Cover:      cover,
					Raw:        raw,
					OtherType:  typ.Value,
					OtherForm:  form.Value,
					OtherGroup: tag.Value,
					OtherTag:   opt.Value,
					Status:     service.TASK_STATUS_PENDING,
				}
				tasks = append(tasks, &item)
			}
			statistics := task.TaskStatistics{
				Count: len(tasks),
				Items: map[int]int{
					service.TASK_STATUS_PENDING:  len(tasks),
					service.TASK_STATUS_RUNNING:  0,
					service.TASK_STATUS_FINISHED: 0,
					service.TASK_STATUS_ERROR:    0,
				},
			}
			job.Statistics, _ = json.Marshal(statistics)
			err = global.DB.Transaction(func(tx *gorm.DB) error {
				if err = tx.Where(&model.Task{OtherId: job.OtherId}).
					Assign(job).FirstOrCreate(job).Error; err != nil {
					return err
				}
				for _, x := range tasks {
					if err = tx.Where(&model.Task{OtherId: x.OtherId}).
						Assign(x).FirstOrCreate(x).Error; err != nil {
						return err
					}
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type Tag struct {
	Option
	Options []Option `json:"options"`
}

type Option struct {
	Label string `json:"label"`
	Value int32  `json:"value"`
}

type TagMap struct {
	Label   string           `json:"label"`
	Options map[int32]string `json:"options"`
}

var (
	ProductTypes = map[int32]Option{
		1: {"体系课", 1},
		4: {"公开课", 4},
	}

	ProductForms = []Option{
		{"图文+音频", 1},
		{"视频", 2},
	}
)

var TagJSON = `
[
  {
    "label":"后端/架构",
    "value": 3,
    "options": [
      {
		"label": "Java",
		"value": 10
	  },
	  {
		"label": "Go",
		"value": 11
	  },
	  {
		"label": "Python",
		"value": 12
	  },
	  {
		"label": ".Net",
		"value": 65
	  },
	  {
		"label": "C语言",
		"value": 64
	  },
	  {
		"label": "基本功",
		"value": 22
	  },
	  {
		"label": "分布式",
		"value": 20
	  },
	  {
		"label": "中间件",
		"value": 17
	  },
	  {
		"label": "区块链",
		"value": 23
	  },
	  {
		"label": "全栈",
		"value": 21
	  },
	  {
		"label": "C++",
		"value": 61
	  },
	  {
		"label": "中台",
		"value": 19
	  },
	  {
		"label": "DDD",
		"value": 18
	  },
	  {
		"label": "案例",
		"value": 16
	  },
	  {
		"label": "微服务",
		"value": 14
	  },
	  {
		"label": "数据库",
		"value": 13
	  },
	  {
		"label": "Rust",
		"value": 98
	  }
    ]
  },
  {
    "label": "前端/移动",
    "value": 5,
    "options": [
      {
		"label": "JavaScript",
		"value": 30
	  },
	  {
		"label": "iOS",
		"value": 34
	  },
	  {
		"label": "Android",
		"value": 33
	  },
	  {
		"label": "Node.js",
		"value": 41
	  },
	  {
		"label": "网络",
		"value": 40
	  },
	  {
		"label": "webpack",
		"value": 39
	  },
	  {
		"label": "音视频",
		"value": 38
	  },
	  {
		"label": "浏览器",
		"value": 37
	  },
	  {
		"label": "Swift",
		"value": 36
	  },
	  {
		"label": "Kotlin",
		"value": 35
	  },
	  {
		"label": "框架",
		"value": 32
	  },
	  {
		"label": "TypeScript",
		"value": 31
	  },
	  {
		"label": "可视化",
		"value": 70
	  }
    ]
  },
  {
    "label": "计算机基础",
    "value": 9,
    "options": [
    	  {
			"label": "算法",
			"value": 75
		  },
		  {
			"label": "网络",
			"value": 76
		  },
		  {
			"label": "数据库",
			"value": 78
		  },
		  {
			"label": "面试专场",
			"value": 102
		  },
		  {
			"label": "组成原理",
			"value": 83
		  },
		  {
			"label": "数学",
			"value": 82
		  },
		  {
			"label": "代码",
			"value": 81
		  },
		  {
			"label": "编译原理",
			"value": 80
		  },
		  {
			"label": "操作系统",
			"value": 79
		  },
		  {
			"label": "工具",
			"value": 77
		  }
      ]
  },
  {
     "label": "AI/大数据",
     "value": 8,
     "options": [
      {
		"label": "NLP",
		"value": 66
	  },
	  {
		"label": "机器学习",
		"value": 56
	  },
	  {
		"label": "数据分析",
		"value": 72
	  },
	  {
		"label": "大数据",
		"value": 60
	  },
	  {
		"label": "推荐系统",
		"value": 59
	  },
	  {
		"label": "AI算法",
		"value": 58
	  },
	  {
		"label": "数学",
		"value": 57
	  },
	  {
		"label": "AIGC应用",
		"value": 99
	  },
	  {
		"label": "大模型",
		"value": 100
	  }
     ]
  },
  {
    "label": "运维/测试",
    "value": 6,
    "options": [
      {
		"label": "CI/CD",
		"value": 45
	  },
	  {
		"label": "云计算",
		"value": 67
	  },
	  {
		"label": "测试",
		"value": 47
	  },
	  {
		"label": "敏捷开发",
		"value": 50
	  },
	  {
		"label": "性能",
		"value": 49
	  },
	  {
		"label": "Linux",
		"value": 48
	  },
	  {
		"label": "安全",
		"value": 46
	  },
	  {
		"label": "运维管理",
		"value": 44
	  },
	  {
		"label": "DevOps",
		"value": 43
	  },
	  {
		"label": "Kubernetes",
		"value": 42
	  }
     ]
  },
  {
     "label": "产品/运营",
     "value": 7,
     "options": [
       {
		"label": "产品创新",
		"value": 55
	  },
	  {
		"label": "增长",
		"value": 54
	  },
	  {
		"label": "运营",
		"value": 53
	  },
	  {
		"label": "市场",
		"value": 74
	  }
     ]
  },
  {
    "label": "管理/成长",
    "value": 4,
    "options": [
      {
		"label": "研发效能",
		"value": 26
	  },
	  {
		"label": "团队管理",
		"value": 27
	  },
	  {
		"label": "个人成长",
		"value": 28
	  },
	  {
		"label": "面试专场",
		"value": 97
	  },
	  {
		"label": "短视频",
		"value": 95
	  },
	  {
		"label": "英语",
		"value": 94
	  },
	  {
		"label": "写作",
		"value": 93
	  },
	  {
		"label": "摄影",
		"value": 92
	  },
	  {
		"label": "音乐",
		"value": 91
	  },
	  {
		"label": "项目管理",
		"value": 29
	  },
	  {
		"label": "新经理",
		"value": 25
	  },
	  {
		"label": "OKR",
		"value": 24
	  }
    ]
  }
]
`
