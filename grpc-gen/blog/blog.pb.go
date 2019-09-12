// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blog.proto

package blog

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Status int32

const (
	Status_DRAFT   Status = 0
	Status_PUBLISH Status = 1
	Status_REMOVE  Status = 2
)

var Status_name = map[int32]string{
	0: "DRAFT",
	1: "PUBLISH",
	2: "REMOVE",
}

var Status_value = map[string]int32{
	"DRAFT":   0,
	"PUBLISH": 1,
	"REMOVE":  2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{0}
}

type ListRequest struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{0}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListResponse struct {
	Code                 int64       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Blogs                []*BlogData `protobuf:"bytes,3,rep,name=blogs,proto3" json:"blogs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{1}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ListResponse) GetBlogs() []*BlogData {
	if m != nil {
		return m.Blogs
	}
	return nil
}

type BlogResponse struct {
	Code                 int64     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Blog                 *BlogData `protobuf:"bytes,3,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BlogResponse) Reset()         { *m = BlogResponse{} }
func (m *BlogResponse) String() string { return proto.CompactTextString(m) }
func (*BlogResponse) ProtoMessage()    {}
func (*BlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{2}
}
func (m *BlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogResponse.Unmarshal(m, b)
}
func (m *BlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogResponse.Marshal(b, m, deterministic)
}
func (m *BlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogResponse.Merge(m, src)
}
func (m *BlogResponse) XXX_Size() int {
	return xxx_messageInfo_BlogResponse.Size(m)
}
func (m *BlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlogResponse proto.InternalMessageInfo

func (m *BlogResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BlogResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BlogResponse) GetBlog() *BlogData {
	if m != nil {
		return m.Blog
	}
	return nil
}

type GetRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{3}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteRequest struct {
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{4}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Comment struct {
	Author               string   `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	CreateTime           int64    `protobuf:"varint,2,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{5}
}
func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Comment) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Comment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type BlogData struct {
	Id                   int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT"`
	Title                string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Des                  string     `protobuf:"bytes,3,opt,name=des,proto3" json:"des,omitempty"`
	Content              string     `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	CreateTime           int64      `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Status               Status     `protobuf:"varint,6,opt,name=status,proto3,enum=blog.Status" json:"status,omitempty"`
	Type                 string     `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Likes                int64      `protobuf:"varint,8,opt,name=likes,proto3" json:"likes,omitempty"`
	Views                int64      `protobuf:"varint,9,opt,name=views,proto3" json:"views,omitempty"`
	Comments             []*Comment `protobuf:"bytes,10,rep,name=comments,proto3" json:"comments,omitempty" gorm:"-"`
	CommentStr           string     `protobuf:"bytes,11,opt,name=commentStr,proto3" json:"commentStr,omitempty" json:"-"`
	Tags                 []string   `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty" gorm:"-"`
	TagStr               string     `protobuf:"bytes,13,opt,name=tagStr,proto3" json:"tagStr,omitempty" json:"-"`
	Images               []string   `protobuf:"bytes,14,rep,name=images,proto3" json:"images,omitempty" gorm:"-"`
	ImagesStr            string     `protobuf:"bytes,15,opt,name=imagesStr,proto3" json:"imagesStr,omitempty" json:"-"`
	UserID               int64      `protobuf:"varint,16,opt,name=userID,proto3" json:"userID,omitempty"`
	UserAvatar           string     `protobuf:"bytes,17,opt,name=userAvatar,proto3" json:"userAvatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BlogData) Reset()         { *m = BlogData{} }
func (m *BlogData) String() string { return proto.CompactTextString(m) }
func (*BlogData) ProtoMessage()    {}
func (*BlogData) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{6}
}
func (m *BlogData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogData.Unmarshal(m, b)
}
func (m *BlogData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogData.Marshal(b, m, deterministic)
}
func (m *BlogData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogData.Merge(m, src)
}
func (m *BlogData) XXX_Size() int {
	return xxx_messageInfo_BlogData.Size(m)
}
func (m *BlogData) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogData.DiscardUnknown(m)
}

var xxx_messageInfo_BlogData proto.InternalMessageInfo

func (m *BlogData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BlogData) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BlogData) GetDes() string {
	if m != nil {
		return m.Des
	}
	return ""
}

func (m *BlogData) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *BlogData) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *BlogData) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_DRAFT
}

func (m *BlogData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *BlogData) GetLikes() int64 {
	if m != nil {
		return m.Likes
	}
	return 0
}

func (m *BlogData) GetViews() int64 {
	if m != nil {
		return m.Views
	}
	return 0
}

func (m *BlogData) GetComments() []*Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

func (m *BlogData) GetCommentStr() string {
	if m != nil {
		return m.CommentStr
	}
	return ""
}

func (m *BlogData) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *BlogData) GetTagStr() string {
	if m != nil {
		return m.TagStr
	}
	return ""
}

func (m *BlogData) GetImages() []string {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *BlogData) GetImagesStr() string {
	if m != nil {
		return m.ImagesStr
	}
	return ""
}

func (m *BlogData) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *BlogData) GetUserAvatar() string {
	if m != nil {
		return m.UserAvatar
	}
	return ""
}

func init() {
	proto.RegisterEnum("blog.Status", Status_name, Status_value)
	proto.RegisterType((*ListRequest)(nil), "blog.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "blog.ListResponse")
	proto.RegisterType((*BlogResponse)(nil), "blog.BlogResponse")
	proto.RegisterType((*GetRequest)(nil), "blog.GetRequest")
	proto.RegisterType((*DeleteRequest)(nil), "blog.DeleteRequest")
	proto.RegisterType((*Comment)(nil), "blog.Comment")
	proto.RegisterType((*BlogData)(nil), "blog.BlogData")
}

func init() { proto.RegisterFile("blog.proto", fileDescriptor_6745b25902462fb1) }

var fileDescriptor_6745b25902462fb1 = []byte{
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x25, 0xaf, 0x49, 0x72, 0xf3, 0x20, 0x75, 0x11, 0xb2, 0xa2, 0x8a, 0x86, 0x51, 0x91, 0xaa,
	0xa8, 0x6d, 0xa4, 0x22, 0x84, 0xd4, 0x0a, 0x41, 0xd2, 0xa4, 0xa5, 0x52, 0x1f, 0x68, 0x92, 0xb2,
	0x61, 0x51, 0x9c, 0xc4, 0x9d, 0x9a, 0x66, 0xc6, 0x61, 0xec, 0x14, 0x55, 0x88, 0x0d, 0xbf, 0xc0,
	0x1f, 0xf1, 0x0b, 0xec, 0x58, 0x74, 0xc5, 0x17, 0xf4, 0x0b, 0x90, 0xed, 0x99, 0x74, 0x52, 0x22,
	0x84, 0xd8, 0xdd, 0xd7, 0x39, 0xbe, 0x3e, 0xbe, 0xd7, 0x00, 0xfd, 0x11, 0x77, 0x37, 0xc6, 0x01,
	0x97, 0x1c, 0xa5, 0x95, 0x5d, 0x5d, 0x72, 0x39, 0x77, 0x47, 0xb4, 0x41, 0xc6, 0xac, 0x41, 0x7c,
	0x9f, 0x4b, 0x22, 0x19, 0xf7, 0x85, 0xa9, 0xa9, 0xae, 0xbb, 0x4c, 0x9e, 0x4f, 0xfa, 0x1b, 0x03,
	0xee, 0x35, 0x5c, 0xee, 0xf2, 0x86, 0x0e, 0xf7, 0x27, 0x67, 0xda, 0xd3, 0x8e, 0xb6, 0x4c, 0xb9,
	0xbd, 0x0d, 0x85, 0x03, 0x26, 0xa4, 0x43, 0x3f, 0x4e, 0xa8, 0x90, 0xe8, 0x01, 0x64, 0x46, 0xcc,
	0x63, 0x12, 0x27, 0x6a, 0x89, 0xd5, 0x94, 0x63, 0x1c, 0xf4, 0x10, 0x2c, 0x7e, 0x76, 0x26, 0xa8,
	0xc4, 0x49, 0x1d, 0x0e, 0x3d, 0xbb, 0x0f, 0x45, 0x03, 0x16, 0x63, 0xee, 0x0b, 0x8a, 0x10, 0xa4,
	0x07, 0x7c, 0x48, 0x43, 0xb0, 0xb6, 0x11, 0x86, 0xac, 0x47, 0x85, 0x20, 0x2e, 0xd5, 0xe0, 0xbc,
	0x13, 0xb9, 0x68, 0x05, 0x32, 0xea, 0x3e, 0x02, 0xa7, 0x6a, 0xa9, 0xd5, 0xc2, 0x66, 0x79, 0x43,
	0xdf, 0xb4, 0x35, 0xe2, 0x6e, 0x9b, 0x48, 0xe2, 0x98, 0xa4, 0xfd, 0x1e, 0x8a, 0x2a, 0xf4, 0x9f,
	0x67, 0xd8, 0xa0, 0x35, 0xc3, 0xa9, 0x5a, 0x62, 0xce, 0x11, 0x3a, 0x67, 0x2f, 0x01, 0xec, 0xd1,
	0xa9, 0x02, 0x65, 0x48, 0xb2, 0x61, 0xc8, 0x9e, 0x64, 0x43, 0x7b, 0x19, 0x4a, 0x6d, 0x3a, 0xa2,
	0x92, 0xce, 0x16, 0x24, 0xa7, 0x05, 0xef, 0x20, 0xbb, 0xc3, 0x3d, 0x8f, 0xfa, 0x5a, 0x27, 0x32,
	0x91, 0xe7, 0x3c, 0xd0, 0xf8, 0xbc, 0x13, 0x7a, 0xe8, 0x11, 0xc0, 0x20, 0xa0, 0x44, 0xd2, 0x1e,
	0xf3, 0x68, 0x08, 0x8d, 0x45, 0x54, 0xff, 0x03, 0xee, 0x4b, 0xea, 0x4b, 0xdd, 0x68, 0xde, 0x89,
	0x5c, 0xfb, 0x7b, 0x1a, 0x72, 0x51, 0xbb, 0xe8, 0xd9, 0x6d, 0x6b, 0xad, 0x27, 0x37, 0xd7, 0xcb,
	0x8f, 0x5d, 0x1e, 0x78, 0x5b, 0xf6, 0x38, 0x60, 0x1e, 0x09, 0xae, 0x4e, 0x2f, 0xe8, 0xd5, 0x76,
	0xf3, 0xa4, 0x77, 0x7c, 0xba, 0x7f, 0xb4, 0xe3, 0x74, 0x0e, 0x3b, 0x47, 0x3d, 0x5b, 0x35, 0xa8,
	0xde, 0x54, 0x32, 0x39, 0x8a, 0xb4, 0x31, 0x0e, 0xaa, 0x40, 0x6a, 0x48, 0x45, 0x78, 0x9e, 0x32,
	0xe3, 0x5d, 0xa4, 0x67, 0xba, 0xb8, 0xd3, 0x7f, 0xe6, 0x8f, 0xfe, 0x57, 0xc0, 0x12, 0x92, 0xc8,
	0x89, 0xc0, 0x56, 0x2d, 0xb1, 0x5a, 0xde, 0x2c, 0x1a, 0x9d, 0xbb, 0x3a, 0xe6, 0x84, 0x39, 0xf5,
	0x72, 0xf2, 0x6a, 0x4c, 0x71, 0x56, 0x93, 0x6b, 0xdb, 0xcc, 0xdb, 0x05, 0x15, 0x38, 0x17, 0xcd,
	0xdb, 0x05, 0x15, 0x2a, 0x7a, 0xc9, 0xe8, 0x27, 0x81, 0xf3, 0x26, 0xaa, 0x1d, 0xf4, 0x1c, 0x72,
	0x03, 0x23, 0xb4, 0xc0, 0xa0, 0x47, 0xa6, 0x64, 0xce, 0x09, 0xe5, 0x6f, 0x15, 0x6f, 0xae, 0x97,
	0x73, 0x46, 0x93, 0x75, 0xdb, 0x99, 0x16, 0xa3, 0x35, 0x80, 0xd0, 0xee, 0xca, 0x00, 0x17, 0xd4,
	0xf1, 0xa6, 0xf6, 0x83, 0xe0, 0xbe, 0xae, 0x8d, 0xe5, 0x51, 0x0d, 0xd2, 0x92, 0xb8, 0x02, 0x17,
	0x6b, 0xa9, 0xa8, 0x6e, 0xca, 0xa9, 0x33, 0xea, 0xba, 0x92, 0xb8, 0x8a, 0xab, 0x34, 0x87, 0x2b,
	0xcc, 0xa9, 0x2a, 0xe6, 0x11, 0x97, 0x0a, 0x5c, 0x9e, 0xc3, 0x14, 0xe6, 0x50, 0x1d, 0xf2, 0xc6,
	0x52, 0x74, 0xf7, 0xe7, 0xd0, 0xdd, 0xa6, 0xd5, 0x78, 0x4d, 0x04, 0x0d, 0xf6, 0xdb, 0xb8, 0x62,
	0xd6, 0xd0, 0x78, 0xea, 0x79, 0x94, 0xd5, 0xbc, 0x24, 0x92, 0x04, 0x78, 0x41, 0xcb, 0x1b, 0x8b,
	0xd4, 0xd7, 0xc0, 0x32, 0x4f, 0x81, 0xf2, 0x90, 0x69, 0x3b, 0xcd, 0xdd, 0x5e, 0xe5, 0x1e, 0x2a,
	0x40, 0xf6, 0xcd, 0x49, 0xeb, 0x60, 0xbf, 0xfb, 0xba, 0x92, 0x40, 0x00, 0x96, 0xd3, 0x39, 0x3c,
	0x7e, 0xdb, 0xa9, 0x24, 0x37, 0x7f, 0x26, 0xa1, 0xa0, 0x46, 0xae, 0x4b, 0x83, 0x4b, 0x36, 0xa0,
	0xe8, 0x05, 0xa4, 0xd5, 0x92, 0xa3, 0x05, 0x23, 0x76, 0xec, 0xb7, 0xa8, 0xa2, 0x78, 0xc8, 0xec,
	0xa7, 0x5d, 0xfe, 0xfa, 0xe3, 0xd7, 0xb7, 0x64, 0x0e, 0x59, 0x0d, 0xbd, 0xbf, 0xe8, 0x25, 0x58,
	0x3b, 0x7a, 0x52, 0xd0, 0x9d, 0xed, 0x8b, 0xd0, 0xf1, 0xed, 0xb6, 0x17, 0x34, 0xba, 0x60, 0x87,
	0xe8, 0xad, 0x44, 0x1d, 0xbd, 0x82, 0xd4, 0x1e, 0x95, 0xa8, 0x62, 0xaa, 0x6f, 0x37, 0x75, 0x2e,
	0x7e, 0x51, 0xe3, 0x4b, 0xa8, 0x60, 0xf0, 0x8d, 0xcf, 0x6c, 0xf8, 0x05, 0x35, 0xc1, 0x3a, 0x19,
	0x0f, 0xff, 0xb5, 0x85, 0x90, 0xa2, 0x3a, 0x43, 0xb1, 0x0b, 0x96, 0xf9, 0x05, 0xd0, 0xa2, 0x81,
	0xcc, 0xfc, 0x09, 0x7f, 0xe3, 0xa9, 0xc7, 0x79, 0xfa, 0x96, 0xfe, 0x75, 0x9f, 0xfe, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0xf0, 0xd5, 0x72, 0x46, 0xd6, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	// List Blog
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Create Blog
	Create(ctx context.Context, in *BlogData, opts ...grpc.CallOption) (*BlogResponse, error)
	// Get Blog
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*BlogResponse, error)
	// Update Blog
	Update(ctx context.Context, in *BlogData, opts ...grpc.CallOption) (*BlogResponse, error)
	// Delete Blog
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*BlogResponse, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Create(ctx context.Context, in *BlogData, opts ...grpc.CallOption) (*BlogResponse, error) {
	out := new(BlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*BlogResponse, error) {
	out := new(BlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Update(ctx context.Context, in *BlogData, opts ...grpc.CallOption) (*BlogResponse, error) {
	out := new(BlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*BlogResponse, error) {
	out := new(BlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	// List Blog
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Create Blog
	Create(context.Context, *BlogData) (*BlogResponse, error)
	// Get Blog
	Get(context.Context, *GetRequest) (*BlogResponse, error)
	// Update Blog
	Update(context.Context, *BlogData) (*BlogResponse, error)
	// Delete Blog
	Delete(context.Context, *DeleteRequest) (*BlogResponse, error)
}

// UnimplementedBlogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (*UnimplementedBlogServiceServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedBlogServiceServer) Create(ctx context.Context, req *BlogData) (*BlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedBlogServiceServer) Get(ctx context.Context, req *GetRequest) (*BlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedBlogServiceServer) Update(ctx context.Context, req *BlogData) (*BlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedBlogServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*BlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Create(ctx, req.(*BlogData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Update(ctx, req.(*BlogData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _BlogService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _BlogService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BlogService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BlogService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BlogService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog.proto",
}
