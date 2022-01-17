// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.19.3
// source: mygrpc/bullet.proto

package mygrpc

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

type SearchAnimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SearchAnimeRequest) Reset() {
	*x = SearchAnimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mygrpc_bullet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchAnimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchAnimeRequest) ProtoMessage() {}

func (x *SearchAnimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mygrpc_bullet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchAnimeRequest.ProtoReflect.Descriptor instead.
func (*SearchAnimeRequest) Descriptor() ([]byte, []int) {
	return file_mygrpc_bullet_proto_rawDescGZIP(), []int{0}
}

func (x *SearchAnimeRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SearchAnimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*AnimeEntity `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SearchAnimeResponse) Reset() {
	*x = SearchAnimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mygrpc_bullet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchAnimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchAnimeResponse) ProtoMessage() {}

func (x *SearchAnimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mygrpc_bullet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchAnimeResponse.ProtoReflect.Descriptor instead.
func (*SearchAnimeResponse) Descriptor() ([]byte, []int) {
	return file_mygrpc_bullet_proto_rawDescGZIP(), []int{1}
}

func (x *SearchAnimeResponse) GetData() []*AnimeEntity {
	if x != nil {
		return x.Data
	}
	return nil
}

type AnimeEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonID int64  `protobuf:"varint,1,opt,name=seasonID,proto3" json:"seasonID,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *AnimeEntity) Reset() {
	*x = AnimeEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mygrpc_bullet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeEntity) ProtoMessage() {}

func (x *AnimeEntity) ProtoReflect() protoreflect.Message {
	mi := &file_mygrpc_bullet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeEntity.ProtoReflect.Descriptor instead.
func (*AnimeEntity) Descriptor() ([]byte, []int) {
	return file_mygrpc_bullet_proto_rawDescGZIP(), []int{2}
}

func (x *AnimeEntity) GetSeasonID() int64 {
	if x != nil {
		return x.SeasonID
	}
	return 0
}

func (x *AnimeEntity) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_mygrpc_bullet_proto protoreflect.FileDescriptor

var file_mygrpc_bullet_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x79, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x28, 0x0a, 0x12, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3c, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x3f, 0x0a, 0x0b, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x32, 0x4e, 0x0a, 0x06, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x44,
	0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x6d, 0x79, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mygrpc_bullet_proto_rawDescOnce sync.Once
	file_mygrpc_bullet_proto_rawDescData = file_mygrpc_bullet_proto_rawDesc
)

func file_mygrpc_bullet_proto_rawDescGZIP() []byte {
	file_mygrpc_bullet_proto_rawDescOnce.Do(func() {
		file_mygrpc_bullet_proto_rawDescData = protoimpl.X.CompressGZIP(file_mygrpc_bullet_proto_rawDescData)
	})
	return file_mygrpc_bullet_proto_rawDescData
}

var file_mygrpc_bullet_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mygrpc_bullet_proto_goTypes = []interface{}{
	(*SearchAnimeRequest)(nil),  // 0: grpc.SearchAnimeRequest
	(*SearchAnimeResponse)(nil), // 1: grpc.SearchAnimeResponse
	(*AnimeEntity)(nil),         // 2: grpc.AnimeEntity
}
var file_mygrpc_bullet_proto_depIdxs = []int32{
	2, // 0: grpc.SearchAnimeResponse.data:type_name -> grpc.AnimeEntity
	0, // 1: grpc.Bullet.SearchAnime:input_type -> grpc.SearchAnimeRequest
	1, // 2: grpc.Bullet.SearchAnime:output_type -> grpc.SearchAnimeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mygrpc_bullet_proto_init() }
func file_mygrpc_bullet_proto_init() {
	if File_mygrpc_bullet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mygrpc_bullet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchAnimeRequest); i {
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
		file_mygrpc_bullet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchAnimeResponse); i {
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
		file_mygrpc_bullet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeEntity); i {
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
			RawDescriptor: file_mygrpc_bullet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mygrpc_bullet_proto_goTypes,
		DependencyIndexes: file_mygrpc_bullet_proto_depIdxs,
		MessageInfos:      file_mygrpc_bullet_proto_msgTypes,
	}.Build()
	File_mygrpc_bullet_proto = out.File
	file_mygrpc_bullet_proto_rawDesc = nil
	file_mygrpc_bullet_proto_goTypes = nil
	file_mygrpc_bullet_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BulletClient is the client API for Bullet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BulletClient interface {
	SearchAnime(ctx context.Context, in *SearchAnimeRequest, opts ...grpc.CallOption) (*SearchAnimeResponse, error)
}

type bulletClient struct {
	cc grpc.ClientConnInterface
}

func NewBulletClient(cc grpc.ClientConnInterface) BulletClient {
	return &bulletClient{cc}
}

func (c *bulletClient) SearchAnime(ctx context.Context, in *SearchAnimeRequest, opts ...grpc.CallOption) (*SearchAnimeResponse, error) {
	out := new(SearchAnimeResponse)
	err := c.cc.Invoke(ctx, "/grpc.Bullet/SearchAnime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BulletServer is the server API for Bullet service.
type BulletServer interface {
	SearchAnime(context.Context, *SearchAnimeRequest) (*SearchAnimeResponse, error)
}

// UnimplementedBulletServer can be embedded to have forward compatible implementations.
type UnimplementedBulletServer struct {
}

func (*UnimplementedBulletServer) SearchAnime(context.Context, *SearchAnimeRequest) (*SearchAnimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAnime not implemented")
}

func RegisterBulletServer(s *grpc.Server, srv BulletServer) {
	s.RegisterService(&_Bullet_serviceDesc, srv)
}

func _Bullet_SearchAnime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAnimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BulletServer).SearchAnime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Bullet/SearchAnime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BulletServer).SearchAnime(ctx, req.(*SearchAnimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bullet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Bullet",
	HandlerType: (*BulletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchAnime",
			Handler:    _Bullet_SearchAnime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mygrpc/bullet.proto",
}