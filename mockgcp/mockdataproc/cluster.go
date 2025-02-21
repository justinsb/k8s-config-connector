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

// +tool:mockgcp-support
// proto.service: google.cloud.dataproc.v1.ClusterController
// proto.message: google.cloud.dataproc.v1.Cluster

package mockdataproc

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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dataproc/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type dataprocClusterService struct {
	*MockService
	pb.UnimplementedClusterControllerServer
}

func (s *dataprocClusterService) GetCluster(ctx context.Context, req *pb.GetClusterRequest) (*pb.Cluster, error) {
	name, err := s.parseClusterName(req.ProjectId, req.Region, req.ClusterName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *dataprocClusterService) CreateCluster(ctx context.Context, req *pb.CreateClusterRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", req.GetProjectId(), req.GetRegion(), req.GetCluster().GetClusterName())
	name, err := s.parseClusterName(req.ProjectId, req.Region, req.Cluster.ClusterName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetCluster()).(*pb.Cluster)
	obj.ProjectId = name.Project.ProjectID
	obj.ClusterName = name.ClusterName
	obj.Status = &pb.ClusterStatus{
		State:        pb.ClusterStatus_CREATING,
		StateStartTime: timestamppb.New(now),
	}

	obj.StatusHistory = append(obj.StatusHistory, obj.Status)

	s.populateDefaultsForCluster(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project.ProjectID, name.Region)
	lroMetadata := &pb.ClusterOperationMetadata{
		ProjectId:   name.Project.ProjectID,
		ClusterName: name.ClusterName,
		ClusterUuid: string(obj.ClusterUuid),
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()

		return mutateObject(ctx, s.storage, fqn, func(obj *pb.Cluster) error {
			obj.Status.State = pb.ClusterStatus_RUNNING
			return nil
		})
	})
}

func (s *dataprocClusterService) populateDefaultsForCluster(obj *pb.Cluster) {

}

func (s *dataprocClusterService) UpdateCluster(ctx context.Context, req *pb.UpdateClusterRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseClusterName(req.ProjectId, req.Region, req.ClusterName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Status = &pb.ClusterStatus{
		State:        pb.ClusterStatus_UPDATING,
		StateStartTime: timestamppb.New(now),
	}

	// TODO: Implement proper fieldmask support.
	updated := proto.Clone(req.GetCluster()).(*pb.Cluster)
	updated.ProjectId = name.Project.ProjectID
	updated.ClusterName = name.ClusterName
	updated.Status = &pb.ClusterStatus{
		State:        pb.ClusterStatus_UPDATING,
		StateStartTime: timestamppb.New(now),
	}
	updated.StatusHistory = append(updated.StatusHistory, updated.Status)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project.ProjectID, name.Region)
	lroMetadata := &pb.ClusterOperationMetadata{
		ProjectId:   name.Project.ProjectID,
		ClusterName: name.ClusterName,
		ClusterUuid: string(obj.ClusterUuid),
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		if err := s.storage.Get(ctx, fqn, obj); err != nil {
			return nil, err
		}

		obj.Status = &pb.ClusterStatus{
			State:          pb.ClusterStatus_RUNNING,
			StateStartTime: timestamppb.New(now),
		}
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *dataprocClusterService) DeleteCluster(ctx context.Context, req *pb.DeleteClusterRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseClusterName(req.ProjectId, req.Region, req.ClusterName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	deleted := &pb.Cluster{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project.ProjectID, name.Region)
	lroMetadata := &pb.ClusterOperationMetadata{
		ProjectId:   name.Project.ProjectID,
		ClusterName: name.ClusterName,
		ClusterUuid: string(deleted.ClusterUuid),
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type clusterName struct {
	Project     *projects.ProjectData
	Region      string
	ClusterName string
}

func (n *clusterName) String() string {
	return fmt.Sprintf("projects/%s/regions/%s/clusters/%s", n.Project.ID, n.Region, n.ClusterName)
}

// parseClusterName parses a string into an clusterName.
// The expected form is `projects/*/regions/*/clusters/*`.
func (s *MockService) parseClusterName(projectID, region, clusterName string) (*clusterName, error) {
	project, err := s.Projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	name := &clusterName{
		Project:     project,
		Region:      region,
		ClusterName: clusterName,
	}

	return name, nil
}

// mutateObject updates the object; it gets the object by fqn, calls mutator, then updates the object
func mutateObject[T proto.Message](ctx context.Context, storage storage.Storage, fqn string, mutator func(obj T) error) (T, error) {
	var nilT T

	typeT := reflect.TypeOf(nilT)
	obj := reflect.New(typeT.Elem()).Interface().(T)
	if err := storage.Get(ctx, fqn, obj); err != nil {
		return nilT, err
	}

	if err := mutator(obj); err != nil {
		return nilT, err
	}

	if err := storage.Update(ctx, fqn, obj); err != nil {
		return nilT, err
	}

	return obj, nil
}


