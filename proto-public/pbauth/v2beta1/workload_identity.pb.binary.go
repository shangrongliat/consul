// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: pbauth/v2beta1/workload_identity.proto

package authv2beta1

import (
	"google.golang.org/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *WorkloadIdentity) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *WorkloadIdentity) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}
