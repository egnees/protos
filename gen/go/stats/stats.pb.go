// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: stats/stats.proto

package statsv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetPostStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId int64 `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *GetPostStatsRequest) Reset() {
	*x = GetPostStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostStatsRequest) ProtoMessage() {}

func (x *GetPostStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stats_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostStatsRequest.ProtoReflect.Descriptor instead.
func (*GetPostStatsRequest) Descriptor() ([]byte, []int) {
	return file_stats_stats_proto_rawDescGZIP(), []int{0}
}

func (x *GetPostStatsRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type GetPostStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId int64 `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Likes  int64 `protobuf:"varint,2,opt,name=likes,proto3" json:"likes,omitempty"`
	Views  int64 `protobuf:"varint,3,opt,name=views,proto3" json:"views,omitempty"`
}

func (x *GetPostStatsResponse) Reset() {
	*x = GetPostStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostStatsResponse) ProtoMessage() {}

func (x *GetPostStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stats_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostStatsResponse.ProtoReflect.Descriptor instead.
func (*GetPostStatsResponse) Descriptor() ([]byte, []int) {
	return file_stats_stats_proto_rawDescGZIP(), []int{1}
}

func (x *GetPostStatsResponse) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *GetPostStatsResponse) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *GetPostStatsResponse) GetViews() int64 {
	if x != nil {
		return x.Views
	}
	return 0
}

type GetTopNPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SortByLikes bool  `protobuf:"varint,1,opt,name=sort_by_likes,json=sortByLikes,proto3" json:"sort_by_likes,omitempty"`
	N           int64 `protobuf:"varint,2,opt,name=N,proto3" json:"N,omitempty"`
}

func (x *GetTopNPostsRequest) Reset() {
	*x = GetTopNPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopNPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopNPostsRequest) ProtoMessage() {}

func (x *GetTopNPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stats_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopNPostsRequest.ProtoReflect.Descriptor instead.
func (*GetTopNPostsRequest) Descriptor() ([]byte, []int) {
	return file_stats_stats_proto_rawDescGZIP(), []int{2}
}

func (x *GetTopNPostsRequest) GetSortByLikes() bool {
	if x != nil {
		return x.SortByLikes
	}
	return false
}

func (x *GetTopNPostsRequest) GetN() int64 {
	if x != nil {
		return x.N
	}
	return 0
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId      int64  `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	AuthorLogin string `protobuf:"bytes,2,opt,name=author_login,json=authorLogin,proto3" json:"author_login,omitempty"`
	Likes       int64  `protobuf:"varint,3,opt,name=likes,proto3" json:"likes,omitempty"`
	Views       int64  `protobuf:"varint,4,opt,name=views,proto3" json:"views,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_stats_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_stats_stats_proto_rawDescGZIP(), []int{3}
}

func (x *Post) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *Post) GetAuthorLogin() string {
	if x != nil {
		return x.AuthorLogin
	}
	return ""
}

func (x *Post) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *Post) GetViews() int64 {
	if x != nil {
		return x.Views
	}
	return 0
}

type GetTopNPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetTopNPostsResponse) Reset() {
	*x = GetTopNPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_stats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopNPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopNPostsResponse) ProtoMessage() {}

func (x *GetTopNPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stats_stats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopNPostsResponse.ProtoReflect.Descriptor instead.
func (*GetTopNPostsResponse) Descriptor() ([]byte, []int) {
	return file_stats_stats_proto_rawDescGZIP(), []int{4}
}

func (x *GetTopNPostsResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type GetTopNUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SortByLikes bool  `protobuf:"varint,1,opt,name=sort_by_likes,json=sortByLikes,proto3" json:"sort_by_likes,omitempty"`
	N           int64 `protobuf:"varint,2,opt,name=N,proto3" json:"N,omitempty"`
}

func (x *GetTopNUsersRequest) Reset() {
	*x = GetTopNUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_stats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopNUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopNUsersRequest) ProtoMessage() {}

func (x *GetTopNUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stats_stats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopNUsersRequest.ProtoReflect.Descriptor instead.
func (*GetTopNUsersRequest) Descriptor() ([]byte, []int) {
	return file_stats_stats_proto_rawDescGZIP(), []int{5}
}

func (x *GetTopNUsersRequest) GetSortByLikes() bool {
	if x != nil {
		return x.SortByLikes
	}
	return false
}

func (x *GetTopNUsersRequest) GetN() int64 {
	if x != nil {
		return x.N
	}
	return 0
}

type UserStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Likes int64  `protobuf:"varint,2,opt,name=likes,proto3" json:"likes,omitempty"`
	Views int64  `protobuf:"varint,3,opt,name=views,proto3" json:"views,omitempty"`
}

func (x *UserStat) Reset() {
	*x = UserStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_stats_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStat) ProtoMessage() {}

func (x *UserStat) ProtoReflect() protoreflect.Message {
	mi := &file_stats_stats_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStat.ProtoReflect.Descriptor instead.
func (*UserStat) Descriptor() ([]byte, []int) {
	return file_stats_stats_proto_rawDescGZIP(), []int{6}
}

func (x *UserStat) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UserStat) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *UserStat) GetViews() int64 {
	if x != nil {
		return x.Views
	}
	return 0
}

type GetTopNUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserStat `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetTopNUsersResponse) Reset() {
	*x = GetTopNUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_stats_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopNUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopNUsersResponse) ProtoMessage() {}

func (x *GetTopNUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stats_stats_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopNUsersResponse.ProtoReflect.Descriptor instead.
func (*GetTopNUsersResponse) Descriptor() ([]byte, []int) {
	return file_stats_stats_proto_rawDescGZIP(), []int{7}
}

func (x *GetTopNUsersResponse) GetUsers() []*UserStat {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_stats_stats_proto protoreflect.FileDescriptor

var file_stats_stats_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x2e, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6b, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x22, 0x47, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x70, 0x4e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0d, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x4c, 0x69, 0x6b,
	0x65, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x4e,
	0x22, 0x6e, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x22, 0x39, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x50, 0x6f, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x47, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x6f, 0x72, 0x74, 0x42,
	0x79, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x01, 0x4e, 0x22, 0x4c, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x22, 0x3d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x32, 0xe2, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x47, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x70, 0x4e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x74, 0x69, 0x6e, 0x79, 0x2d, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stats_stats_proto_rawDescOnce sync.Once
	file_stats_stats_proto_rawDescData = file_stats_stats_proto_rawDesc
)

func file_stats_stats_proto_rawDescGZIP() []byte {
	file_stats_stats_proto_rawDescOnce.Do(func() {
		file_stats_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_stats_stats_proto_rawDescData)
	})
	return file_stats_stats_proto_rawDescData
}

var file_stats_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_stats_stats_proto_goTypes = []interface{}{
	(*GetPostStatsRequest)(nil),  // 0: stats.GetPostStatsRequest
	(*GetPostStatsResponse)(nil), // 1: stats.GetPostStatsResponse
	(*GetTopNPostsRequest)(nil),  // 2: stats.GetTopNPostsRequest
	(*Post)(nil),                 // 3: stats.Post
	(*GetTopNPostsResponse)(nil), // 4: stats.GetTopNPostsResponse
	(*GetTopNUsersRequest)(nil),  // 5: stats.GetTopNUsersRequest
	(*UserStat)(nil),             // 6: stats.UserStat
	(*GetTopNUsersResponse)(nil), // 7: stats.GetTopNUsersResponse
}
var file_stats_stats_proto_depIdxs = []int32{
	3, // 0: stats.GetTopNPostsResponse.posts:type_name -> stats.Post
	6, // 1: stats.GetTopNUsersResponse.users:type_name -> stats.UserStat
	0, // 2: stats.Stats.GetPostStats:input_type -> stats.GetPostStatsRequest
	2, // 3: stats.Stats.GetTopNPosts:input_type -> stats.GetTopNPostsRequest
	5, // 4: stats.Stats.GetTopNUsers:input_type -> stats.GetTopNUsersRequest
	1, // 5: stats.Stats.GetPostStats:output_type -> stats.GetPostStatsResponse
	4, // 6: stats.Stats.GetTopNPosts:output_type -> stats.GetTopNPostsResponse
	7, // 7: stats.Stats.GetTopNUsers:output_type -> stats.GetTopNUsersResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stats_stats_proto_init() }
func file_stats_stats_proto_init() {
	if File_stats_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stats_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostStatsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stats_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostStatsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stats_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopNPostsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stats_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stats_stats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopNPostsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stats_stats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopNUsersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stats_stats_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_stats_stats_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopNUsersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stats_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stats_stats_proto_goTypes,
		DependencyIndexes: file_stats_stats_proto_depIdxs,
		MessageInfos:      file_stats_stats_proto_msgTypes,
	}.Build()
	File_stats_stats_proto = out.File
	file_stats_stats_proto_rawDesc = nil
	file_stats_stats_proto_goTypes = nil
	file_stats_stats_proto_depIdxs = nil
}