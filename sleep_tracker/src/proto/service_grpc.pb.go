// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// SleepTrackerServiceClient is the client API for SleepTrackerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SleepTrackerServiceClient interface {
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*ServerResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ServerResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*ServerResponse, error)
}

type sleepTrackerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSleepTrackerServiceClient(cc grpc.ClientConnInterface) SleepTrackerServiceClient {
	return &sleepTrackerServiceClient{cc}
}

func (c *sleepTrackerServiceClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/proto.SleepTrackerService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sleepTrackerServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/proto.SleepTrackerService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sleepTrackerServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/proto.SleepTrackerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SleepTrackerServiceServer is the server API for SleepTrackerService service.
// All implementations must embed UnimplementedSleepTrackerServiceServer
// for forward compatibility
type SleepTrackerServiceServer interface {
	Set(context.Context, *SetRequest) (*ServerResponse, error)
	Get(context.Context, *GetRequest) (*ServerResponse, error)
	Delete(context.Context, *DeleteRequest) (*ServerResponse, error)
	mustEmbedUnimplementedSleepTrackerServiceServer()
}

// UnimplementedSleepTrackerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSleepTrackerServiceServer struct {
}

func (UnimplementedSleepTrackerServiceServer) Set(context.Context, *SetRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedSleepTrackerServiceServer) Get(context.Context, *GetRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSleepTrackerServiceServer) Delete(context.Context, *DeleteRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSleepTrackerServiceServer) mustEmbedUnimplementedSleepTrackerServiceServer() {}

// UnsafeSleepTrackerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SleepTrackerServiceServer will
// result in compilation errors.
type UnsafeSleepTrackerServiceServer interface {
	mustEmbedUnimplementedSleepTrackerServiceServer()
}

func RegisterSleepTrackerServiceServer(s grpc.ServiceRegistrar, srv SleepTrackerServiceServer) {
	s.RegisterService(&SleepTrackerService_ServiceDesc, srv)
}

func _SleepTrackerService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SleepTrackerServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SleepTrackerService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SleepTrackerServiceServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SleepTrackerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SleepTrackerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SleepTrackerService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SleepTrackerServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SleepTrackerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SleepTrackerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SleepTrackerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SleepTrackerServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SleepTrackerService_ServiceDesc is the grpc.ServiceDesc for SleepTrackerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SleepTrackerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SleepTrackerService",
	HandlerType: (*SleepTrackerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _SleepTrackerService_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SleepTrackerService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SleepTrackerService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/proto/service.proto",
}
