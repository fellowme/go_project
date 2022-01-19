// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: role.proto

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

type UserRoleMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds string `protobuf:"bytes,1,opt,name=UserIds,proto3" json:"UserIds,omitempty"`
}

func (x *UserRoleMenuRequest) Reset() {
	*x = UserRoleMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRoleMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRoleMenuRequest) ProtoMessage() {}

func (x *UserRoleMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRoleMenuRequest.ProtoReflect.Descriptor instead.
func (*UserRoleMenuRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{0}
}

func (x *UserRoleMenuRequest) GetUserIds() string {
	if x != nil {
		return x.UserIds
	}
	return ""
}

type UserRoleMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MenuName string `protobuf:"bytes,2,opt,name=MenuName,proto3" json:"MenuName,omitempty"`
	MenuPath string `protobuf:"bytes,3,opt,name=MenuPath,proto3" json:"MenuPath,omitempty"`
	MenuType int32  `protobuf:"varint,5,opt,name=MenuType,proto3" json:"MenuType,omitempty"`
}

func (x *UserRoleMenuResponse) Reset() {
	*x = UserRoleMenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRoleMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRoleMenuResponse) ProtoMessage() {}

func (x *UserRoleMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRoleMenuResponse.ProtoReflect.Descriptor instead.
func (*UserRoleMenuResponse) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{1}
}

func (x *UserRoleMenuResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRoleMenuResponse) GetMenuName() string {
	if x != nil {
		return x.MenuName
	}
	return ""
}

func (x *UserRoleMenuResponse) GetMenuPath() string {
	if x != nil {
		return x.MenuPath
	}
	return ""
}

func (x *UserRoleMenuResponse) GetMenuType() int32 {
	if x != nil {
		return x.MenuType
	}
	return 0
}

type UserRoleMenuListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuList []*UserRoleMenuResponse `protobuf:"bytes,1,rep,name=menuList,proto3" json:"menuList,omitempty"`
}

func (x *UserRoleMenuListResponse) Reset() {
	*x = UserRoleMenuListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRoleMenuListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRoleMenuListResponse) ProtoMessage() {}

func (x *UserRoleMenuListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRoleMenuListResponse.ProtoReflect.Descriptor instead.
func (*UserRoleMenuListResponse) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{2}
}

func (x *UserRoleMenuListResponse) GetMenuList() []*UserRoleMenuResponse {
	if x != nil {
		return x.MenuList
	}
	return nil
}

var File_role_proto protoreflect.FileDescriptor

var file_role_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2f, 0x0a, 0x13, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x7a, 0x0a, 0x14, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x65,
	0x6e, 0x75, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4d, 0x65,
	0x6e, 0x75, 0x54, 0x79, 0x70, 0x65, 0x22, 0x59, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73,
	0x74, 0x32, 0x76, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e,
	0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x12, 0x20, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x3b,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_role_proto_rawDescOnce sync.Once
	file_role_proto_rawDescData = file_role_proto_rawDesc
)

func file_role_proto_rawDescGZIP() []byte {
	file_role_proto_rawDescOnce.Do(func() {
		file_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_proto_rawDescData)
	})
	return file_role_proto_rawDescData
}

var file_role_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_role_proto_goTypes = []interface{}{
	(*UserRoleMenuRequest)(nil),      // 0: rpc_service.UserRoleMenuRequest
	(*UserRoleMenuResponse)(nil),     // 1: rpc_service.UserRoleMenuResponse
	(*UserRoleMenuListResponse)(nil), // 2: rpc_service.UserRoleMenuListResponse
}
var file_role_proto_depIdxs = []int32{
	1, // 0: rpc_service.UserRoleMenuListResponse.menuList:type_name -> rpc_service.UserRoleMenuResponse
	0, // 1: rpc_service.UserRoleMenuService.GetRoleMenuByUserIds:input_type -> rpc_service.UserRoleMenuRequest
	2, // 2: rpc_service.UserRoleMenuService.GetRoleMenuByUserIds:output_type -> rpc_service.UserRoleMenuListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_role_proto_init() }
func file_role_proto_init() {
	if File_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRoleMenuRequest); i {
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
		file_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRoleMenuResponse); i {
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
		file_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRoleMenuListResponse); i {
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
			RawDescriptor: file_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_proto_goTypes,
		DependencyIndexes: file_role_proto_depIdxs,
		MessageInfos:      file_role_proto_msgTypes,
	}.Build()
	File_role_proto = out.File
	file_role_proto_rawDesc = nil
	file_role_proto_goTypes = nil
	file_role_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserRoleMenuServiceClient is the client API for UserRoleMenuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserRoleMenuServiceClient interface {
	GetRoleMenuByUserIds(ctx context.Context, in *UserRoleMenuRequest, opts ...grpc.CallOption) (*UserRoleMenuListResponse, error)
}

type userRoleMenuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRoleMenuServiceClient(cc grpc.ClientConnInterface) UserRoleMenuServiceClient {
	return &userRoleMenuServiceClient{cc}
}

func (c *userRoleMenuServiceClient) GetRoleMenuByUserIds(ctx context.Context, in *UserRoleMenuRequest, opts ...grpc.CallOption) (*UserRoleMenuListResponse, error) {
	out := new(UserRoleMenuListResponse)
	err := c.cc.Invoke(ctx, "/rpc_service.UserRoleMenuService/GetRoleMenuByUserIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRoleMenuServiceServer is the server API for UserRoleMenuService service.
type UserRoleMenuServiceServer interface {
	GetRoleMenuByUserIds(context.Context, *UserRoleMenuRequest) (*UserRoleMenuListResponse, error)
}

// UnimplementedUserRoleMenuServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserRoleMenuServiceServer struct {
}

func (*UnimplementedUserRoleMenuServiceServer) GetRoleMenuByUserIds(context.Context, *UserRoleMenuRequest) (*UserRoleMenuListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleMenuByUserIds not implemented")
}

func RegisterUserRoleMenuServiceServer(s *grpc.Server, srv UserRoleMenuServiceServer) {
	s.RegisterService(&_UserRoleMenuService_serviceDesc, srv)
}

func _UserRoleMenuService_GetRoleMenuByUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRoleMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRoleMenuServiceServer).GetRoleMenuByUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc_service.UserRoleMenuService/GetRoleMenuByUserIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRoleMenuServiceServer).GetRoleMenuByUserIds(ctx, req.(*UserRoleMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserRoleMenuService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc_service.UserRoleMenuService",
	HandlerType: (*UserRoleMenuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoleMenuByUserIds",
			Handler:    _UserRoleMenuService_GetRoleMenuByUserIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role.proto",
}
