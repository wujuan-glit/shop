// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: seckill_activity.proto

package seckill_activity

import (
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

type SeckillActivityEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SeckillActivityEmpty) Reset() {
	*x = SeckillActivityEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seckill_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeckillActivityEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeckillActivityEmpty) ProtoMessage() {}

func (x *SeckillActivityEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_seckill_activity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeckillActivityEmpty.ProtoReflect.Descriptor instead.
func (*SeckillActivityEmpty) Descriptor() ([]byte, []int) {
	return file_seckill_activity_proto_rawDescGZIP(), []int{0}
}

type SeckillActivityListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Page     int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *SeckillActivityListReq) Reset() {
	*x = SeckillActivityListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seckill_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeckillActivityListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeckillActivityListReq) ProtoMessage() {}

func (x *SeckillActivityListReq) ProtoReflect() protoreflect.Message {
	mi := &file_seckill_activity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeckillActivityListReq.ProtoReflect.Descriptor instead.
func (*SeckillActivityListReq) Descriptor() ([]byte, []int) {
	return file_seckill_activity_proto_rawDescGZIP(), []int{1}
}

func (x *SeckillActivityListReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SeckillActivityListReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SeckillActivityListReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type SeckillActivityData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ActivityStartTime int32  `protobuf:"varint,3,opt,name=activity_start_time,json=activityStartTime,proto3" json:"activity_start_time,omitempty"`
	ActivityEndTime   int32  `protobuf:"varint,4,opt,name=activity_end_time,json=activityEndTime,proto3" json:"activity_end_time,omitempty"`
	Status            int32  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SeckillActivityData) Reset() {
	*x = SeckillActivityData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seckill_activity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeckillActivityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeckillActivityData) ProtoMessage() {}

func (x *SeckillActivityData) ProtoReflect() protoreflect.Message {
	mi := &file_seckill_activity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeckillActivityData.ProtoReflect.Descriptor instead.
func (*SeckillActivityData) Descriptor() ([]byte, []int) {
	return file_seckill_activity_proto_rawDescGZIP(), []int{2}
}

func (x *SeckillActivityData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SeckillActivityData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SeckillActivityData) GetActivityStartTime() int32 {
	if x != nil {
		return x.ActivityStartTime
	}
	return 0
}

func (x *SeckillActivityData) GetActivityEndTime() int32 {
	if x != nil {
		return x.ActivityEndTime
	}
	return 0
}

func (x *SeckillActivityData) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type SeckillActivityListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*SeckillActivityData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SeckillActivityListResp) Reset() {
	*x = SeckillActivityListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seckill_activity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeckillActivityListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeckillActivityListResp) ProtoMessage() {}

func (x *SeckillActivityListResp) ProtoReflect() protoreflect.Message {
	mi := &file_seckill_activity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeckillActivityListResp.ProtoReflect.Descriptor instead.
func (*SeckillActivityListResp) Descriptor() ([]byte, []int) {
	return file_seckill_activity_proto_rawDescGZIP(), []int{3}
}

func (x *SeckillActivityListResp) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SeckillActivityListResp) GetData() []*SeckillActivityData {
	if x != nil {
		return x.Data
	}
	return nil
}

type SeckillRecordCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeckillActivityId      int64 `protobuf:"varint,1,opt,name=seckill_activity_id,json=seckillActivityId,proto3" json:"seckill_activity_id,omitempty"`
	SeckillActivityGoodsId int64 `protobuf:"varint,2,opt,name=seckill_activity_goods_id,json=seckillActivityGoodsId,proto3" json:"seckill_activity_goods_id,omitempty"`
	UserId                 int64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *SeckillRecordCreateRequest) Reset() {
	*x = SeckillRecordCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seckill_activity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeckillRecordCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeckillRecordCreateRequest) ProtoMessage() {}

func (x *SeckillRecordCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seckill_activity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeckillRecordCreateRequest.ProtoReflect.Descriptor instead.
func (*SeckillRecordCreateRequest) Descriptor() ([]byte, []int) {
	return file_seckill_activity_proto_rawDescGZIP(), []int{4}
}

func (x *SeckillRecordCreateRequest) GetSeckillActivityId() int64 {
	if x != nil {
		return x.SeckillActivityId
	}
	return 0
}

func (x *SeckillRecordCreateRequest) GetSeckillActivityGoodsId() int64 {
	if x != nil {
		return x.SeckillActivityGoodsId
	}
	return 0
}

func (x *SeckillRecordCreateRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type SeckillActivityGoodsListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeckillActivityId int64 `protobuf:"varint,1,opt,name=seckill_activity_id,json=seckillActivityId,proto3" json:"seckill_activity_id,omitempty"`
}

func (x *SeckillActivityGoodsListReq) Reset() {
	*x = SeckillActivityGoodsListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seckill_activity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeckillActivityGoodsListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeckillActivityGoodsListReq) ProtoMessage() {}

func (x *SeckillActivityGoodsListReq) ProtoReflect() protoreflect.Message {
	mi := &file_seckill_activity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeckillActivityGoodsListReq.ProtoReflect.Descriptor instead.
func (*SeckillActivityGoodsListReq) Descriptor() ([]byte, []int) {
	return file_seckill_activity_proto_rawDescGZIP(), []int{5}
}

func (x *SeckillActivityGoodsListReq) GetSeckillActivityId() int64 {
	if x != nil {
		return x.SeckillActivityId
	}
	return 0
}

type SeckillGoods struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ActivityId string `protobuf:"bytes,2,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	GoodsId    int32  `protobuf:"varint,3,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	Stock      int32  `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
	Status     int32  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SeckillGoods) Reset() {
	*x = SeckillGoods{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seckill_activity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeckillGoods) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeckillGoods) ProtoMessage() {}

func (x *SeckillGoods) ProtoReflect() protoreflect.Message {
	mi := &file_seckill_activity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeckillGoods.ProtoReflect.Descriptor instead.
func (*SeckillGoods) Descriptor() ([]byte, []int) {
	return file_seckill_activity_proto_rawDescGZIP(), []int{6}
}

func (x *SeckillGoods) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SeckillGoods) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

func (x *SeckillGoods) GetGoodsId() int32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *SeckillGoods) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *SeckillGoods) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type SeckillActivityGoodsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32           `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*SeckillGoods `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SeckillActivityGoodsData) Reset() {
	*x = SeckillActivityGoodsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seckill_activity_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeckillActivityGoodsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeckillActivityGoodsData) ProtoMessage() {}

func (x *SeckillActivityGoodsData) ProtoReflect() protoreflect.Message {
	mi := &file_seckill_activity_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeckillActivityGoodsData.ProtoReflect.Descriptor instead.
func (*SeckillActivityGoodsData) Descriptor() ([]byte, []int) {
	return file_seckill_activity_proto_rawDescGZIP(), []int{7}
}

func (x *SeckillActivityGoodsData) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SeckillActivityGoodsData) GetData() []*SeckillGoods {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_seckill_activity_proto protoreflect.FileDescriptor

var file_seckill_activity_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65, 0x63, 0x6b,
	0x69, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x61, 0x0a, 0x16, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2e, 0x0a, 0x13, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x59, 0x0a, 0x17, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa0,
	0x01, 0x0a, 0x1a, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a,
	0x13, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x65, 0x63, 0x6b,
	0x69, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a,
	0x19, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x5f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x16, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x4d, 0x0a, 0x1b, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x2e, 0x0a, 0x13, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73,
	0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x22, 0x88, 0x01, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x53, 0x0a, 0x18, 0x53,
	0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x21, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65,
	0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0xfb, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x13, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x53, 0x65,
	0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x53,
	0x0a, 0x18, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x53, 0x65, 0x63,
	0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x53, 0x65, 0x63, 0x6b, 0x69,
	0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x49, 0x0a, 0x13, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x53, 0x65, 0x63,
	0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c,
	0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x14,
	0x5a, 0x12, 0x2e, 0x3b, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_seckill_activity_proto_rawDescOnce sync.Once
	file_seckill_activity_proto_rawDescData = file_seckill_activity_proto_rawDesc
)

func file_seckill_activity_proto_rawDescGZIP() []byte {
	file_seckill_activity_proto_rawDescOnce.Do(func() {
		file_seckill_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_seckill_activity_proto_rawDescData)
	})
	return file_seckill_activity_proto_rawDescData
}

var file_seckill_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_seckill_activity_proto_goTypes = []any{
	(*SeckillActivityEmpty)(nil),        // 0: SeckillActivityEmpty
	(*SeckillActivityListReq)(nil),      // 1: SeckillActivityListReq
	(*SeckillActivityData)(nil),         // 2: SeckillActivityData
	(*SeckillActivityListResp)(nil),     // 3: SeckillActivityListResp
	(*SeckillRecordCreateRequest)(nil),  // 4: SeckillRecordCreateRequest
	(*SeckillActivityGoodsListReq)(nil), // 5: SeckillActivityGoodsListReq
	(*SeckillGoods)(nil),                // 6: seckillGoods
	(*SeckillActivityGoodsData)(nil),    // 7: SeckillActivityGoodsData
}
var file_seckill_activity_proto_depIdxs = []int32{
	2, // 0: SeckillActivityListResp.data:type_name -> SeckillActivityData
	6, // 1: SeckillActivityGoodsData.data:type_name -> seckillGoods
	1, // 2: SeckillActivity.SeckillActivityList:input_type -> SeckillActivityListReq
	5, // 3: SeckillActivity.SeckillActivityGoodsList:input_type -> SeckillActivityGoodsListReq
	4, // 4: SeckillActivity.SeckillRecordCreate:input_type -> SeckillRecordCreateRequest
	3, // 5: SeckillActivity.SeckillActivityList:output_type -> SeckillActivityListResp
	7, // 6: SeckillActivity.SeckillActivityGoodsList:output_type -> SeckillActivityGoodsData
	0, // 7: SeckillActivity.SeckillRecordCreate:output_type -> SeckillActivityEmpty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_seckill_activity_proto_init() }
func file_seckill_activity_proto_init() {
	if File_seckill_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_seckill_activity_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SeckillActivityEmpty); i {
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
		file_seckill_activity_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SeckillActivityListReq); i {
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
		file_seckill_activity_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SeckillActivityData); i {
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
		file_seckill_activity_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SeckillActivityListResp); i {
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
		file_seckill_activity_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SeckillRecordCreateRequest); i {
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
		file_seckill_activity_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SeckillActivityGoodsListReq); i {
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
		file_seckill_activity_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SeckillGoods); i {
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
		file_seckill_activity_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*SeckillActivityGoodsData); i {
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
			RawDescriptor: file_seckill_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_seckill_activity_proto_goTypes,
		DependencyIndexes: file_seckill_activity_proto_depIdxs,
		MessageInfos:      file_seckill_activity_proto_msgTypes,
	}.Build()
	File_seckill_activity_proto = out.File
	file_seckill_activity_proto_rawDesc = nil
	file_seckill_activity_proto_goTypes = nil
	file_seckill_activity_proto_depIdxs = nil
}
