// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: api.proto

package api

import (
	context "context"
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

// The CollectorIDRequest message for a collector ID.
type CollectorIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orchestrator string `protobuf:"bytes,2,opt,name=orchestrator,proto3" json:"orchestrator,omitempty"`
}

func (x *CollectorIDRequest) Reset() {
	*x = CollectorIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectorIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectorIDRequest) ProtoMessage() {}

func (x *CollectorIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectorIDRequest.ProtoReflect.Descriptor instead.
func (*CollectorIDRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *CollectorIDRequest) GetOrchestrator() string {
	if x != nil {
		return x.Orchestrator
	}
	return ""
}

// The CollectorIDReply response message from a collector ID request.
type CollectorIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CollectorIDReply) Reset() {
	*x = CollectorIDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectorIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectorIDReply) ProtoMessage() {}

func (x *CollectorIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectorIDReply.ProtoReflect.Descriptor instead.
func (*CollectorIDReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *CollectorIDReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x22, 0x38, 0x0a, 0x12, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72,
	0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0x59,
	0x0a, 0x16, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0b, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x2f, 0x6b, 0x38, 0x73, 0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x65, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_goTypes = []interface{}{
	(*CollectorIDRequest)(nil), // 0: api.CollectorIDRequest
	(*CollectorIDReply)(nil),   // 1: api.CollectorIDReply
}
var file_api_proto_depIdxs = []int32{
	0, // 0: api.CollectorSetController.CollectorID:input_type -> api.CollectorIDRequest
	1, // 1: api.CollectorSetController.CollectorID:output_type -> api.CollectorIDReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectorIDRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectorIDReply); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CollectorSetControllerClient is the client API for CollectorSetController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CollectorSetControllerClient interface {
	// Retrieves a collector ID.
	CollectorID(ctx context.Context, in *CollectorIDRequest, opts ...grpc.CallOption) (*CollectorIDReply, error)
}

type collectorSetControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectorSetControllerClient(cc grpc.ClientConnInterface) CollectorSetControllerClient {
	return &collectorSetControllerClient{cc}
}

func (c *collectorSetControllerClient) CollectorID(ctx context.Context, in *CollectorIDRequest, opts ...grpc.CallOption) (*CollectorIDReply, error) {
	out := new(CollectorIDReply)
	err := c.cc.Invoke(ctx, "/api.CollectorSetController/CollectorID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectorSetControllerServer is the server API for CollectorSetController service.
type CollectorSetControllerServer interface {
	// Retrieves a collector ID.
	CollectorID(context.Context, *CollectorIDRequest) (*CollectorIDReply, error)
}

// UnimplementedCollectorSetControllerServer can be embedded to have forward compatible implementations.
type UnimplementedCollectorSetControllerServer struct {
}

func (*UnimplementedCollectorSetControllerServer) CollectorID(context.Context, *CollectorIDRequest) (*CollectorIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectorID not implemented")
}

func RegisterCollectorSetControllerServer(s *grpc.Server, srv CollectorSetControllerServer) {
	s.RegisterService(&_CollectorSetController_serviceDesc, srv)
}

func _CollectorSetController_CollectorID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectorIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorSetControllerServer).CollectorID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CollectorSetController/CollectorID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorSetControllerServer).CollectorID(ctx, req.(*CollectorIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CollectorSetController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.CollectorSetController",
	HandlerType: (*CollectorSetControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CollectorID",
			Handler:    _CollectorSetController_CollectorID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
