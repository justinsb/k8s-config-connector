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

package mockartifactregistry

import (
	"context"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/devtools/artifactregistry/v1"
)

type ArtifactRegistryV1 struct {
	*MockService
	pb.UnimplementedArtifactRegistryServer
}

func (s *ArtifactRegistryV1) GetRepository(ctx context.Context, req *pb.GetRepositoryRequest) (*pb.Repository, error) {
	name, err := s.parseArtifactRegistryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Repository{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ArtifactRegistryV1) CreateRepository(ctx context.Context, req *pb.CreateRepositoryRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/repositories/" + req.RepositoryId
	name, err := s.parseArtifactRegistryName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Repository).(*pb.Repository)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *ArtifactRegistryV1) UpdateRepository(ctx context.Context, req *pb.UpdateRepositoryRequest) (*pb.Repository, error) {
	reqName := req.GetRepository().GetName()

	name, err := s.parseArtifactRegistryName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Repository{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.Repository.GetDescription()
		case "labels":
			obj.Labels = req.Repository.GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ArtifactRegistryV1) DeleteRepository(ctx context.Context, req *pb.DeleteRepositoryRequest) (*longrunning.Operation, error) {
	name, err := s.parseArtifactRegistryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Repository{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}
