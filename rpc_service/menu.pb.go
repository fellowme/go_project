// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: menu.proto

package rpc_service

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

type MenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuIds string `protobuf:"bytes,1,opt,name=MenuIds,proto3" json:"MenuIds,omitempty"`
}

func (x *MenuRequest) Reset() {
	*x = MenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuRequest) ProtoMessage() {}

func (x *MenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuRequest.ProtoReflect.Descriptor instead.
func (*MenuRequest) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{0}
}

func (x *MenuRequest) GetMenuIds() string {
	if x != nil {
		return x.MenuIds
	}
	return ""
}

type MenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MenuName string `protobuf:"bytes,2,opt,name=MenuName,proto3" json:"MenuName,omitempty"`
	Path     string `protobuf:"bytes,3,opt,name=Path,proto3" json:"Path,omitempty"`
	Remark   string `protobuf:"bytes,4,opt,name=Remark,proto3" json:"Remark,omitempty"`
	Method   string `protobuf:"bytes,5,opt,name=Method,proto3" json:"Method,omitempty"`
	Handler  string `protobuf:"bytes,6,opt,name=Handler,proto3" json:"Handler,omitempty"`
}

func (x *MenuResponse) Reset() {
	*x = MenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuResponse) ProtoMessage() {}

func (x *MenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuResponse.ProtoReflect.Descriptor instead.
func (*MenuResponse) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{1}
}

func (x *MenuResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MenuResponse) GetMenuName() string {
	if x != nil {
		return x.MenuName
	}
	return ""
}

func (x *MenuResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *MenuResponse) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *MenuResponse) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *MenuResponse) GetHandler() string {
	if x != nil {
		return x.Handler
	}
	return ""
}

type MenuListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuList []*MenuResponse `protobuf:"bytes,1,rep,name=menuList,proto3" json:"menuList,omitempty"`
}

func (x *MenuListResponse) Reset() {
	*x = MenuListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuListResponse) ProtoMessage() {}

func (x *MenuListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuListResponse.ProtoReflect.Descriptor instead.
func (*MenuListResponse) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{2}
}

func (x *MenuListResponse) GetMenuList() []*MenuResponse {
	if x != nil {
		return x.MenuList
	}
	return nil
}

var File_menu_proto protoreflect.FileDescriptor

var file_menu_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x4d, 0x65, 0x6e,
	0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x6e, 0x75,
	0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x6e, 0x75, 0x49,
	0x64, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x22, 0x49, 0x0a,
	0x10, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08,
	0x6d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x56, 0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x6e, 0x75, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x3b, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_menu_proto_rawDescOnce sync.Once
	file_menu_proto_rawDescData = file_menu_proto_rawDesc
)

func file_menu_proto_rawDescGZIP() []byte {
	file_menu_proto_rawDescOnce.Do(func() {
		file_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_menu_proto_rawDescData)
	})
	return file_menu_proto_rawDescData
}

var file_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_menu_proto_goTypes = []interface{}{
	(*MenuRequest)(nil),      // 0: rpc_service.MenuRequest
	(*MenuResponse)(nil),     // 1: rpc_service.MenuResponse
	(*MenuListResponse)(nil), // 2: rpc_service.MenuListResponse
}
var file_menu_proto_depIdxs = []int32{
	1, // 0: rpc_service.MenuListResponse.menuList:type_name -> rpc_service.MenuResponse
	0, // 1: rpc_service.MenuService.GetMenuByIds:input_type -> rpc_service.MenuRequest
	2, // 2: rpc_service.MenuService.GetMenuByIds:output_type -> rpc_service.MenuListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_menu_proto_init() }
func file_menu_proto_init() {
	if File_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuRequest); i {
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
		file_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuResponse); i {
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
		file_menu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuListResponse); i {
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
			RawDescriptor: file_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_menu_proto_goTypes,
		DependencyIndexes: file_menu_proto_depIdxs,
		MessageInfos:      file_menu_proto_msgTypes,
	}.Build()
	File_menu_proto = out.File
	file_menu_proto_rawDesc = nil
	file_menu_proto_goTypes = nil
	file_menu_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MenuServiceClient is the client API for MenuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MenuServiceClient interface {
	GetMenuByIds(ctx context.Context, in *MenuRequest, opts ...grpc.CallOption) (*MenuListResponse, error)
}

type menuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMenuServiceClient(cc grpc.ClientConnInterface) MenuServiceClient {
	return &menuServiceClient{cc}
}

func (c *menuServiceClient) GetMenuByIds(ctx context.Context, in *MenuRequest, opts ...grpc.CallOption) (*MenuListResponse, error) {
	out := new(MenuListResponse)
	err := c.cc.Invoke(ctx, "/rpc_service.MenuService/GetMenuByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuServiceServer is the server API for MenuService service.
type MenuServiceServer interface {
	GetMenuByIds(context.Context, *MenuRequest) (*MenuListResponse, error)
}

// UnimplementedMenuServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMenuServiceServer struct {
}

func (*UnimplementedMenuServiceServer) GetMenuByIds(context.Context, *MenuRequest) (*MenuListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenuByIds not implemented")
}

func RegisterMenuServiceServer(s *grpc.Server, srv MenuServiceServer) {
	s.RegisterService(&_MenuService_serviceDesc, srv)
}

func _MenuService_GetMenuByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).GetMenuByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc_service.MenuService/GetMenuByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).GetMenuByIds(ctx, req.(*MenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MenuService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc_service.MenuService",
	HandlerType: (*MenuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMenuByIds",
			Handler:    _MenuService_GetMenuByIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "menu.proto",
}
