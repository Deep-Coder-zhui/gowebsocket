// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/custom_interest_status.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum containing possible custom interest types.
type CustomInterestStatusEnum_CustomInterestStatus int32

const (
	// Not specified.
	CustomInterestStatusEnum_UNSPECIFIED CustomInterestStatusEnum_CustomInterestStatus = 0
	// Used for return value only. Represents value unknown in this version.
	CustomInterestStatusEnum_UNKNOWN CustomInterestStatusEnum_CustomInterestStatus = 1
	// Enabled status - custom interest is enabled and can be targeted to.
	CustomInterestStatusEnum_ENABLED CustomInterestStatusEnum_CustomInterestStatus = 2
	// Removed status - custom interest is removed and cannot be used for
	// targeting.
	CustomInterestStatusEnum_REMOVED CustomInterestStatusEnum_CustomInterestStatus = 3
)

var CustomInterestStatusEnum_CustomInterestStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}

var CustomInterestStatusEnum_CustomInterestStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x CustomInterestStatusEnum_CustomInterestStatus) String() string {
	return proto.EnumName(CustomInterestStatusEnum_CustomInterestStatus_name, int32(x))
}

func (CustomInterestStatusEnum_CustomInterestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_65be17a4ddbb4c40, []int{0, 0}
}

// The status of custom interest.
type CustomInterestStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomInterestStatusEnum) Reset()         { *m = CustomInterestStatusEnum{} }
func (m *CustomInterestStatusEnum) String() string { return proto.CompactTextString(m) }
func (*CustomInterestStatusEnum) ProtoMessage()    {}
func (*CustomInterestStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_65be17a4ddbb4c40, []int{0}
}

func (m *CustomInterestStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomInterestStatusEnum.Unmarshal(m, b)
}
func (m *CustomInterestStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomInterestStatusEnum.Marshal(b, m, deterministic)
}
func (m *CustomInterestStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomInterestStatusEnum.Merge(m, src)
}
func (m *CustomInterestStatusEnum) XXX_Size() int {
	return xxx_messageInfo_CustomInterestStatusEnum.Size(m)
}
func (m *CustomInterestStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomInterestStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomInterestStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.CustomInterestStatusEnum_CustomInterestStatus", CustomInterestStatusEnum_CustomInterestStatus_name, CustomInterestStatusEnum_CustomInterestStatus_value)
	proto.RegisterType((*CustomInterestStatusEnum)(nil), "google.ads.googleads.v2.enums.CustomInterestStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/custom_interest_status.proto", fileDescriptor_65be17a4ddbb4c40)
}

var fileDescriptor_65be17a4ddbb4c40 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0xb5, 0x5b, 0x50, 0x48, 0x0f, 0x2e, 0xc5, 0x83, 0x8a, 0x3d, 0xb4, 0x1f, 0x90, 0xc0, 0x7a,
	0x8b, 0xa7, 0x6c, 0x1b, 0x4b, 0x51, 0xd3, 0x62, 0xe9, 0x0a, 0xb2, 0x50, 0x62, 0x77, 0x09, 0x2b,
	0xdd, 0xa4, 0x74, 0xb2, 0xfd, 0x20, 0x8f, 0x7e, 0x8a, 0x9f, 0xe2, 0xc1, 0x6f, 0x90, 0x64, 0xdb,
	0x3d, 0x55, 0x2f, 0xe1, 0x4d, 0xde, 0xbc, 0x37, 0x6f, 0x06, 0x51, 0x65, 0x8c, 0x5a, 0xe7, 0x44,
	0x66, 0x40, 0x6a, 0xe8, 0xd0, 0x2e, 0x22, 0xb9, 0xae, 0x4a, 0x20, 0xab, 0x0a, 0xac, 0x29, 0x97,
	0x85, 0xb6, 0xf9, 0x36, 0x07, 0xbb, 0x04, 0x2b, 0x6d, 0x05, 0x78, 0xb3, 0x35, 0xd6, 0x74, 0x7b,
	0xb5, 0x00, 0xcb, 0x0c, 0x70, 0xa3, 0xc5, 0xbb, 0x08, 0x7b, 0xed, 0xf5, 0xcd, 0xc1, 0x7a, 0x53,
	0x10, 0xa9, 0xb5, 0xb1, 0xd2, 0x16, 0x46, 0xef, 0xc5, 0x83, 0x77, 0x74, 0x39, 0xf4, 0xe6, 0x93,
	0xbd, 0xf7, 0xdc, 0x5b, 0x73, 0x5d, 0x95, 0x03, 0x81, 0x2e, 0x8e, 0x71, 0xdd, 0x73, 0xd4, 0x59,
	0x88, 0xf9, 0x8c, 0x0f, 0x27, 0xf7, 0x13, 0x3e, 0x0a, 0x4f, 0xba, 0x1d, 0x74, 0xb6, 0x10, 0x0f,
	0x62, 0xfa, 0x22, 0xc2, 0x96, 0x2b, 0xb8, 0x60, 0xf1, 0x23, 0x1f, 0x85, 0x81, 0x2b, 0x9e, 0xf9,
	0xd3, 0x34, 0xe1, 0xa3, 0xb0, 0x1d, 0xff, 0xb4, 0x50, 0x7f, 0x65, 0x4a, 0xfc, 0x6f, 0xde, 0xf8,
	0xea, 0xd8, 0xcc, 0x99, 0x0b, 0x3b, 0x6b, 0xbd, 0xc6, 0x7b, 0xad, 0x32, 0x6b, 0xa9, 0x15, 0x36,
	0x5b, 0x45, 0x54, 0xae, 0xfd, 0x2a, 0x87, 0xbb, 0x6d, 0x0a, 0xf8, 0xe3, 0x8c, 0x77, 0xfe, 0xfd,
	0x08, 0xda, 0x63, 0xc6, 0x3e, 0x83, 0xde, 0xb8, 0xb6, 0x62, 0x19, 0xe0, 0x1a, 0x3a, 0x94, 0x44,
	0xd8, 0xed, 0x0e, 0x5f, 0x07, 0x3e, 0x65, 0x19, 0xa4, 0x0d, 0x9f, 0x26, 0x51, 0xea, 0xf9, 0xef,
	0xa0, 0x5f, 0x7f, 0x52, 0xca, 0x32, 0xa0, 0xb4, 0xe9, 0xa0, 0x34, 0x89, 0x28, 0xf5, 0x3d, 0x6f,
	0xa7, 0x3e, 0xd8, 0xed, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0x23, 0x10, 0x2f, 0xde, 0x01,
	0x00, 0x00,
}
