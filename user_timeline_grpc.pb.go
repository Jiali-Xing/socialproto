// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: user_timeline.proto

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

// UserTimelineClient is the client API for UserTimeline service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserTimelineClient interface {
	ReadUserTimeline(ctx context.Context, in *ReadUserTimelineRequest, opts ...grpc.CallOption) (*ReadUserTimelineResponse, error)
	WriteUserTimeline(ctx context.Context, in *WriteUserTimelineRequest, opts ...grpc.CallOption) (*WriteUserTimelineResponse, error)
}

type userTimelineClient struct {
	cc grpc.ClientConnInterface
}

func NewUserTimelineClient(cc grpc.ClientConnInterface) UserTimelineClient {
	return &userTimelineClient{cc}
}

func (c *userTimelineClient) ReadUserTimeline(ctx context.Context, in *ReadUserTimelineRequest, opts ...grpc.CallOption) (*ReadUserTimelineResponse, error) {
	out := new(ReadUserTimelineResponse)
	err := c.cc.Invoke(ctx, "/socialproto.UserTimeline/ReadUserTimeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTimelineClient) WriteUserTimeline(ctx context.Context, in *WriteUserTimelineRequest, opts ...grpc.CallOption) (*WriteUserTimelineResponse, error) {
	out := new(WriteUserTimelineResponse)
	err := c.cc.Invoke(ctx, "/socialproto.UserTimeline/WriteUserTimeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserTimelineServer is the server API for UserTimeline service.
// All implementations must embed UnimplementedUserTimelineServer
// for forward compatibility
type UserTimelineServer interface {
	ReadUserTimeline(context.Context, *ReadUserTimelineRequest) (*ReadUserTimelineResponse, error)
	WriteUserTimeline(context.Context, *WriteUserTimelineRequest) (*WriteUserTimelineResponse, error)
	mustEmbedUnimplementedUserTimelineServer()
}

// UnimplementedUserTimelineServer must be embedded to have forward compatible implementations.
type UnimplementedUserTimelineServer struct {
}

func (UnimplementedUserTimelineServer) ReadUserTimeline(context.Context, *ReadUserTimelineRequest) (*ReadUserTimelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUserTimeline not implemented")
}
func (UnimplementedUserTimelineServer) WriteUserTimeline(context.Context, *WriteUserTimelineRequest) (*WriteUserTimelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteUserTimeline not implemented")
}
func (UnimplementedUserTimelineServer) mustEmbedUnimplementedUserTimelineServer() {}

// UnsafeUserTimelineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserTimelineServer will
// result in compilation errors.
type UnsafeUserTimelineServer interface {
	mustEmbedUnimplementedUserTimelineServer()
}

func RegisterUserTimelineServer(s grpc.ServiceRegistrar, srv UserTimelineServer) {
	s.RegisterService(&UserTimeline_ServiceDesc, srv)
}

func _UserTimeline_ReadUserTimeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadUserTimelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTimelineServer).ReadUserTimeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialproto.UserTimeline/ReadUserTimeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTimelineServer).ReadUserTimeline(ctx, req.(*ReadUserTimelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserTimeline_WriteUserTimeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteUserTimelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTimelineServer).WriteUserTimeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialproto.UserTimeline/WriteUserTimeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTimelineServer).WriteUserTimeline(ctx, req.(*WriteUserTimelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserTimeline_ServiceDesc is the grpc.ServiceDesc for UserTimeline service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserTimeline_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "socialproto.UserTimeline",
	HandlerType: (*UserTimelineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadUserTimeline",
			Handler:    _UserTimeline_ReadUserTimeline_Handler,
		},
		{
			MethodName: "WriteUserTimeline",
			Handler:    _UserTimeline_WriteUserTimeline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_timeline.proto",
}
