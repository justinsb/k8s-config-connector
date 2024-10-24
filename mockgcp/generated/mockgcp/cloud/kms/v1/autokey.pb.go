// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mockgcp/cloud/kms/v1/autokey.proto

package kmspb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for
// [Autokey.CreateKeyHandle][mockgcp.cloud.kms.v1.Autokey.CreateKeyHandle].
type CreateKeyHandleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the resource project and location to create the
	// [KeyHandle][mockgcp.cloud.kms.v1.KeyHandle] in, e.g.
	// `projects/{PROJECT_ID}/locations/{LOCATION}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. Id of the [KeyHandle][mockgcp.cloud.kms.v1.KeyHandle]. Must be
	// unique to the resource project and location. If not provided by the caller,
	// a new UUID is used.
	KeyHandleId string `protobuf:"bytes,2,opt,name=key_handle_id,json=keyHandleId,proto3" json:"key_handle_id,omitempty"`
	// Required. [KeyHandle][mockgcp.cloud.kms.v1.KeyHandle] to create.
	KeyHandle *KeyHandle `protobuf:"bytes,3,opt,name=key_handle,json=keyHandle,proto3" json:"key_handle,omitempty"`
}

func (x *CreateKeyHandleRequest) Reset() {
	*x = CreateKeyHandleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKeyHandleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKeyHandleRequest) ProtoMessage() {}

func (x *CreateKeyHandleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKeyHandleRequest.ProtoReflect.Descriptor instead.
func (*CreateKeyHandleRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_kms_v1_autokey_proto_rawDescGZIP(), []int{0}
}

func (x *CreateKeyHandleRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateKeyHandleRequest) GetKeyHandleId() string {
	if x != nil {
		return x.KeyHandleId
	}
	return ""
}

func (x *CreateKeyHandleRequest) GetKeyHandle() *KeyHandle {
	if x != nil {
		return x.KeyHandle
	}
	return nil
}

// Request message for [GetKeyHandle][mockgcp.cloud.kms.v1.Autokey.GetKeyHandle].
type GetKeyHandleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the [KeyHandle][mockgcp.cloud.kms.v1.KeyHandle] resource,
	// e.g.
	// `projects/{PROJECT_ID}/locations/{LOCATION}/keyHandles/{KEY_HANDLE_ID}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetKeyHandleRequest) Reset() {
	*x = GetKeyHandleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyHandleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyHandleRequest) ProtoMessage() {}

func (x *GetKeyHandleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyHandleRequest.ProtoReflect.Descriptor instead.
func (*GetKeyHandleRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_kms_v1_autokey_proto_rawDescGZIP(), []int{1}
}

func (x *GetKeyHandleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Resource-oriented representation of a request to Cloud KMS Autokey and the
// resulting provisioning of a [CryptoKey][mockgcp.cloud.kms.v1.CryptoKey].
type KeyHandle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. Name of the [KeyHandle][mockgcp.cloud.kms.v1.KeyHandle]
	// resource, e.g.
	// `projects/{PROJECT_ID}/locations/{LOCATION}/keyHandles/{KEY_HANDLE_ID}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Name of a [CryptoKey][mockgcp.cloud.kms.v1.CryptoKey] that has
	// been provisioned for Customer Managed Encryption Key (CMEK) use in the
	// [KeyHandle][mockgcp.cloud.kms.v1.KeyHandle] project and location for the
	// requested resource type. The [CryptoKey][mockgcp.cloud.kms.v1.CryptoKey]
	// project will reflect the value configured in the
	// [AutokeyConfig][mockgcp.cloud.kms.v1.AutokeyConfig] on the resource
	// project's ancestor folder at the time of the
	// [KeyHandle][mockgcp.cloud.kms.v1.KeyHandle] creation. If more than one
	// ancestor folder has a configured
	// [AutokeyConfig][mockgcp.cloud.kms.v1.AutokeyConfig], the nearest of these
	// configurations is used.
	KmsKey string `protobuf:"bytes,3,opt,name=kms_key,json=kmsKey,proto3" json:"kms_key,omitempty"`
	// Required. Indicates the resource type that the resulting
	// [CryptoKey][mockgcp.cloud.kms.v1.CryptoKey] is meant to protect, e.g.
	// `{SERVICE}.googleapis.com/{TYPE}`. See documentation for supported resource
	// types.
	ResourceTypeSelector string `protobuf:"bytes,4,opt,name=resource_type_selector,json=resourceTypeSelector,proto3" json:"resource_type_selector,omitempty"`
}

func (x *KeyHandle) Reset() {
	*x = KeyHandle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyHandle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyHandle) ProtoMessage() {}

func (x *KeyHandle) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyHandle.ProtoReflect.Descriptor instead.
func (*KeyHandle) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_kms_v1_autokey_proto_rawDescGZIP(), []int{2}
}

func (x *KeyHandle) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KeyHandle) GetKmsKey() string {
	if x != nil {
		return x.KmsKey
	}
	return ""
}

func (x *KeyHandle) GetResourceTypeSelector() string {
	if x != nil {
		return x.ResourceTypeSelector
	}
	return ""
}

// Metadata message for
// [CreateKeyHandle][mockgcp.cloud.kms.v1.Autokey.CreateKeyHandle] long-running
// operation response.
type CreateKeyHandleMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateKeyHandleMetadata) Reset() {
	*x = CreateKeyHandleMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKeyHandleMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKeyHandleMetadata) ProtoMessage() {}

func (x *CreateKeyHandleMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKeyHandleMetadata.ProtoReflect.Descriptor instead.
func (*CreateKeyHandleMetadata) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_kms_v1_autokey_proto_rawDescGZIP(), []int{3}
}

// Request message for
// [Autokey.ListKeyHandles][mockgcp.cloud.kms.v1.Autokey.ListKeyHandles].
type ListKeyHandlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the resource project and location from which to list
	// [KeyHandles][mockgcp.cloud.kms.v1.KeyHandle], e.g.
	// `projects/{PROJECT_ID}/locations/{LOCATION}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. Optional limit on the number of
	// [KeyHandles][mockgcp.cloud.kms.v1.KeyHandle] to include in the response. The
	// service may return fewer than this value. Further
	// [KeyHandles][mockgcp.cloud.kms.v1.KeyHandle] can subsequently be obtained by
	// including the
	// [ListKeyHandlesResponse.next_page_token][mockgcp.cloud.kms.v1.ListKeyHandlesResponse.next_page_token]
	// in a subsequent request.  If unspecified, at most
	// 100 [KeyHandles][mockgcp.cloud.kms.v1.KeyHandle] will be returned.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. Optional pagination token, returned earlier via
	// [ListKeyHandlesResponse.next_page_token][mockgcp.cloud.kms.v1.ListKeyHandlesResponse.next_page_token].
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. Filter to apply when listing
	// [KeyHandles][mockgcp.cloud.kms.v1.KeyHandle], e.g.
	// `resource_type_selector="{SERVICE}.googleapis.com/{TYPE}"`.
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListKeyHandlesRequest) Reset() {
	*x = ListKeyHandlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKeyHandlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKeyHandlesRequest) ProtoMessage() {}

func (x *ListKeyHandlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKeyHandlesRequest.ProtoReflect.Descriptor instead.
func (*ListKeyHandlesRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_kms_v1_autokey_proto_rawDescGZIP(), []int{4}
}

func (x *ListKeyHandlesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListKeyHandlesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListKeyHandlesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListKeyHandlesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// Response message for
// [Autokey.ListKeyHandles][mockgcp.cloud.kms.v1.Autokey.ListKeyHandles].
type ListKeyHandlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resulting [KeyHandles][mockgcp.cloud.kms.v1.KeyHandle].
	KeyHandles []*KeyHandle `protobuf:"bytes,1,rep,name=key_handles,json=keyHandles,proto3" json:"key_handles,omitempty"`
	// A token to retrieve next page of results. Pass this value in
	// [ListKeyHandlesRequest.page_token][mockgcp.cloud.kms.v1.ListKeyHandlesRequest.page_token]
	// to retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListKeyHandlesResponse) Reset() {
	*x = ListKeyHandlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKeyHandlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKeyHandlesResponse) ProtoMessage() {}

func (x *ListKeyHandlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKeyHandlesResponse.ProtoReflect.Descriptor instead.
func (*ListKeyHandlesResponse) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_kms_v1_autokey_proto_rawDescGZIP(), []int{5}
}

func (x *ListKeyHandlesResponse) GetKeyHandles() []*KeyHandle {
	if x != nil {
		return x.KeyHandles
	}
	return nil
}

func (x *ListKeyHandlesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_mockgcp_cloud_kms_v1_autokey_proto protoreflect.FileDescriptor

var file_mockgcp_cloud_kms_v1_autokey_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6b, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0,
	0x41, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x27, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x6b, 0x65,
	0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x0a, 0x6b, 0x65, 0x79,
	0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x54,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x6b, 0x6d, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa3, 0x02, 0x0a, 0x09, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x6b,
	0x6d, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41,
	0x03, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x6d, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x52, 0x06, 0x6b, 0x6d, 0x73, 0x4b, 0x65, 0x79, 0x12,
	0x39, 0x0a, 0x16, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x14, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x3a, 0x7e, 0xea, 0x41, 0x7b, 0x0a,
	0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x6d, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x12, 0x3f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6b, 0x65, 0x79, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x5f, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x7d, 0x2a, 0x0a, 0x6b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x32,
	0x09, 0x6b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xbd, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65,
	0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x41, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x82, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65,
	0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x40, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x0a, 0x6b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xb9, 0x05, 0x0a, 0x07, 0x41,
	0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x12, 0xec, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x2c, 0x2e, 0x6d, 0x6f, 0x63,
	0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8b, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c,
	0x22, 0x2e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73,
	0x3a, 0x0a, 0x6b, 0x65, 0x79, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0xda, 0x41, 0x1f, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x6b, 0x65, 0x79, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x2c, 0x6b, 0x65, 0x79, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0xca, 0x41,
	0x24, 0x0a, 0x09, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x99, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x29, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x22, 0x3d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x6b, 0x65, 0x79,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0xac, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x73, 0x12, 0x2b, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4b, 0x65, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6b, 0x65, 0x79,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x1a, 0x74, 0xca, 0x41, 0x17, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x6d, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x57, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x6b, 0x6d, 0x73, 0x42, 0x55, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f,
	0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x0c, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x29, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6b, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x2f, 0x6b, 0x6d, 0x73, 0x70, 0x62, 0x3b, 0x6b, 0x6d, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mockgcp_cloud_kms_v1_autokey_proto_rawDescOnce sync.Once
	file_mockgcp_cloud_kms_v1_autokey_proto_rawDescData = file_mockgcp_cloud_kms_v1_autokey_proto_rawDesc
)

func file_mockgcp_cloud_kms_v1_autokey_proto_rawDescGZIP() []byte {
	file_mockgcp_cloud_kms_v1_autokey_proto_rawDescOnce.Do(func() {
		file_mockgcp_cloud_kms_v1_autokey_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgcp_cloud_kms_v1_autokey_proto_rawDescData)
	})
	return file_mockgcp_cloud_kms_v1_autokey_proto_rawDescData
}

var file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mockgcp_cloud_kms_v1_autokey_proto_goTypes = []interface{}{
	(*CreateKeyHandleRequest)(nil),  // 0: mockgcp.cloud.kms.v1.CreateKeyHandleRequest
	(*GetKeyHandleRequest)(nil),     // 1: mockgcp.cloud.kms.v1.GetKeyHandleRequest
	(*KeyHandle)(nil),               // 2: mockgcp.cloud.kms.v1.KeyHandle
	(*CreateKeyHandleMetadata)(nil), // 3: mockgcp.cloud.kms.v1.CreateKeyHandleMetadata
	(*ListKeyHandlesRequest)(nil),   // 4: mockgcp.cloud.kms.v1.ListKeyHandlesRequest
	(*ListKeyHandlesResponse)(nil),  // 5: mockgcp.cloud.kms.v1.ListKeyHandlesResponse
	(*longrunningpb.Operation)(nil), // 6: google.longrunning.Operation
}
var file_mockgcp_cloud_kms_v1_autokey_proto_depIdxs = []int32{
	2, // 0: mockgcp.cloud.kms.v1.CreateKeyHandleRequest.key_handle:type_name -> mockgcp.cloud.kms.v1.KeyHandle
	2, // 1: mockgcp.cloud.kms.v1.ListKeyHandlesResponse.key_handles:type_name -> mockgcp.cloud.kms.v1.KeyHandle
	0, // 2: mockgcp.cloud.kms.v1.Autokey.CreateKeyHandle:input_type -> mockgcp.cloud.kms.v1.CreateKeyHandleRequest
	1, // 3: mockgcp.cloud.kms.v1.Autokey.GetKeyHandle:input_type -> mockgcp.cloud.kms.v1.GetKeyHandleRequest
	4, // 4: mockgcp.cloud.kms.v1.Autokey.ListKeyHandles:input_type -> mockgcp.cloud.kms.v1.ListKeyHandlesRequest
	6, // 5: mockgcp.cloud.kms.v1.Autokey.CreateKeyHandle:output_type -> google.longrunning.Operation
	2, // 6: mockgcp.cloud.kms.v1.Autokey.GetKeyHandle:output_type -> mockgcp.cloud.kms.v1.KeyHandle
	5, // 7: mockgcp.cloud.kms.v1.Autokey.ListKeyHandles:output_type -> mockgcp.cloud.kms.v1.ListKeyHandlesResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mockgcp_cloud_kms_v1_autokey_proto_init() }
func file_mockgcp_cloud_kms_v1_autokey_proto_init() {
	if File_mockgcp_cloud_kms_v1_autokey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKeyHandleRequest); i {
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
		file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyHandleRequest); i {
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
		file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyHandle); i {
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
		file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKeyHandleMetadata); i {
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
		file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKeyHandlesRequest); i {
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
		file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKeyHandlesResponse); i {
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
			RawDescriptor: file_mockgcp_cloud_kms_v1_autokey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mockgcp_cloud_kms_v1_autokey_proto_goTypes,
		DependencyIndexes: file_mockgcp_cloud_kms_v1_autokey_proto_depIdxs,
		MessageInfos:      file_mockgcp_cloud_kms_v1_autokey_proto_msgTypes,
	}.Build()
	File_mockgcp_cloud_kms_v1_autokey_proto = out.File
	file_mockgcp_cloud_kms_v1_autokey_proto_rawDesc = nil
	file_mockgcp_cloud_kms_v1_autokey_proto_goTypes = nil
	file_mockgcp_cloud_kms_v1_autokey_proto_depIdxs = nil
}