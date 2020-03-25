// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type DescribeUsersRequest struct {
	// query key, support these fields(user_id, email, phone_number, status)
	SearchWord *wrappers.StringValue `protobuf:"bytes,1,opt,name=search_word,json=searchWord,proto3" json:"search_word,omitempty"`
	// sort key, order by sort_key, default create_time
	SortKey *wrappers.StringValue `protobuf:"bytes,2,opt,name=sort_key,json=sortKey,proto3" json:"sort_key,omitempty"`
	// value = 0 sort ASC, value = 1 sort DESC
	Reverse *wrappers.BoolValue `protobuf:"bytes,3,opt,name=reverse,proto3" json:"reverse,omitempty"`
	// data limit, default 20, max 200
	Limit uint32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// data offset, default 0
	Offset uint32 `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	// use root group ids to get all group ids
	RootGroupId []string `protobuf:"bytes,6,rep,name=root_group_id,json=rootGroupId,proto3" json:"root_group_id,omitempty"`
	// group ids
	GroupId []string `protobuf:"bytes,7,rep,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// user ids
	UserId []string `protobuf:"bytes,8,rep,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// status eg.[active|deleted]
	Status []string `protobuf:"bytes,9,rep,name=status,proto3" json:"status,omitempty"`
	// role ids
	RoleId []string `protobuf:"bytes,10,rep,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// username
	Username []string `protobuf:"bytes,11,rep,name=username,proto3" json:"username,omitempty"`
	// email, eg.op@yunify.com
	Email []string `protobuf:"bytes,12,rep,name=email,proto3" json:"email,omitempty"`
	// phone number, string of 11 digital
	PhoneNumber          []string `protobuf:"bytes,13,rep,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeUsersRequest) Reset()         { *m = DescribeUsersRequest{} }
func (m *DescribeUsersRequest) String() string { return proto.CompactTextString(m) }
func (*DescribeUsersRequest) ProtoMessage()    {}
func (*DescribeUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *DescribeUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeUsersRequest.Unmarshal(m, b)
}
func (m *DescribeUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeUsersRequest.Marshal(b, m, deterministic)
}
func (m *DescribeUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeUsersRequest.Merge(m, src)
}
func (m *DescribeUsersRequest) XXX_Size() int {
	return xxx_messageInfo_DescribeUsersRequest.Size(m)
}
func (m *DescribeUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeUsersRequest proto.InternalMessageInfo

func (m *DescribeUsersRequest) GetSearchWord() *wrappers.StringValue {
	if m != nil {
		return m.SearchWord
	}
	return nil
}

func (m *DescribeUsersRequest) GetSortKey() *wrappers.StringValue {
	if m != nil {
		return m.SortKey
	}
	return nil
}

func (m *DescribeUsersRequest) GetReverse() *wrappers.BoolValue {
	if m != nil {
		return m.Reverse
	}
	return nil
}

func (m *DescribeUsersRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeUsersRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeUsersRequest) GetRootGroupId() []string {
	if m != nil {
		return m.RootGroupId
	}
	return nil
}

func (m *DescribeUsersRequest) GetGroupId() []string {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *DescribeUsersRequest) GetUserId() []string {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *DescribeUsersRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeUsersRequest) GetRoleId() []string {
	if m != nil {
		return m.RoleId
	}
	return nil
}

func (m *DescribeUsersRequest) GetUsername() []string {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *DescribeUsersRequest) GetEmail() []string {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *DescribeUsersRequest) GetPhoneNumber() []string {
	if m != nil {
		return m.PhoneNumber
	}
	return nil
}

type DescribeUsersResponse struct {
	// total count of qualified user
	TotalCount uint32 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// list of user
	UserSet              []*User  `protobuf:"bytes,2,rep,name=user_set,json=userSet,proto3" json:"user_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeUsersResponse) Reset()         { *m = DescribeUsersResponse{} }
func (m *DescribeUsersResponse) String() string { return proto.CompactTextString(m) }
func (*DescribeUsersResponse) ProtoMessage()    {}
func (*DescribeUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *DescribeUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeUsersResponse.Unmarshal(m, b)
}
func (m *DescribeUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeUsersResponse.Marshal(b, m, deterministic)
}
func (m *DescribeUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeUsersResponse.Merge(m, src)
}
func (m *DescribeUsersResponse) XXX_Size() int {
	return xxx_messageInfo_DescribeUsersResponse.Size(m)
}
func (m *DescribeUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeUsersResponse proto.InternalMessageInfo

func (m *DescribeUsersResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeUsersResponse) GetUserSet() []*User {
	if m != nil {
		return m.UserSet
	}
	return nil
}

type GetUserRequest struct {
	// user id
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FloatValue           float32  `protobuf:"fixed32,2,opt,name=float_value,json=floatValue,proto3" json:"float_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{2}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GetUserRequest) GetFloatValue() float32 {
	if m != nil {
		return m.FloatValue
	}
	return 0
}

type User struct {
	// user id, user belong to different group and role, has different permissions
	UserId *wrappers.StringValue `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// user name
	Username *wrappers.StringValue `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// user email
	Email *wrappers.StringValue `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// user phone number
	PhoneNumber *wrappers.StringValue `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// user description
	Description *wrappers.StringValue `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// user status eg.[active|deleted]
	Status *wrappers.StringValue `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	// the time when user create
	CreateTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// the time when user update
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// record changed time of status
	StatusTime           *timestamp.Timestamp `protobuf:"bytes,9,opt,name=status_time,json=statusTime,proto3" json:"status_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{3}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() *wrappers.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *User) GetUsername() *wrappers.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *User) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *User) GetPhoneNumber() *wrappers.StringValue {
	if m != nil {
		return m.PhoneNumber
	}
	return nil
}

func (m *User) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *User) GetStatus() *wrappers.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *User) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *User) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *User) GetStatusTime() *timestamp.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

type CreateUserRequest struct {
	// required, user email
	Email *wrappers.StringValue `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// user phone number
	PhoneNumber *wrappers.StringValue `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// required, user password
	Password *wrappers.StringValue `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// required, user role_id
	RoleId *wrappers.StringValue `protobuf:"bytes,4,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// user description
	Description          *wrappers.StringValue `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{4}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *CreateUserRequest) GetPhoneNumber() *wrappers.StringValue {
	if m != nil {
		return m.PhoneNumber
	}
	return nil
}

func (m *CreateUserRequest) GetPassword() *wrappers.StringValue {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *CreateUserRequest) GetRoleId() *wrappers.StringValue {
	if m != nil {
		return m.RoleId
	}
	return nil
}

func (m *CreateUserRequest) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

type CreateUserResponse struct {
	// id of user created
	UserId               *wrappers.StringValue `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{5}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUserId() *wrappers.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func init() {
	proto.RegisterType((*DescribeUsersRequest)(nil), "hal9000.DescribeUsersRequest")
	proto.RegisterType((*DescribeUsersResponse)(nil), "hal9000.DescribeUsersResponse")
	proto.RegisterType((*GetUserRequest)(nil), "hal9000.GetUserRequest")
	proto.RegisterType((*User)(nil), "hal9000.User")
	proto.RegisterType((*CreateUserRequest)(nil), "hal9000.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "hal9000.CreateUserResponse")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 980 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x5b, 0x6f, 0x1b, 0x45,
	0x14, 0xd6, 0xda, 0x6e, 0xec, 0x8c, 0x63, 0xa4, 0x0e, 0x0d, 0x59, 0xb6, 0x97, 0x4c, 0xf7, 0xc9,
	0x54, 0xbe, 0xc4, 0x4b, 0x2a, 0x52, 0x73, 0x93, 0x5d, 0xd4, 0x28, 0x8a, 0x04, 0x92, 0x5b, 0xa8,
	0x54, 0x14, 0xad, 0xc6, 0xde, 0xf1, 0x7a, 0xc4, 0x7a, 0x67, 0x99, 0x99, 0x75, 0x54, 0x45, 0x11,
	0x88, 0x47, 0x1e, 0x8d, 0xc4, 0x7f, 0x41, 0xfc, 0x08, 0x10, 0xbc, 0xf1, 0x88, 0xf8, 0x21, 0x68,
	0x66, 0xd6, 0x1b, 0xe7, 0x52, 0xc5, 0x81, 0x27, 0xfb, 0x7c, 0xe7, 0x7c, 0x3b, 0x67, 0xce, 0xf9,
	0xce, 0x19, 0x00, 0x24, 0x11, 0xb2, 0x95, 0x70, 0x26, 0x19, 0x2c, 0x4f, 0x70, 0xf4, 0x64, 0x67,
	0x67, 0xc7, 0x79, 0x10, 0x32, 0x16, 0x46, 0xa4, 0xad, 0xe1, 0x61, 0x3a, 0x6e, 0x1f, 0x73, 0x9c,
	0x24, 0x84, 0x0b, 0x13, 0xe8, 0x6c, 0x5f, 0xf4, 0x4b, 0x3a, 0x25, 0x42, 0xe2, 0x69, 0x92, 0x05,
	0xdc, 0xcb, 0x02, 0x70, 0x42, 0xdb, 0x38, 0x8e, 0x99, 0xc4, 0x92, 0xb2, 0x78, 0x41, 0x6f, 0xe8,
	0x9f, 0x51, 0x33, 0x24, 0x71, 0x53, 0x1c, 0xe3, 0x30, 0x24, 0xbc, 0xcd, 0x12, 0x1d, 0x71, 0x39,
	0xda, 0xfd, 0xbd, 0x08, 0xee, 0x7c, 0x46, 0xc4, 0x88, 0xd3, 0x21, 0xf9, 0x52, 0x10, 0x2e, 0x06,
	0xe4, 0xdb, 0x94, 0x08, 0x09, 0x3f, 0x06, 0x55, 0x41, 0x30, 0x1f, 0x4d, 0xfc, 0x63, 0xc6, 0x03,
	0xdb, 0x42, 0x56, 0xbd, 0xea, 0xdd, 0x6b, 0x99, 0xa3, 0x5b, 0x8b, 0xdc, 0x5a, 0xcf, 0x25, 0xa7,
	0x71, 0xf8, 0x15, 0x8e, 0x52, 0x32, 0x00, 0x86, 0xf0, 0x92, 0xf1, 0x00, 0x7e, 0x00, 0x2a, 0x82,
	0x71, 0xe9, 0x7f, 0x43, 0x5e, 0xdb, 0x85, 0x15, 0xb8, 0x65, 0x15, 0x7d, 0x48, 0x5e, 0xc3, 0x5d,
	0x50, 0xe6, 0x64, 0x46, 0xb8, 0x20, 0x76, 0x51, 0xf3, 0x9c, 0x4b, 0xbc, 0x3e, 0x63, 0x51, 0xc6,
	0xca, 0x42, 0xe1, 0x1d, 0x70, 0x2b, 0xa2, 0x53, 0x2a, 0xed, 0x12, 0xb2, 0xea, 0xb5, 0x81, 0x31,
	0xe0, 0x3b, 0x60, 0x8d, 0x8d, 0xc7, 0x82, 0x48, 0xfb, 0x96, 0x86, 0x33, 0x0b, 0xba, 0xa0, 0xc6,
	0x19, 0x93, 0x7e, 0xc8, 0x59, 0x9a, 0xf8, 0x34, 0xb0, 0xd7, 0x50, 0xb1, 0xbe, 0x3e, 0xa8, 0x2a,
	0x70, 0x5f, 0x61, 0x07, 0x01, 0x7c, 0x17, 0x54, 0x72, 0x77, 0x59, 0xbb, 0xcb, 0x61, 0xe6, 0xda,
	0x02, 0xe5, 0x54, 0x10, 0xae, 0x3c, 0x15, 0xed, 0x59, 0x53, 0xe6, 0x41, 0xa0, 0xce, 0x13, 0x12,
	0xcb, 0x54, 0xd8, 0xeb, 0x06, 0x37, 0x96, 0x22, 0x70, 0x16, 0x11, 0x45, 0x00, 0xc6, 0xa1, 0xcc,
	0x83, 0x00, 0x3a, 0xa0, 0xa2, 0xa8, 0x31, 0x9e, 0x12, 0xbb, 0xaa, 0x3d, 0xb9, 0xad, 0xae, 0x44,
	0xa6, 0x98, 0x46, 0xf6, 0x86, 0x76, 0x18, 0x03, 0x3e, 0x04, 0x1b, 0xc9, 0x84, 0xc5, 0xc4, 0x8f,
	0xd3, 0xe9, 0x90, 0x70, 0xbb, 0x66, 0x32, 0xd7, 0xd8, 0xe7, 0x1a, 0x72, 0x87, 0x60, 0xf3, 0x42,
	0x47, 0x45, 0xc2, 0x62, 0x41, 0xe0, 0x36, 0xa8, 0x4a, 0x26, 0x71, 0xe4, 0x8f, 0x58, 0x1a, 0x4b,
	0xdd, 0xd2, 0xda, 0x00, 0x68, 0xe8, 0xa9, 0x42, 0x60, 0xdd, 0xa4, 0xe3, 0xab, 0x8a, 0x15, 0x50,
	0xb1, 0x5e, 0xf5, 0x6a, 0xad, 0x4c, 0xb5, 0x2d, 0xf5, 0xa9, 0x81, 0xbe, 0xf7, 0x73, 0x22, 0xdd,
	0xbf, 0x2d, 0xf0, 0xd6, 0x3e, 0x91, 0x1a, 0xcc, 0x04, 0x33, 0x3b, 0xab, 0x8a, 0xfa, 0xf2, 0x7a,
	0xff, 0x68, 0xde, 0x7b, 0xe5, 0x69, 0x0c, 0xd1, 0xe0, 0x7b, 0xcb, 0xfa, 0xd1, 0x3a, 0xfc, 0x1a,
	0x37, 0xc7, 0xbd, 0xe6, 0xb3, 0x9d, 0xe6, 0x93, 0xa3, 0x93, 0xbd, 0xd3, 0xe6, 0xb2, 0xb9, 0x7b,
	0x13, 0xb3, 0xe3, 0x9d, 0xe6, 0x45, 0xff, 0x08, 0x54, 0xc7, 0x11, 0xc3, 0xd2, 0x9f, 0x29, 0x49,
	0x68, 0xb1, 0x15, 0xfa, 0x77, 0xe7, 0x3d, 0xdb, 0xbb, 0xfd, 0x4c, 0xe1, 0x48, 0xe3, 0x68, 0x4c,
	0x49, 0x14, 0x74, 0x8b, 0x3b, 0x2d, 0x6f, 0x00, 0x74, 0xbc, 0x56, 0x50, 0xd7, 0x99, 0xf7, 0xb6,
	0xbc, 0x4d, 0xf8, 0xf6, 0x09, 0x72, 0xb3, 0xec, 0xdd, 0x2e, 0x72, 0x3d, 0xcf, 0x45, 0xa7, 0xee,
	0xaf, 0x25, 0x50, 0x52, 0x37, 0x84, 0x2f, 0xcf, 0x5f, 0xed, 0x1a, 0x2d, 0xf7, 0xd1, 0xbc, 0x77,
	0x3f, 0xbf, 0x78, 0xb7, 0xd8, 0xe9, 0x74, 0xfe, 0xb0, 0x36, 0x84, 0x0e, 0x30, 0x49, 0xe6, 0xb9,
	0xef, 0x2d, 0xf5, 0x7f, 0x95, 0x29, 0x39, 0x53, 0x87, 0xb7, 0x50, 0x47, 0x71, 0x05, 0x5a, 0xa6,
	0x9d, 0x4f, 0x2f, 0x68, 0xa7, 0xb4, 0x02, 0x75, 0x59, 0x59, 0xf0, 0x13, 0x50, 0x0d, 0xb4, 0xb2,
	0xf4, 0x3a, 0xd1, 0x43, 0x75, 0x2d, 0x7f, 0x89, 0x00, 0x77, 0xf3, 0xf9, 0x58, 0x5b, 0x81, 0xba,
	0x98, 0x9e, 0x0f, 0x41, 0x75, 0xc4, 0x09, 0x96, 0xc4, 0x57, 0x8b, 0xd0, 0x2e, 0xbf, 0x61, 0x2b,
	0xbc, 0x58, 0x6c, 0xc9, 0x01, 0x30, 0xe1, 0x0a, 0x50, 0xe4, 0x34, 0x09, 0x72, 0x72, 0xe5, 0x7a,
	0xb2, 0x09, 0x5f, 0x90, 0x4d, 0x0e, 0x86, 0xbc, 0x7e, 0x3d, 0xd9, 0x84, 0x2b, 0xc0, 0xfd, 0xa5,
	0x00, 0x6e, 0x3f, 0xd5, 0x89, 0x2c, 0x4f, 0x49, 0xde, 0x37, 0xeb, 0xbf, 0xf7, 0xad, 0x70, 0xd3,
	0xbe, 0xed, 0x81, 0x4a, 0x82, 0x85, 0xd0, 0x8b, 0x7c, 0x15, 0xbd, 0xe4, 0xd1, 0xf0, 0xf1, 0xd9,
	0xe6, 0x5a, 0x45, 0x2d, 0x8b, 0xbd, 0xf6, 0x3f, 0x85, 0xe2, 0x1e, 0x02, 0xb8, 0x5c, 0xba, 0x6c,
	0x7f, 0x3d, 0xbe, 0xd1, 0x18, 0x2e, 0x86, 0xcc, 0xfb, 0xab, 0x08, 0x4a, 0x2f, 0x54, 0xed, 0x7f,
	0xb3, 0x40, 0xed, 0xdc, 0x66, 0x84, 0xf7, 0xf3, 0xf5, 0x76, 0xd5, 0x1b, 0xe8, 0x3c, 0x78, 0x93,
	0xdb, 0x24, 0xe4, 0x7e, 0x37, 0xef, 0xf9, 0xf0, 0x68, 0x9f, 0x48, 0xa4, 0x8e, 0x12, 0x0d, 0x34,
	0xa6, 0x91, 0x24, 0x1c, 0x1d, 0x53, 0x39, 0x31, 0x0b, 0x47, 0xd4, 0xb3, 0x8c, 0x1b, 0x48, 0xf7,
	0xb2, 0x81, 0x96, 0x3b, 0xd9, 0x40, 0x46, 0x2f, 0xef, 0x35, 0x50, 0x40, 0xc6, 0x38, 0x8d, 0x24,
	0xe2, 0x44, 0xa6, 0x3c, 0x46, 0x38, 0x8a, 0xcc, 0x37, 0x7f, 0xf8, 0xf3, 0x9f, 0x9f, 0x0a, 0x5b,
	0x70, 0xb3, 0x3d, 0x52, 0xaf, 0x3d, 0xc5, 0xd3, 0x16, 0x65, 0xed, 0x59, 0xa7, 0xad, 0x9d, 0xf0,
	0x67, 0x0b, 0x80, 0xb3, 0x42, 0x41, 0x27, 0xcf, 0xf7, 0x92, 0xf0, 0x9c, 0xbb, 0x57, 0xfa, 0xb2,
	0x8b, 0x7c, 0x31, 0xef, 0x35, 0xe0, 0x23, 0xe3, 0xd0, 0xe7, 0x36, 0x10, 0x1d, 0xeb, 0x3f, 0x68,
	0x82, 0x67, 0x04, 0xe1, 0x60, 0x4a, 0x63, 0x94, 0x10, 0x3e, 0xa5, 0x42, 0x50, 0x16, 0xeb, 0xac,
	0x1c, 0xf7, 0xea, 0xac, 0xba, 0xd6, 0x23, 0x18, 0x82, 0x72, 0xf6, 0x3c, 0xc0, 0xad, 0xfc, 0xe0,
	0xf3, 0x0f, 0x86, 0x73, 0xfe, 0x6d, 0x71, 0x77, 0xe7, 0x3d, 0x00, 0x2b, 0x61, 0x56, 0x4c, 0x7d,
	0xc2, 0x43, 0xb8, 0x7d, 0xe5, 0x09, 0xed, 0x93, 0xac, 0xaa, 0xa7, 0xfd, 0xd2, 0xab, 0x42, 0x32,
	0x1c, 0xae, 0x69, 0x01, 0xbc, 0xff, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0x43, 0x02, 0x84,
	0x70, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	// Get users, filter with fields(user_id, email, phone_number, status), default return all users
	DescribeUsers(ctx context.Context, in *DescribeUsersRequest, opts ...grpc.CallOption) (*DescribeUsersResponse, error)
	// Create user, if user have admin permission
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	//Get user
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) DescribeUsers(ctx context.Context, in *DescribeUsersRequest, opts ...grpc.CallOption) (*DescribeUsersResponse, error) {
	out := new(DescribeUsersResponse)
	err := c.cc.Invoke(ctx, "/hal9000.Test/DescribeUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/hal9000.Test/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/hal9000.Test/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	// Get users, filter with fields(user_id, email, phone_number, status), default return all users
	DescribeUsers(context.Context, *DescribeUsersRequest) (*DescribeUsersResponse, error)
	// Create user, if user have admin permission
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	//Get user
	GetUser(context.Context, *GetUserRequest) (*User, error)
}

// UnimplementedTestServer can be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (*UnimplementedTestServer) DescribeUsers(ctx context.Context, req *DescribeUsersRequest) (*DescribeUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeUsers not implemented")
}
func (*UnimplementedTestServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedTestServer) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_DescribeUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).DescribeUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hal9000.Test/DescribeUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).DescribeUsers(ctx, req.(*DescribeUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hal9000.Test/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hal9000.Test/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hal9000.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeUsers",
			Handler:    _Test_DescribeUsers_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Test_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Test_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
