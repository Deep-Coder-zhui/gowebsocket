// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grafeas/v1/package.proto

package grafeas

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Instruction set architectures supported by various package managers.
type Architecture int32

const (
	// Unknown architecture.
	Architecture_ARCHITECTURE_UNSPECIFIED Architecture = 0
	// X86 architecture.
	Architecture_X86 Architecture = 1
	// X64 architecture.
	Architecture_X64 Architecture = 2
)

var Architecture_name = map[int32]string{
	0: "ARCHITECTURE_UNSPECIFIED",
	1: "X86",
	2: "X64",
}

var Architecture_value = map[string]int32{
	"ARCHITECTURE_UNSPECIFIED": 0,
	"X86":                      1,
	"X64":                      2,
}

func (x Architecture) String() string {
	return proto.EnumName(Architecture_name, int32(x))
}

func (Architecture) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6152b3fff9015bb3, []int{0}
}

// Whether this is an ordinary package version or a sentinel MIN/MAX version.
type Version_VersionKind int32

const (
	// Unknown.
	Version_VERSION_KIND_UNSPECIFIED Version_VersionKind = 0
	// A standard package version.
	Version_NORMAL Version_VersionKind = 1
	// A special version representing negative infinity.
	Version_MINIMUM Version_VersionKind = 2
	// A special version representing positive infinity.
	Version_MAXIMUM Version_VersionKind = 3
)

var Version_VersionKind_name = map[int32]string{
	0: "VERSION_KIND_UNSPECIFIED",
	1: "NORMAL",
	2: "MINIMUM",
	3: "MAXIMUM",
}

var Version_VersionKind_value = map[string]int32{
	"VERSION_KIND_UNSPECIFIED": 0,
	"NORMAL":                   1,
	"MINIMUM":                  2,
	"MAXIMUM":                  3,
}

func (x Version_VersionKind) String() string {
	return proto.EnumName(Version_VersionKind_name, int32(x))
}

func (Version_VersionKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6152b3fff9015bb3, []int{4, 0}
}

// This represents a particular channel of distribution for a given package.
// E.g., Debian's jessie-backports dpkg mirror.
type Distribution struct {
	// Required. The cpe_uri in [CPE format](https://cpe.mitre.org/specification/)
	// denoting the package manager version distributing a package.
	CpeUri string `protobuf:"bytes,1,opt,name=cpe_uri,json=cpeUri,proto3" json:"cpe_uri,omitempty"`
	// The CPU architecture for which packages in this distribution channel were
	// built.
	Architecture Architecture `protobuf:"varint,2,opt,name=architecture,proto3,enum=grafeas.v1.Architecture" json:"architecture,omitempty"`
	// The latest available version of this package in this distribution channel.
	LatestVersion *Version `protobuf:"bytes,3,opt,name=latest_version,json=latestVersion,proto3" json:"latest_version,omitempty"`
	// A freeform string denoting the maintainer of this package.
	Maintainer string `protobuf:"bytes,4,opt,name=maintainer,proto3" json:"maintainer,omitempty"`
	// The distribution channel-specific homepage for this package.
	Url string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	// The distribution channel-specific description of this package.
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Distribution) Reset()         { *m = Distribution{} }
func (m *Distribution) String() string { return proto.CompactTextString(m) }
func (*Distribution) ProtoMessage()    {}
func (*Distribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_6152b3fff9015bb3, []int{0}
}

func (m *Distribution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Distribution.Unmarshal(m, b)
}
func (m *Distribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Distribution.Marshal(b, m, deterministic)
}
func (m *Distribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Distribution.Merge(m, src)
}
func (m *Distribution) XXX_Size() int {
	return xxx_messageInfo_Distribution.Size(m)
}
func (m *Distribution) XXX_DiscardUnknown() {
	xxx_messageInfo_Distribution.DiscardUnknown(m)
}

var xxx_messageInfo_Distribution proto.InternalMessageInfo

func (m *Distribution) GetCpeUri() string {
	if m != nil {
		return m.CpeUri
	}
	return ""
}

func (m *Distribution) GetArchitecture() Architecture {
	if m != nil {
		return m.Architecture
	}
	return Architecture_ARCHITECTURE_UNSPECIFIED
}

func (m *Distribution) GetLatestVersion() *Version {
	if m != nil {
		return m.LatestVersion
	}
	return nil
}

func (m *Distribution) GetMaintainer() string {
	if m != nil {
		return m.Maintainer
	}
	return ""
}

func (m *Distribution) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Distribution) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// An occurrence of a particular package installation found within a system's
// filesystem. E.g., glibc was found in `/var/lib/dpkg/status`.
type Location struct {
	// Required. The CPE URI in [CPE format](https://cpe.mitre.org/specification/)
	// denoting the package manager version distributing a package.
	CpeUri string `protobuf:"bytes,1,opt,name=cpe_uri,json=cpeUri,proto3" json:"cpe_uri,omitempty"`
	// The version installed at this location.
	Version *Version `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// The path from which we gathered that this package/version is installed.
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_6152b3fff9015bb3, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetCpeUri() string {
	if m != nil {
		return m.CpeUri
	}
	return ""
}

func (m *Location) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *Location) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// This represents a particular package that is distributed over various
// channels. E.g., glibc (aka libc6) is distributed by many, at various
// versions.
type PackageNote struct {
	// Required. Immutable. The name of the package.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The various channels by which a package is distributed.
	Distribution         []*Distribution `protobuf:"bytes,10,rep,name=distribution,proto3" json:"distribution,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PackageNote) Reset()         { *m = PackageNote{} }
func (m *PackageNote) String() string { return proto.CompactTextString(m) }
func (*PackageNote) ProtoMessage()    {}
func (*PackageNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_6152b3fff9015bb3, []int{2}
}

func (m *PackageNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageNote.Unmarshal(m, b)
}
func (m *PackageNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageNote.Marshal(b, m, deterministic)
}
func (m *PackageNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageNote.Merge(m, src)
}
func (m *PackageNote) XXX_Size() int {
	return xxx_messageInfo_PackageNote.Size(m)
}
func (m *PackageNote) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageNote.DiscardUnknown(m)
}

var xxx_messageInfo_PackageNote proto.InternalMessageInfo

func (m *PackageNote) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PackageNote) GetDistribution() []*Distribution {
	if m != nil {
		return m.Distribution
	}
	return nil
}

// Details on how a particular software package was installed on a system.
type PackageOccurrence struct {
	// Output only. The name of the installed package.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. All of the places within the filesystem versions of this package
	// have been found.
	Location             []*Location `protobuf:"bytes,2,rep,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PackageOccurrence) Reset()         { *m = PackageOccurrence{} }
func (m *PackageOccurrence) String() string { return proto.CompactTextString(m) }
func (*PackageOccurrence) ProtoMessage()    {}
func (*PackageOccurrence) Descriptor() ([]byte, []int) {
	return fileDescriptor_6152b3fff9015bb3, []int{3}
}

func (m *PackageOccurrence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageOccurrence.Unmarshal(m, b)
}
func (m *PackageOccurrence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageOccurrence.Marshal(b, m, deterministic)
}
func (m *PackageOccurrence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageOccurrence.Merge(m, src)
}
func (m *PackageOccurrence) XXX_Size() int {
	return xxx_messageInfo_PackageOccurrence.Size(m)
}
func (m *PackageOccurrence) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageOccurrence.DiscardUnknown(m)
}

var xxx_messageInfo_PackageOccurrence proto.InternalMessageInfo

func (m *PackageOccurrence) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PackageOccurrence) GetLocation() []*Location {
	if m != nil {
		return m.Location
	}
	return nil
}

// Version contains structured information about the version of a package.
type Version struct {
	// Used to correct mistakes in the version numbering scheme.
	Epoch int32 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// Required only when version kind is NORMAL. The main part of the version
	// name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The iteration of the package build from the above version.
	Revision string `protobuf:"bytes,3,opt,name=revision,proto3" json:"revision,omitempty"`
	// Required. Distinguishes between sentinel MIN/MAX versions and normal
	// versions.
	Kind Version_VersionKind `protobuf:"varint,4,opt,name=kind,proto3,enum=grafeas.v1.Version_VersionKind" json:"kind,omitempty"`
	// Human readable version string. This string is of the form
	// <epoch>:<name>-<revision> and is only set when kind is NORMAL.
	FullName             string   `protobuf:"bytes,5,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_6152b3fff9015bb3, []int{4}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetEpoch() int32 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *Version) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Version) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *Version) GetKind() Version_VersionKind {
	if m != nil {
		return m.Kind
	}
	return Version_VERSION_KIND_UNSPECIFIED
}

func (m *Version) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func init() {
	proto.RegisterEnum("grafeas.v1.Architecture", Architecture_name, Architecture_value)
	proto.RegisterEnum("grafeas.v1.Version_VersionKind", Version_VersionKind_name, Version_VersionKind_value)
	proto.RegisterType((*Distribution)(nil), "grafeas.v1.Distribution")
	proto.RegisterType((*Location)(nil), "grafeas.v1.Location")
	proto.RegisterType((*PackageNote)(nil), "grafeas.v1.PackageNote")
	proto.RegisterType((*PackageOccurrence)(nil), "grafeas.v1.PackageOccurrence")
	proto.RegisterType((*Version)(nil), "grafeas.v1.Version")
}

func init() { proto.RegisterFile("grafeas/v1/package.proto", fileDescriptor_6152b3fff9015bb3) }

var fileDescriptor_6152b3fff9015bb3 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0x26, 0xc9, 0xd6, 0x1f, 0x97, 0x6d, 0x0a, 0x66, 0x12, 0x11, 0x20, 0xa8, 0xf2, 0x54, 0x21,
	0x91, 0xb2, 0x0e, 0x4d, 0x13, 0x20, 0xa4, 0xd2, 0x16, 0x88, 0xb6, 0xa6, 0x9d, 0xb7, 0x4e, 0x83,
	0x97, 0xc8, 0x4b, 0xdd, 0xd4, 0x5a, 0x1a, 0x47, 0x4e, 0xd2, 0x3f, 0x88, 0xbf, 0x11, 0xde, 0x51,
	0x9c, 0xb4, 0x4b, 0xd1, 0xe0, 0x29, 0x77, 0xfe, 0x3e, 0xdf, 0x77, 0x77, 0x5f, 0x0c, 0x66, 0x20,
	0xc8, 0x9c, 0x92, 0xa4, 0xb3, 0x3a, 0xea, 0xc4, 0xc4, 0xbf, 0x23, 0x01, 0xb5, 0x63, 0xc1, 0x53,
	0x8e, 0xa0, 0x44, 0xec, 0xd5, 0x91, 0xf5, 0x5b, 0x81, 0xbd, 0x01, 0x4b, 0x52, 0xc1, 0x6e, 0xb3,
	0x94, 0xf1, 0x08, 0x3d, 0x85, 0xba, 0x1f, 0x53, 0x2f, 0x13, 0xcc, 0x54, 0x5a, 0x4a, 0xbb, 0x89,
	0x6b, 0x7e, 0x4c, 0xa7, 0x82, 0xa1, 0x8f, 0xb0, 0x47, 0x84, 0xbf, 0x60, 0x29, 0xf5, 0xd3, 0x4c,
	0x50, 0x53, 0x6d, 0x29, 0xed, 0x83, 0xae, 0x69, 0xdf, 0x17, 0xb3, 0x7b, 0x15, 0x1c, 0x6f, 0xb1,
	0xd1, 0x7b, 0x38, 0x08, 0x49, 0x4a, 0x93, 0xd4, 0x5b, 0x51, 0x91, 0x30, 0x1e, 0x99, 0x5a, 0x4b,
	0x69, 0xeb, 0xdd, 0x27, 0xd5, 0xfb, 0xd7, 0x05, 0x84, 0xf7, 0x0b, 0x6a, 0x99, 0xa2, 0x97, 0x00,
	0x4b, 0xc2, 0xa2, 0x94, 0xb0, 0x88, 0x0a, 0x73, 0x47, 0x76, 0x55, 0x39, 0x41, 0x06, 0x68, 0x99,
	0x08, 0xcd, 0x5d, 0x09, 0xe4, 0x21, 0x6a, 0x81, 0x3e, 0xa3, 0x89, 0x2f, 0x58, 0x9c, 0xcf, 0x64,
	0xd6, 0x24, 0x52, 0x3d, 0xb2, 0xe6, 0xd0, 0x38, 0xe7, 0x3e, 0xf9, 0xff, 0xc8, 0x6f, 0xa0, 0xbe,
	0xee, 0x56, 0xfd, 0x77, 0xb7, 0x6b, 0x0e, 0x42, 0xb0, 0x13, 0x93, 0x74, 0x21, 0x27, 0x6b, 0x62,
	0x19, 0x5b, 0x1e, 0xe8, 0x93, 0x62, 0xf9, 0x2e, 0x4f, 0x69, 0x4e, 0x89, 0xc8, 0x92, 0x96, 0x3a,
	0x32, 0xce, 0x17, 0x3b, 0xab, 0x38, 0x60, 0x42, 0x4b, 0x6b, 0xeb, 0xdb, 0x8b, 0xad, 0x3a, 0x84,
	0xb7, 0xd8, 0xd6, 0x77, 0x78, 0x5c, 0x0a, 0x8c, 0x7d, 0x3f, 0x13, 0x82, 0x46, 0xfe, 0xc3, 0x32,
	0x6f, 0xa1, 0x11, 0x96, 0x13, 0x9b, 0xaa, 0x94, 0x38, 0xac, 0x4a, 0xac, 0xb7, 0x81, 0x37, 0x2c,
	0xeb, 0x97, 0x02, 0xf5, 0xb5, 0x07, 0x87, 0xb0, 0x4b, 0x63, 0xee, 0x2f, 0x64, 0xc9, 0x5d, 0x5c,
	0x24, 0x1b, 0x1d, 0xb5, 0xa2, 0xf3, 0x0c, 0x1a, 0x82, 0xae, 0xd8, 0xc6, 0xe3, 0x26, 0xde, 0xe4,
	0xe8, 0x18, 0x76, 0xee, 0x58, 0x34, 0x93, 0x1e, 0x1e, 0x74, 0x5f, 0x3d, 0xb0, 0xcd, 0xf5, 0xf7,
	0x8c, 0x45, 0x33, 0x2c, 0xc9, 0xe8, 0x39, 0x34, 0xe7, 0x59, 0x18, 0x7a, 0x52, 0xa9, 0x30, 0xb9,
	0x91, 0x1f, 0xb8, 0x64, 0x49, 0xad, 0x0b, 0xd0, 0x2b, 0x37, 0xd0, 0x0b, 0x30, 0xaf, 0x87, 0xf8,
	0xd2, 0x19, 0xbb, 0xde, 0x99, 0xe3, 0x0e, 0xbc, 0xa9, 0x7b, 0x39, 0x19, 0xf6, 0x9d, 0x2f, 0xce,
	0x70, 0x60, 0x3c, 0x42, 0x00, 0x35, 0x77, 0x8c, 0x47, 0xbd, 0x73, 0x43, 0x41, 0x3a, 0xd4, 0x47,
	0x8e, 0xeb, 0x8c, 0xa6, 0x23, 0x43, 0x95, 0x49, 0xef, 0x46, 0x26, 0xda, 0xeb, 0x4f, 0xb0, 0x57,
	0xfd, 0x91, 0xf3, 0x9a, 0x3d, 0xdc, 0xff, 0xe6, 0x5c, 0x0d, 0xfb, 0x57, 0x53, 0x3c, 0xfc, 0xab,
	0x66, 0x1d, 0xb4, 0x9b, 0xd3, 0x13, 0x43, 0x91, 0xc1, 0xc9, 0x3b, 0x43, 0xfd, 0x7c, 0x01, 0xfb,
	0x8c, 0x57, 0x46, 0x9b, 0x28, 0x3f, 0x4e, 0x03, 0xce, 0x83, 0x90, 0xda, 0x01, 0x0f, 0x49, 0x14,
	0xd8, 0x5c, 0x04, 0x9d, 0x80, 0x46, 0xf2, 0x35, 0x76, 0x0a, 0x88, 0xc4, 0x2c, 0xe9, 0xdc, 0xbf,
	0xd8, 0x0f, 0x65, 0xf8, 0x53, 0xd5, 0xbe, 0xe2, 0xde, 0x6d, 0x4d, 0x52, 0x8f, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xd8, 0x11, 0xfa, 0xb8, 0xd4, 0x03, 0x00, 0x00,
}
