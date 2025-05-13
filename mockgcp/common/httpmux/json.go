// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package httpmux

import (
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"k8s.io/klog/v2"
)

func MarshalAsJSON(obj proto.Message) ([]byte, error) {
	return protojson.MarshalOptions{Resolver: &protoResolver{}}.Marshal(obj)
}

type protoResolver struct {
}

var _ protoregistry.ExtensionTypeResolver = &protoResolver{}

func (r *protoResolver) FindExtensionByName(message protoreflect.FullName) (protoreflect.ExtensionType, error) {
	v, err := protoregistry.GlobalTypes.FindExtensionByName(r.remapName(message))
	if err != nil {
		klog.Warningf("error resolving FindExtensionByName(%v): %v", message, err)
	}
	return v, err
}

func (r *protoResolver) FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error) {
	v, err := protoregistry.GlobalTypes.FindExtensionByNumber(r.remapName(message), field)
	if err != nil {
		klog.Warningf("error resolving FindExtensionByNumber(%v, %v): %v", message, field, err)
	}
	return v, err
}

var _ protoregistry.MessageTypeResolver = &protoResolver{}

func (r *protoResolver) FindMessageByName(message protoreflect.FullName) (protoreflect.MessageType, error) {
	v, err := protoregistry.GlobalTypes.FindMessageByName(r.remapName(message))
	if err != nil {
		klog.Warningf("error resolving FindMessageByName(%v): %v", message, err)
	}
	return v, err
}

func (r *protoResolver) FindMessageByURL(url string) (protoreflect.MessageType, error) {
	var aliases []string

	if strings.HasPrefix(url, "type.googleapis.com/google.") {
		aliases = append(aliases, "type.googleapis.com/mockgcp."+strings.TrimPrefix(url, "type.googleapis.com/google."))
	}

	switch url {
	case "type.googleapis.com/google.cloud.apigee.v1.OperationMetadata":
		aliases = append(aliases, "type.googleapis.com/mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1OperationMetadata")
	}

	for _, alias := range aliases {
		mt, err := protoregistry.GlobalTypes.FindMessageByURL(alias)
		if err != nil {
			klog.Warningf("alias lookup for FindMessageByURL(%q) failed: %v", alias, err)
		} else {
			return mt, nil
		}
	}

	v, err := protoregistry.GlobalTypes.FindMessageByURL(url)
	if err != nil {
		klog.Warningf("error resolving FindMessageByName(%v): %v", url, err)
	}
	return v, err
}

func (r *protoResolver) remapName(name protoreflect.FullName) protoreflect.FullName {
	// Remap names with a prefix of "google."" to be "mockgcp.", so we can find them.

	s := string(name)
	if strings.HasPrefix(s, "google.") {
		s = "mockgcp." + strings.TrimPrefix(s, "google.")
		return protoreflect.FullName(s)
	}

	switch s {
	case "google.cloud.apigee.v1.OperationMetadata":
		return protoreflect.FullName("mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1OperationMetadata")
	}

	return name
}
