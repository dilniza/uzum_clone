// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: order_status_notes.proto

package order_service

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

type CreateStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId   string `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status    string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	UserId    string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Reason    string `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt int32  `protobuf:"varint,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *CreateStatusRequest) Reset() {
	*x = CreateStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_status_notes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStatusRequest) ProtoMessage() {}

func (x *CreateStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_status_notes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStatusRequest.ProtoReflect.Descriptor instead.
func (*CreateStatusRequest) Descriptor() ([]byte, []int) {
	return file_order_status_notes_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateStatusRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateStatusRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateStatusRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *CreateStatusRequest) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CreateStatusRequest) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *CreateStatusRequest) GetDeletedAt() int32 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type OrderPrimaryKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrderPrimaryKeyRequest) Reset() {
	*x = OrderPrimaryKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_status_notes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPrimaryKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPrimaryKeyRequest) ProtoMessage() {}

func (x *OrderPrimaryKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_status_notes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPrimaryKeyRequest.ProtoReflect.Descriptor instead.
func (*OrderPrimaryKeyRequest) Descriptor() ([]byte, []int) {
	return file_order_status_notes_proto_rawDescGZIP(), []int{1}
}

func (x *OrderPrimaryKeyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type OrderPrimaryStatusKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OrderPrimaryStatusKeyRequest) Reset() {
	*x = OrderPrimaryStatusKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_status_notes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPrimaryStatusKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPrimaryStatusKeyRequest) ProtoMessage() {}

func (x *OrderPrimaryStatusKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_status_notes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPrimaryStatusKeyRequest.ProtoReflect.Descriptor instead.
func (*OrderPrimaryStatusKeyRequest) Descriptor() ([]byte, []int) {
	return file_order_status_notes_proto_rawDescGZIP(), []int{2}
}

func (x *OrderPrimaryStatusKeyRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderPrimaryStatusKeyRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetOrderStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId string `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status  string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	UserId  string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Reason  string `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *GetOrderStatusResponse) Reset() {
	*x = GetOrderStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_status_notes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderStatusResponse) ProtoMessage() {}

func (x *GetOrderStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_status_notes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderStatusResponse.ProtoReflect.Descriptor instead.
func (*GetOrderStatusResponse) Descriptor() ([]byte, []int) {
	return file_order_status_notes_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderStatusResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetOrderStatusResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *GetOrderStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetOrderStatusResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetOrderStatusResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_order_status_notes_proto protoreflect.FileDescriptor

var file_order_status_notes_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x28, 0x0a,
	0x16, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x1c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x32, 0xdc, 0x02, 0x0a, 0x12, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x49, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x20, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x05, 0x50, 0x55,
	0x54, 0x43, 0x48, 0x12, 0x26, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x79, 0x49, 0x44, 0x12, 0x26,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_status_notes_proto_rawDescOnce sync.Once
	file_order_status_notes_proto_rawDescData = file_order_status_notes_proto_rawDesc
)

func file_order_status_notes_proto_rawDescGZIP() []byte {
	file_order_status_notes_proto_rawDescOnce.Do(func() {
		file_order_status_notes_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_status_notes_proto_rawDescData)
	})
	return file_order_status_notes_proto_rawDescData
}

var file_order_status_notes_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_order_status_notes_proto_goTypes = []interface{}{
	(*CreateStatusRequest)(nil),          // 0: genproto.CreateStatusRequest
	(*OrderPrimaryKeyRequest)(nil),       // 1: genproto.OrderPrimaryKeyRequest
	(*OrderPrimaryStatusKeyRequest)(nil), // 2: genproto.OrderPrimaryStatusKeyRequest
	(*GetOrderStatusResponse)(nil),       // 3: genproto.GetOrderStatusResponse
}
var file_order_status_notes_proto_depIdxs = []int32{
	0, // 0: genproto.OrderStatusService.Create:input_type -> genproto.CreateStatusRequest
	1, // 1: genproto.OrderStatusService.GetByID:input_type -> genproto.OrderPrimaryKeyRequest
	2, // 2: genproto.OrderStatusService.PUTCH:input_type -> genproto.OrderPrimaryStatusKeyRequest
	2, // 3: genproto.OrderStatusService.GetStatusByID:input_type -> genproto.OrderPrimaryStatusKeyRequest
	3, // 4: genproto.OrderStatusService.Create:output_type -> genproto.GetOrderStatusResponse
	3, // 5: genproto.OrderStatusService.GetByID:output_type -> genproto.GetOrderStatusResponse
	3, // 6: genproto.OrderStatusService.PUTCH:output_type -> genproto.GetOrderStatusResponse
	3, // 7: genproto.OrderStatusService.GetStatusByID:output_type -> genproto.GetOrderStatusResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_status_notes_proto_init() }
func file_order_status_notes_proto_init() {
	if File_order_status_notes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_status_notes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStatusRequest); i {
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
		file_order_status_notes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPrimaryKeyRequest); i {
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
		file_order_status_notes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPrimaryStatusKeyRequest); i {
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
		file_order_status_notes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderStatusResponse); i {
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
			RawDescriptor: file_order_status_notes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_status_notes_proto_goTypes,
		DependencyIndexes: file_order_status_notes_proto_depIdxs,
		MessageInfos:      file_order_status_notes_proto_msgTypes,
	}.Build()
	File_order_status_notes_proto = out.File
	file_order_status_notes_proto_rawDesc = nil
	file_order_status_notes_proto_goTypes = nil
	file_order_status_notes_proto_depIdxs = nil
}