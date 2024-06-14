// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: social_graph.proto

package socialproto

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

type InsertUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *InsertUserRequest) Reset() {
	*x = InsertUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertUserRequest) ProtoMessage() {}

func (x *InsertUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertUserRequest.ProtoReflect.Descriptor instead.
func (*InsertUserRequest) Descriptor() ([]byte, []int) {
	return file_social_graph_proto_rawDescGZIP(), []int{0}
}

func (x *InsertUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type InsertUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *InsertUserResponse) Reset() {
	*x = InsertUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertUserResponse) ProtoMessage() {}

func (x *InsertUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertUserResponse.ProtoReflect.Descriptor instead.
func (*InsertUserResponse) Descriptor() ([]byte, []int) {
	return file_social_graph_proto_rawDescGZIP(), []int{1}
}

func (x *InsertUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetFollowersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetFollowersRequest) Reset() {
	*x = GetFollowersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersRequest) ProtoMessage() {}

func (x *GetFollowersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersRequest.ProtoReflect.Descriptor instead.
func (*GetFollowersRequest) Descriptor() ([]byte, []int) {
	return file_social_graph_proto_rawDescGZIP(), []int{2}
}

func (x *GetFollowersRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetFollowersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followers []string `protobuf:"bytes,1,rep,name=followers,proto3" json:"followers,omitempty"`
}

func (x *GetFollowersResponse) Reset() {
	*x = GetFollowersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersResponse) ProtoMessage() {}

func (x *GetFollowersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersResponse.ProtoReflect.Descriptor instead.
func (*GetFollowersResponse) Descriptor() ([]byte, []int) {
	return file_social_graph_proto_rawDescGZIP(), []int{3}
}

func (x *GetFollowersResponse) GetFollowers() []string {
	if x != nil {
		return x.Followers
	}
	return nil
}

type GetFolloweesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetFolloweesRequest) Reset() {
	*x = GetFolloweesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFolloweesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFolloweesRequest) ProtoMessage() {}

func (x *GetFolloweesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFolloweesRequest.ProtoReflect.Descriptor instead.
func (*GetFolloweesRequest) Descriptor() ([]byte, []int) {
	return file_social_graph_proto_rawDescGZIP(), []int{4}
}

func (x *GetFolloweesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetFolloweesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followees []string `protobuf:"bytes,1,rep,name=followees,proto3" json:"followees,omitempty"`
}

func (x *GetFolloweesResponse) Reset() {
	*x = GetFolloweesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFolloweesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFolloweesResponse) ProtoMessage() {}

func (x *GetFolloweesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFolloweesResponse.ProtoReflect.Descriptor instead.
func (*GetFolloweesResponse) Descriptor() ([]byte, []int) {
	return file_social_graph_proto_rawDescGZIP(), []int{5}
}

func (x *GetFolloweesResponse) GetFollowees() []string {
	if x != nil {
		return x.Followees
	}
	return nil
}

type FollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerId string `protobuf:"bytes,1,opt,name=follower_id,json=followerId,proto3" json:"follower_id,omitempty"`
	FolloweeId string `protobuf:"bytes,2,opt,name=followee_id,json=followeeId,proto3" json:"followee_id,omitempty"`
}

func (x *FollowRequest) Reset() {
	*x = FollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequest) ProtoMessage() {}

func (x *FollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequest.ProtoReflect.Descriptor instead.
func (*FollowRequest) Descriptor() ([]byte, []int) {
	return file_social_graph_proto_rawDescGZIP(), []int{6}
}

func (x *FollowRequest) GetFollowerId() string {
	if x != nil {
		return x.FollowerId
	}
	return ""
}

func (x *FollowRequest) GetFolloweeId() string {
	if x != nil {
		return x.FolloweeId
	}
	return ""
}

type FollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *FollowResponse) Reset() {
	*x = FollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowResponse) ProtoMessage() {}

func (x *FollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowResponse.ProtoReflect.Descriptor instead.
func (*FollowResponse) Descriptor() ([]byte, []int) {
	return file_social_graph_proto_rawDescGZIP(), []int{7}
}

func (x *FollowResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type FollowManyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FollowerIds []string `protobuf:"bytes,2,rep,name=follower_ids,json=followerIds,proto3" json:"follower_ids,omitempty"`
	FolloweeIds []string `protobuf:"bytes,3,rep,name=followee_ids,json=followeeIds,proto3" json:"followee_ids,omitempty"`
}

func (x *FollowManyRequest) Reset() {
	*x = FollowManyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowManyRequest) ProtoMessage() {}

func (x *FollowManyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowManyRequest.ProtoReflect.Descriptor instead.
func (*FollowManyRequest) Descriptor() ([]byte, []int) {
	return file_social_graph_proto_rawDescGZIP(), []int{8}
}

func (x *FollowManyRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FollowManyRequest) GetFollowerIds() []string {
	if x != nil {
		return x.FollowerIds
	}
	return nil
}

func (x *FollowManyRequest) GetFolloweeIds() []string {
	if x != nil {
		return x.FolloweeIds
	}
	return nil
}

type FollowManyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *FollowManyResponse) Reset() {
	*x = FollowManyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowManyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowManyResponse) ProtoMessage() {}

func (x *FollowManyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowManyResponse.ProtoReflect.Descriptor instead.
func (*FollowManyResponse) Descriptor() ([]byte, []int) {
	return file_social_graph_proto_rawDescGZIP(), []int{9}
}

func (x *FollowManyResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_social_graph_proto protoreflect.FileDescriptor

var file_social_graph_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2c, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x2e, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x2e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x34, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x73, 0x22, 0x2e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x73, 0x22, 0x51, 0x0a, 0x0d, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x49, 0x64, 0x22, 0x2a,
	0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x72, 0x0a, 0x11, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x49, 0x64, 0x73, 0x22, 0x2e,
	0x0a, 0x12, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x98,
	0x03, 0x0a, 0x0b, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x4d,
	0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x20, 0x2e,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x65, 0x73, 0x12, 0x20, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x12, 0x1a, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x1e, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x69, 0x61, 0x6c, 0x69, 0x2d, 0x58, 0x69,
	0x6e, 0x67, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_social_graph_proto_rawDescOnce sync.Once
	file_social_graph_proto_rawDescData = file_social_graph_proto_rawDesc
)

func file_social_graph_proto_rawDescGZIP() []byte {
	file_social_graph_proto_rawDescOnce.Do(func() {
		file_social_graph_proto_rawDescData = protoimpl.X.CompressGZIP(file_social_graph_proto_rawDescData)
	})
	return file_social_graph_proto_rawDescData
}

var file_social_graph_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_social_graph_proto_goTypes = []interface{}{
	(*InsertUserRequest)(nil),    // 0: socialproto.InsertUserRequest
	(*InsertUserResponse)(nil),   // 1: socialproto.InsertUserResponse
	(*GetFollowersRequest)(nil),  // 2: socialproto.GetFollowersRequest
	(*GetFollowersResponse)(nil), // 3: socialproto.GetFollowersResponse
	(*GetFolloweesRequest)(nil),  // 4: socialproto.GetFolloweesRequest
	(*GetFolloweesResponse)(nil), // 5: socialproto.GetFolloweesResponse
	(*FollowRequest)(nil),        // 6: socialproto.FollowRequest
	(*FollowResponse)(nil),       // 7: socialproto.FollowResponse
	(*FollowManyRequest)(nil),    // 8: socialproto.FollowManyRequest
	(*FollowManyResponse)(nil),   // 9: socialproto.FollowManyResponse
}
var file_social_graph_proto_depIdxs = []int32{
	0, // 0: socialproto.SocialGraph.InsertUser:input_type -> socialproto.InsertUserRequest
	2, // 1: socialproto.SocialGraph.GetFollowers:input_type -> socialproto.GetFollowersRequest
	4, // 2: socialproto.SocialGraph.GetFollowees:input_type -> socialproto.GetFolloweesRequest
	6, // 3: socialproto.SocialGraph.Follow:input_type -> socialproto.FollowRequest
	8, // 4: socialproto.SocialGraph.FollowMany:input_type -> socialproto.FollowManyRequest
	1, // 5: socialproto.SocialGraph.InsertUser:output_type -> socialproto.InsertUserResponse
	3, // 6: socialproto.SocialGraph.GetFollowers:output_type -> socialproto.GetFollowersResponse
	5, // 7: socialproto.SocialGraph.GetFollowees:output_type -> socialproto.GetFolloweesResponse
	7, // 8: socialproto.SocialGraph.Follow:output_type -> socialproto.FollowResponse
	9, // 9: socialproto.SocialGraph.FollowMany:output_type -> socialproto.FollowManyResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_social_graph_proto_init() }
func file_social_graph_proto_init() {
	if File_social_graph_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_social_graph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertUserRequest); i {
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
		file_social_graph_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertUserResponse); i {
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
		file_social_graph_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowersRequest); i {
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
		file_social_graph_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowersResponse); i {
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
		file_social_graph_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFolloweesRequest); i {
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
		file_social_graph_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFolloweesResponse); i {
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
		file_social_graph_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRequest); i {
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
		file_social_graph_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowResponse); i {
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
		file_social_graph_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowManyRequest); i {
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
		file_social_graph_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowManyResponse); i {
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
			RawDescriptor: file_social_graph_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_social_graph_proto_goTypes,
		DependencyIndexes: file_social_graph_proto_depIdxs,
		MessageInfos:      file_social_graph_proto_msgTypes,
	}.Build()
	File_social_graph_proto = out.File
	file_social_graph_proto_rawDesc = nil
	file_social_graph_proto_goTypes = nil
	file_social_graph_proto_depIdxs = nil
}
