// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cloudinstance

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

// CloudInstanceServiceClient is the client API for CloudInstanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudInstanceServiceClient interface {
	// 查询云服务示例信息 edge = s driver = c
	QueryCloudInstanceByPlatform(ctx context.Context, in *QueryCloudInstanceByPlatformRequest, opts ...grpc.CallOption) (*QueryCloudInstanceByPlatformResponse, error)
	// 上报驱动使用云实例
	DriverReportPlatformInfo(ctx context.Context, in *DriverReportPlatformInfoRequest, opts ...grpc.CallOption) (*DriverReportPlatformInfoResponse, error)
}

type cloudInstanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudInstanceServiceClient(cc grpc.ClientConnInterface) CloudInstanceServiceClient {
	return &cloudInstanceServiceClient{cc}
}

func (c *cloudInstanceServiceClient) QueryCloudInstanceByPlatform(ctx context.Context, in *QueryCloudInstanceByPlatformRequest, opts ...grpc.CallOption) (*QueryCloudInstanceByPlatformResponse, error) {
	out := new(QueryCloudInstanceByPlatformResponse)
	err := c.cc.Invoke(ctx, "/cloudinstance.CloudInstanceService/QueryCloudInstanceByPlatform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudInstanceServiceClient) DriverReportPlatformInfo(ctx context.Context, in *DriverReportPlatformInfoRequest, opts ...grpc.CallOption) (*DriverReportPlatformInfoResponse, error) {
	out := new(DriverReportPlatformInfoResponse)
	err := c.cc.Invoke(ctx, "/cloudinstance.CloudInstanceService/DriverReportPlatformInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudInstanceServiceServer is the server API for CloudInstanceService service.
// All implementations must embed UnimplementedCloudInstanceServiceServer
// for forward compatibility
type CloudInstanceServiceServer interface {
	// 查询云服务示例信息 edge = s driver = c
	QueryCloudInstanceByPlatform(context.Context, *QueryCloudInstanceByPlatformRequest) (*QueryCloudInstanceByPlatformResponse, error)
	// 上报驱动使用云实例
	DriverReportPlatformInfo(context.Context, *DriverReportPlatformInfoRequest) (*DriverReportPlatformInfoResponse, error)
	mustEmbedUnimplementedCloudInstanceServiceServer()
}

// UnimplementedCloudInstanceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCloudInstanceServiceServer struct {
}

func (UnimplementedCloudInstanceServiceServer) QueryCloudInstanceByPlatform(context.Context, *QueryCloudInstanceByPlatformRequest) (*QueryCloudInstanceByPlatformResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCloudInstanceByPlatform not implemented")
}
func (UnimplementedCloudInstanceServiceServer) DriverReportPlatformInfo(context.Context, *DriverReportPlatformInfoRequest) (*DriverReportPlatformInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DriverReportPlatformInfo not implemented")
}
func (UnimplementedCloudInstanceServiceServer) mustEmbedUnimplementedCloudInstanceServiceServer() {}

// UnsafeCloudInstanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudInstanceServiceServer will
// result in compilation errors.
type UnsafeCloudInstanceServiceServer interface {
	mustEmbedUnimplementedCloudInstanceServiceServer()
}

func RegisterCloudInstanceServiceServer(s grpc.ServiceRegistrar, srv CloudInstanceServiceServer) {
	s.RegisterService(&CloudInstanceService_ServiceDesc, srv)
}

func _CloudInstanceService_QueryCloudInstanceByPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCloudInstanceByPlatformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudInstanceServiceServer).QueryCloudInstanceByPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudinstance.CloudInstanceService/QueryCloudInstanceByPlatform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudInstanceServiceServer).QueryCloudInstanceByPlatform(ctx, req.(*QueryCloudInstanceByPlatformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudInstanceService_DriverReportPlatformInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverReportPlatformInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudInstanceServiceServer).DriverReportPlatformInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudinstance.CloudInstanceService/DriverReportPlatformInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudInstanceServiceServer).DriverReportPlatformInfo(ctx, req.(*DriverReportPlatformInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudInstanceService_ServiceDesc is the grpc.ServiceDesc for CloudInstanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudInstanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloudinstance.CloudInstanceService",
	HandlerType: (*CloudInstanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryCloudInstanceByPlatform",
			Handler:    _CloudInstanceService_QueryCloudInstanceByPlatform_Handler,
		},
		{
			MethodName: "DriverReportPlatformInfo",
			Handler:    _CloudInstanceService_DriverReportPlatformInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloudinstance/cloudinstance.proto",
}
