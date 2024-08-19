// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: userop.proto

package userop

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Message_MessageList_FullMethodName   = "/Message/MessageList"
	Message_CreateMessage_FullMethodName = "/Message/CreateMessage"
)

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 用户留言
type MessageClient interface {
	MessageList(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageListResponse, error)
	CreateMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) MessageList(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageListResponse)
	err := c.cc.Invoke(ctx, Message_MessageList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) CreateMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, Message_CreateMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
// All implementations must embed UnimplementedMessageServer
// for forward compatibility
//
// 用户留言
type MessageServer interface {
	MessageList(context.Context, *MessageRequest) (*MessageListResponse, error)
	CreateMessage(context.Context, *MessageRequest) (*MessageResponse, error)
	mustEmbedUnimplementedMessageServer()
}

// UnimplementedMessageServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (UnimplementedMessageServer) MessageList(context.Context, *MessageRequest) (*MessageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageList not implemented")
}
func (UnimplementedMessageServer) CreateMessage(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedMessageServer) mustEmbedUnimplementedMessageServer() {}

// UnsafeMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServer will
// result in compilation errors.
type UnsafeMessageServer interface {
	mustEmbedUnimplementedMessageServer()
}

func RegisterMessageServer(s grpc.ServiceRegistrar, srv MessageServer) {
	s.RegisterService(&Message_ServiceDesc, srv)
}

func _Message_MessageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).MessageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_MessageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).MessageList(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_CreateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).CreateMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Message_ServiceDesc is the grpc.ServiceDesc for Message service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Message_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessageList",
			Handler:    _Message_MessageList_Handler,
		},
		{
			MethodName: "CreateMessage",
			Handler:    _Message_CreateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userop.proto",
}

const (
	Address_GetAddressList_FullMethodName = "/Address/GetAddressList"
	Address_CreateAddress_FullMethodName  = "/Address/CreateAddress"
	Address_DeleteAddress_FullMethodName  = "/Address/DeleteAddress"
	Address_UpdateAddress_FullMethodName  = "/Address/UpdateAddress"
)

// AddressClient is the client API for Address service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 地址
type AddressClient interface {
	GetAddressList(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressListResponse, error)
	CreateAddress(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressResponse, error)
	DeleteAddress(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressEmpty, error)
	UpdateAddress(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressEmpty, error)
}

type addressClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressClient(cc grpc.ClientConnInterface) AddressClient {
	return &addressClient{cc}
}

func (c *addressClient) GetAddressList(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddressListResponse)
	err := c.cc.Invoke(ctx, Address_GetAddressList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressClient) CreateAddress(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddressResponse)
	err := c.cc.Invoke(ctx, Address_CreateAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressClient) DeleteAddress(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddressEmpty)
	err := c.cc.Invoke(ctx, Address_DeleteAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressClient) UpdateAddress(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddressEmpty)
	err := c.cc.Invoke(ctx, Address_UpdateAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServer is the server API for Address service.
// All implementations must embed UnimplementedAddressServer
// for forward compatibility
//
// 地址
type AddressServer interface {
	GetAddressList(context.Context, *AddressRequest) (*AddressListResponse, error)
	CreateAddress(context.Context, *AddressRequest) (*AddressResponse, error)
	DeleteAddress(context.Context, *AddressRequest) (*AddressEmpty, error)
	UpdateAddress(context.Context, *AddressRequest) (*AddressEmpty, error)
	mustEmbedUnimplementedAddressServer()
}

// UnimplementedAddressServer must be embedded to have forward compatible implementations.
type UnimplementedAddressServer struct {
}

func (UnimplementedAddressServer) GetAddressList(context.Context, *AddressRequest) (*AddressListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressList not implemented")
}
func (UnimplementedAddressServer) CreateAddress(context.Context, *AddressRequest) (*AddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddress not implemented")
}
func (UnimplementedAddressServer) DeleteAddress(context.Context, *AddressRequest) (*AddressEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAddress not implemented")
}
func (UnimplementedAddressServer) UpdateAddress(context.Context, *AddressRequest) (*AddressEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddress not implemented")
}
func (UnimplementedAddressServer) mustEmbedUnimplementedAddressServer() {}

// UnsafeAddressServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressServer will
// result in compilation errors.
type UnsafeAddressServer interface {
	mustEmbedUnimplementedAddressServer()
}

func RegisterAddressServer(s grpc.ServiceRegistrar, srv AddressServer) {
	s.RegisterService(&Address_ServiceDesc, srv)
}

func _Address_GetAddressList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServer).GetAddressList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Address_GetAddressList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServer).GetAddressList(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Address_CreateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServer).CreateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Address_CreateAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServer).CreateAddress(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Address_DeleteAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServer).DeleteAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Address_DeleteAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServer).DeleteAddress(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Address_UpdateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServer).UpdateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Address_UpdateAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServer).UpdateAddress(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Address_ServiceDesc is the grpc.ServiceDesc for Address service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Address_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Address",
	HandlerType: (*AddressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddressList",
			Handler:    _Address_GetAddressList_Handler,
		},
		{
			MethodName: "CreateAddress",
			Handler:    _Address_CreateAddress_Handler,
		},
		{
			MethodName: "DeleteAddress",
			Handler:    _Address_DeleteAddress_Handler,
		},
		{
			MethodName: "UpdateAddress",
			Handler:    _Address_UpdateAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userop.proto",
}

const (
	UserFav_GetFavList_FullMethodName       = "/UserFav/GetFavList"
	UserFav_AddUserFav_FullMethodName       = "/UserFav/AddUserFav"
	UserFav_DeleteUserFav_FullMethodName    = "/UserFav/DeleteUserFav"
	UserFav_GetUserFavDetail_FullMethodName = "/UserFav/GetUserFavDetail"
)

// UserFavClient is the client API for UserFav service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 用户收藏
type UserFavClient interface {
	GetFavList(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*UserFavListResponse, error)
	AddUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*UserFavEmpty, error)
	DeleteUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*UserFavEmpty, error)
	GetUserFavDetail(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*UserFavEmpty, error)
}

type userFavClient struct {
	cc grpc.ClientConnInterface
}

func NewUserFavClient(cc grpc.ClientConnInterface) UserFavClient {
	return &userFavClient{cc}
}

func (c *userFavClient) GetFavList(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*UserFavListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserFavListResponse)
	err := c.cc.Invoke(ctx, UserFav_GetFavList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userFavClient) AddUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*UserFavEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserFavEmpty)
	err := c.cc.Invoke(ctx, UserFav_AddUserFav_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userFavClient) DeleteUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*UserFavEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserFavEmpty)
	err := c.cc.Invoke(ctx, UserFav_DeleteUserFav_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userFavClient) GetUserFavDetail(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*UserFavEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserFavEmpty)
	err := c.cc.Invoke(ctx, UserFav_GetUserFavDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserFavServer is the server API for UserFav service.
// All implementations must embed UnimplementedUserFavServer
// for forward compatibility
//
// 用户收藏
type UserFavServer interface {
	GetFavList(context.Context, *UserFavRequest) (*UserFavListResponse, error)
	AddUserFav(context.Context, *UserFavRequest) (*UserFavEmpty, error)
	DeleteUserFav(context.Context, *UserFavRequest) (*UserFavEmpty, error)
	GetUserFavDetail(context.Context, *UserFavRequest) (*UserFavEmpty, error)
	mustEmbedUnimplementedUserFavServer()
}

// UnimplementedUserFavServer must be embedded to have forward compatible implementations.
type UnimplementedUserFavServer struct {
}

func (UnimplementedUserFavServer) GetFavList(context.Context, *UserFavRequest) (*UserFavListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavList not implemented")
}
func (UnimplementedUserFavServer) AddUserFav(context.Context, *UserFavRequest) (*UserFavEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserFav not implemented")
}
func (UnimplementedUserFavServer) DeleteUserFav(context.Context, *UserFavRequest) (*UserFavEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserFav not implemented")
}
func (UnimplementedUserFavServer) GetUserFavDetail(context.Context, *UserFavRequest) (*UserFavEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFavDetail not implemented")
}
func (UnimplementedUserFavServer) mustEmbedUnimplementedUserFavServer() {}

// UnsafeUserFavServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserFavServer will
// result in compilation errors.
type UnsafeUserFavServer interface {
	mustEmbedUnimplementedUserFavServer()
}

func RegisterUserFavServer(s grpc.ServiceRegistrar, srv UserFavServer) {
	s.RegisterService(&UserFav_ServiceDesc, srv)
}

func _UserFav_GetFavList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFavServer).GetFavList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserFav_GetFavList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFavServer).GetFavList(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserFav_AddUserFav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFavServer).AddUserFav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserFav_AddUserFav_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFavServer).AddUserFav(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserFav_DeleteUserFav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFavServer).DeleteUserFav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserFav_DeleteUserFav_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFavServer).DeleteUserFav(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserFav_GetUserFavDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFavServer).GetUserFavDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserFav_GetUserFavDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFavServer).GetUserFavDetail(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserFav_ServiceDesc is the grpc.ServiceDesc for UserFav service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserFav_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserFav",
	HandlerType: (*UserFavServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFavList",
			Handler:    _UserFav_GetFavList_Handler,
		},
		{
			MethodName: "AddUserFav",
			Handler:    _UserFav_AddUserFav_Handler,
		},
		{
			MethodName: "DeleteUserFav",
			Handler:    _UserFav_DeleteUserFav_Handler,
		},
		{
			MethodName: "GetUserFavDetail",
			Handler:    _UserFav_GetUserFavDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userop.proto",
}
