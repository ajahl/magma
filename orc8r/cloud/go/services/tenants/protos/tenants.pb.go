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
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: orc8r/cloud/go/services/tenants/protos/tenants.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	protos "magma/orc8r/lib/go/protos"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTenantRequest) Reset() {
	*x = GetTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantRequest) ProtoMessage() {}

func (x *GetTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantRequest.ProtoReflect.Descriptor instead.
func (*GetTenantRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescGZIP(), []int{0}
}

func (x *GetTenantRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Tenant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Networks []string `protobuf:"bytes,2,rep,name=networks,proto3" json:"networks,omitempty"`
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenant.ProtoReflect.Descriptor instead.
func (*Tenant) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescGZIP(), []int{1}
}

func (x *Tenant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tenant) GetNetworks() []string {
	if x != nil {
		return x.Networks
	}
	return nil
}

type TenantList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenants []*IDAndTenant `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
}

func (x *TenantList) Reset() {
	*x = TenantList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantList) ProtoMessage() {}

func (x *TenantList) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantList.ProtoReflect.Descriptor instead.
func (*TenantList) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescGZIP(), []int{2}
}

func (x *TenantList) GetTenants() []*IDAndTenant {
	if x != nil {
		return x.Tenants
	}
	return nil
}

type IDAndTenant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Tenant *Tenant `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *IDAndTenant) Reset() {
	*x = IDAndTenant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDAndTenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDAndTenant) ProtoMessage() {}

func (x *IDAndTenant) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDAndTenant.ProtoReflect.Descriptor instead.
func (*IDAndTenant) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescGZIP(), []int{3}
}

func (x *IDAndTenant) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IDAndTenant) GetTenant() *Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type GetControlProxyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetControlProxyRequest) Reset() {
	*x = GetControlProxyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetControlProxyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetControlProxyRequest) ProtoMessage() {}

func (x *GetControlProxyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetControlProxyRequest.ProtoReflect.Descriptor instead.
func (*GetControlProxyRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescGZIP(), []int{4}
}

func (x *GetControlProxyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetControlProxyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ControlProxy string `protobuf:"bytes,2,opt,name=controlProxy,proto3" json:"controlProxy,omitempty"`
}

func (x *GetControlProxyResponse) Reset() {
	*x = GetControlProxyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetControlProxyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetControlProxyResponse) ProtoMessage() {}

func (x *GetControlProxyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetControlProxyResponse.ProtoReflect.Descriptor instead.
func (*GetControlProxyResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescGZIP(), []int{5}
}

func (x *GetControlProxyResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetControlProxyResponse) GetControlProxy() string {
	if x != nil {
		return x.ControlProxy
	}
	return ""
}

type CreateOrUpdateControlProxyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ControlProxy string `protobuf:"bytes,2,opt,name=controlProxy,proto3" json:"controlProxy,omitempty"`
}

func (x *CreateOrUpdateControlProxyRequest) Reset() {
	*x = CreateOrUpdateControlProxyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateControlProxyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateControlProxyRequest) ProtoMessage() {}

func (x *CreateOrUpdateControlProxyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateControlProxyRequest.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateControlProxyRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescGZIP(), []int{6}
}

func (x *CreateOrUpdateControlProxyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateOrUpdateControlProxyRequest) GetControlProxy() string {
	if x != nil {
		return x.ControlProxy
	}
	return ""
}

type CreateOrUpdateControlProxyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateOrUpdateControlProxyResponse) Reset() {
	*x = CreateOrUpdateControlProxyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateControlProxyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateControlProxyResponse) ProtoMessage() {}

func (x *CreateOrUpdateControlProxyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateControlProxyResponse.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateControlProxyResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescGZIP(), []int{7}
}

var File_orc8r_cloud_go_services_tenants_protos_tenants_proto protoreflect.FileDescriptor

var file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72,
	0x63, 0x38, 0x72, 0x1a, 0x19, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x38, 0x0a, 0x06, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x22, 0x40, 0x0a, 0x0a,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x61,
	0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x49, 0x44, 0x41, 0x6e, 0x64, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x4a,
	0x0a, 0x0b, 0x49, 0x44, 0x41, 0x6e, 0x64, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a,
	0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x22, 0x57, 0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x22, 0x24, 0x0a, 0x22,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xb2, 0x04, 0x0a, 0x0e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f,
	0x72, 0x63, 0x38, 0x72, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e,
	0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x49, 0x44, 0x41, 0x6e, 0x64, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38,
	0x72, 0x2e, 0x49, 0x44, 0x41, 0x6e, 0x64, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x1a, 0x11, 0x2e,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x22, 0x00, 0x12, 0x42, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x23, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x12, 0x2e, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63,
	0x38, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63,
	0x38, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x6d, 0x61, 0x67, 0x6d, 0x61,
	0x2f, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescOnce sync.Once
	file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescData = file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDesc
)

func file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescGZIP() []byte {
	file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescOnce.Do(func() {
		file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescData = protoimpl.X.CompressGZIP(file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescData)
	})
	return file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDescData
}

var file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_orc8r_cloud_go_services_tenants_protos_tenants_proto_goTypes = []interface{}{
	(*GetTenantRequest)(nil),                   // 0: magma.orc8r.GetTenantRequest
	(*Tenant)(nil),                             // 1: magma.orc8r.Tenant
	(*TenantList)(nil),                         // 2: magma.orc8r.TenantList
	(*IDAndTenant)(nil),                        // 3: magma.orc8r.IDAndTenant
	(*GetControlProxyRequest)(nil),             // 4: magma.orc8r.GetControlProxyRequest
	(*GetControlProxyResponse)(nil),            // 5: magma.orc8r.GetControlProxyResponse
	(*CreateOrUpdateControlProxyRequest)(nil),  // 6: magma.orc8r.CreateOrUpdateControlProxyRequest
	(*CreateOrUpdateControlProxyResponse)(nil), // 7: magma.orc8r.CreateOrUpdateControlProxyResponse
	(*protos.Void)(nil),                        // 8: magma.orc8r.Void
}
var file_orc8r_cloud_go_services_tenants_protos_tenants_proto_depIdxs = []int32{
	3, // 0: magma.orc8r.TenantList.tenants:type_name -> magma.orc8r.IDAndTenant
	1, // 1: magma.orc8r.IDAndTenant.tenant:type_name -> magma.orc8r.Tenant
	8, // 2: magma.orc8r.TenantsService.GetAllTenants:input_type -> magma.orc8r.Void
	0, // 3: magma.orc8r.TenantsService.GetTenant:input_type -> magma.orc8r.GetTenantRequest
	3, // 4: magma.orc8r.TenantsService.CreateTenant:input_type -> magma.orc8r.IDAndTenant
	3, // 5: magma.orc8r.TenantsService.SetTenant:input_type -> magma.orc8r.IDAndTenant
	0, // 6: magma.orc8r.TenantsService.DeleteTenant:input_type -> magma.orc8r.GetTenantRequest
	4, // 7: magma.orc8r.TenantsService.GetControlProxy:input_type -> magma.orc8r.GetControlProxyRequest
	6, // 8: magma.orc8r.TenantsService.CreateOrUpdateControlProxy:input_type -> magma.orc8r.CreateOrUpdateControlProxyRequest
	2, // 9: magma.orc8r.TenantsService.GetAllTenants:output_type -> magma.orc8r.TenantList
	1, // 10: magma.orc8r.TenantsService.GetTenant:output_type -> magma.orc8r.Tenant
	8, // 11: magma.orc8r.TenantsService.CreateTenant:output_type -> magma.orc8r.Void
	8, // 12: magma.orc8r.TenantsService.SetTenant:output_type -> magma.orc8r.Void
	8, // 13: magma.orc8r.TenantsService.DeleteTenant:output_type -> magma.orc8r.Void
	5, // 14: magma.orc8r.TenantsService.GetControlProxy:output_type -> magma.orc8r.GetControlProxyResponse
	7, // 15: magma.orc8r.TenantsService.CreateOrUpdateControlProxy:output_type -> magma.orc8r.CreateOrUpdateControlProxyResponse
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_orc8r_cloud_go_services_tenants_protos_tenants_proto_init() }
func file_orc8r_cloud_go_services_tenants_protos_tenants_proto_init() {
	if File_orc8r_cloud_go_services_tenants_protos_tenants_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTenantRequest); i {
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
		file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tenant); i {
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
		file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantList); i {
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
		file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDAndTenant); i {
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
		file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetControlProxyRequest); i {
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
		file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetControlProxyResponse); i {
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
		file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateControlProxyRequest); i {
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
		file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateControlProxyResponse); i {
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
			RawDescriptor: file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orc8r_cloud_go_services_tenants_protos_tenants_proto_goTypes,
		DependencyIndexes: file_orc8r_cloud_go_services_tenants_protos_tenants_proto_depIdxs,
		MessageInfos:      file_orc8r_cloud_go_services_tenants_protos_tenants_proto_msgTypes,
	}.Build()
	File_orc8r_cloud_go_services_tenants_protos_tenants_proto = out.File
	file_orc8r_cloud_go_services_tenants_protos_tenants_proto_rawDesc = nil
	file_orc8r_cloud_go_services_tenants_protos_tenants_proto_goTypes = nil
	file_orc8r_cloud_go_services_tenants_protos_tenants_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TenantsServiceClient is the client API for TenantsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TenantsServiceClient interface {
	GetAllTenants(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*TenantList, error)
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	CreateTenant(ctx context.Context, in *IDAndTenant, opts ...grpc.CallOption) (*protos.Void, error)
	SetTenant(ctx context.Context, in *IDAndTenant, opts ...grpc.CallOption) (*protos.Void, error)
	DeleteTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*protos.Void, error)
	GetControlProxy(ctx context.Context, in *GetControlProxyRequest, opts ...grpc.CallOption) (*GetControlProxyResponse, error)
	CreateOrUpdateControlProxy(ctx context.Context, in *CreateOrUpdateControlProxyRequest, opts ...grpc.CallOption) (*CreateOrUpdateControlProxyResponse, error)
}

type tenantsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantsServiceClient(cc grpc.ClientConnInterface) TenantsServiceClient {
	return &tenantsServiceClient{cc}
}

func (c *tenantsServiceClient) GetAllTenants(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*TenantList, error) {
	out := new(TenantList)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/GetAllTenants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/GetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) CreateTenant(ctx context.Context, in *IDAndTenant, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/CreateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) SetTenant(ctx context.Context, in *IDAndTenant, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/SetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) DeleteTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/DeleteTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) GetControlProxy(ctx context.Context, in *GetControlProxyRequest, opts ...grpc.CallOption) (*GetControlProxyResponse, error) {
	out := new(GetControlProxyResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/GetControlProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) CreateOrUpdateControlProxy(ctx context.Context, in *CreateOrUpdateControlProxyRequest, opts ...grpc.CallOption) (*CreateOrUpdateControlProxyResponse, error) {
	out := new(CreateOrUpdateControlProxyResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/CreateOrUpdateControlProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantsServiceServer is the server API for TenantsService service.
type TenantsServiceServer interface {
	GetAllTenants(context.Context, *protos.Void) (*TenantList, error)
	GetTenant(context.Context, *GetTenantRequest) (*Tenant, error)
	CreateTenant(context.Context, *IDAndTenant) (*protos.Void, error)
	SetTenant(context.Context, *IDAndTenant) (*protos.Void, error)
	DeleteTenant(context.Context, *GetTenantRequest) (*protos.Void, error)
	GetControlProxy(context.Context, *GetControlProxyRequest) (*GetControlProxyResponse, error)
	CreateOrUpdateControlProxy(context.Context, *CreateOrUpdateControlProxyRequest) (*CreateOrUpdateControlProxyResponse, error)
}

// UnimplementedTenantsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTenantsServiceServer struct {
}

func (*UnimplementedTenantsServiceServer) GetAllTenants(context.Context, *protos.Void) (*TenantList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTenants not implemented")
}
func (*UnimplementedTenantsServiceServer) GetTenant(context.Context, *GetTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (*UnimplementedTenantsServiceServer) CreateTenant(context.Context, *IDAndTenant) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (*UnimplementedTenantsServiceServer) SetTenant(context.Context, *IDAndTenant) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTenant not implemented")
}
func (*UnimplementedTenantsServiceServer) DeleteTenant(context.Context, *GetTenantRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenant not implemented")
}
func (*UnimplementedTenantsServiceServer) GetControlProxy(context.Context, *GetControlProxyRequest) (*GetControlProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetControlProxy not implemented")
}
func (*UnimplementedTenantsServiceServer) CreateOrUpdateControlProxy(context.Context, *CreateOrUpdateControlProxyRequest) (*CreateOrUpdateControlProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateControlProxy not implemented")
}

func RegisterTenantsServiceServer(s *grpc.Server, srv TenantsServiceServer) {
	s.RegisterService(&_TenantsService_serviceDesc, srv)
}

func _TenantsService_GetAllTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).GetAllTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/GetAllTenants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).GetAllTenants(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/GetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDAndTenant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/CreateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).CreateTenant(ctx, req.(*IDAndTenant))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_SetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDAndTenant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).SetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/SetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).SetTenant(ctx, req.(*IDAndTenant))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_DeleteTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).DeleteTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/DeleteTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).DeleteTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_GetControlProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetControlProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).GetControlProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/GetControlProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).GetControlProxy(ctx, req.(*GetControlProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_CreateOrUpdateControlProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateControlProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).CreateOrUpdateControlProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/CreateOrUpdateControlProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).CreateOrUpdateControlProxy(ctx, req.(*CreateOrUpdateControlProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TenantsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.TenantsService",
	HandlerType: (*TenantsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllTenants",
			Handler:    _TenantsService_GetAllTenants_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _TenantsService_GetTenant_Handler,
		},
		{
			MethodName: "CreateTenant",
			Handler:    _TenantsService_CreateTenant_Handler,
		},
		{
			MethodName: "SetTenant",
			Handler:    _TenantsService_SetTenant_Handler,
		},
		{
			MethodName: "DeleteTenant",
			Handler:    _TenantsService_DeleteTenant_Handler,
		},
		{
			MethodName: "GetControlProxy",
			Handler:    _TenantsService_GetControlProxy_Handler,
		},
		{
			MethodName: "CreateOrUpdateControlProxy",
			Handler:    _TenantsService_CreateOrUpdateControlProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/cloud/go/services/tenants/protos/tenants.proto",
}
