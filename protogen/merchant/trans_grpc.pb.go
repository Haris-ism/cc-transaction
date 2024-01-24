// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/merchant/trans.proto

package merchant

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

// TransServicesClient is the client API for TransServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransServicesClient interface {
	TransItems(ctx context.Context, in *ReqTransItems, opts ...grpc.CallOption) (*ResMerchantTransModel, error)
	CallbackTransItems(ctx context.Context, in *ReqCallbackItems, opts ...grpc.CallOption) (*ResMerchantCallbackModel, error)
}

type transServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewTransServicesClient(cc grpc.ClientConnInterface) TransServicesClient {
	return &transServicesClient{cc}
}

func (c *transServicesClient) TransItems(ctx context.Context, in *ReqTransItems, opts ...grpc.CallOption) (*ResMerchantTransModel, error) {
	out := new(ResMerchantTransModel)
	err := c.cc.Invoke(ctx, "/TransServices/TransItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transServicesClient) CallbackTransItems(ctx context.Context, in *ReqCallbackItems, opts ...grpc.CallOption) (*ResMerchantCallbackModel, error) {
	out := new(ResMerchantCallbackModel)
	err := c.cc.Invoke(ctx, "/TransServices/CallbackTransItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransServicesServer is the server API for TransServices service.
// All implementations must embed UnimplementedTransServicesServer
// for forward compatibility
type TransServicesServer interface {
	TransItems(context.Context, *ReqTransItems) (*ResMerchantTransModel, error)
	CallbackTransItems(context.Context, *ReqCallbackItems) (*ResMerchantCallbackModel, error)
	mustEmbedUnimplementedTransServicesServer()
}

// UnimplementedTransServicesServer must be embedded to have forward compatible implementations.
type UnimplementedTransServicesServer struct {
}

func (UnimplementedTransServicesServer) TransItems(context.Context, *ReqTransItems) (*ResMerchantTransModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransItems not implemented")
}
func (UnimplementedTransServicesServer) CallbackTransItems(context.Context, *ReqCallbackItems) (*ResMerchantCallbackModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallbackTransItems not implemented")
}
func (UnimplementedTransServicesServer) mustEmbedUnimplementedTransServicesServer() {}

// UnsafeTransServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransServicesServer will
// result in compilation errors.
type UnsafeTransServicesServer interface {
	mustEmbedUnimplementedTransServicesServer()
}

func RegisterTransServicesServer(s grpc.ServiceRegistrar, srv TransServicesServer) {
	s.RegisterService(&TransServices_ServiceDesc, srv)
}

func _TransServices_TransItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqTransItems)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransServicesServer).TransItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransServices/TransItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransServicesServer).TransItems(ctx, req.(*ReqTransItems))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransServices_CallbackTransItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCallbackItems)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransServicesServer).CallbackTransItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransServices/CallbackTransItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransServicesServer).CallbackTransItems(ctx, req.(*ReqCallbackItems))
	}
	return interceptor(ctx, in, info, handler)
}

// TransServices_ServiceDesc is the grpc.ServiceDesc for TransServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TransServices",
	HandlerType: (*TransServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransItems",
			Handler:    _TransServices_TransItems_Handler,
		},
		{
			MethodName: "CallbackTransItems",
			Handler:    _TransServices_CallbackTransItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/merchant/trans.proto",
}