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
	"reflect"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *clusterControllerServer) GetCluster(ctx context.Context, req *pb.GetClusterRequest) (*pb.Cluster, error) {
	name, err := s.buildClusterName(req.ProjectId, req.Region, req.ClusterName)
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

func (s *clusterControllerServer) ListClusters(ctx context.Context, req *pb.ListClustersRequest) (*pb.ListClustersResponse, error) {
	name, err := s.buildClusterName(req.ProjectId, req.Region, "")
	if err != nil {
		return nil, err
	}

	findPrefix := name.String()

	var clusters []*pb.Cluster

	findKind := (&pb.Cluster{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		cluster := obj.(*pb.Cluster)
		clusters = append(clusters, cluster)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListClustersResponse{
		Clusters: clusters,
	}, nil

}

func (s *clusterControllerServer) CreateCluster(ctx context.Context, req *pb.CreateClusterRequest) (*longrunningpb.Operation, error) {
	name, err := s.buildClusterName(req.ProjectId, req.Region, req.Cluster.ClusterName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetCluster()).(*pb.Cluster)
	obj.ProjectId = name.Project.ID
	obj.ClusterName = name.ClusterName
	obj.Status = &pb.ClusterStatus{
		State:          pb.ClusterStatus_CREATING,
		StateStartTime: timestamppb.New(now),
	}

	obj.StatusHistory = append(obj.StatusHistory, obj.Status)

	s.populateDefaultsForCluster(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)
	lroMetadata := &pb.ClusterOperationMetadata{
		// ProjectId:   name.Project.ID,
		ClusterName:   name.ClusterName,
		ClusterUuid:   obj.ClusterUuid,
		OperationType: "CREATE",
		Description:   "Create cluster with 2 workers",
		Status: &pb.ClusterOperationStatus{
			InnerState:     "PENDING",
			State:          pb.ClusterOperationStatus_PENDING,
			StateStartTime: timestamppb.New(now),
		},
		Warnings: []string{
			"The firewall rules for specified network or subnetwork would allow ingress traffic from 0.0.0.0/0, which could be a security risk.",
			"The specified custom staging bucket 'dataproc-staging-us-central1-${projectNumber}-ch70stme' is not using uniform bucket level access IAM configuration. It is recommended to update bucket to enable the same. See https://cloud.google.com/storage/docs/uniform-bucket-level-access.",
			"No image specified. Using the default image version. It is recommended to select a specific image version in production, as the default image version may change at any time.",
		},
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		// lroMetadata.EndTime = timestamppb.Now()

		return mutateObject(ctx, s.storage, fqn, func(obj *pb.Cluster) error {
			obj.Status.State = pb.ClusterStatus_RUNNING
			obj.Config.ConfigBucket = "dataproc-staging-us-central1-${projectNumber}-ch70stme"
			obj.Config.EndpointConfig = &pb.EndpointConfig{}
			obj.Config.GceClusterConfig.InternalIpOnly = PtrTo(true)
			obj.Config.GceClusterConfig.NetworkUri = "https://www.googleapis.com/compute/v1/projects/${projectId}/global/networks/default"
			obj.Config.GceClusterConfig.ServiceAccountScopes = []string{"https://www.googleapis.com/auth/cloud-platform"}
			obj.Config.MasterConfig.DiskConfig.BootDiskSizeGb = 1000
			return nil
		})
	})
}

func (s *clusterControllerServer) populateDefaultsForCluster(obj *pb.Cluster) {

}

func (s *clusterControllerServer) UpdateCluster(ctx context.Context, req *pb.UpdateClusterRequest) (*longrunningpb.Operation, error) {
	name, err := s.buildClusterName(req.ProjectId, req.Region, req.ClusterName)
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
		State:          pb.ClusterStatus_UPDATING,
		StateStartTime: timestamppb.New(now),
	}

	// TODO: Implement proper fieldmask support.
	updated := proto.Clone(req.GetCluster()).(*pb.Cluster)
	updated.ProjectId = name.Project.ID
	updated.ClusterName = name.ClusterName
	updated.Status = &pb.ClusterStatus{
		State:          pb.ClusterStatus_UPDATING,
		StateStartTime: timestamppb.New(now),
	}
	updated.StatusHistory = append(updated.StatusHistory, updated.Status)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)
	lroMetadata := &pb.ClusterOperationMetadata{
		// ProjectId:   name.Project.ID,
		ClusterName: name.ClusterName,
		ClusterUuid: string(obj.ClusterUuid),
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		// lroMetadata.EndTime = timestamppb.Now()
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

func (s *clusterControllerServer) DeleteCluster(ctx context.Context, req *pb.DeleteClusterRequest) (*longrunningpb.Operation, error) {
	name, err := s.buildClusterName(req.ProjectId, req.Region, req.ClusterName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	// now := time.Now()

	deleted := &pb.Cluster{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)
	lroMetadata := &pb.ClusterOperationMetadata{
		// ProjectId:   name.Project.ID,
		ClusterName: name.ClusterName,
		ClusterUuid: string(deleted.ClusterUuid),
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		// lroMetadata.EndTime = timestamppb.Now()
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
func (s *MockService) parseClusterName(name string) (*clusterName, error) {

	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "clusters" {
		return s.buildClusterName(tokens[1], tokens[3], tokens[5])
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}

// buildClusterName builds a clusterName from the components.
func (s *MockService) buildClusterName(projectName, region, cluster string) (*clusterName, error) {

	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	return &clusterName{
		Project:     project,
		Region:      region,
		ClusterName: cluster,
	}, nil
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
