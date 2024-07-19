// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: order/v1/order.proto

package __order

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId     int64                  `protobuf:"varint,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	AddresseeId int64                  `protobuf:"varint,2,opt,name=AddresseeId,proto3" json:"AddresseeId,omitempty"`
	Weight      *int64                 `protobuf:"varint,3,opt,name=Weight,proto3,oneof" json:"Weight,omitempty"`
	Price       *int64                 `protobuf:"varint,4,opt,name=Price,proto3,oneof" json:"Price,omitempty"`
	WrapType    *int64                 `protobuf:"varint,5,opt,name=WrapType,proto3,oneof" json:"WrapType,omitempty"`
	ShelfLife   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=ShelfLife,proto3" json:"ShelfLife,omitempty"`
}

func (x *AddOrderRequest) Reset() {
	*x = AddOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_v1_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrderRequest) ProtoMessage() {}

func (x *AddOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrderRequest.ProtoReflect.Descriptor instead.
func (*AddOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{0}
}

func (x *AddOrderRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *AddOrderRequest) GetAddresseeId() int64 {
	if x != nil {
		return x.AddresseeId
	}
	return 0
}

func (x *AddOrderRequest) GetWeight() int64 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

func (x *AddOrderRequest) GetPrice() int64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *AddOrderRequest) GetWrapType() int64 {
	if x != nil && x.WrapType != nil {
		return *x.WrapType
	}
	return 0
}

func (x *AddOrderRequest) GetShelfLife() *timestamppb.Timestamp {
	if x != nil {
		return x.ShelfLife
	}
	return nil
}

type ReturnToDelivererRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
}

func (x *ReturnToDelivererRequest) Reset() {
	*x = ReturnToDelivererRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_v1_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnToDelivererRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnToDelivererRequest) ProtoMessage() {}

func (x *ReturnToDelivererRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnToDelivererRequest.ProtoReflect.Descriptor instead.
func (*ReturnToDelivererRequest) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{1}
}

func (x *ReturnToDelivererRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type GiveToAddresseeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId     int64 `protobuf:"varint,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	AddresseeId int64 `protobuf:"varint,2,opt,name=AddresseeId,proto3" json:"AddresseeId,omitempty"`
}

func (x *GiveToAddresseeRequest) Reset() {
	*x = GiveToAddresseeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_v1_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiveToAddresseeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiveToAddresseeRequest) ProtoMessage() {}

func (x *GiveToAddresseeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiveToAddresseeRequest.ProtoReflect.Descriptor instead.
func (*GiveToAddresseeRequest) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{2}
}

func (x *GiveToAddresseeRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *GiveToAddresseeRequest) GetAddresseeId() int64 {
	if x != nil {
		return x.AddresseeId
	}
	return 0
}

type ListRefundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageLen    *int64 `protobuf:"varint,1,opt,name=PageLen,proto3,oneof" json:"PageLen,omitempty"`
	PageNumber *int64 `protobuf:"varint,2,opt,name=PageNumber,proto3,oneof" json:"PageNumber,omitempty"`
}

func (x *ListRefundRequest) Reset() {
	*x = ListRefundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_v1_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRefundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRefundRequest) ProtoMessage() {}

func (x *ListRefundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRefundRequest.ProtoReflect.Descriptor instead.
func (*ListRefundRequest) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{3}
}

func (x *ListRefundRequest) GetPageLen() int64 {
	if x != nil && x.PageLen != nil {
		return *x.PageLen
	}
	return 0
}

func (x *ListRefundRequest) GetPageNumber() int64 {
	if x != nil && x.PageNumber != nil {
		return *x.PageNumber
	}
	return 0
}

type ListOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId    int64  `protobuf:"varint,1,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	CountOrders *int64 `protobuf:"varint,2,opt,name=CountOrders,proto3,oneof" json:"CountOrders,omitempty"`
}

func (x *ListOrderRequest) Reset() {
	*x = ListOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_v1_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderRequest) ProtoMessage() {}

func (x *ListOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderRequest.ProtoReflect.Descriptor instead.
func (*ListOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{4}
}

func (x *ListOrderRequest) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *ListOrderRequest) GetCountOrders() int64 {
	if x != nil && x.CountOrders != nil {
		return *x.CountOrders
	}
	return 0
}

type ReturnFromAddresseeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderIds []int64 `protobuf:"varint,1,rep,packed,name=OrderIds,proto3" json:"OrderIds,omitempty"`
}

func (x *ReturnFromAddresseeRequest) Reset() {
	*x = ReturnFromAddresseeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_v1_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnFromAddresseeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnFromAddresseeRequest) ProtoMessage() {}

func (x *ReturnFromAddresseeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnFromAddresseeRequest.ProtoReflect.Descriptor instead.
func (*ReturnFromAddresseeRequest) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{5}
}

func (x *ReturnFromAddresseeRequest) GetOrderIds() []int64 {
	if x != nil {
		return x.OrderIds
	}
	return nil
}

type ListOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []string `protobuf:"bytes,1,rep,name=Orders,proto3" json:"Orders,omitempty"`
}

func (x *ListOrderResponse) Reset() {
	*x = ListOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_v1_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderResponse) ProtoMessage() {}

func (x *ListOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderResponse.ProtoReflect.Descriptor instead.
func (*ListOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{6}
}

func (x *ListOrderResponse) GetOrders() []string {
	if x != nil {
		return x.Orders
	}
	return nil
}

type ListRefundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Refunds []string `protobuf:"bytes,1,rep,name=Refunds,proto3" json:"Refunds,omitempty"`
}

func (x *ListRefundResponse) Reset() {
	*x = ListRefundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_v1_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRefundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRefundResponse) ProtoMessage() {}

func (x *ListRefundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRefundResponse.ProtoReflect.Descriptor instead.
func (*ListRefundResponse) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{7}
}

func (x *ListRefundResponse) GetRefunds() []string {
	if x != nil {
		return x.Refunds
	}
	return nil
}

var File_order_v1_order_proto protoreflect.FileDescriptor

var file_order_v1_order_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x82, 0x02, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x19, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x57,
	0x72, 0x61, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52,
	0x08, 0x57, 0x72, 0x61, 0x70, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x09,
	0x53, 0x68, 0x65, 0x6c, 0x66, 0x4c, 0x69, 0x66, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x53, 0x68, 0x65,
	0x6c, 0x66, 0x4c, 0x69, 0x66, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x57, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x57, 0x72, 0x61, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0x34, 0x0a, 0x18, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x54, 0x6f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x54,
	0x0a, 0x16, 0x47, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x65, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x65, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x07, 0x50, 0x61, 0x67,
	0x65, 0x4c, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x07, 0x50, 0x61,
	0x67, 0x65, 0x4c, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0a,
	0x50, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x65, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x50, 0x61,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x65, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x88, 0x01, 0x01, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22,
	0x38, 0x0a, 0x1a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x2b, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x52,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x32, 0xcc, 0x03, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x3a, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4c, 0x0a, 0x11,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x54, 0x6f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65,
	0x72, 0x12, 0x1f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x54, 0x6f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3e, 0x0a, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x47, 0x69,
	0x76, 0x65, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x65, 0x12, 0x1d, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x50, 0x0a, 0x13, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x46, 0x72,
	0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x65, 0x12, 0x21, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x1a, 0x92, 0x41, 0x17, 0x12, 0x15,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_v1_order_proto_rawDescOnce sync.Once
	file_order_v1_order_proto_rawDescData = file_order_v1_order_proto_rawDesc
)

func file_order_v1_order_proto_rawDescGZIP() []byte {
	file_order_v1_order_proto_rawDescOnce.Do(func() {
		file_order_v1_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_v1_order_proto_rawDescData)
	})
	return file_order_v1_order_proto_rawDescData
}

var file_order_v1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_order_v1_order_proto_goTypes = []any{
	(*AddOrderRequest)(nil),            // 0: order.AddOrderRequest
	(*ReturnToDelivererRequest)(nil),   // 1: order.ReturnToDelivererRequest
	(*GiveToAddresseeRequest)(nil),     // 2: order.GiveToAddresseeRequest
	(*ListRefundRequest)(nil),          // 3: order.ListRefundRequest
	(*ListOrderRequest)(nil),           // 4: order.ListOrderRequest
	(*ReturnFromAddresseeRequest)(nil), // 5: order.ReturnFromAddresseeRequest
	(*ListOrderResponse)(nil),          // 6: order.ListOrderResponse
	(*ListRefundResponse)(nil),         // 7: order.ListRefundResponse
	(*timestamppb.Timestamp)(nil),      // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),              // 9: google.protobuf.Empty
}
var file_order_v1_order_proto_depIdxs = []int32{
	8, // 0: order.AddOrderRequest.ShelfLife:type_name -> google.protobuf.Timestamp
	0, // 1: order.Order.AddOrder:input_type -> order.AddOrderRequest
	1, // 2: order.Order.ReturnToDeliverer:input_type -> order.ReturnToDelivererRequest
	4, // 3: order.Order.ListOrder:input_type -> order.ListOrderRequest
	2, // 4: order.Order.GiveToAddressee:input_type -> order.GiveToAddresseeRequest
	5, // 5: order.Order.ReturnFromAddressee:input_type -> order.ReturnFromAddresseeRequest
	3, // 6: order.Order.ListRefund:input_type -> order.ListRefundRequest
	9, // 7: order.Order.AddOrder:output_type -> google.protobuf.Empty
	9, // 8: order.Order.ReturnToDeliverer:output_type -> google.protobuf.Empty
	6, // 9: order.Order.ListOrder:output_type -> order.ListOrderResponse
	9, // 10: order.Order.GiveToAddressee:output_type -> google.protobuf.Empty
	9, // 11: order.Order.ReturnFromAddressee:output_type -> google.protobuf.Empty
	7, // 12: order.Order.ListRefund:output_type -> order.ListRefundResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_order_v1_order_proto_init() }
func file_order_v1_order_proto_init() {
	if File_order_v1_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_v1_order_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddOrderRequest); i {
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
		file_order_v1_order_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ReturnToDelivererRequest); i {
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
		file_order_v1_order_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GiveToAddresseeRequest); i {
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
		file_order_v1_order_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListRefundRequest); i {
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
		file_order_v1_order_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListOrderRequest); i {
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
		file_order_v1_order_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ReturnFromAddresseeRequest); i {
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
		file_order_v1_order_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListOrderResponse); i {
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
		file_order_v1_order_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ListRefundResponse); i {
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
	file_order_v1_order_proto_msgTypes[0].OneofWrappers = []any{}
	file_order_v1_order_proto_msgTypes[3].OneofWrappers = []any{}
	file_order_v1_order_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_v1_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_v1_order_proto_goTypes,
		DependencyIndexes: file_order_v1_order_proto_depIdxs,
		MessageInfos:      file_order_v1_order_proto_msgTypes,
	}.Build()
	File_order_v1_order_proto = out.File
	file_order_v1_order_proto_rawDesc = nil
	file_order_v1_order_proto_goTypes = nil
	file_order_v1_order_proto_depIdxs = nil
}