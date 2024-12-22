// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user.proto

package model

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// User
type User struct {
	// id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;autoIncrement;comment:id"`
	// uid
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty" gorm:"unique;size:128;comment:uid"`
	// email
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty" gorm:"index;size:20;comment:email"`
	// user_name
	UserName string `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty" gorm:"index;size:255;comment:user name"`
	// nike_name
	NikeName string `protobuf:"bytes,5,opt,name=nike_name,json=nikeName,proto3" json:"nike_name,omitempty" gorm:"size:255;comment:nike name"`
	// password
	Password string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty" gorm:"size:128;comment:password"`
	// avatar
	Avatar string `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty" gorm:"size:255;comment:avatar"`
	// access_token
	AccessToken string `protobuf:"bytes,8,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty" gorm:"comment:access_token"`
	// role_id
	RoleId int32 `protobuf:"varint,9,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty" gorm:"default:2;size:2;comment:role_id"`
	// status
	Status int32 `protobuf:"varint,19,opt,name=status,proto3" json:"status,omitempty" gorm:"default:1;size:2;comment:status 1 normal"`
	// created_at
	CreatedAt int64 `protobuf:"varint,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" gorm:"comment:created at"`
	// updated_at
	UpdatedAt int64 `protobuf:"varint,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" gorm:"comment:updated at"`
	// deleted_at
	DeletedAt int64 `protobuf:"varint,22,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty" gorm:"comment:deleted at"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetNikeName() string {
	if m != nil {
		return m.NikeName
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *User) GetRoleId() int32 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *User) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *User) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *User) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *User) GetDeletedAt() int64 {
	if m != nil {
		return m.DeletedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x86, 0xeb, 0xa4, 0x71, 0x93, 0xf9, 0xbe, 0x95, 0xf9, 0x91, 0x01, 0x61, 0x9b, 0xe1, 0x2f,
	0x02, 0x91, 0x90, 0x54, 0x95, 0xaa, 0xa4, 0x42, 0x34, 0x62, 0x41, 0x55, 0x89, 0x85, 0x05, 0x1b,
	0x36, 0xd1, 0xc4, 0x73, 0x1a, 0x46, 0xb5, 0x3d, 0x61, 0x3c, 0x06, 0xca, 0x55, 0x70, 0x59, 0x2c,
	0xbb, 0x44, 0x42, 0xb2, 0x20, 0xb9, 0x03, 0xaf, 0x58, 0x22, 0xcf, 0x4c, 0x42, 0xda, 0x2a, 0xab,
	0x9c, 0x37, 0x7a, 0x9f, 0xc7, 0x33, 0x47, 0x36, 0x42, 0x79, 0x06, 0xa2, 0x33, 0x13, 0x5c, 0xf2,
	0xdb, 0xcf, 0xa6, 0x4c, 0x7e, 0xc8, 0x27, 0x9d, 0x88, 0x27, 0xdd, 0x29, 0x9f, 0xf2, 0xae, 0xfa,
	0x7b, 0x92, 0x9f, 0xa8, 0xa4, 0x82, 0x9a, 0x74, 0x1d, 0xff, 0xb4, 0xd1, 0xf6, 0xbb, 0x0c, 0x84,
	0xf3, 0x02, 0xd5, 0x18, 0x75, 0xad, 0xc0, 0x6a, 0xd7, 0x47, 0x9d, 0xb2, 0xf0, 0x9f, 0x4c, 0xb9,
	0x48, 0x06, 0x78, 0x26, 0x58, 0x42, 0xc4, 0xd9, 0x31, 0x9c, 0x0d, 0x49, 0x2e, 0xf9, 0x51, 0x1a,
	0x09, 0x48, 0x20, 0x95, 0xc3, 0x88, 0x27, 0xd5, 0xef, 0x80, 0x51, 0x1c, 0xd6, 0x18, 0x75, 0xf6,
	0x51, 0x3d, 0x67, 0xd4, 0xad, 0x05, 0x56, 0xbb, 0x35, 0x7a, 0x54, 0x16, 0x3e, 0xd6, 0x82, 0x3c,
	0x65, 0x1f, 0x73, 0x18, 0x66, 0xec, 0x2b, 0x0c, 0x7a, 0xfd, 0xfd, 0x15, 0x97, 0x57, 0x60, 0x85,
	0x38, 0x07, 0xa8, 0x01, 0x09, 0x61, 0xb1, 0x5b, 0xbf, 0xcc, 0xb2, 0x94, 0xc2, 0x17, 0x8d, 0xf6,
	0x9f, 0xaf, 0x48, 0x55, 0xc6, 0xa1, 0x86, 0x9c, 0xd7, 0xa8, 0x55, 0xdd, 0x7e, 0x9c, 0x92, 0x04,
	0xdc, 0x6d, 0x65, 0x78, 0x5a, 0x16, 0xfe, 0xe3, 0xab, 0x86, 0xbd, 0xbd, 0x7f, 0x0f, 0xcf, 0x40,
	0x04, 0x15, 0x81, 0xc3, 0x66, 0x35, 0xbf, 0x21, 0x09, 0x38, 0x23, 0xd4, 0x4a, 0xd9, 0x29, 0x68,
	0x53, 0x43, 0x99, 0x1e, 0x96, 0x85, 0x7f, 0x4f, 0x9b, 0xae, 0x38, 0xaa, 0xee, 0xd2, 0x51, 0xcd,
	0xca, 0xf1, 0x12, 0x35, 0x67, 0x24, 0xcb, 0x3e, 0x73, 0x41, 0x5d, 0x5b, 0x29, 0x1e, 0x94, 0x85,
	0x1f, 0xac, 0x29, 0xd6, 0x77, 0xb0, 0xac, 0xe2, 0x70, 0x45, 0x39, 0x03, 0x64, 0x93, 0x4f, 0x44,
	0x12, 0xe1, 0xee, 0x28, 0x1e, 0x97, 0x85, 0xef, 0x6d, 0x38, 0x82, 0x2e, 0xe2, 0xd0, 0x10, 0xce,
	0x08, 0xfd, 0x4f, 0xa2, 0x08, 0xb2, 0x6c, 0x2c, 0xf9, 0x29, 0xa4, 0x6e, 0x53, 0x19, 0xfc, 0xb2,
	0xf0, 0xef, 0x68, 0xc3, 0x0a, 0x5c, 0x6b, 0xe1, 0xf0, 0x3f, 0x1d, 0xdf, 0x56, 0xc9, 0x79, 0x85,
	0x76, 0x04, 0x8f, 0x61, 0xcc, 0xa8, 0xdb, 0x0a, 0xac, 0x76, 0x63, 0x7d, 0x9b, 0x14, 0x4e, 0x48,
	0x1e, 0xcb, 0x41, 0xdf, 0x6c, 0x74, 0x75, 0x10, 0x43, 0xe0, 0xd0, 0xae, 0xa6, 0x23, 0xea, 0x1c,
	0x23, 0x3b, 0x93, 0x44, 0xe6, 0x99, 0x7b, 0x4d, 0x49, 0x76, 0xcb, 0xc2, 0xef, 0x5e, 0x94, 0xf4,
	0x2e, 0x4b, 0x34, 0x10, 0xf4, 0x82, 0x94, 0x8b, 0x84, 0xc4, 0x38, 0x34, 0x0a, 0xe7, 0x00, 0xa1,
	0x48, 0x00, 0x91, 0x40, 0xc7, 0x44, 0xba, 0xd7, 0xd5, 0x2b, 0x7a, 0xb7, 0x2c, 0xfc, 0x5b, 0x17,
	0x2f, 0x65, 0x3a, 0x01, 0x91, 0x38, 0x6c, 0x99, 0x70, 0x28, 0x2b, 0x3a, 0x9f, 0xd1, 0x25, 0x7d,
	0x63, 0x13, 0x6d, 0x3a, 0x9a, 0x36, 0x41, 0xd3, 0x14, 0x62, 0x30, 0xf4, 0xcd, 0x4d, 0xb4, 0xe9,
	0x68, 0xda, 0x84, 0x43, 0x39, 0xba, 0xff, 0xe7, 0xb7, 0x67, 0x7d, 0x9f, 0x7b, 0xd6, 0xf9, 0xdc,
	0xb3, 0x7e, 0xcd, 0x3d, 0xeb, 0xdb, 0xc2, 0xdb, 0x3a, 0x5f, 0x78, 0x5b, 0x3f, 0x16, 0xde, 0xd6,
	0xfb, 0x46, 0xc2, 0x29, 0xc4, 0x13, 0x5b, 0x7d, 0x89, 0xbb, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x20, 0x8a, 0xb9, 0xa5, 0xc6, 0x03, 0x00, 0x00,
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DeletedAt != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.DeletedAt))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb0
	}
	if m.UpdatedAt != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa8
	}
	if m.CreatedAt != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa0
	}
	if m.Status != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x98
	}
	if m.RoleId != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.RoleId))
		i--
		dAtA[i] = 0x48
	}
	if len(m.AccessToken) > 0 {
		i -= len(m.AccessToken)
		copy(dAtA[i:], m.AccessToken)
		i = encodeVarintUser(dAtA, i, uint64(len(m.AccessToken)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Avatar) > 0 {
		i -= len(m.Avatar)
		copy(dAtA[i:], m.Avatar)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Avatar)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.NikeName) > 0 {
		i -= len(m.NikeName)
		copy(dAtA[i:], m.NikeName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.NikeName)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.UserName) > 0 {
		i -= len(m.UserName)
		copy(dAtA[i:], m.UserName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.UserName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedUser(r randyUser, easy bool) *User {
	this := &User{}
	this.Id = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Id *= -1
	}
	this.Uid = string(randStringUser(r))
	this.Email = string(randStringUser(r))
	this.UserName = string(randStringUser(r))
	this.NikeName = string(randStringUser(r))
	this.Password = string(randStringUser(r))
	this.Avatar = string(randStringUser(r))
	this.AccessToken = string(randStringUser(r))
	this.RoleId = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.RoleId *= -1
	}
	this.Status = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Status *= -1
	}
	this.CreatedAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.CreatedAt *= -1
	}
	this.UpdatedAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.UpdatedAt *= -1
	}
	this.DeletedAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.DeletedAt *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyUser interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneUser(r randyUser) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringUser(r randyUser) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneUser(r)
	}
	return string(tmps)
}
func randUnrecognizedUser(r randyUser, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldUser(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldUser(dAtA []byte, r randyUser, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateUser(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateUser(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateUser(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUser(uint64(m.Id))
	}
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.UserName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.NikeName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Avatar)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.AccessToken)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.RoleId != 0 {
		n += 1 + sovUser(uint64(m.RoleId))
	}
	if m.Status != 0 {
		n += 2 + sovUser(uint64(m.Status))
	}
	if m.CreatedAt != 0 {
		n += 2 + sovUser(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 2 + sovUser(uint64(m.UpdatedAt))
	}
	if m.DeletedAt != 0 {
		n += 2 + sovUser(uint64(m.DeletedAt))
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NikeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NikeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Avatar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Avatar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleId", wireType)
			}
			m.RoleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoleId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 19:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 22:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedAt", wireType)
			}
			m.DeletedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeletedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUser
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUser
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUser = fmt.Errorf("proto: unexpected end of group")
)
