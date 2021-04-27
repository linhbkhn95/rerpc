// Code generated by protoc-gen-go. DO NOT EDIT.
// source: status.proto

package status

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// The `Status` type defines a logical error model that is suitable for
// different programming environments, including REST APIs and RPC APIs. It is
// used by [gRPC](https://github.com/grpc). The error model is designed to be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message, and error details. The error code should be an enum value of
// [google.rpc.Code][google.rpc.Code], but it may accept additional error codes
// if needed.  The error message should be a developer-facing English message
// that helps developers *understand* and *resolve* the error. If a localized
// user-facing error message is needed, put the localized message in the error
// details or localize it in the client. The optional error details may contain
// arbitrary information about the error. There is a predefined set of error
// detail types in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error model, but it
// is not necessarily the actual wire format. When the `Status` message is
// exposed in different client libraries and different wire protocols, it can be
// mapped differently. For example, it will likely be mapped to some exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety of
// environments, either with or without APIs, to provide a
// consistent developer experience across different environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the client,
//     it may embed the `Status` in the normal response to indicate the partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch response, the
//     `Status` message should be used directly inside batch response, one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous operation
//     results in its response, the status of those operations should be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message `Status` could
//     be used directly after any stripping needed for security/privacy reasons.
type Status struct {
	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// A developer-facing error message, which should be in English. Any
	// user-facing error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	// by the client.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// A list of messages that carry the error details.  There is a common set of
	// message types for APIs to use.
	Details              []*any.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfe4fce6682daf5b, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Status) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "google.rpc.Status")
}

func init() { proto.RegisterFile("status.proto", fileDescriptor_dfe4fce6682daf5b) }

var fileDescriptor_dfe4fce6682daf5b = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x49, 0x2c,
	0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0xcf, 0xcf, 0x4f, 0xcf, 0x49,
	0xd5, 0x2b, 0x2a, 0x48, 0x96, 0x92, 0x84, 0xb0, 0xf5, 0xc1, 0x32, 0x49, 0xa5, 0x69, 0xfa, 0x89,
	0x79, 0x95, 0x10, 0x65, 0x4a, 0x69, 0x5c, 0x6c, 0xc1, 0x60, 0x6d, 0x42, 0x42, 0x5c, 0x2c, 0xc9,
	0xf9, 0x29, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x60, 0xb6, 0x90, 0x04, 0x17, 0x7b,
	0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b,
	0xa4, 0xc7, 0xc5, 0x9e, 0x92, 0x5a, 0x92, 0x98, 0x99, 0x53, 0x2c, 0xc1, 0xac, 0xc0, 0xac, 0xc1,
	0x6d, 0x24, 0xa2, 0x07, 0xb5, 0x10, 0x66, 0x89, 0x9e, 0x63, 0x5e, 0x65, 0x10, 0x4c, 0x91, 0x53,
	0x1c, 0x17, 0x5f, 0x72, 0x7e, 0xae, 0x1e, 0xc2, 0x51, 0x4e, 0xdc, 0x10, 0x7b, 0x03, 0x40, 0xca,
	0x03, 0x18, 0xa3, 0xcc, 0xa1, 0x52, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9,
	0xfa, 0xe9, 0xa9, 0x79, 0x60, 0xc3, 0xf4, 0x21, 0x52, 0x89, 0x05, 0x99, 0xc5, 0xfa, 0x45, 0x05,
	0xc9, 0xfa, 0x10, 0x7f, 0x5a, 0x43, 0xa8, 0x45, 0x4c, 0xcc, 0x41, 0x01, 0xce, 0x49, 0x6c, 0x60,
	0x95, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0x27, 0xdb, 0x17, 0x05, 0x01, 0x00, 0x00,
}
