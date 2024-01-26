// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gateway

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RpcGatewayClient is the client API for RpcGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcGatewayClient interface {
	GetGatewayInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GateWayInfoResponse, error)
}

type rpcGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcGatewayClient(cc grpc.ClientConnInterface) RpcGatewayClient {
	return &rpcGatewayClient{cc}
}

func (c *rpcGatewayClient) GetGatewayInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GateWayInfoResponse, error) {
	out := new(GateWayInfoResponse)
	err := c.cc.Invoke(ctx, "/gateway.RpcGateway/GetGatewayInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcGatewayServer is the server API for RpcGateway service.
// All implementations must embed UnimplementedRpcGatewayServer
// for forward compatibility
type RpcGatewayServer interface {
	GetGatewayInfo(context.Context, *emptypb.Empty) (*GateWayInfoResponse, error)
	mustEmbedUnimplementedRpcGatewayServer()
}

// UnimplementedRpcGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedRpcGatewayServer struct {
}

func (UnimplementedRpcGatewayServer) GetGatewayInfo(context.Context, *emptypb.Empty) (*GateWayInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGatewayInfo not implemented")
}
func (UnimplementedRpcGatewayServer) mustEmbedUnimplementedRpcGatewayServer() {}

// UnsafeRpcGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcGatewayServer will
// result in compilation errors.
type UnsafeRpcGatewayServer interface {
	mustEmbedUnimplementedRpcGatewayServer()
}

func RegisterRpcGatewayServer(s grpc.ServiceRegistrar, srv RpcGatewayServer) {
	s.RegisterService(&RpcGateway_ServiceDesc, srv)
}

func _RpcGateway_GetGatewayInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcGatewayServer).GetGatewayInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.RpcGateway/GetGatewayInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcGatewayServer).GetGatewayInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcGateway_ServiceDesc is the grpc.ServiceDesc for RpcGateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcGateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.RpcGateway",
	HandlerType: (*RpcGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGatewayInfo",
			Handler:    _RpcGateway_GetGatewayInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway/gateway.proto",
}