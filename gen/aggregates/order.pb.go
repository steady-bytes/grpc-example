// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aggregates/order.proto

package order_v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
// `Order` is an aggregate type, meaning that it's the top level collection
// of many different fields that can also include other value objects. A native
// type can be validated, and so can value objects.
type Order struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// items a user wishes to order
	Items []*Item `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	// users payment method of their choice for this order
	PaymentMethod        *PaymentMethod `protobuf:"bytes,4,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_428bf5ef2a905b9d, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Order) GetPaymentMethod() *PaymentMethod {
	if m != nil {
		return m.PaymentMethod
	}
	return nil
}

type Item struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_428bf5ef2a905b9d, []int{1}
}
func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PaymentMethod struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentMethod) Reset()         { *m = PaymentMethod{} }
func (m *PaymentMethod) String() string { return proto.CompactTextString(m) }
func (*PaymentMethod) ProtoMessage()    {}
func (*PaymentMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_428bf5ef2a905b9d, []int{2}
}
func (m *PaymentMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentMethod.Unmarshal(m, b)
}
func (m *PaymentMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentMethod.Marshal(b, m, deterministic)
}
func (m *PaymentMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentMethod.Merge(m, src)
}
func (m *PaymentMethod) XXX_Size() int {
	return xxx_messageInfo_PaymentMethod.Size(m)
}
func (m *PaymentMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentMethod.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentMethod proto.InternalMessageInfo

func (m *PaymentMethod) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Order)(nil), "order_v1.Order")
	proto.RegisterType((*Item)(nil), "order_v1.Item")
	proto.RegisterType((*PaymentMethod)(nil), "order_v1.PaymentMethod")
}

func init() { proto.RegisterFile("aggregates/order.proto", fileDescriptor_428bf5ef2a905b9d) }

var fileDescriptor_428bf5ef2a905b9d = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x4c, 0x4f, 0x2f,
	0x4a, 0x4d, 0x4f, 0x2c, 0x49, 0x2d, 0xd6, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x73, 0xe2, 0xcb, 0x0c, 0x95, 0x6a, 0xb9, 0x58, 0xfd, 0x41, 0x6c,
	0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x14,
	0x21, 0x15, 0x2e, 0xd6, 0xcc, 0x92, 0xd4, 0xdc, 0x62, 0x09, 0x66, 0x05, 0x66, 0x0d, 0x6e, 0x23,
	0x3e, 0x3d, 0x98, 0x16, 0x3d, 0xcf, 0x92, 0xd4, 0xdc, 0x20, 0x88, 0xa4, 0x90, 0x1d, 0x17, 0x5f,
	0x41, 0x62, 0x65, 0x6e, 0x6a, 0x5e, 0x49, 0x7c, 0x6e, 0x6a, 0x49, 0x46, 0x7e, 0x8a, 0x04, 0x8b,
	0x02, 0xa3, 0x06, 0xb7, 0x91, 0x38, 0x42, 0x79, 0x00, 0x44, 0xde, 0x17, 0x2c, 0x1d, 0xc4, 0x5b,
	0x80, 0xcc, 0x55, 0xd2, 0xe2, 0x62, 0x01, 0x19, 0x87, 0x61, 0xbb, 0x10, 0x17, 0x4b, 0x5e, 0x62,
	0x6e, 0xaa, 0x04, 0x13, 0x58, 0x04, 0xcc, 0x56, 0x92, 0xe7, 0xe2, 0x45, 0x31, 0x0b, 0x5d, 0x53,
	0x12, 0x1b, 0xd8, 0x73, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x70, 0xec, 0xb3, 0x4b, 0xf6,
	0x00, 0x00, 0x00,
}
