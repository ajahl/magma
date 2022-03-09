//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.10.0
// source: orc8r/protos/mconfig.proto

package protos

import (
	any "github.com/golang/protobuf/ptypes/any"
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

// --------------------------------------------------------------------------
// GatewayConfigs structure is a container for all Access Gateway's (AG) Cloud
// Managed Configs (CMC). Each and every field of GatewayConfigs represents
// one AG service config
// --------------------------------------------------------------------------
// NOTE: a service config field name (control_proxy, enodebd, etc.) must match
//       the corresponding gateway service's name exactly
type GatewayConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigsByKey map[string]*any.Any     `protobuf:"bytes,10,rep,name=configs_by_key,json=configsByKey,proto3" json:"configs_by_key,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Metadata     *GatewayConfigsMetadata `protobuf:"bytes,11,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *GatewayConfigs) Reset() {
	*x = GatewayConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_mconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayConfigs) ProtoMessage() {}

func (x *GatewayConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_mconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayConfigs.ProtoReflect.Descriptor instead.
func (*GatewayConfigs) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_mconfig_proto_rawDescGZIP(), []int{0}
}

func (x *GatewayConfigs) GetConfigsByKey() map[string]*any.Any {
	if x != nil {
		return x.ConfigsByKey
	}
	return nil
}

func (x *GatewayConfigs) GetMetadata() *GatewayConfigsMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Deterministic hash of a serialized GatewayConfigs proto
type GatewayConfigsDigest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hexadecimal MD5 hash of the UTF-8-encoded stringified full mconfigs
	Md5HexDigest string `protobuf:"bytes,1,opt,name=md5_hex_digest,json=md5HexDigest,proto3" json:"md5_hex_digest,omitempty"`
}

func (x *GatewayConfigsDigest) Reset() {
	*x = GatewayConfigsDigest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_mconfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayConfigsDigest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayConfigsDigest) ProtoMessage() {}

func (x *GatewayConfigsDigest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_mconfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayConfigsDigest.ProtoReflect.Descriptor instead.
func (*GatewayConfigsDigest) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_mconfig_proto_rawDescGZIP(), []int{1}
}

func (x *GatewayConfigsDigest) GetMd5HexDigest() string {
	if x != nil {
		return x.Md5HexDigest
	}
	return ""
}

// Metadata about the configs.
type GatewayConfigsMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix timestamp of Cloud at the time of config generation.
	CreatedAt uint64                `protobuf:"varint,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Digest    *GatewayConfigsDigest `protobuf:"bytes,12,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *GatewayConfigsMetadata) Reset() {
	*x = GatewayConfigsMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_mconfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayConfigsMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayConfigsMetadata) ProtoMessage() {}

func (x *GatewayConfigsMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_mconfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayConfigsMetadata.ProtoReflect.Descriptor instead.
func (*GatewayConfigsMetadata) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_mconfig_proto_rawDescGZIP(), []int{2}
}

func (x *GatewayConfigsMetadata) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *GatewayConfigsMetadata) GetDigest() *GatewayConfigsDigest {
	if x != nil {
		return x.Digest
	}
	return nil
}

// Wraps a gateway config and a stream offset that the config was computed
// from
type OffsetGatewayConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configs *GatewayConfigs `protobuf:"bytes,1,opt,name=configs,proto3" json:"configs,omitempty"`
	Offset  int64           `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *OffsetGatewayConfigs) Reset() {
	*x = OffsetGatewayConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_mconfig_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OffsetGatewayConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OffsetGatewayConfigs) ProtoMessage() {}

func (x *OffsetGatewayConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_mconfig_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OffsetGatewayConfigs.ProtoReflect.Descriptor instead.
func (*OffsetGatewayConfigs) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_mconfig_proto_rawDescGZIP(), []int{3}
}

func (x *OffsetGatewayConfigs) GetConfigs() *GatewayConfigs {
	if x != nil {
		return x.Configs
	}
	return nil
}

func (x *OffsetGatewayConfigs) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Stream request passed as extra args to the streaming mconfig streamer policy.
// Contains a single field, the offset of the mconfig currently stored on
// the device.
type MconfigStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *MconfigStreamRequest) Reset() {
	*x = MconfigStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_mconfig_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MconfigStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MconfigStreamRequest) ProtoMessage() {}

func (x *MconfigStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_mconfig_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MconfigStreamRequest.ProtoReflect.Descriptor instead.
func (*MconfigStreamRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_mconfig_proto_rawDescGZIP(), []int{4}
}

func (x *MconfigStreamRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_orc8r_protos_mconfig_proto protoreflect.FileDescriptor

var file_orc8r_protos_mconfig_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61,
	0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x0e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x53, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x3f, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x55, 0x0a,
	0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x3c, 0x0a, 0x14, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e,
	0x6d, 0x64, 0x35, 0x5f, 0x68, 0x65, 0x78, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x64, 0x35, 0x48, 0x65, 0x78, 0x44, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x22, 0x72, 0x0a, 0x16, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x61,
	0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x22, 0x65, 0x0a, 0x14, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x35,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x2e, 0x0a,
	0x14, 0x4d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x1b, 0x5a,
	0x19, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x6c, 0x69, 0x62,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_orc8r_protos_mconfig_proto_rawDescOnce sync.Once
	file_orc8r_protos_mconfig_proto_rawDescData = file_orc8r_protos_mconfig_proto_rawDesc
)

func file_orc8r_protos_mconfig_proto_rawDescGZIP() []byte {
	file_orc8r_protos_mconfig_proto_rawDescOnce.Do(func() {
		file_orc8r_protos_mconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_orc8r_protos_mconfig_proto_rawDescData)
	})
	return file_orc8r_protos_mconfig_proto_rawDescData
}

var file_orc8r_protos_mconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_orc8r_protos_mconfig_proto_goTypes = []interface{}{
	(*GatewayConfigs)(nil),         // 0: magma.orc8r.GatewayConfigs
	(*GatewayConfigsDigest)(nil),   // 1: magma.orc8r.GatewayConfigsDigest
	(*GatewayConfigsMetadata)(nil), // 2: magma.orc8r.GatewayConfigsMetadata
	(*OffsetGatewayConfigs)(nil),   // 3: magma.orc8r.OffsetGatewayConfigs
	(*MconfigStreamRequest)(nil),   // 4: magma.orc8r.MconfigStreamRequest
	nil,                            // 5: magma.orc8r.GatewayConfigs.ConfigsByKeyEntry
	(*any.Any)(nil),                // 6: google.protobuf.Any
}
var file_orc8r_protos_mconfig_proto_depIdxs = []int32{
	5, // 0: magma.orc8r.GatewayConfigs.configs_by_key:type_name -> magma.orc8r.GatewayConfigs.ConfigsByKeyEntry
	2, // 1: magma.orc8r.GatewayConfigs.metadata:type_name -> magma.orc8r.GatewayConfigsMetadata
	1, // 2: magma.orc8r.GatewayConfigsMetadata.digest:type_name -> magma.orc8r.GatewayConfigsDigest
	0, // 3: magma.orc8r.OffsetGatewayConfigs.configs:type_name -> magma.orc8r.GatewayConfigs
	6, // 4: magma.orc8r.GatewayConfigs.ConfigsByKeyEntry.value:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_orc8r_protos_mconfig_proto_init() }
func file_orc8r_protos_mconfig_proto_init() {
	if File_orc8r_protos_mconfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orc8r_protos_mconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayConfigs); i {
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
		file_orc8r_protos_mconfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayConfigsDigest); i {
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
		file_orc8r_protos_mconfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayConfigsMetadata); i {
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
		file_orc8r_protos_mconfig_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OffsetGatewayConfigs); i {
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
		file_orc8r_protos_mconfig_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MconfigStreamRequest); i {
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
			RawDescriptor: file_orc8r_protos_mconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orc8r_protos_mconfig_proto_goTypes,
		DependencyIndexes: file_orc8r_protos_mconfig_proto_depIdxs,
		MessageInfos:      file_orc8r_protos_mconfig_proto_msgTypes,
	}.Build()
	File_orc8r_protos_mconfig_proto = out.File
	file_orc8r_protos_mconfig_proto_rawDesc = nil
	file_orc8r_protos_mconfig_proto_goTypes = nil
	file_orc8r_protos_mconfig_proto_depIdxs = nil
}
