// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: api/echo/proto/echo.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_echo_proto_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_echo_proto_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_api_echo_proto_echo_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_echo_proto_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_echo_proto_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_api_echo_proto_echo_proto_rawDescGZIP(), []int{1}
}

func (x *EchoResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_api_echo_proto_echo_proto protoreflect.FileDescriptor

var file_api_echo_proto_echo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x63, 0x68,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01,
	0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0x92, 0x41,
	0x13, 0x32, 0x11, 0x54, 0x68, 0x65, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x74, 0x6f, 0x20,
	0x65, 0x63, 0x68, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x46, 0x92, 0x41, 0x43,
	0x0a, 0x28, 0x2a, 0x0c, 0x45, 0x63, 0x68, 0x6f, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x32, 0x10, 0x54, 0x68, 0x65, 0x20, 0x65, 0x63, 0x68, 0x6f, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0xd2, 0x01, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x17, 0x12, 0x15, 0x7b, 0x20,
	0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x21,
	0x22, 0x20, 0x7d, 0x22, 0x85, 0x01, 0x0a, 0x0c, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x15, 0x92, 0x41, 0x12, 0x32, 0x10, 0x54, 0x68, 0x65, 0x20, 0x65, 0x63,
	0x68, 0x6f, 0x65, 0x64, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x48, 0x92, 0x41, 0x45, 0x0a, 0x2a, 0x2a, 0x0d, 0x45, 0x63, 0x68, 0x6f, 0x20, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x11, 0x54, 0x68, 0x65, 0x20, 0x65, 0x63, 0x68,
	0x6f, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0xd2, 0x01, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x32, 0x17, 0x12, 0x15, 0x7b, 0x20, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3a,
	0x20, 0x22, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x21, 0x22, 0x20, 0x7d, 0x32, 0x8f, 0x03, 0x0a, 0x04,
	0x45, 0x63, 0x68, 0x6f, 0x12, 0xca, 0x01, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x11, 0x2e,
	0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9a, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x83, 0x01, 0x12, 0x04,
	0x45, 0x63, 0x68, 0x6f, 0x1a, 0x67, 0x54, 0x68, 0x65, 0x20, 0x45, 0x63, 0x68, 0x6f, 0x20, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20,
	0x61, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x68, 0x61, 0x76, 0x65,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x65, 0x64, 0x20, 0x76, 0x65, 0x72, 0x62, 0x61, 0x74, 0x69, 0x6d, 0x2e, 0x62, 0x12, 0x0a,
	0x10, 0x0a, 0x06, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x12, 0x06, 0x0a, 0x04, 0x65, 0x63, 0x68,
	0x6f, 0x1a, 0xb9, 0x01, 0x92, 0x41, 0xb5, 0x01, 0x12, 0x64, 0x54, 0x68, 0x65, 0x20, 0x45, 0x63,
	0x68, 0x6f, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x73, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x77, 0x69, 0x74, 0x68,
	0x20, 0x61, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x68, 0x61, 0x76,
	0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x65, 0x64, 0x20, 0x76, 0x65, 0x72, 0x62, 0x61, 0x74, 0x69, 0x6d, 0x2e, 0x1a, 0x4d,
	0x0a, 0x27, 0x52, 0x65, 0x61, 0x64, 0x20, 0x75, 0x70, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x65, 0x63, 0x68, 0x6f, 0x20, 0x61, 0x63, 0x6f, 0x75, 0x73, 0x74, 0x69, 0x63, 0x20, 0x70,
	0x68, 0x65, 0x6e, 0x6f, 0x6d, 0x65, 0x6e, 0x6f, 0x6e, 0x12, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x65, 0x6e, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x77, 0x69, 0x6b, 0x69, 0x2f, 0x45, 0x63, 0x68, 0x6f, 0x42, 0x10, 0x5a,
	0x0e, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_echo_proto_echo_proto_rawDescOnce sync.Once
	file_api_echo_proto_echo_proto_rawDescData = file_api_echo_proto_echo_proto_rawDesc
)

func file_api_echo_proto_echo_proto_rawDescGZIP() []byte {
	file_api_echo_proto_echo_proto_rawDescOnce.Do(func() {
		file_api_echo_proto_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_echo_proto_echo_proto_rawDescData)
	})
	return file_api_echo_proto_echo_proto_rawDescData
}

var file_api_echo_proto_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_echo_proto_echo_proto_goTypes = []interface{}{
	(*EchoRequest)(nil),  // 0: echo.EchoRequest
	(*EchoResponse)(nil), // 1: echo.EchoResponse
}
var file_api_echo_proto_echo_proto_depIdxs = []int32{
	0, // 0: echo.Echo.Echo:input_type -> echo.EchoRequest
	1, // 1: echo.Echo.Echo:output_type -> echo.EchoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_echo_proto_echo_proto_init() }
func file_api_echo_proto_echo_proto_init() {
	if File_api_echo_proto_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_echo_proto_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
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
		file_api_echo_proto_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResponse); i {
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
			RawDescriptor: file_api_echo_proto_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_echo_proto_echo_proto_goTypes,
		DependencyIndexes: file_api_echo_proto_echo_proto_depIdxs,
		MessageInfos:      file_api_echo_proto_echo_proto_msgTypes,
	}.Build()
	File_api_echo_proto_echo_proto = out.File
	file_api_echo_proto_echo_proto_rawDesc = nil
	file_api_echo_proto_echo_proto_goTypes = nil
	file_api_echo_proto_echo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	// The Echo operation enables users to probe the server with a value and have
	// the value returned verbatim.
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type echoClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoClient(cc grpc.ClientConnInterface) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/echo.Echo/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	// The Echo operation enables users to probe the server with a value and have
	// the value returned verbatim.
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

// UnimplementedEchoServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (*UnimplementedEchoServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Echo_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/echo/proto/echo.proto",
}