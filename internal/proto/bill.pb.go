// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bill.proto

package internal

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetBillsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBillsRequest) Reset()         { *m = GetBillsRequest{} }
func (m *GetBillsRequest) String() string { return proto.CompactTextString(m) }
func (*GetBillsRequest) ProtoMessage()    {}
func (*GetBillsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_24141385411f973b, []int{0}
}

func (m *GetBillsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBillsRequest.Unmarshal(m, b)
}
func (m *GetBillsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBillsRequest.Marshal(b, m, deterministic)
}
func (m *GetBillsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBillsRequest.Merge(m, src)
}
func (m *GetBillsRequest) XXX_Size() int {
	return xxx_messageInfo_GetBillsRequest.Size(m)
}
func (m *GetBillsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBillsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBillsRequest proto.InternalMessageInfo

type BillInfo struct {
	BillId               string   `protobuf:"bytes,1,opt,name=billId,proto3" json:"billId,omitempty"`
	ReceiveAddr          string   `protobuf:"bytes,2,opt,name=receiveAddr,proto3" json:"receiveAddr,omitempty"`
	BuyTime              int64    `protobuf:"varint,3,opt,name=buyTime,proto3" json:"buyTime,omitempty"`
	Status               int32    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	GoodsId              int32    `protobuf:"varint,5,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
	GoodsNum             int32    `protobuf:"varint,6,opt,name=goodsNum,proto3" json:"goodsNum,omitempty"`
	GoodsName            string   `protobuf:"bytes,7,opt,name=goodsName,proto3" json:"goodsName,omitempty"`
	PayTime              int64    `protobuf:"varint,8,opt,name=payTime,proto3" json:"payTime,omitempty"`
	PostId               string   `protobuf:"bytes,9,opt,name=postId,proto3" json:"postId,omitempty"`
	PostTime             int64    `protobuf:"varint,10,opt,name=postTime,proto3" json:"postTime,omitempty"`
	ReceiveTime          int64    `protobuf:"varint,11,opt,name=receiveTime,proto3" json:"receiveTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BillInfo) Reset()         { *m = BillInfo{} }
func (m *BillInfo) String() string { return proto.CompactTextString(m) }
func (*BillInfo) ProtoMessage()    {}
func (*BillInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_24141385411f973b, []int{1}
}

func (m *BillInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillInfo.Unmarshal(m, b)
}
func (m *BillInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillInfo.Marshal(b, m, deterministic)
}
func (m *BillInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillInfo.Merge(m, src)
}
func (m *BillInfo) XXX_Size() int {
	return xxx_messageInfo_BillInfo.Size(m)
}
func (m *BillInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BillInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BillInfo proto.InternalMessageInfo

func (m *BillInfo) GetBillId() string {
	if m != nil {
		return m.BillId
	}
	return ""
}

func (m *BillInfo) GetReceiveAddr() string {
	if m != nil {
		return m.ReceiveAddr
	}
	return ""
}

func (m *BillInfo) GetBuyTime() int64 {
	if m != nil {
		return m.BuyTime
	}
	return 0
}

func (m *BillInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *BillInfo) GetGoodsId() int32 {
	if m != nil {
		return m.GoodsId
	}
	return 0
}

func (m *BillInfo) GetGoodsNum() int32 {
	if m != nil {
		return m.GoodsNum
	}
	return 0
}

func (m *BillInfo) GetGoodsName() string {
	if m != nil {
		return m.GoodsName
	}
	return ""
}

func (m *BillInfo) GetPayTime() int64 {
	if m != nil {
		return m.PayTime
	}
	return 0
}

func (m *BillInfo) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

func (m *BillInfo) GetPostTime() int64 {
	if m != nil {
		return m.PostTime
	}
	return 0
}

func (m *BillInfo) GetReceiveTime() int64 {
	if m != nil {
		return m.ReceiveTime
	}
	return 0
}

type GetBillsResponse struct {
	Code                 int32       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Bills                []*BillInfo `protobuf:"bytes,2,rep,name=bills,proto3" json:"bills,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetBillsResponse) Reset()         { *m = GetBillsResponse{} }
func (m *GetBillsResponse) String() string { return proto.CompactTextString(m) }
func (*GetBillsResponse) ProtoMessage()    {}
func (*GetBillsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_24141385411f973b, []int{2}
}

func (m *GetBillsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBillsResponse.Unmarshal(m, b)
}
func (m *GetBillsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBillsResponse.Marshal(b, m, deterministic)
}
func (m *GetBillsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBillsResponse.Merge(m, src)
}
func (m *GetBillsResponse) XXX_Size() int {
	return xxx_messageInfo_GetBillsResponse.Size(m)
}
func (m *GetBillsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBillsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBillsResponse proto.InternalMessageInfo

func (m *GetBillsResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetBillsResponse) GetBills() []*BillInfo {
	if m != nil {
		return m.Bills
	}
	return nil
}

func init() {
	proto.RegisterType((*GetBillsRequest)(nil), "internal.GetBillsRequest")
	proto.RegisterType((*BillInfo)(nil), "internal.BillInfo")
	proto.RegisterType((*GetBillsResponse)(nil), "internal.GetBillsResponse")
}

func init() { proto.RegisterFile("bill.proto", fileDescriptor_24141385411f973b) }

var fileDescriptor_24141385411f973b = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x41, 0x4f, 0xb4, 0x30,
	0x10, 0xfd, 0xd8, 0x5d, 0x76, 0x61, 0x38, 0x7c, 0x3a, 0x07, 0x53, 0x89, 0x07, 0xc2, 0x89, 0x13,
	0x87, 0xf5, 0x17, 0xa8, 0x07, 0x43, 0x4c, 0x8c, 0x69, 0xfc, 0x03, 0xb0, 0xad, 0x86, 0x04, 0x28,
	0xd2, 0x62, 0xe2, 0x5f, 0xf3, 0xd7, 0x99, 0x4e, 0x61, 0x21, 0xc6, 0xdb, 0xbc, 0xf7, 0x3a, 0xd3,
	0x37, 0x6f, 0x00, 0xaa, 0xba, 0x69, 0xf2, 0x7e, 0x50, 0x46, 0x61, 0x50, 0x77, 0x46, 0x0e, 0x5d,
	0xd9, 0xa4, 0x97, 0xf0, 0xff, 0x51, 0x9a, 0xfb, 0xba, 0x69, 0x34, 0x97, 0x1f, 0xa3, 0xd4, 0x26,
	0xfd, 0xde, 0x40, 0x60, 0x89, 0xa2, 0x7b, 0x53, 0x78, 0x05, 0x7b, 0xdb, 0x57, 0x08, 0xe6, 0x25,
	0x5e, 0x16, 0xf2, 0x09, 0x61, 0x02, 0xd1, 0x20, 0x4f, 0xb2, 0xfe, 0x94, 0x77, 0x42, 0x0c, 0x6c,
	0x43, 0xe2, 0x9a, 0x42, 0x06, 0x87, 0x6a, 0xfc, 0x7a, 0xad, 0x5b, 0xc9, 0xb6, 0x89, 0x97, 0x6d,
	0xf9, 0x0c, 0xed, 0x4c, 0x6d, 0x4a, 0x33, 0x6a, 0xb6, 0x4b, 0xbc, 0xcc, 0xe7, 0x13, 0xb2, 0x1d,
	0xef, 0x4a, 0x09, 0x5d, 0x08, 0xe6, 0x93, 0x30, 0x43, 0x8c, 0x21, 0xa0, 0xf2, 0x79, 0x6c, 0xd9,
	0x9e, 0xa4, 0x33, 0xc6, 0x1b, 0x08, 0x5d, 0x5d, 0xb6, 0x92, 0x1d, 0xc8, 0xc7, 0x42, 0xd8, 0x99,
	0x7d, 0xe9, 0x5c, 0x04, 0xce, 0xc5, 0x04, 0xad, 0x8b, 0x5e, 0x69, 0x53, 0x08, 0x16, 0xba, 0xcd,
	0x1c, 0xb2, 0x7f, 0xd9, 0x8a, 0x5a, 0x80, 0x5a, 0xce, 0x78, 0xb5, 0x35, 0xc9, 0x11, 0xc9, 0x6b,
	0x2a, 0x7d, 0x81, 0x8b, 0x25, 0x4f, 0xdd, 0xab, 0x4e, 0x4b, 0x44, 0xd8, 0x9d, 0x94, 0x90, 0x94,
	0xa0, 0xcf, 0xa9, 0xc6, 0x0c, 0x7c, 0x9b, 0xa4, 0x66, 0x9b, 0x64, 0x9b, 0x45, 0x47, 0xcc, 0xe7,
	0x8b, 0xe4, 0x73, 0xf4, 0xdc, 0x3d, 0x38, 0x3e, 0xc1, 0xce, 0x52, 0xf8, 0x00, 0xc1, 0x3c, 0x19,
	0xaf, 0x97, 0xe7, 0xbf, 0xae, 0x17, 0xc7, 0x7f, 0x49, 0xce, 0x48, 0xfa, 0xaf, 0xda, 0xd3, 0xfd,
	0x6f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x9c, 0x04, 0x55, 0x0d, 0x02, 0x00, 0x00,
}
