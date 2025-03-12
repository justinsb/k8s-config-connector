// Copyright 2025 Google LLC
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

// +tool:mockgcp-support
// proto.service: google.cloud.discoveryengine.v1.SiteSearchEngineService
// proto.message: google.cloud.discoveryengine.v1.TargetSite

package mockdiscoveryengine

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *siteSearchEngineService) GetTargetSite(ctx context.Context, req *pb.GetTargetSiteRequest) (*pb.TargetSite, error) {
	name, err := s.parseTargetSiteName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TargetSite{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "TargetSite %v not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *siteSearchEngineService) CreateTargetSite(ctx context.Context, req *pb.CreateTargetSiteRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/targetSites/%s", req.GetParent(), "id")
	name, err := s.parseTargetSiteName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetTargetSite()).(*pb.TargetSite)
	obj.Name = fqn
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", name.Project.ID, name.Location, name.Collection, name.DataStore)

	return s.operations.StartLRO(ctx, prefix, nil, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *siteSearchEngineService) UpdateTargetSite(ctx context.Context, req *pb.UpdateTargetSiteRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseTargetSiteName(req.TargetSite.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.TargetSite{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// paths := req.GetUpdateMask().GetPaths()
	// if len(paths) == 0 {
	// 	return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	// }

	// TODO: support update mask

	proto.Merge(obj, req.GetTargetSite())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", name.Project.ID, name.Location, name.Collection, name.DataStore)

	return s.operations.StartLRO(ctx, prefix, nil, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *siteSearchEngineService) DeleteTargetSite(ctx context.Context, req *pb.DeleteTargetSiteRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseTargetSiteName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.TargetSite{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", name.Project.ID, name.Location, name.Collection, name.DataStore)

	return s.operations.StartLRO(ctx, prefix, nil, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

type targetSiteName struct {
	Project    *projects.ProjectData
	Location   string
	Collection string
	DataStore  string
	TargetSite string
}

func (n *targetSiteName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s/siteSearchEngine/targetSites/%s", n.Project.Number, n.Location, n.Collection, n.DataStore, n.TargetSite)
}

func (s *MockService) parseTargetSiteName(name string) (*targetSiteName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" && tokens[8] == "targetSites" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &targetSiteName{
			Project:    project,
			Location:   tokens[3],
			Collection: tokens[5],
			DataStore:  tokens[7],
			TargetSite: tokens[9],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}
