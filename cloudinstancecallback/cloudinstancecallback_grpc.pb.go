// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cloudinstancecallback

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

// CloudInstanceCallBackServiceClient is the client API for CloudInstanceCallBackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudInstanceCallBackServiceClient interface {
	// 云实例状态回调 edge = c driver = s
	CloudInstanceStatueCallback(ctx context.Context, in *CloudInstanceStatueCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type cloudInstanceCallBackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudInstanceCallBackServiceClient(cc grpc.ClientConnInterface) CloudInstanceCallBackServiceClient {
	return &cloudInstanceCallBackServiceClient{cc}
}

func (c *cloudInstanceCallBackServiceClient) CloudInstanceStatueCallback(ctx context.Context, in *CloudInstanceStatueCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloudinstancecallback.CloudInstanceCallBackService/CloudInstanceStatueCallback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudInstanceCallBackServiceServer is the server API for CloudInstanceCallBackService service.
// All implementations must embed UnimplementedCloudInstanceCallBackServiceServer
// for forward compatibility
type CloudInstanceCallBackServiceServer interface {
	// 云实例状态回调 edge = c driver = s
	CloudInstanceStatueCallback(context.Context, *CloudInstanceStatueCallbackRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCloudInstanceCallBackServiceServer()
}

// UnimplementedCloudInstanceCallBackServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCloudInstanceCallBackServiceServer struct {
}

func (UnimplementedCloudInstanceCallBackServiceServer) CloudInstanceStatueCallback(context.Context, *CloudInstanceStatueCallbackRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloudInstanceStatueCallback not implemented")
}
func (UnimplementedCloudInstanceCallBackServiceServer) mustEmbedUnimplementedCloudInstanceCallBackServiceServer() {
}

// UnsafeCloudInstanceCallBackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudInstanceCallBackServiceServer will
// result in compilation errors.
type UnsafeCloudInstanceCallBackServiceServer interface {
	mustEmbedUnimplementedCloudInstanceCallBackServiceServer()
}

func RegisterCloudInstanceCallBackServiceServer(s grpc.ServiceRegistrar, srv CloudInstanceCallBackServiceServer) {
	s.RegisterService(&CloudInstanceCallBackService_ServiceDesc, srv)
}

func _CloudInstanceCallBackService_CloudInstanceStatueCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudInstanceStatueCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudInstanceCallBackServiceServer).CloudInstanceStatueCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudinstancecallback.CloudInstanceCallBackService/CloudInstanceStatueCallback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudInstanceCallBackServiceServer).CloudInstanceStatueCallback(ctx, req.(*CloudInstanceStatueCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudInstanceCallBackService_ServiceDesc is the grpc.ServiceDesc for CloudInstanceCallBackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudInstanceCallBackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloudinstancecallback.CloudInstanceCallBackService",
	HandlerType: (*CloudInstanceCallBackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CloudInstanceStatueCallback",
			Handler:    _CloudInstanceCallBackService_CloudInstanceStatueCallback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloudinstancecallback/cloudinstancecallback.proto",
}