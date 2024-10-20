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

package securesourcemanager

import (
	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securesourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Instance_HostConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_HostConfig) *krm.Instance_HostConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_HostConfig{}
	out.Html = direct.LazyPtr(in.GetHtml())
	out.Api = direct.LazyPtr(in.GetApi())
	out.GitHttp = direct.LazyPtr(in.GetGitHttp())
	out.GitSsh = direct.LazyPtr(in.GetGitSsh())
	return out
}
func Instance_HostConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_HostConfig) *pb.Instance_HostConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_HostConfig{}
	out.Html = direct.ValueOf(in.Html)
	out.Api = direct.ValueOf(in.Api)
	out.GitHttp = direct.ValueOf(in.GitHttp)
	out.GitSsh = direct.ValueOf(in.GitSsh)
	return out
}
func OperationMetadata_FromProto(mapCtx *direct.MapContext, in *pb.OperationMetadata) *krm.OperationMetadata {
	if in == nil {
		return nil
	}
	out := &krm.OperationMetadata{}
	out.CreateTime = OperationMetadata_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.EndTime = OperationMetadata_EndTime_FromProto(mapCtx, in.GetEndTime())
	out.Target = direct.LazyPtr(in.GetTarget())
	out.Verb = direct.LazyPtr(in.GetVerb())
	out.StatusMessage = direct.LazyPtr(in.GetStatusMessage())
	out.RequestedCancellation = direct.LazyPtr(in.GetRequestedCancellation())
	out.ApiVersion = direct.LazyPtr(in.GetApiVersion())
	return out
}
func OperationMetadata_ToProto(mapCtx *direct.MapContext, in *krm.OperationMetadata) *pb.OperationMetadata {
	if in == nil {
		return nil
	}
	out := &pb.OperationMetadata{}
	out.CreateTime = OperationMetadata_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.EndTime = OperationMetadata_EndTime_ToProto(mapCtx, in.EndTime)
	out.Target = direct.ValueOf(in.Target)
	out.Verb = direct.ValueOf(in.Verb)
	out.StatusMessage = direct.ValueOf(in.StatusMessage)
	out.RequestedCancellation = direct.ValueOf(in.RequestedCancellation)
	out.ApiVersion = direct.ValueOf(in.ApiVersion)
	return out
}
func Repository_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.Repository {
	if in == nil {
		return nil
	}
	out := &krm.Repository{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Instance = direct.LazyPtr(in.GetInstance())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = Repository_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = Repository_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Uris = Repository_URIs_FromProto(mapCtx, in.GetUris())
	out.InitialConfig = Repository_InitialConfig_FromProto(mapCtx, in.GetInitialConfig())
	return out
}
func Repository_ToProto(mapCtx *direct.MapContext, in *krm.Repository) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.Instance = direct.ValueOf(in.Instance)
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = Repository_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = Repository_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	out.Etag = direct.ValueOf(in.Etag)
	out.Uris = Repository_URIs_ToProto(mapCtx, in.Uris)
	out.InitialConfig = Repository_InitialConfig_ToProto(mapCtx, in.InitialConfig)
	return out
}
func Repository_InitialConfig_FromProto(mapCtx *direct.MapContext, in *pb.Repository_InitialConfig) *krm.Repository_InitialConfig {
	if in == nil {
		return nil
	}
	out := &krm.Repository_InitialConfig{}
	out.DefaultBranch = direct.LazyPtr(in.GetDefaultBranch())
	out.Gitignores = in.Gitignores
	out.License = direct.LazyPtr(in.GetLicense())
	out.Readme = direct.LazyPtr(in.GetReadme())
	return out
}
func Repository_InitialConfig_ToProto(mapCtx *direct.MapContext, in *krm.Repository_InitialConfig) *pb.Repository_InitialConfig {
	if in == nil {
		return nil
	}
	out := &pb.Repository_InitialConfig{}
	out.DefaultBranch = direct.ValueOf(in.DefaultBranch)
	out.Gitignores = in.Gitignores
	out.License = direct.ValueOf(in.License)
	out.Readme = direct.ValueOf(in.Readme)
	return out
}
func Repository_URIs_FromProto(mapCtx *direct.MapContext, in *pb.Repository_URIs) *krm.Repository_URIs {
	if in == nil {
		return nil
	}
	out := &krm.Repository_URIs{}
	out.Html = direct.LazyPtr(in.GetHtml())
	out.GitHttps = direct.LazyPtr(in.GetGitHttps())
	out.Api = direct.LazyPtr(in.GetApi())
	return out
}
func Repository_URIs_ToProto(mapCtx *direct.MapContext, in *krm.Repository_URIs) *pb.Repository_URIs {
	if in == nil {
		return nil
	}
	out := &pb.Repository_URIs{}
	out.Html = direct.ValueOf(in.Html)
	out.GitHttps = direct.ValueOf(in.GitHttps)
	out.Api = direct.ValueOf(in.Api)
	return out
}
func SecureSourceManagerInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.SecureSourceManagerInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerInstanceObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.State = direct.Enum_FromProto(mapCtx, in.State)
	out.StateNote = direct.Enum_FromProto(mapCtx, in.StateNote)
	// MISSING: KmsKey
	out.HostConfig = Instance_HostConfig_FromProto(mapCtx, in.GetHostConfig())
	return out
}
func SecureSourceManagerInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecureSourceManagerInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StateNote = direct.Enum_ToProto[pb.Instance_StateNote](mapCtx, in.StateNote)
	// MISSING: KmsKey
	out.HostConfig = Instance_HostConfig_ToProto(mapCtx, in.HostConfig)
	return out
}
func SecureSourceManagerInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.SecureSourceManagerInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerInstanceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: State
	// MISSING: StateNote
	out.KmsKey = direct.LazyPtr(in.GetKmsKey())
	// MISSING: HostConfig
	return out
}
func SecureSourceManagerInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecureSourceManagerInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: State
	// MISSING: StateNote
	out.KmsKey = direct.ValueOf(in.KmsKey)
	// MISSING: HostConfig
	return out
}
