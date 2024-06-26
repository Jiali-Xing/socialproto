// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: post_storage.proto

package socialproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PostStorageClient is the client API for PostStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostStorageClient interface {
	StorePost(ctx context.Context, in *StorePostRequest, opts ...grpc.CallOption) (*StorePostResponse, error)
	StorePostMulti(ctx context.Context, in *StorePostMultiRequest, opts ...grpc.CallOption) (*StorePostMultiResponse, error)
	ReadPost(ctx context.Context, in *ReadPostRequest, opts ...grpc.CallOption) (*ReadPostResponse, error)
	ReadPosts(ctx context.Context, in *ReadPostsRequest, opts ...grpc.CallOption) (*ReadPostsResponse, error)
}

type postStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewPostStorageClient(cc grpc.ClientConnInterface) PostStorageClient {
	return &postStorageClient{cc}
}

func (c *postStorageClient) StorePost(ctx context.Context, in *StorePostRequest, opts ...grpc.CallOption) (*StorePostResponse, error) {
	out := new(StorePostResponse)
	err := c.cc.Invoke(ctx, "/socialproto.PostStorage/StorePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postStorageClient) StorePostMulti(ctx context.Context, in *StorePostMultiRequest, opts ...grpc.CallOption) (*StorePostMultiResponse, error) {
	out := new(StorePostMultiResponse)
	err := c.cc.Invoke(ctx, "/socialproto.PostStorage/StorePostMulti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postStorageClient) ReadPost(ctx context.Context, in *ReadPostRequest, opts ...grpc.CallOption) (*ReadPostResponse, error) {
	out := new(ReadPostResponse)
	err := c.cc.Invoke(ctx, "/socialproto.PostStorage/ReadPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postStorageClient) ReadPosts(ctx context.Context, in *ReadPostsRequest, opts ...grpc.CallOption) (*ReadPostsResponse, error) {
	out := new(ReadPostsResponse)
	err := c.cc.Invoke(ctx, "/socialproto.PostStorage/ReadPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostStorageServer is the server API for PostStorage service.
// All implementations must embed UnimplementedPostStorageServer
// for forward compatibility
type PostStorageServer interface {
	StorePost(context.Context, *StorePostRequest) (*StorePostResponse, error)
	StorePostMulti(context.Context, *StorePostMultiRequest) (*StorePostMultiResponse, error)
	ReadPost(context.Context, *ReadPostRequest) (*ReadPostResponse, error)
	ReadPosts(context.Context, *ReadPostsRequest) (*ReadPostsResponse, error)
	mustEmbedUnimplementedPostStorageServer()
}

// UnimplementedPostStorageServer must be embedded to have forward compatible implementations.
type UnimplementedPostStorageServer struct {
}

func (UnimplementedPostStorageServer) StorePost(context.Context, *StorePostRequest) (*StorePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorePost not implemented")
}
func (UnimplementedPostStorageServer) StorePostMulti(context.Context, *StorePostMultiRequest) (*StorePostMultiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorePostMulti not implemented")
}
func (UnimplementedPostStorageServer) ReadPost(context.Context, *ReadPostRequest) (*ReadPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPost not implemented")
}
func (UnimplementedPostStorageServer) ReadPosts(context.Context, *ReadPostsRequest) (*ReadPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPosts not implemented")
}
func (UnimplementedPostStorageServer) mustEmbedUnimplementedPostStorageServer() {}

// UnsafePostStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostStorageServer will
// result in compilation errors.
type UnsafePostStorageServer interface {
	mustEmbedUnimplementedPostStorageServer()
}

func RegisterPostStorageServer(s grpc.ServiceRegistrar, srv PostStorageServer) {
	s.RegisterService(&PostStorage_ServiceDesc, srv)
}

func _PostStorage_StorePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostStorageServer).StorePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialproto.PostStorage/StorePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostStorageServer).StorePost(ctx, req.(*StorePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostStorage_StorePostMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorePostMultiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostStorageServer).StorePostMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialproto.PostStorage/StorePostMulti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostStorageServer).StorePostMulti(ctx, req.(*StorePostMultiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostStorage_ReadPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostStorageServer).ReadPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialproto.PostStorage/ReadPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostStorageServer).ReadPost(ctx, req.(*ReadPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostStorage_ReadPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostStorageServer).ReadPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialproto.PostStorage/ReadPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostStorageServer).ReadPosts(ctx, req.(*ReadPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostStorage_ServiceDesc is the grpc.ServiceDesc for PostStorage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostStorage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "socialproto.PostStorage",
	HandlerType: (*PostStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StorePost",
			Handler:    _PostStorage_StorePost_Handler,
		},
		{
			MethodName: "StorePostMulti",
			Handler:    _PostStorage_StorePostMulti_Handler,
		},
		{
			MethodName: "ReadPost",
			Handler:    _PostStorage_ReadPost_Handler,
		},
		{
			MethodName: "ReadPosts",
			Handler:    _PostStorage_ReadPosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post_storage.proto",
}
