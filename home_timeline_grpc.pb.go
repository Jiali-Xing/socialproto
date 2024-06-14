// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: home_timeline.proto

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

// HomeTimelineClient is the client API for HomeTimeline service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomeTimelineClient interface {
	ReadHomeTimeline(ctx context.Context, in *ReadHomeTimelineRequest, opts ...grpc.CallOption) (*ReadHomeTimelineResponse, error)
	WriteHomeTimeline(ctx context.Context, in *WriteHomeTimelineRequest, opts ...grpc.CallOption) (*WriteHomeTimelineResponse, error)
}

type homeTimelineClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeTimelineClient(cc grpc.ClientConnInterface) HomeTimelineClient {
	return &homeTimelineClient{cc}
}

func (c *homeTimelineClient) ReadHomeTimeline(ctx context.Context, in *ReadHomeTimelineRequest, opts ...grpc.CallOption) (*ReadHomeTimelineResponse, error) {
	out := new(ReadHomeTimelineResponse)
	err := c.cc.Invoke(ctx, "/socialproto.HomeTimeline/ReadHomeTimeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeTimelineClient) WriteHomeTimeline(ctx context.Context, in *WriteHomeTimelineRequest, opts ...grpc.CallOption) (*WriteHomeTimelineResponse, error) {
	out := new(WriteHomeTimelineResponse)
	err := c.cc.Invoke(ctx, "/socialproto.HomeTimeline/WriteHomeTimeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeTimelineServer is the server API for HomeTimeline service.
// All implementations must embed UnimplementedHomeTimelineServer
// for forward compatibility
type HomeTimelineServer interface {
	ReadHomeTimeline(context.Context, *ReadHomeTimelineRequest) (*ReadHomeTimelineResponse, error)
	WriteHomeTimeline(context.Context, *WriteHomeTimelineRequest) (*WriteHomeTimelineResponse, error)
	mustEmbedUnimplementedHomeTimelineServer()
}

// UnimplementedHomeTimelineServer must be embedded to have forward compatible implementations.
type UnimplementedHomeTimelineServer struct {
}

func (UnimplementedHomeTimelineServer) ReadHomeTimeline(context.Context, *ReadHomeTimelineRequest) (*ReadHomeTimelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadHomeTimeline not implemented")
}
func (UnimplementedHomeTimelineServer) WriteHomeTimeline(context.Context, *WriteHomeTimelineRequest) (*WriteHomeTimelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteHomeTimeline not implemented")
}
func (UnimplementedHomeTimelineServer) mustEmbedUnimplementedHomeTimelineServer() {}

// UnsafeHomeTimelineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeTimelineServer will
// result in compilation errors.
type UnsafeHomeTimelineServer interface {
	mustEmbedUnimplementedHomeTimelineServer()
}

func RegisterHomeTimelineServer(s grpc.ServiceRegistrar, srv HomeTimelineServer) {
	s.RegisterService(&HomeTimeline_ServiceDesc, srv)
}

func _HomeTimeline_ReadHomeTimeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadHomeTimelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeTimelineServer).ReadHomeTimeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialproto.HomeTimeline/ReadHomeTimeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeTimelineServer).ReadHomeTimeline(ctx, req.(*ReadHomeTimelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeTimeline_WriteHomeTimeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteHomeTimelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeTimelineServer).WriteHomeTimeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/socialproto.HomeTimeline/WriteHomeTimeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeTimelineServer).WriteHomeTimeline(ctx, req.(*WriteHomeTimelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HomeTimeline_ServiceDesc is the grpc.ServiceDesc for HomeTimeline service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HomeTimeline_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "socialproto.HomeTimeline",
	HandlerType: (*HomeTimelineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadHomeTimeline",
			Handler:    _HomeTimeline_ReadHomeTimeline_Handler,
		},
		{
			MethodName: "WriteHomeTimeline",
			Handler:    _HomeTimeline_WriteHomeTimeline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "home_timeline.proto",
}