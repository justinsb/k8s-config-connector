// Copyright 2022 Google LLC
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

package mockcomposer

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/klog/v2"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/orchestration/airflow/service/v1"
)

type ComposerV1 struct {
	*MockService
	pb.UnimplementedEnvironmentsServer
}

func (s *ComposerV1) GetEnvironment(ctx context.Context, req *pb.GetEnvironmentRequest) (*pb.Environment, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "environment %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading environment: %v", err)
		}
	}

	return obj, nil
}
func (s *ComposerV1) CreateEnvironment(ctx context.Context, req *pb.CreateEnvironmentRequest) (*longrunning.Operation, error) {
	reqName := req.GetEnvironment().GetName()
	name, err := s.parseEnvironmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	if req.GetParent()+"/environments/"+name.EnvironmentName != fqn {
		return nil, status.Errorf(codes.InvalidArgument, "name %q does not match parent", reqName)
	}

	now := timestamppb.New(time.Now().Round(time.Microsecond))
	obj := proto.Clone(req.Environment).(*pb.Environment)
	obj.Name = fqn
	obj.Uuid = string(uuid.NewUUID())

	uuidNoDashes := strings.ReplaceAll(obj.Uuid, "-", "")

	shortEnvironmentName := fmt.Sprintf("%.15s", name.EnvironmentName)
	prefix := name.Location + "-" + shortEnvironmentName + "-" + uuidNoDashes[:8]

	if obj.Config == nil {
		obj.Config = &pb.EnvironmentConfig{}
	}
	obj.Config.AirflowByoidUri = fmt.Sprintf("https://%s-dot-%s.composer.byoid.googleusercontent.com", uuidNoDashes, name.Location)
	obj.Config.AirflowUri = fmt.Sprintf("https://%s-dot-%s.composer.googleusercontent.com", uuidNoDashes, name.Location)
	obj.Config.GkeCluster = fmt.Sprintf("projects/%s/locations/%s/clusters/%s-gke", name.Project.ID, name.Location, prefix)
	obj.Config.DagGcsPrefix = fmt.Sprintf("gs://%s-bucket/dags", prefix)

	if obj.Config.DatabaseConfig == nil {
		obj.Config.DatabaseConfig = &pb.DatabaseConfig{}
	}
	if obj.Config.EncryptionConfig == nil {
		obj.Config.EncryptionConfig = &pb.EncryptionConfig{}
	}
	if obj.Config.EnvironmentSize == pb.EnvironmentConfig_ENVIRONMENT_SIZE_UNSPECIFIED {
		obj.Config.EnvironmentSize = pb.EnvironmentConfig_ENVIRONMENT_SIZE_SMALL
	}

	if obj.Config.MaintenanceWindow == nil {
		obj.Config.MaintenanceWindow = &pb.MaintenanceWindow{}
	}
	if obj.Config.MaintenanceWindow.EndTime == nil {
		// Unclear for more hours, maybe a bug?
		obj.Config.MaintenanceWindow.EndTime = timestamppb.New(time.Unix(0, 0).Add(4 * time.Hour))
	}
	if obj.Config.MaintenanceWindow.StartTime == nil {
		obj.Config.MaintenanceWindow.StartTime = timestamppb.New(time.Unix(0, 0))
	}
	if obj.Config.MaintenanceWindow.Recurrence == "" {
		obj.Config.MaintenanceWindow.Recurrence = "FREQ=WEEKLY;BYDAY=FR,SA,SU"
	}

	if obj.Config.NodeConfig == nil {
		obj.Config.NodeConfig = &pb.NodeConfig{}
	}
	if obj.Config.NodeConfig.IpAllocationPolicy == nil {
		obj.Config.NodeConfig.IpAllocationPolicy = &pb.IPAllocationPolicy{}
	}
	if obj.Config.NodeConfig.Network == "" {
		obj.Config.NodeConfig.Network = "projects/" + name.Project.ID + "/global/networks/default"
	}

	if obj.Config.PrivateEnvironmentConfig == nil {
		obj.Config.PrivateEnvironmentConfig = &pb.PrivateEnvironmentConfig{}
	}
	if obj.Config.PrivateEnvironmentConfig.CloudComposerNetworkIpv4CidrBlock == "" {
		obj.Config.PrivateEnvironmentConfig.CloudComposerNetworkIpv4CidrBlock = "172.31.245.0/24"
	}
	if obj.Config.PrivateEnvironmentConfig.CloudSqlIpv4CidrBlock == "" {
		obj.Config.PrivateEnvironmentConfig.CloudSqlIpv4CidrBlock = "10.0.0.0/12"
	}
	if obj.Config.PrivateEnvironmentConfig.PrivateClusterConfig == nil {
		obj.Config.PrivateEnvironmentConfig.PrivateClusterConfig = &pb.PrivateClusterConfig{}
	}

	if obj.Config.SoftwareConfig == nil {
		obj.Config.SoftwareConfig = &pb.SoftwareConfig{}
	}
	if obj.Config.SoftwareConfig.ImageVersion == "" {
		obj.Config.SoftwareConfig.ImageVersion = "composer-2.4.6-airflow-2.5.3"
	}

	if obj.Config.WebServerNetworkAccessControl == nil {
		obj.Config.WebServerNetworkAccessControl = &pb.WebServerNetworkAccessControl{}
	}
	if obj.Config.WebServerNetworkAccessControl.AllowedIpRanges == nil {
		obj.Config.WebServerNetworkAccessControl.AllowedIpRanges = []*pb.WebServerNetworkAccessControl_AllowedIpRange{
			{Description: "Allows access from all IPv4 addresses (default value)", Value: "0.0.0.0/0"},
			{Description: "Allows access from all IPv6 addresses (default value)", Value: "::0/0"},
		}
	}

	if obj.Config.WorkloadsConfig == nil {
		obj.Config.WorkloadsConfig = &pb.WorkloadsConfig{
			Scheduler: &pb.WorkloadsConfig_SchedulerResource{
				Count:     1,
				Cpu:       0.5,
				MemoryGb:  2,
				StorageGb: 1,
			},
			WebServer: &pb.WorkloadsConfig_WebServerResource{
				Cpu:       0.5,
				MemoryGb:  2,
				StorageGb: 1,
			},
			Worker: &pb.WorkloadsConfig_WorkerResource{
				Cpu:       0.5,
				MemoryGb:  2,
				StorageGb: 1,
				MinCount:  1,
				MaxCount:  3,
			},
		}
	}

	obj.CreateTime = now
	obj.State = pb.Environment_RUNNING
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating environment: %v", err)
	}

	lro, err := s.operations.NewLRO(ctx)
	if err != nil {
		return nil, err
	}
	lro.Name = "projects/" + name.Project.ID + "/locations/" + name.Location + "/" + lro.Name

	metadata := &pb.OperationMetadata{
		CreateTime:    now,
		OperationType: pb.OperationMetadata_CREATE,
		Resource:      fqn,
		ResourceUuid:  obj.Uuid,
		State:         pb.OperationMetadata_PENDING, // TODO: Match with done
	}
	any, err := anypb.New(metadata)
	if err != nil {
		return nil, err
	}
	lro.Metadata = any
	return lro, nil
}

func (s *ComposerV1) UpdateEnvironment(ctx context.Context, req *pb.UpdateEnvironmentRequest) (*longrunning.Operation, error) {
	reqName := req.GetEnvironment().GetName()

	name, err := s.parseEnvironmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "environment %q not found", reqName)
		}
		return nil, status.Errorf(codes.Internal, "error reading environment: %v", err)
	}

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		klog.Warningf("update_mask was not provided in request, should be required")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.GetEnvironment().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating environment: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *ComposerV1) DeleteEnvironment(ctx context.Context, req *pb.DeleteEnvironmentRequest) (*longrunning.Operation, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := timestamppb.Now()

	kind := (&pb.Environment{}).ProtoReflect().Descriptor()
	if err := s.storage.Delete(ctx, kind, fqn); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "environment %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting environment: %v", err)
		}
	}

	lro, err := s.operations.NewLRO(ctx)
	if err != nil {
		return nil, err
	}
	lro.Name = "projects/" + name.Project.ID + "/locations/" + name.Location + "/" + lro.Name

	metadata := &pb.OperationMetadata{
		CreateTime:    now,
		OperationType: pb.OperationMetadata_DELETE,
		Resource:      fqn,
		//ResourceUuid:  deleted.Uuid,
		State: pb.OperationMetadata_PENDING, // TODO: Match with done
	}
	any, err := anypb.New(metadata)
	if err != nil {
		return nil, err
	}
	lro.Metadata = any
	return lro, nil
}
