// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: cache.proto

package cache

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

// Enum value maps for UploadStatusCode.
var (
	UploadStatusCode_name = map[int32]string{
		0: "Unknown",
		1: "Ok",
		2: "Failed",
	}
	UploadStatusCode_value = map[string]int32{
		"Unknown": 0,
		"Ok":      1,
		"Failed":  2,
	}
)

func (x UploadStatusCode) Enum() *UploadStatusCode {
	p := new(UploadStatusCode)
	*p = x
	return p
}

func (x UploadStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_cache_proto_enumTypes[0].Descriptor()
}

func (UploadStatusCode) Type() protoreflect.EnumType {
	return &file_cache_proto_enumTypes[0]
}

func (x UploadStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadStatusCode.Descriptor instead.
func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{0}
}

// Used for our healthcheck/status check
type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{0}
}

func (x *StatusRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{1}
}

func (x *StatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Used for file transer
type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{2}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type UploadStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string           `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code    UploadStatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=cache.UploadStatusCode" json:"Code,omitempty"`
}

func (x *UploadStatus) Reset() {
	*x = UploadStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadStatus) ProtoMessage() {}

func (x *UploadStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadStatus.ProtoReflect.Descriptor instead.
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return file_cache_proto_rawDescGZIP(), []int{3}
}

func (x *UploadStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UploadStatus) GetCode() UploadStatusCode {
	if x != nil {
		return x.Code
	}
	return UploadStatusCode_Unknown
}

var File_cache_proto protoreflect.FileDescriptor

var file_cache_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x21, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x55, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x2a,
	0x33, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00,
	0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x10, 0x02, 0x32, 0x77, 0x0a, 0x08, 0x43, 0x61, 0x63, 0x68, 0x65, 0x48, 0x75, 0x62,
	0x12, 0x3a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x2e,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0c, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x13, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cache_proto_rawDescOnce sync.Once
	file_cache_proto_rawDescData = file_cache_proto_rawDesc
)

func file_cache_proto_rawDescGZIP() []byte {
	file_cache_proto_rawDescOnce.Do(func() {
		file_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_cache_proto_rawDescData)
	})
	return file_cache_proto_rawDescData
}

var file_cache_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cache_proto_goTypes = []interface{}{
	(UploadStatusCode)(0),  // 0: cache.UploadStatusCode
	(*StatusRequest)(nil),  // 1: cache.StatusRequest
	(*StatusResponse)(nil), // 2: cache.StatusResponse
	(*Chunk)(nil),          // 3: cache.Chunk
	(*UploadStatus)(nil),   // 4: cache.UploadStatus
}
var file_cache_proto_depIdxs = []int32{
	0, // 0: cache.UploadStatus.Code:type_name -> cache.UploadStatusCode
	1, // 1: cache.CacheHub.GetStatus:input_type -> cache.StatusRequest
	3, // 2: cache.CacheHub.Upload:input_type -> cache.Chunk
	2, // 3: cache.CacheHub.GetStatus:output_type -> cache.StatusResponse
	4, // 4: cache.CacheHub.Upload:output_type -> cache.UploadStatus
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cache_proto_init() }
func file_cache_proto_init() {
	if File_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_cache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_cache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadStatus); i {
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
			RawDescriptor: file_cache_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cache_proto_goTypes,
		DependencyIndexes: file_cache_proto_depIdxs,
		EnumInfos:         file_cache_proto_enumTypes,
		MessageInfos:      file_cache_proto_msgTypes,
	}.Build()
	File_cache_proto = out.File
	file_cache_proto_rawDesc = nil
	file_cache_proto_goTypes = nil
	file_cache_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CacheHubClient is the client API for CacheHub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CacheHubClient interface {
	GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Upload(ctx context.Context, opts ...grpc.CallOption) (CacheHub_UploadClient, error)
}

type cacheHubClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheHubClient(cc grpc.ClientConnInterface) CacheHubClient {
	return &cacheHubClient{cc}
}

func (c *cacheHubClient) GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/cache.CacheHub/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheHubClient) Upload(ctx context.Context, opts ...grpc.CallOption) (CacheHub_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CacheHub_serviceDesc.Streams[0], "/cache.CacheHub/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &cacheHubUploadClient{stream}
	return x, nil
}

type CacheHub_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type cacheHubUploadClient struct {
	grpc.ClientStream
}

func (x *cacheHubUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cacheHubUploadClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CacheHubServer is the server API for CacheHub service.
type CacheHubServer interface {
	GetStatus(context.Context, *StatusRequest) (*StatusResponse, error)
	Upload(CacheHub_UploadServer) error
}

// UnimplementedCacheHubServer can be embedded to have forward compatible implementations.
type UnimplementedCacheHubServer struct {
}

func (*UnimplementedCacheHubServer) GetStatus(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (*UnimplementedCacheHubServer) Upload(CacheHub_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterCacheHubServer(s *grpc.Server, srv CacheHubServer) {
	s.RegisterService(&_CacheHub_serviceDesc, srv)
}

func _CacheHub_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheHubServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache.CacheHub/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheHubServer).GetStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheHub_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CacheHubServer).Upload(&cacheHubUploadServer{stream})
}

type CacheHub_UploadServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type cacheHubUploadServer struct {
	grpc.ServerStream
}

func (x *cacheHubUploadServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cacheHubUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CacheHub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cache.CacheHub",
	HandlerType: (*CacheHubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _CacheHub_GetStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _CacheHub_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "cache.proto",
}