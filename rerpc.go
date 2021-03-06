package rerpc

import (
	"fmt"
	"runtime"
)

// Version is the semantic version of the reRPC module.
const Version = "0.0.1"

// MaxHeaderBytes is 8KiB, gRPC's recommended maximum header size. To enforce
// this limit, set MaxHeaderBytes on your http.Server.
const MaxHeaderBytes = 1024 * 8

// ReRPC's supported HTTP Content-Types. Servers decide whether to use the gRPC
// or Twirp protocol based on the request's Content-Type. See the protocol
// documentation at https://rerpc.github.io for more information.
const (
	TypeDefaultGRPC = "application/grpc"
	TypeProtoGRPC   = "application/grpc+proto"
	TypeProtoTwirp  = "application/protobuf"
	TypeJSON        = "application/json"
)

// ReRPC's supported compression methods.
const (
	CompressionIdentity = "identity"
	CompressionGzip     = "gzip"
)

// These constants are used in compile-time handshakes with reRPC's generated
// code.
const (
	SupportsCodeGenV0 = iota
)

var userAgent = fmt.Sprintf("grpc-go-rerpc/%s (%s)", Version, runtime.Version())

// UserAgent describes reRPC to servers, following the convention outlined in
// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#user-agents.
// The output will resemble "grpc-go-rerpc/1.2.3 (go1.16.6)".
func UserAgent() string {
	return userAgent
}
