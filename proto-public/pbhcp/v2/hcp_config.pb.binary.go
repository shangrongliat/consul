// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: pbhcp/v2/hcp_config.proto

package hcpv2

import (
	"google.golang.org/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *HCPConfig) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *HCPConfig) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}
