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
// source: orc8r/protos/digest.proto

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

// DigestTree contains the full set of digest information for a particular network.
// DigestTree is similar to a depth=2 Merkle tree.
type DigestTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// root_digest is the amalgum of all leaf digests.
	RootDigest *Digest `protobuf:"bytes,1,opt,name=root_digest,json=rootDigest,proto3" json:"root_digest,omitempty"`
	// leaf_digests contains per-object digests, along with the object IDs, sorted by ID.
	LeafDigests []*LeafDigest `protobuf:"bytes,2,rep,name=leaf_digests,json=leafDigests,proto3" json:"leaf_digests,omitempty"`
}

func (x *DigestTree) Reset() {
	*x = DigestTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_digest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DigestTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DigestTree) ProtoMessage() {}

func (x *DigestTree) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_digest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DigestTree.ProtoReflect.Descriptor instead.
func (*DigestTree) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_digest_proto_rawDescGZIP(), []int{0}
}

func (x *DigestTree) GetRootDigest() *Digest {
	if x != nil {
		return x.RootDigest
	}
	return nil
}

func (x *DigestTree) GetLeafDigests() []*LeafDigest {
	if x != nil {
		return x.LeafDigests
	}
	return nil
}

// Digest contains the digest (hash) of some object.
type Digest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// md5_base64_digest is a base64-encoded MD5 digest.
	Md5Base64Digest string `protobuf:"bytes,1,opt,name=md5_base64_digest,json=md5Base64Digest,proto3" json:"md5_base64_digest,omitempty"`
}

func (x *Digest) Reset() {
	*x = Digest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_digest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Digest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Digest) ProtoMessage() {}

func (x *Digest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_digest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Digest.ProtoReflect.Descriptor instead.
func (*Digest) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_digest_proto_rawDescGZIP(), []int{1}
}

func (x *Digest) GetMd5Base64Digest() string {
	if x != nil {
		return x.Md5Base64Digest
	}
	return ""
}

type LeafDigest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the network-wide unique identifier of the object.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// digest is the deterministic digest of the object.
	Digest *Digest `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *LeafDigest) Reset() {
	*x = LeafDigest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_digest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeafDigest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeafDigest) ProtoMessage() {}

func (x *LeafDigest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_digest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeafDigest.ProtoReflect.Descriptor instead.
func (*LeafDigest) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_digest_proto_rawDescGZIP(), []int{2}
}

func (x *LeafDigest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LeafDigest) GetDigest() *Digest {
	if x != nil {
		return x.Digest
	}
	return nil
}

// LeafDigests is used to encapsulate a list of leaf digests exclusively for
// serialization en masse.
// NOTE: In a proto message used by gRPC endpoints (e.g. DigestTree), the leaf
// digests should still be represented directly as a list for simplicity.
type LeafDigests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digests []*LeafDigest `protobuf:"bytes,1,rep,name=digests,proto3" json:"digests,omitempty"`
}

func (x *LeafDigests) Reset() {
	*x = LeafDigests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_digest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeafDigests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeafDigests) ProtoMessage() {}

func (x *LeafDigests) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_digest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeafDigests.ProtoReflect.Descriptor instead.
func (*LeafDigests) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_digest_proto_rawDescGZIP(), []int{3}
}

func (x *LeafDigests) GetDigests() []*LeafDigest {
	if x != nil {
		return x.Digests
	}
	return nil
}

// Changeset contains the set of differences between two lists of objects.
type Changeset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// to_renew contains all objects which were added or updated, sorted by ID.
	ToRenew []*any.Any `protobuf:"bytes,1,rep,name=to_renew,json=toRenew,proto3" json:"to_renew,omitempty"`
	// deleted lists the IDs for all objects which were deleted.
	Deleted []string `protobuf:"bytes,2,rep,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *Changeset) Reset() {
	*x = Changeset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_digest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Changeset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Changeset) ProtoMessage() {}

func (x *Changeset) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_digest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Changeset.ProtoReflect.Descriptor instead.
func (*Changeset) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_digest_proto_rawDescGZIP(), []int{4}
}

func (x *Changeset) GetToRenew() []*any.Any {
	if x != nil {
		return x.ToRenew
	}
	return nil
}

func (x *Changeset) GetDeleted() []string {
	if x != nil {
		return x.Deleted
	}
	return nil
}

var File_orc8r_protos_digest_proto protoreflect.FileDescriptor

var file_orc8r_protos_digest_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x0a, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x54, 0x72, 0x65,
	0x65, 0x12, 0x34, 0x0a, 0x0b, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f,
	0x72, 0x63, 0x38, 0x72, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x72, 0x6f, 0x6f,
	0x74, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x66, 0x5f,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x4c, 0x65, 0x61, 0x66,
	0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x6c, 0x65, 0x61, 0x66, 0x44, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x73, 0x22, 0x34, 0x0a, 0x06, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x11, 0x6d, 0x64, 0x35, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x5f, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x64, 0x35, 0x42, 0x61, 0x73,
	0x65, 0x36, 0x34, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x0a, 0x4c, 0x65, 0x61,
	0x66, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e,
	0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x0b, 0x4c, 0x65, 0x61, 0x66, 0x44, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63,
	0x38, 0x72, 0x2e, 0x4c, 0x65, 0x61, 0x66, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x07, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x22, 0x56, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x65, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x74, 0x6f, 0x52,
	0x65, 0x6e, 0x65, 0x77, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x1b,
	0x5a, 0x19, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x6c, 0x69,
	0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_orc8r_protos_digest_proto_rawDescOnce sync.Once
	file_orc8r_protos_digest_proto_rawDescData = file_orc8r_protos_digest_proto_rawDesc
)

func file_orc8r_protos_digest_proto_rawDescGZIP() []byte {
	file_orc8r_protos_digest_proto_rawDescOnce.Do(func() {
		file_orc8r_protos_digest_proto_rawDescData = protoimpl.X.CompressGZIP(file_orc8r_protos_digest_proto_rawDescData)
	})
	return file_orc8r_protos_digest_proto_rawDescData
}

var file_orc8r_protos_digest_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_orc8r_protos_digest_proto_goTypes = []interface{}{
	(*DigestTree)(nil),  // 0: magma.orc8r.DigestTree
	(*Digest)(nil),      // 1: magma.orc8r.Digest
	(*LeafDigest)(nil),  // 2: magma.orc8r.LeafDigest
	(*LeafDigests)(nil), // 3: magma.orc8r.LeafDigests
	(*Changeset)(nil),   // 4: magma.orc8r.Changeset
	(*any.Any)(nil),     // 5: google.protobuf.Any
}
var file_orc8r_protos_digest_proto_depIdxs = []int32{
	1, // 0: magma.orc8r.DigestTree.root_digest:type_name -> magma.orc8r.Digest
	2, // 1: magma.orc8r.DigestTree.leaf_digests:type_name -> magma.orc8r.LeafDigest
	1, // 2: magma.orc8r.LeafDigest.digest:type_name -> magma.orc8r.Digest
	2, // 3: magma.orc8r.LeafDigests.digests:type_name -> magma.orc8r.LeafDigest
	5, // 4: magma.orc8r.Changeset.to_renew:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_orc8r_protos_digest_proto_init() }
func file_orc8r_protos_digest_proto_init() {
	if File_orc8r_protos_digest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orc8r_protos_digest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DigestTree); i {
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
		file_orc8r_protos_digest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Digest); i {
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
		file_orc8r_protos_digest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeafDigest); i {
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
		file_orc8r_protos_digest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeafDigests); i {
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
		file_orc8r_protos_digest_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Changeset); i {
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
			RawDescriptor: file_orc8r_protos_digest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orc8r_protos_digest_proto_goTypes,
		DependencyIndexes: file_orc8r_protos_digest_proto_depIdxs,
		MessageInfos:      file_orc8r_protos_digest_proto_msgTypes,
	}.Build()
	File_orc8r_protos_digest_proto = out.File
	file_orc8r_protos_digest_proto_rawDesc = nil
	file_orc8r_protos_digest_proto_goTypes = nil
	file_orc8r_protos_digest_proto_depIdxs = nil
}
