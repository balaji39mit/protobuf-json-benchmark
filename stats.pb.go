// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stats.proto

package proto_files

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AggrData struct {
	TotalAmount          float32           `protobuf:"fixed32,1,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
	Count                int32             `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	Ids                  map[string]string `protobuf:"bytes,3,rep,name=Ids,proto3" json:"Ids,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MerchantId           map[string]string `protobuf:"bytes,4,rep,name=MerchantId,proto3" json:"MerchantId,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AggrData) Reset()         { *m = AggrData{} }
func (m *AggrData) String() string { return proto.CompactTextString(m) }
func (*AggrData) ProtoMessage()    {}
func (*AggrData) Descriptor() ([]byte, []int) {
	return fileDescriptor_stats_1017553db8d8bcf8, []int{0}
}
func (m *AggrData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggrData.Unmarshal(m, b)
}
func (m *AggrData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggrData.Marshal(b, m, deterministic)
}
func (dst *AggrData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggrData.Merge(dst, src)
}
func (m *AggrData) XXX_Size() int {
	return xxx_messageInfo_AggrData.Size(m)
}
func (m *AggrData) XXX_DiscardUnknown() {
	xxx_messageInfo_AggrData.DiscardUnknown(m)
}

var xxx_messageInfo_AggrData proto.InternalMessageInfo

func (m *AggrData) GetTotalAmount() float32 {
	if m != nil {
		return m.TotalAmount
	}
	return 0
}

func (m *AggrData) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *AggrData) GetIds() map[string]string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *AggrData) GetMerchantId() map[string]string {
	if m != nil {
		return m.MerchantId
	}
	return nil
}

type Stats struct {
	Year                 int32     `protobuf:"varint,1,opt,name=Year,proto3" json:"Year,omitempty"`
	Month                int32     `protobuf:"varint,2,opt,name=Month,proto3" json:"Month,omitempty"`
	Day                  int32     `protobuf:"varint,3,opt,name=Day,proto3" json:"Day,omitempty"`
	Hour                 int32     `protobuf:"varint,4,opt,name=Hour,proto3" json:"Hour,omitempty"`
	AggrData             *AggrData `protobuf:"bytes,5,opt,name=AggrData,proto3" json:"AggrData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_stats_1017553db8d8bcf8, []int{1}
}
func (m *Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stats.Unmarshal(m, b)
}
func (m *Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stats.Marshal(b, m, deterministic)
}
func (dst *Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats.Merge(dst, src)
}
func (m *Stats) XXX_Size() int {
	return xxx_messageInfo_Stats.Size(m)
}
func (m *Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Stats proto.InternalMessageInfo

func (m *Stats) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Stats) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Stats) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *Stats) GetHour() int32 {
	if m != nil {
		return m.Hour
	}
	return 0
}

func (m *Stats) GetAggrData() *AggrData {
	if m != nil {
		return m.AggrData
	}
	return nil
}

type StatsArray struct {
	StatsArray           []*Stats `protobuf:"bytes,1,rep,name=statsArray,proto3" json:"statsArray,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsArray) Reset()         { *m = StatsArray{} }
func (m *StatsArray) String() string { return proto.CompactTextString(m) }
func (*StatsArray) ProtoMessage()    {}
func (*StatsArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_stats_1017553db8d8bcf8, []int{2}
}
func (m *StatsArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsArray.Unmarshal(m, b)
}
func (m *StatsArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsArray.Marshal(b, m, deterministic)
}
func (dst *StatsArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsArray.Merge(dst, src)
}
func (m *StatsArray) XXX_Size() int {
	return xxx_messageInfo_StatsArray.Size(m)
}
func (m *StatsArray) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsArray.DiscardUnknown(m)
}

var xxx_messageInfo_StatsArray proto.InternalMessageInfo

func (m *StatsArray) GetStatsArray() []*Stats {
	if m != nil {
		return m.StatsArray
	}
	return nil
}

func init() {
	proto.RegisterType((*AggrData)(nil), "proto_files.AggrData")
	proto.RegisterMapType((map[string]string)(nil), "proto_files.AggrData.IdsEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto_files.AggrData.MerchantIdEntry")
	proto.RegisterType((*Stats)(nil), "proto_files.Stats")
	proto.RegisterType((*StatsArray)(nil), "proto_files.StatsArray")
}

func init() { proto.RegisterFile("stats.proto", fileDescriptor_stats_1017553db8d8bcf8) }

var fileDescriptor_stats_1017553db8d8bcf8 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xd9, 0xa4, 0xf9, 0xd3, 0x4e, 0x0e, 0x7f, 0x19, 0x14, 0x96, 0x1e, 0x24, 0x14, 0x84,
	0x9c, 0x82, 0x46, 0x10, 0x11, 0x04, 0x83, 0x2d, 0x98, 0x43, 0x2f, 0xab, 0x17, 0x4f, 0xb2, 0x9a,
	0xd8, 0x8a, 0x31, 0x2b, 0xbb, 0x1b, 0x21, 0x5f, 0xc1, 0x8f, 0xe3, 0x27, 0x94, 0x9d, 0x90, 0x36,
	0x4a, 0x2f, 0x9e, 0x32, 0xb3, 0xf3, 0x7e, 0x8f, 0xc7, 0x0b, 0x84, 0xc6, 0x4a, 0x6b, 0x92, 0x77,
	0xad, 0xac, 0xc2, 0x90, 0x3e, 0x0f, 0xcf, 0x2f, 0x55, 0x69, 0x66, 0x5f, 0x1e, 0x8c, 0xb3, 0xd5,
	0x4a, 0xcf, 0xa5, 0x95, 0x18, 0x41, 0x78, 0xa7, 0xac, 0xac, 0xb2, 0x37, 0xd5, 0xd4, 0x96, 0xb3,
	0x88, 0xc5, 0x9e, 0x18, 0x3e, 0xe1, 0x3e, 0x04, 0xd7, 0x74, 0xf3, 0x22, 0x16, 0x07, 0xa2, 0x5b,
	0xf0, 0x18, 0xfc, 0xbc, 0x30, 0xdc, 0x8f, 0xfc, 0x38, 0x4c, 0x0f, 0x93, 0x81, 0x7f, 0xd2, 0x7b,
	0x27, 0x79, 0x61, 0x16, 0xb5, 0xd5, 0xad, 0x70, 0x52, 0x5c, 0x00, 0x2c, 0x4b, 0xfd, 0xb4, 0x96,
	0xb5, 0xcd, 0x0b, 0x3e, 0x22, 0xf0, 0x68, 0x37, 0xb8, 0xd5, 0x75, 0xfc, 0x00, 0x9c, 0x9e, 0xc1,
	0xb8, 0xf7, 0xc5, 0x3d, 0xf0, 0x5f, 0xcb, 0x96, 0x42, 0x4f, 0x84, 0x1b, 0x5d, 0xd8, 0x0f, 0x59,
	0x35, 0x25, 0x85, 0x9d, 0x88, 0x6e, 0xb9, 0xf0, 0xce, 0xd9, 0xf4, 0x12, 0xfe, 0xff, 0xb2, 0xfd,
	0x0b, 0x3e, 0xfb, 0x64, 0x10, 0xdc, 0xba, 0x46, 0x11, 0x61, 0x74, 0x5f, 0x4a, 0x4d, 0x58, 0x20,
	0x68, 0x76, 0xdc, 0x52, 0xd5, 0x76, 0xdd, 0x77, 0x44, 0x8b, 0xf3, 0x9f, 0xcb, 0x96, 0xfb, 0xf4,
	0xe6, 0x46, 0xc7, 0xde, 0xa8, 0x46, 0xf3, 0x51, 0xc7, 0xba, 0x19, 0x4f, 0xb6, 0x7f, 0x83, 0x07,
	0x11, 0x8b, 0xc3, 0xf4, 0x60, 0x67, 0x2b, 0x62, 0x23, 0x9b, 0x5d, 0x01, 0x50, 0x96, 0x4c, 0x6b,
	0xd9, 0x62, 0x0a, 0x60, 0x36, 0x1b, 0x67, 0x54, 0x2c, 0xfe, 0xb0, 0x20, 0xb1, 0x18, 0xa8, 0x1e,
	0xff, 0xd1, 0xf9, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xae, 0x14, 0xa6, 0x26, 0x02, 0x00,
	0x00,
}