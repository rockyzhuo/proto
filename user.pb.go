// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PasswordCheckInfo struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	EncryptedPassword    string   `protobuf:"bytes,2,opt,name=encryptedPassword,proto3" json:"encryptedPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordCheckInfo) Reset()         { *m = PasswordCheckInfo{} }
func (m *PasswordCheckInfo) String() string { return proto.CompactTextString(m) }
func (*PasswordCheckInfo) ProtoMessage()    {}
func (*PasswordCheckInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *PasswordCheckInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordCheckInfo.Unmarshal(m, b)
}
func (m *PasswordCheckInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordCheckInfo.Marshal(b, m, deterministic)
}
func (m *PasswordCheckInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordCheckInfo.Merge(m, src)
}
func (m *PasswordCheckInfo) XXX_Size() int {
	return xxx_messageInfo_PasswordCheckInfo.Size(m)
}
func (m *PasswordCheckInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordCheckInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordCheckInfo proto.InternalMessageInfo

func (m *PasswordCheckInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *PasswordCheckInfo) GetEncryptedPassword() string {
	if m != nil {
		return m.EncryptedPassword
	}
	return ""
}

type CheckResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (m *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(m, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type PageInfo struct {
	Pn                   uint32   `protobuf:"varint,1,opt,name=pn,proto3" json:"pn,omitempty"`
	PSize                uint32   `protobuf:"varint,2,opt,name=pSize,proto3" json:"pSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageInfo) Reset()         { *m = PageInfo{} }
func (m *PageInfo) String() string { return proto.CompactTextString(m) }
func (*PageInfo) ProtoMessage()    {}
func (*PageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *PageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageInfo.Unmarshal(m, b)
}
func (m *PageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageInfo.Marshal(b, m, deterministic)
}
func (m *PageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageInfo.Merge(m, src)
}
func (m *PageInfo) XXX_Size() int {
	return xxx_messageInfo_PageInfo.Size(m)
}
func (m *PageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PageInfo proto.InternalMessageInfo

func (m *PageInfo) GetPn() uint32 {
	if m != nil {
		return m.Pn
	}
	return 0
}

func (m *PageInfo) GetPSize() uint32 {
	if m != nil {
		return m.PSize
	}
	return 0
}

type MobileRequest struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MobileRequest) Reset()         { *m = MobileRequest{} }
func (m *MobileRequest) String() string { return proto.CompactTextString(m) }
func (*MobileRequest) ProtoMessage()    {}
func (*MobileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *MobileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileRequest.Unmarshal(m, b)
}
func (m *MobileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileRequest.Marshal(b, m, deterministic)
}
func (m *MobileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileRequest.Merge(m, src)
}
func (m *MobileRequest) XXX_Size() int {
	return xxx_messageInfo_MobileRequest.Size(m)
}
func (m *MobileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MobileRequest proto.InternalMessageInfo

func (m *MobileRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type IdRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdRequest) Reset()         { *m = IdRequest{} }
func (m *IdRequest) String() string { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()    {}
func (*IdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *IdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdRequest.Unmarshal(m, b)
}
func (m *IdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdRequest.Marshal(b, m, deterministic)
}
func (m *IdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdRequest.Merge(m, src)
}
func (m *IdRequest) XXX_Size() int {
	return xxx_messageInfo_IdRequest.Size(m)
}
func (m *IdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdRequest proto.InternalMessageInfo

func (m *IdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateUserInfo struct {
	NikeName             string   `protobuf:"bytes,1,opt,name=nikeName,proto3" json:"nikeName,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserInfo) Reset()         { *m = CreateUserInfo{} }
func (m *CreateUserInfo) String() string { return proto.CompactTextString(m) }
func (*CreateUserInfo) ProtoMessage()    {}
func (*CreateUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *CreateUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserInfo.Unmarshal(m, b)
}
func (m *CreateUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserInfo.Marshal(b, m, deterministic)
}
func (m *CreateUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserInfo.Merge(m, src)
}
func (m *CreateUserInfo) XXX_Size() int {
	return xxx_messageInfo_CreateUserInfo.Size(m)
}
func (m *CreateUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserInfo proto.InternalMessageInfo

func (m *CreateUserInfo) GetNikeName() string {
	if m != nil {
		return m.NikeName
	}
	return ""
}

func (m *CreateUserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserInfo) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type UpdateUserInfo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NikeName             string   `protobuf:"bytes,2,opt,name=nikeName,proto3" json:"nikeName,omitempty"`
	Gender               string   `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthday             uint64   `protobuf:"varint,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserInfo) Reset()         { *m = UpdateUserInfo{} }
func (m *UpdateUserInfo) String() string { return proto.CompactTextString(m) }
func (*UpdateUserInfo) ProtoMessage()    {}
func (*UpdateUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UpdateUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserInfo.Unmarshal(m, b)
}
func (m *UpdateUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserInfo.Marshal(b, m, deterministic)
}
func (m *UpdateUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserInfo.Merge(m, src)
}
func (m *UpdateUserInfo) XXX_Size() int {
	return xxx_messageInfo_UpdateUserInfo.Size(m)
}
func (m *UpdateUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserInfo proto.InternalMessageInfo

func (m *UpdateUserInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateUserInfo) GetNikeName() string {
	if m != nil {
		return m.NikeName
	}
	return ""
}

func (m *UpdateUserInfo) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UpdateUserInfo) GetBirthday() uint64 {
	if m != nil {
		return m.Birthday
	}
	return 0
}

type UserInfoResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PassWord             string   `protobuf:"bytes,2,opt,name=passWord,proto3" json:"passWord,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	NikeName             string   `protobuf:"bytes,4,opt,name=nikeName,proto3" json:"nikeName,omitempty"`
	BirthDay             uint64   `protobuf:"varint,5,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
	Gender               string   `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Role                 int32    `protobuf:"varint,7,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(m, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfoResponse) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

func (m *UserInfoResponse) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserInfoResponse) GetNikeName() string {
	if m != nil {
		return m.NikeName
	}
	return ""
}

func (m *UserInfoResponse) GetBirthDay() uint64 {
	if m != nil {
		return m.BirthDay
	}
	return 0
}

func (m *UserInfoResponse) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UserInfoResponse) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

type UserListResponse struct {
	Total                int32               `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data                 []*UserInfoResponse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (m *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(m, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *UserListResponse) GetData() []*UserInfoResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*PasswordCheckInfo)(nil), "PasswordCheckInfo")
	proto.RegisterType((*CheckResponse)(nil), "CheckResponse")
	proto.RegisterType((*EmptyResponse)(nil), "EmptyResponse")
	proto.RegisterType((*PageInfo)(nil), "PageInfo")
	proto.RegisterType((*MobileRequest)(nil), "MobileRequest")
	proto.RegisterType((*IdRequest)(nil), "IdRequest")
	proto.RegisterType((*CreateUserInfo)(nil), "CreateUserInfo")
	proto.RegisterType((*UpdateUserInfo)(nil), "UpdateUserInfo")
	proto.RegisterType((*UserInfoResponse)(nil), "UserInfoResponse")
	proto.RegisterType((*UserListResponse)(nil), "UserListResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x6d, 0x6b, 0xdb, 0x30,
	0x10, 0x26, 0xae, 0xf3, 0x76, 0x25, 0xce, 0x22, 0x4a, 0x31, 0xde, 0x97, 0x60, 0x18, 0xcb, 0x58,
	0x27, 0x4a, 0xbb, 0x6f, 0xfb, 0x96, 0x76, 0x8c, 0xc0, 0x5e, 0x8a, 0x47, 0x29, 0x0c, 0x06, 0x53,
	0xa2, 0x5b, 0x6b, 0x9a, 0xd8, 0x9a, 0xa4, 0x30, 0xbc, 0xff, 0xb4, 0x3f, 0xb0, 0x5f, 0x37, 0x22,
	0x5b, 0x4e, 0xe4, 0x84, 0x7e, 0xb2, 0x1f, 0xe9, 0xee, 0x9e, 0xe7, 0xce, 0xcf, 0x19, 0x60, 0xad,
	0x50, 0x52, 0x21, 0x73, 0x9d, 0xc7, 0xdf, 0x61, 0x74, 0xc3, 0x94, 0xfa, 0x9d, 0x4b, 0x7e, 0xf5,
	0x80, 0x8b, 0xc7, 0x59, 0xf6, 0x33, 0x27, 0x11, 0xf4, 0x44, 0x75, 0x18, 0xb6, 0xc6, 0xad, 0x49,
	0x3f, 0xa9, 0x31, 0x39, 0x83, 0x11, 0x66, 0x0b, 0x59, 0x08, 0x8d, 0xdc, 0x66, 0x86, 0x9e, 0x09,
	0xda, 0xbf, 0x88, 0x5f, 0xc1, 0xc0, 0x94, 0x4d, 0x50, 0x89, 0x3c, 0x53, 0x48, 0x42, 0xe8, 0xaa,
	0xf5, 0x62, 0x81, 0x4a, 0x99, 0xca, 0xbd, 0xc4, 0xc2, 0x78, 0x08, 0x83, 0xf7, 0x2b, 0xa1, 0x0b,
	0x1b, 0x1a, 0x9f, 0x43, 0xef, 0x86, 0xdd, 0xa3, 0x51, 0x14, 0x80, 0x27, 0x32, 0x93, 0x31, 0x48,
	0x3c, 0x91, 0x91, 0x13, 0x68, 0x8b, 0xaf, 0xe9, 0x1f, 0x34, 0xcc, 0x83, 0xa4, 0x04, 0xf1, 0x4b,
	0x18, 0x7c, 0xca, 0xe7, 0xe9, 0x12, 0x13, 0xfc, 0xb5, 0x46, 0xa5, 0xc9, 0x29, 0x74, 0x56, 0xe6,
	0xa0, 0x6a, 0xa3, 0x42, 0xf1, 0x73, 0xe8, 0xcf, 0xb8, 0x0d, 0x0a, 0xc0, 0x4b, 0xcb, 0x3e, 0xdb,
	0x89, 0x97, 0xf2, 0xf8, 0x07, 0x04, 0x57, 0x12, 0x99, 0xc6, 0x5b, 0x85, 0xd2, 0xce, 0x23, 0x4b,
	0x1f, 0xf1, 0x33, 0x5b, 0xd9, 0x42, 0x35, 0x76, 0x66, 0xe5, 0x35, 0x66, 0xb5, 0xa5, 0x3f, 0x72,
	0xe8, 0x05, 0x04, 0xb7, 0x82, 0xef, 0x32, 0x34, 0x34, 0x38, 0x8c, 0x5e, 0x83, 0xf1, 0x14, 0x3a,
	0xf7, 0x98, 0x71, 0x94, 0xb6, 0x6a, 0x89, 0x36, 0x39, 0xf3, 0x54, 0xea, 0x07, 0xce, 0x8a, 0xd0,
	0x1f, 0xb7, 0x26, 0x7e, 0x52, 0xe3, 0xf8, 0x5f, 0x0b, 0x9e, 0x59, 0xb2, 0xfa, 0x5b, 0x1c, 0x20,
	0xdd, 0x48, 0xbf, 0x6b, 0xb4, 0x72, 0xf7, 0x44, 0x2b, 0x8e, 0x50, 0x7f, 0x7f, 0x34, 0x46, 0xc0,
	0x35, 0x2b, 0xc2, 0xf6, 0x8e, 0xa0, 0x6b, 0x56, 0xec, 0x34, 0xd1, 0x71, 0x9a, 0x20, 0xe0, 0xcb,
	0x7c, 0x89, 0x61, 0xd7, 0xa8, 0x32, 0xef, 0xf1, 0x97, 0x52, 0xfb, 0xc7, 0x54, 0xe9, 0x5a, 0xfb,
	0x09, 0xb4, 0x75, 0xae, 0xd9, 0xb2, 0x92, 0x5f, 0x02, 0xf2, 0x02, 0x7c, 0xce, 0x34, 0x0b, 0xbd,
	0xf1, 0xd1, 0xe4, 0xf8, 0x62, 0x44, 0x9b, 0x2d, 0x27, 0xe6, 0xfa, 0xe2, 0xaf, 0x07, 0xfe, 0xe6,
	0x8a, 0xbc, 0x86, 0xe3, 0x0f, 0xa8, 0x6d, 0x71, 0xd2, 0xa7, 0xd6, 0x70, 0x51, 0x99, 0xeb, 0x50,
	0xbe, 0x85, 0x61, 0x15, 0x3c, 0x2d, 0x4a, 0x9b, 0x91, 0x80, 0x3a, 0x7e, 0x8b, 0xf6, 0x19, 0xc9,
	0x59, 0x4d, 0x31, 0x2d, 0x66, 0x9c, 0x00, 0xad, 0x8d, 0x77, 0x28, 0xfa, 0x1c, 0x60, 0xeb, 0x3d,
	0x32, 0xa4, 0xae, 0x11, 0x0f, 0x65, 0xbc, 0x01, 0xd8, 0x7a, 0x89, 0x0c, 0xa9, 0x6b, 0xac, 0x28,
	0xa0, 0xce, 0x52, 0x91, 0xcb, 0x6a, 0x21, 0xed, 0x86, 0x12, 0x42, 0xf7, 0xf6, 0x3f, 0x0a, 0xa8,
	0xb3, 0xb4, 0xd3, 0xfe, 0xb7, 0x2e, 0x7d, 0x67, 0xfe, 0x17, 0xf3, 0x8e, 0x79, 0x5c, 0xfe, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0x76, 0x17, 0x9d, 0xc6, 0x44, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResponse, error)
	GetUserByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	GetUserById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*EmptyResponse, error)
	CheckPassword(ctx context.Context, in *PasswordCheckInfo, opts ...grpc.CallOption) (*CheckResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserByMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckPassword(ctx context.Context, in *PasswordCheckInfo, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/User/CheckPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetUserList(context.Context, *PageInfo) (*UserListResponse, error)
	GetUserByMobile(context.Context, *MobileRequest) (*UserInfoResponse, error)
	GetUserById(context.Context, *IdRequest) (*UserInfoResponse, error)
	CreateUser(context.Context, *CreateUserInfo) (*UserInfoResponse, error)
	UpdateUser(context.Context, *UpdateUserInfo) (*EmptyResponse, error)
	CheckPassword(context.Context, *PasswordCheckInfo) (*CheckResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUserList(ctx context.Context, req *PageInfo) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (*UnimplementedUserServer) GetUserByMobile(ctx context.Context, req *MobileRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByMobile not implemented")
}
func (*UnimplementedUserServer) GetUserById(ctx context.Context, req *IdRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (*UnimplementedUserServer) CreateUser(ctx context.Context, req *CreateUserInfo) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServer) UpdateUser(ctx context.Context, req *UpdateUserInfo) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUserServer) CheckPassword(ctx context.Context, req *PasswordCheckInfo) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassword not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserList(ctx, req.(*PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserByMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByMobile(ctx, req.(*MobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordCheckInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CheckPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckPassword(ctx, req.(*PasswordCheckInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _User_GetUserList_Handler,
		},
		{
			MethodName: "GetUserByMobile",
			Handler:    _User_GetUserByMobile_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _User_GetUserById_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "CheckPassword",
			Handler:    _User_CheckPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
