// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: seckill_activity.proto

package seckill_activity

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
	SeckillActivity_SeckillActivityList_FullMethodName      = "/SeckillActivity/SeckillActivityList"
	SeckillActivity_SeckillActivityGoodsList_FullMethodName = "/SeckillActivity/SeckillActivityGoodsList"
	SeckillActivity_SeckillRecordCreate_FullMethodName      = "/SeckillActivity/SeckillRecordCreate"
)

// SeckillActivityClient is the client API for SeckillActivity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义商品服务
type SeckillActivityClient interface {
	// 秒杀活动列表
	SeckillActivityList(ctx context.Context, in *SeckillActivityListReq, opts ...grpc.CallOption) (*SeckillActivityListResp, error)
	// 秒杀活动商品列表
	SeckillActivityGoodsList(ctx context.Context, in *SeckillActivityGoodsListReq, opts ...grpc.CallOption) (*SeckillActivityGoodsData, error)
	// 秒杀购买记录表
	//
	//	rpc SeckillRecordList() returns();
	//
	// 秒杀购买记录详情
	//
	//	rpc SeckillRecordDetail() returns();
	//
	// 秒杀购买记录创建
	SeckillRecordCreate(ctx context.Context, in *SeckillRecordCreateRequest, opts ...grpc.CallOption) (*SeckillActivityEmpty, error)
}

type seckillActivityClient struct {
	cc grpc.ClientConnInterface
}

func NewSeckillActivityClient(cc grpc.ClientConnInterface) SeckillActivityClient {
	return &seckillActivityClient{cc}
}

func (c *seckillActivityClient) SeckillActivityList(ctx context.Context, in *SeckillActivityListReq, opts ...grpc.CallOption) (*SeckillActivityListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SeckillActivityListResp)
	err := c.cc.Invoke(ctx, SeckillActivity_SeckillActivityList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seckillActivityClient) SeckillActivityGoodsList(ctx context.Context, in *SeckillActivityGoodsListReq, opts ...grpc.CallOption) (*SeckillActivityGoodsData, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SeckillActivityGoodsData)
	err := c.cc.Invoke(ctx, SeckillActivity_SeckillActivityGoodsList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seckillActivityClient) SeckillRecordCreate(ctx context.Context, in *SeckillRecordCreateRequest, opts ...grpc.CallOption) (*SeckillActivityEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SeckillActivityEmpty)
	err := c.cc.Invoke(ctx, SeckillActivity_SeckillRecordCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeckillActivityServer is the server API for SeckillActivity service.
// All implementations must embed UnimplementedSeckillActivityServer
// for forward compatibility
//
// 定义商品服务
type SeckillActivityServer interface {
	// 秒杀活动列表
	SeckillActivityList(context.Context, *SeckillActivityListReq) (*SeckillActivityListResp, error)
	// 秒杀活动商品列表
	SeckillActivityGoodsList(context.Context, *SeckillActivityGoodsListReq) (*SeckillActivityGoodsData, error)
	// 秒杀购买记录表
	//
	//	rpc SeckillRecordList() returns();
	//
	// 秒杀购买记录详情
	//
	//	rpc SeckillRecordDetail() returns();
	//
	// 秒杀购买记录创建
	SeckillRecordCreate(context.Context, *SeckillRecordCreateRequest) (*SeckillActivityEmpty, error)
	mustEmbedUnimplementedSeckillActivityServer()
}

// UnimplementedSeckillActivityServer must be embedded to have forward compatible implementations.
type UnimplementedSeckillActivityServer struct {
}

func (UnimplementedSeckillActivityServer) SeckillActivityList(context.Context, *SeckillActivityListReq) (*SeckillActivityListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeckillActivityList not implemented")
}
func (UnimplementedSeckillActivityServer) SeckillActivityGoodsList(context.Context, *SeckillActivityGoodsListReq) (*SeckillActivityGoodsData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeckillActivityGoodsList not implemented")
}
func (UnimplementedSeckillActivityServer) SeckillRecordCreate(context.Context, *SeckillRecordCreateRequest) (*SeckillActivityEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeckillRecordCreate not implemented")
}
func (UnimplementedSeckillActivityServer) mustEmbedUnimplementedSeckillActivityServer() {}

// UnsafeSeckillActivityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SeckillActivityServer will
// result in compilation errors.
type UnsafeSeckillActivityServer interface {
	mustEmbedUnimplementedSeckillActivityServer()
}

func RegisterSeckillActivityServer(s grpc.ServiceRegistrar, srv SeckillActivityServer) {
	s.RegisterService(&SeckillActivity_ServiceDesc, srv)
}

func _SeckillActivity_SeckillActivityList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeckillActivityListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeckillActivityServer).SeckillActivityList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeckillActivity_SeckillActivityList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeckillActivityServer).SeckillActivityList(ctx, req.(*SeckillActivityListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeckillActivity_SeckillActivityGoodsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeckillActivityGoodsListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeckillActivityServer).SeckillActivityGoodsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeckillActivity_SeckillActivityGoodsList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeckillActivityServer).SeckillActivityGoodsList(ctx, req.(*SeckillActivityGoodsListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeckillActivity_SeckillRecordCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeckillRecordCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeckillActivityServer).SeckillRecordCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeckillActivity_SeckillRecordCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeckillActivityServer).SeckillRecordCreate(ctx, req.(*SeckillRecordCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SeckillActivity_ServiceDesc is the grpc.ServiceDesc for SeckillActivity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SeckillActivity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SeckillActivity",
	HandlerType: (*SeckillActivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SeckillActivityList",
			Handler:    _SeckillActivity_SeckillActivityList_Handler,
		},
		{
			MethodName: "SeckillActivityGoodsList",
			Handler:    _SeckillActivity_SeckillActivityGoodsList_Handler,
		},
		{
			MethodName: "SeckillRecordCreate",
			Handler:    _SeckillActivity_SeckillRecordCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "seckill_activity.proto",
}
