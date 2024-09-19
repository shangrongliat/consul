// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: pbhcp/v2/link.proto

package hcpv2

import (
	_ "github.com/hashicorp/consul/proto-public/pbresource"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccessLevel int32

const (
	AccessLevel_ACCESS_LEVEL_UNSPECIFIED       AccessLevel = 0
	AccessLevel_ACCESS_LEVEL_GLOBAL_READ_WRITE AccessLevel = 1
	AccessLevel_ACCESS_LEVEL_GLOBAL_READ_ONLY  AccessLevel = 2
)

// Enum value maps for AccessLevel.
var (
	AccessLevel_name = map[int32]string{
		0: "ACCESS_LEVEL_UNSPECIFIED",
		1: "ACCESS_LEVEL_GLOBAL_READ_WRITE",
		2: "ACCESS_LEVEL_GLOBAL_READ_ONLY",
	}
	AccessLevel_value = map[string]int32{
		"ACCESS_LEVEL_UNSPECIFIED":       0,
		"ACCESS_LEVEL_GLOBAL_READ_WRITE": 1,
		"ACCESS_LEVEL_GLOBAL_READ_ONLY":  2,
	}
)

func (x AccessLevel) Enum() *AccessLevel {
	p := new(AccessLevel)
	*p = x
	return p
}

func (x AccessLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccessLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_pbhcp_v2_link_proto_enumTypes[0].Descriptor()
}

func (AccessLevel) Type() protoreflect.EnumType {
	return &file_pbhcp_v2_link_proto_enumTypes[0]
}

func (x AccessLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccessLevel.Descriptor instead.
func (AccessLevel) EnumDescriptor() ([]byte, []int) {
	return file_pbhcp_v2_link_proto_rawDescGZIP(), []int{0}
}

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId    string      `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ClientId      string      `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret  string      `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	HcpClusterUrl string      `protobuf:"bytes,4,opt,name=hcp_cluster_url,json=hcpClusterUrl,proto3" json:"hcp_cluster_url,omitempty"`
	AccessLevel   AccessLevel `protobuf:"varint,5,opt,name=access_level,json=accessLevel,proto3,enum=hashicorp.consul.hcp.v2.AccessLevel" json:"access_level,omitempty"`
	HcpConfig     *HCPConfig  `protobuf:"bytes,6,opt,name=hcp_config,json=hcpConfig,proto3" json:"hcp_config,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbhcp_v2_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_pbhcp_v2_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_pbhcp_v2_link_proto_rawDescGZIP(), []int{0}
}

func (x *Link) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *Link) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Link) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *Link) GetHcpClusterUrl() string {
	if x != nil {
		return x.HcpClusterUrl
	}
	return ""
}

func (x *Link) GetAccessLevel() AccessLevel {
	if x != nil {
		return x.AccessLevel
	}
	return AccessLevel_ACCESS_LEVEL_UNSPECIFIED
}

func (x *Link) GetHcpConfig() *HCPConfig {
	if x != nil {
		return x.HcpConfig
	}
	return nil
}

var File_pbhcp_v2_link_proto protoreflect.FileDescriptor

var file_pbhcp_v2_link_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x68, 0x63, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x68, 0x63, 0x70, 0x2e, 0x76, 0x32, 0x1a, 0x19,
	0x70, 0x62, 0x68, 0x63, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x68, 0x63, 0x70, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x62, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x02, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x68, 0x63, 0x70, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x63,
	0x70, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x47, 0x0a, 0x0c, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x68, 0x63, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x41, 0x0a, 0x0a, 0x68, 0x63, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x68, 0x63, 0x70, 0x2e,
	0x76, 0x32, 0x2e, 0x48, 0x43, 0x50, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x68, 0x63,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x06, 0xa2, 0x93, 0x04, 0x02, 0x08, 0x01, 0x2a,
	0x72, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c,
	0x0a, 0x18, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x47, 0x4c, 0x4f,
	0x42, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x01,
	0x12, 0x21, 0x0a, 0x1d, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x5f, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x4f, 0x4e, 0x4c,
	0x59, 0x10, 0x02, 0x42, 0xe0, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x68, 0x63, 0x70,
	0x2e, 0x76, 0x32, 0x42, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x68, 0x63, 0x70,
	0x2f, 0x76, 0x32, 0x3b, 0x68, 0x63, 0x70, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x48, 0xaa,
	0x02, 0x17, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x48, 0x63, 0x70, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x17, 0x48, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x48, 0x63, 0x70,
	0x5c, 0x56, 0x32, 0xe2, 0x02, 0x23, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x48, 0x63, 0x70, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x48, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x48,
	0x63, 0x70, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbhcp_v2_link_proto_rawDescOnce sync.Once
	file_pbhcp_v2_link_proto_rawDescData = file_pbhcp_v2_link_proto_rawDesc
)

func file_pbhcp_v2_link_proto_rawDescGZIP() []byte {
	file_pbhcp_v2_link_proto_rawDescOnce.Do(func() {
		file_pbhcp_v2_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbhcp_v2_link_proto_rawDescData)
	})
	return file_pbhcp_v2_link_proto_rawDescData
}

var file_pbhcp_v2_link_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pbhcp_v2_link_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pbhcp_v2_link_proto_goTypes = []interface{}{
	(AccessLevel)(0),  // 0: hashicorp.consul.hcp.v2.AccessLevel
	(*Link)(nil),      // 1: hashicorp.consul.hcp.v2.Link
	(*HCPConfig)(nil), // 2: hashicorp.consul.hcp.v2.HCPConfig
}
var file_pbhcp_v2_link_proto_depIdxs = []int32{
	0, // 0: hashicorp.consul.hcp.v2.Link.access_level:type_name -> hashicorp.consul.hcp.v2.AccessLevel
	2, // 1: hashicorp.consul.hcp.v2.Link.hcp_config:type_name -> hashicorp.consul.hcp.v2.HCPConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pbhcp_v2_link_proto_init() }
func file_pbhcp_v2_link_proto_init() {
	if File_pbhcp_v2_link_proto != nil {
		return
	}
	file_pbhcp_v2_hcp_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pbhcp_v2_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbhcp_v2_link_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbhcp_v2_link_proto_goTypes,
		DependencyIndexes: file_pbhcp_v2_link_proto_depIdxs,
		EnumInfos:         file_pbhcp_v2_link_proto_enumTypes,
		MessageInfos:      file_pbhcp_v2_link_proto_msgTypes,
	}.Build()
	File_pbhcp_v2_link_proto = out.File
	file_pbhcp_v2_link_proto_rawDesc = nil
	file_pbhcp_v2_link_proto_goTypes = nil
	file_pbhcp_v2_link_proto_depIdxs = nil
}
