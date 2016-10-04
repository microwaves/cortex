// Code generated by protoc-gen-go.
// source: prism.proto
// DO NOT EDIT!

/*
Package prism is a generated protocol buffer package.

It is generated from these files:
	prism.proto

It has these top-level messages:
	Sample
	LabelPair
	TimeSeries
	LabelMatcher
	ReadRequest
	ReadResponse
	LabelValuesRequest
	LabelValuesResponse
*/
package prism

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

type MatchType int32

const (
	MatchType_EQUAL          MatchType = 0
	MatchType_NOT_EQUAL      MatchType = 1
	MatchType_REGEX_MATCH    MatchType = 2
	MatchType_REGEX_NO_MATCH MatchType = 3
)

var MatchType_name = map[int32]string{
	0: "EQUAL",
	1: "NOT_EQUAL",
	2: "REGEX_MATCH",
	3: "REGEX_NO_MATCH",
}
var MatchType_value = map[string]int32{
	"EQUAL":          0,
	"NOT_EQUAL":      1,
	"REGEX_MATCH":    2,
	"REGEX_NO_MATCH": 3,
}

func (x MatchType) String() string {
	return proto.EnumName(MatchType_name, int32(x))
}
func (MatchType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Sample struct {
	Value       float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	TimestampMs int64   `protobuf:"varint,2,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
}

func (m *Sample) Reset()                    { *m = Sample{} }
func (m *Sample) String() string            { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()               {}
func (*Sample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LabelPair struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *LabelPair) Reset()                    { *m = LabelPair{} }
func (m *LabelPair) String() string            { return proto.CompactTextString(m) }
func (*LabelPair) ProtoMessage()               {}
func (*LabelPair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type TimeSeries struct {
	Labels []*LabelPair `protobuf:"bytes,1,rep,name=labels" json:"labels,omitempty"`
	// Sorted by time, oldest sample first.
	Samples []*Sample `protobuf:"bytes,2,rep,name=samples" json:"samples,omitempty"`
}

func (m *TimeSeries) Reset()                    { *m = TimeSeries{} }
func (m *TimeSeries) String() string            { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()               {}
func (*TimeSeries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TimeSeries) GetLabels() []*LabelPair {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TimeSeries) GetSamples() []*Sample {
	if m != nil {
		return m.Samples
	}
	return nil
}

type LabelMatcher struct {
	Type  MatchType `protobuf:"varint,1,opt,name=type,enum=prism.MatchType" json:"type,omitempty"`
	Name  string    `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Value string    `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *LabelMatcher) Reset()                    { *m = LabelMatcher{} }
func (m *LabelMatcher) String() string            { return proto.CompactTextString(m) }
func (*LabelMatcher) ProtoMessage()               {}
func (*LabelMatcher) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ReadRequest struct {
	StartTimestampMs int64           `protobuf:"varint,1,opt,name=start_timestamp_ms,json=startTimestampMs" json:"start_timestamp_ms,omitempty"`
	EndTimestampMs   int64           `protobuf:"varint,2,opt,name=end_timestamp_ms,json=endTimestampMs" json:"end_timestamp_ms,omitempty"`
	Matchers         []*LabelMatcher `protobuf:"bytes,3,rep,name=matchers" json:"matchers,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadRequest) GetMatchers() []*LabelMatcher {
	if m != nil {
		return m.Matchers
	}
	return nil
}

type ReadResponse struct {
	Timeseries []*TimeSeries `protobuf:"bytes,1,rep,name=timeseries" json:"timeseries,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReadResponse) GetTimeseries() []*TimeSeries {
	if m != nil {
		return m.Timeseries
	}
	return nil
}

type LabelValuesRequest struct {
	LabelName string `protobuf:"bytes,1,opt,name=label_name,json=labelName" json:"label_name,omitempty"`
}

func (m *LabelValuesRequest) Reset()                    { *m = LabelValuesRequest{} }
func (m *LabelValuesRequest) String() string            { return proto.CompactTextString(m) }
func (*LabelValuesRequest) ProtoMessage()               {}
func (*LabelValuesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type LabelValuesResponse struct {
	LabelValues []string `protobuf:"bytes,1,rep,name=label_values,json=labelValues" json:"label_values,omitempty"`
}

func (m *LabelValuesResponse) Reset()                    { *m = LabelValuesResponse{} }
func (m *LabelValuesResponse) String() string            { return proto.CompactTextString(m) }
func (*LabelValuesResponse) ProtoMessage()               {}
func (*LabelValuesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*Sample)(nil), "prism.Sample")
	proto.RegisterType((*LabelPair)(nil), "prism.LabelPair")
	proto.RegisterType((*TimeSeries)(nil), "prism.TimeSeries")
	proto.RegisterType((*LabelMatcher)(nil), "prism.LabelMatcher")
	proto.RegisterType((*ReadRequest)(nil), "prism.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "prism.ReadResponse")
	proto.RegisterType((*LabelValuesRequest)(nil), "prism.LabelValuesRequest")
	proto.RegisterType((*LabelValuesResponse)(nil), "prism.LabelValuesResponse")
	proto.RegisterEnum("prism.MatchType", MatchType_name, MatchType_value)
}

func init() { proto.RegisterFile("prism.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x5d, 0x4b, 0xe3, 0x50,
	0x10, 0xdd, 0x24, 0x6d, 0x77, 0x33, 0x69, 0xbb, 0xd9, 0xdb, 0x7d, 0xe8, 0x8b, 0xa0, 0x41, 0xb0,
	0x88, 0x54, 0xb4, 0x08, 0xbe, 0x16, 0x29, 0x8a, 0xf4, 0x43, 0x6f, 0xa3, 0xf8, 0x64, 0x48, 0xed,
	0x05, 0x03, 0x49, 0x1b, 0x73, 0x53, 0xc1, 0x3f, 0xe2, 0xef, 0x75, 0x32, 0x37, 0x4d, 0x52, 0xe8,
	0x5b, 0xe6, 0xcc, 0xcc, 0xb9, 0xe7, 0x9c, 0x09, 0x58, 0x71, 0x12, 0xc8, 0xa8, 0x1f, 0x27, 0xeb,
	0x74, 0xcd, 0xea, 0x54, 0x38, 0x43, 0x68, 0xcc, 0xfd, 0x28, 0x0e, 0x05, 0xfb, 0x0f, 0xf5, 0x4f,
	0x3f, 0xdc, 0x88, 0xae, 0x76, 0xa8, 0xf5, 0x34, 0xae, 0x0a, 0x76, 0x04, 0xcd, 0x34, 0x88, 0x84,
	0x4c, 0x71, 0xc8, 0x8b, 0x64, 0x57, 0xc7, 0xa6, 0xc1, 0xad, 0x02, 0x9b, 0x48, 0xe7, 0x0a, 0xcc,
	0xb1, 0xbf, 0x10, 0xe1, 0x83, 0x1f, 0x24, 0x8c, 0x41, 0x6d, 0xe5, 0x47, 0x8a, 0xc4, 0xe4, 0xf4,
	0x5d, 0x32, 0xeb, 0x04, 0xaa, 0xc2, 0xf1, 0x00, 0x5c, 0x64, 0x99, 0x8b, 0x24, 0x10, 0x92, 0xf5,
	0xa0, 0x11, 0x66, 0x24, 0x12, 0x37, 0x8d, 0x9e, 0x75, 0x69, 0xf7, 0x95, 0xd8, 0x82, 0x99, 0xe7,
	0x7d, 0x76, 0x02, 0xbf, 0x25, 0x29, 0xce, 0xc4, 0x64, 0xa3, 0xad, 0x7c, 0x54, 0xf9, 0xe0, 0xdb,
	0xae, 0xf3, 0x0a, 0x4d, 0xda, 0x9e, 0xf8, 0xe9, 0xdb, 0xbb, 0x48, 0xd8, 0x31, 0xd4, 0xd2, 0xaf,
	0x58, 0x49, 0x6b, 0x17, 0x0f, 0x50, 0xd7, 0x45, 0x9c, 0x53, 0xb7, 0x30, 0xa0, 0xef, 0x33, 0x60,
	0x54, 0x0d, 0x7c, 0x6b, 0x60, 0x71, 0xe1, 0x2f, 0xb9, 0xf8, 0xd8, 0x60, 0x18, 0xec, 0x0c, 0x18,
	0x46, 0x92, 0xa4, 0xde, 0x4e, 0x60, 0x1a, 0x05, 0x66, 0x53, 0xc7, 0x2d, 0x53, 0x43, 0xc3, 0xb6,
	0x58, 0x2d, 0xbd, 0x3d, 0xe1, 0xb6, 0x11, 0xaf, 0x4e, 0x9e, 0xc3, 0x9f, 0x48, 0x59, 0x90, 0x28,
	0x20, 0x73, 0xdc, 0xa9, 0x86, 0x93, 0xdb, 0xe3, 0xc5, 0x10, 0xde, 0xb4, 0xa9, 0x74, 0xc9, 0x78,
	0xbd, 0x92, 0x82, 0x5d, 0x00, 0xd0, 0x33, 0x94, 0x74, 0x9e, 0xef, 0xbf, 0x9c, 0xa2, 0x3c, 0x01,
	0xaf, 0x0c, 0x39, 0x03, 0x60, 0x44, 0xfe, 0x9c, 0x39, 0x95, 0x5b, 0x87, 0x07, 0x00, 0x74, 0x04,
	0xaf, 0x72, 0x62, 0x93, 0x90, 0x29, 0x02, 0xce, 0x35, 0x74, 0x76, 0x96, 0xf2, 0xe7, 0xf1, 0x17,
	0x52, 0x5b, 0x14, 0x9b, 0x12, 0x60, 0x72, 0x2b, 0x2c, 0x47, 0x4f, 0xef, 0xc1, 0x2c, 0xee, 0xc0,
	0x4c, 0xa8, 0x8f, 0x1e, 0x9f, 0x86, 0x63, 0xfb, 0x17, 0x6b, 0x81, 0x39, 0x9d, 0xb9, 0x9e, 0x2a,
	0x35, 0xf6, 0x17, 0x03, 0x1f, 0xdd, 0x8e, 0x5e, 0xbc, 0xc9, 0xd0, 0xbd, 0xb9, 0xb3, 0x75, 0x3c,
	0x56, 0x5b, 0x01, 0xd3, 0x59, 0x8e, 0x19, 0x8b, 0x06, 0xfd, 0xdf, 0x83, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x31, 0xd4, 0x4b, 0x4c, 0xee, 0x02, 0x00, 0x00,
}
