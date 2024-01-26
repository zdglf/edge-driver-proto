// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: cloudinstance/cloudinstance.proto

package cloudinstance

import (
	drivercommon "github.com/zdglf/edge-driver-proto/drivercommon"
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

type CloudInstanceStatus int32

const (
	CloudInstanceStatus_Stop  CloudInstanceStatus = 0
	CloudInstanceStatus_Start CloudInstanceStatus = 1
	CloudInstanceStatus_Error CloudInstanceStatus = 2
)

// Enum value maps for CloudInstanceStatus.
var (
	CloudInstanceStatus_name = map[int32]string{
		0: "Stop",
		1: "Start",
		2: "Error",
	}
	CloudInstanceStatus_value = map[string]int32{
		"Stop":  0,
		"Start": 1,
		"Error": 2,
	}
)

func (x CloudInstanceStatus) Enum() *CloudInstanceStatus {
	p := new(CloudInstanceStatus)
	*p = x
	return p
}

func (x CloudInstanceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CloudInstanceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_cloudinstance_cloudinstance_proto_enumTypes[0].Descriptor()
}

func (CloudInstanceStatus) Type() protoreflect.EnumType {
	return &file_cloudinstance_cloudinstance_proto_enumTypes[0]
}

func (x CloudInstanceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CloudInstanceStatus.Descriptor instead.
func (CloudInstanceStatus) EnumDescriptor() ([]byte, []int) {
	return file_cloudinstance_cloudinstance_proto_rawDescGZIP(), []int{0}
}

type QueryCloudInstanceByPlatformRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IotPlatform drivercommon.IotPlatform `protobuf:"varint,1,opt,name=iotPlatform,proto3,enum=drivercommon.IotPlatform" json:"iotPlatform,omitempty"`
}

func (x *QueryCloudInstanceByPlatformRequest) Reset() {
	*x = QueryCloudInstanceByPlatformRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudinstance_cloudinstance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCloudInstanceByPlatformRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCloudInstanceByPlatformRequest) ProtoMessage() {}

func (x *QueryCloudInstanceByPlatformRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudinstance_cloudinstance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryCloudInstanceByPlatformRequest.ProtoReflect.Descriptor instead.
func (*QueryCloudInstanceByPlatformRequest) Descriptor() ([]byte, []int) {
	return file_cloudinstance_cloudinstance_proto_rawDescGZIP(), []int{0}
}

func (x *QueryCloudInstanceByPlatformRequest) GetIotPlatform() drivercommon.IotPlatform {
	if x != nil {
		return x.IotPlatform
	}
	return drivercommon.IotPlatform(0)
}

type QueryCloudInstanceByPlatformResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudInstanceId   string              `protobuf:"bytes,1,opt,name=CloudInstanceId,proto3" json:"CloudInstanceId,omitempty"`
	Address           string              `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	CloudInstanceName string              `protobuf:"bytes,3,opt,name=CloudInstanceName,proto3" json:"CloudInstanceName,omitempty"`
	Status            CloudInstanceStatus `protobuf:"varint,4,opt,name=status,proto3,enum=cloudinstance.CloudInstanceStatus" json:"status,omitempty"`
}

func (x *QueryCloudInstanceByPlatformResponse) Reset() {
	*x = QueryCloudInstanceByPlatformResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudinstance_cloudinstance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCloudInstanceByPlatformResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCloudInstanceByPlatformResponse) ProtoMessage() {}

func (x *QueryCloudInstanceByPlatformResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloudinstance_cloudinstance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryCloudInstanceByPlatformResponse.ProtoReflect.Descriptor instead.
func (*QueryCloudInstanceByPlatformResponse) Descriptor() ([]byte, []int) {
	return file_cloudinstance_cloudinstance_proto_rawDescGZIP(), []int{1}
}

func (x *QueryCloudInstanceByPlatformResponse) GetCloudInstanceId() string {
	if x != nil {
		return x.CloudInstanceId
	}
	return ""
}

func (x *QueryCloudInstanceByPlatformResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *QueryCloudInstanceByPlatformResponse) GetCloudInstanceName() string {
	if x != nil {
		return x.CloudInstanceName
	}
	return ""
}

func (x *QueryCloudInstanceByPlatformResponse) GetStatus() CloudInstanceStatus {
	if x != nil {
		return x.Status
	}
	return CloudInstanceStatus_Stop
}

type DriverReportPlatformInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverInstanceId string                   `protobuf:"bytes,1,opt,name=driverInstanceId,proto3" json:"driverInstanceId,omitempty"`
	IotPlatform      drivercommon.IotPlatform `protobuf:"varint,2,opt,name=iotPlatform,proto3,enum=drivercommon.IotPlatform" json:"iotPlatform,omitempty"`
}

func (x *DriverReportPlatformInfoRequest) Reset() {
	*x = DriverReportPlatformInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudinstance_cloudinstance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverReportPlatformInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverReportPlatformInfoRequest) ProtoMessage() {}

func (x *DriverReportPlatformInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudinstance_cloudinstance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverReportPlatformInfoRequest.ProtoReflect.Descriptor instead.
func (*DriverReportPlatformInfoRequest) Descriptor() ([]byte, []int) {
	return file_cloudinstance_cloudinstance_proto_rawDescGZIP(), []int{2}
}

func (x *DriverReportPlatformInfoRequest) GetDriverInstanceId() string {
	if x != nil {
		return x.DriverInstanceId
	}
	return ""
}

func (x *DriverReportPlatformInfoRequest) GetIotPlatform() drivercommon.IotPlatform {
	if x != nil {
		return x.IotPlatform
	}
	return drivercommon.IotPlatform(0)
}

type DriverReportPlatformInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *drivercommon.CommonResponse `protobuf:"bytes,1,opt,name=baseResponse,proto3" json:"baseResponse,omitempty"`
}

func (x *DriverReportPlatformInfoResponse) Reset() {
	*x = DriverReportPlatformInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudinstance_cloudinstance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverReportPlatformInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverReportPlatformInfoResponse) ProtoMessage() {}

func (x *DriverReportPlatformInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloudinstance_cloudinstance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverReportPlatformInfoResponse.ProtoReflect.Descriptor instead.
func (*DriverReportPlatformInfoResponse) Descriptor() ([]byte, []int) {
	return file_cloudinstance_cloudinstance_proto_rawDescGZIP(), []int{3}
}

func (x *DriverReportPlatformInfoResponse) GetBaseResponse() *drivercommon.CommonResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

var File_cloudinstance_cloudinstance_proto protoreflect.FileDescriptor

var file_cloudinstance_cloudinstance_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x1a, 0x19, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a,
	0x23, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x69, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x52, 0x0b, 0x69, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x22, 0xd4, 0x01, 0x0a, 0x24, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c,
	0x0a, 0x11, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x1f, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x69, 0x6f, 0x74, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6f, 0x74,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x0b, 0x69, 0x6f, 0x74, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x64, 0x0a, 0x20, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x35, 0x0a, 0x13, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x10, 0x02, 0x32, 0xa1, 0x02, 0x0a, 0x14, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x1c,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x32, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42,
	0x79, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x33, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7d, 0x0a, 0x18, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x64, 0x67, 0x6c, 0x66, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2d,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cloudinstance_cloudinstance_proto_rawDescOnce sync.Once
	file_cloudinstance_cloudinstance_proto_rawDescData = file_cloudinstance_cloudinstance_proto_rawDesc
)

func file_cloudinstance_cloudinstance_proto_rawDescGZIP() []byte {
	file_cloudinstance_cloudinstance_proto_rawDescOnce.Do(func() {
		file_cloudinstance_cloudinstance_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudinstance_cloudinstance_proto_rawDescData)
	})
	return file_cloudinstance_cloudinstance_proto_rawDescData
}

var file_cloudinstance_cloudinstance_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloudinstance_cloudinstance_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloudinstance_cloudinstance_proto_goTypes = []interface{}{
	(CloudInstanceStatus)(0),                     // 0: cloudinstance.CloudInstanceStatus
	(*QueryCloudInstanceByPlatformRequest)(nil),  // 1: cloudinstance.QueryCloudInstanceByPlatformRequest
	(*QueryCloudInstanceByPlatformResponse)(nil), // 2: cloudinstance.QueryCloudInstanceByPlatformResponse
	(*DriverReportPlatformInfoRequest)(nil),      // 3: cloudinstance.DriverReportPlatformInfoRequest
	(*DriverReportPlatformInfoResponse)(nil),     // 4: cloudinstance.DriverReportPlatformInfoResponse
	(drivercommon.IotPlatform)(0),                // 5: drivercommon.IotPlatform
	(*drivercommon.CommonResponse)(nil),          // 6: drivercommon.CommonResponse
}
var file_cloudinstance_cloudinstance_proto_depIdxs = []int32{
	5, // 0: cloudinstance.QueryCloudInstanceByPlatformRequest.iotPlatform:type_name -> drivercommon.IotPlatform
	0, // 1: cloudinstance.QueryCloudInstanceByPlatformResponse.status:type_name -> cloudinstance.CloudInstanceStatus
	5, // 2: cloudinstance.DriverReportPlatformInfoRequest.iotPlatform:type_name -> drivercommon.IotPlatform
	6, // 3: cloudinstance.DriverReportPlatformInfoResponse.baseResponse:type_name -> drivercommon.CommonResponse
	1, // 4: cloudinstance.CloudInstanceService.QueryCloudInstanceByPlatform:input_type -> cloudinstance.QueryCloudInstanceByPlatformRequest
	3, // 5: cloudinstance.CloudInstanceService.DriverReportPlatformInfo:input_type -> cloudinstance.DriverReportPlatformInfoRequest
	2, // 6: cloudinstance.CloudInstanceService.QueryCloudInstanceByPlatform:output_type -> cloudinstance.QueryCloudInstanceByPlatformResponse
	4, // 7: cloudinstance.CloudInstanceService.DriverReportPlatformInfo:output_type -> cloudinstance.DriverReportPlatformInfoResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cloudinstance_cloudinstance_proto_init() }
func file_cloudinstance_cloudinstance_proto_init() {
	if File_cloudinstance_cloudinstance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudinstance_cloudinstance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCloudInstanceByPlatformRequest); i {
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
		file_cloudinstance_cloudinstance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCloudInstanceByPlatformResponse); i {
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
		file_cloudinstance_cloudinstance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverReportPlatformInfoRequest); i {
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
		file_cloudinstance_cloudinstance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverReportPlatformInfoResponse); i {
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
			RawDescriptor: file_cloudinstance_cloudinstance_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloudinstance_cloudinstance_proto_goTypes,
		DependencyIndexes: file_cloudinstance_cloudinstance_proto_depIdxs,
		EnumInfos:         file_cloudinstance_cloudinstance_proto_enumTypes,
		MessageInfos:      file_cloudinstance_cloudinstance_proto_msgTypes,
	}.Build()
	File_cloudinstance_cloudinstance_proto = out.File
	file_cloudinstance_cloudinstance_proto_rawDesc = nil
	file_cloudinstance_cloudinstance_proto_goTypes = nil
	file_cloudinstance_cloudinstance_proto_depIdxs = nil
}
