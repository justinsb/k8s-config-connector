// Copyright 2023 Google LLC
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

package mockstorage

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

type bucketsServer struct {
	*MockService
	pb.UnimplementedBucketsServerServer
}

func (s *bucketsServer) GetBucket(ctx context.Context, req *pb.GetBucketRequest) (*pb.Bucket, error) {
	fqn := req.GetName()

	obj := &pb.Bucket{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *bucketsServer) InsertBucket(ctx context.Context, req *pb.InsertBucketRequest) (*pb.Bucket, error) {
	fqn := req.GetBucket().GetName()

	now := timestamppb.Now()

	obj := proto.Clone(req.GetBucket()).(*pb.Bucket)
	obj.Name = fqn
	obj.Id = fqn
	obj.Kind = "storage#bucket"
	obj.Name = fqn
	obj.ProjectNumber = 0 // todo
	obj.Location = "US"
	obj.LocationType = "multi-region"
	obj.Rpo = "DEFAULT"
	obj.SelfLink = fmt.Sprintf("https://www.googleapis.com/storage/v1/b/%s", obj.Name)
	obj.StorageClass = "STANDARD"
	obj.TimeCreated = now
	obj.Updated = now

	iamConfiguration := obj.IamConfiguration
	if iamConfiguration == nil {
		iamConfiguration = &pb.BucketIamConfiguration{}
		obj.IamConfiguration = iamConfiguration
	}
	if iamConfiguration.PublicAccessPrevention == "" {
		iamConfiguration.PublicAccessPrevention = "inherited"
	}
	bucketPolicyOnly := iamConfiguration.BucketPolicyOnly
	if bucketPolicyOnly == nil {
		bucketPolicyOnly = &pb.BucketPolicyOnly{}
		iamConfiguration.BucketPolicyOnly = bucketPolicyOnly
	}
	ubla := iamConfiguration.UniformBucketLevelAccess
	if ubla == nil {
		ubla = &pb.UniformBucketLevelAccess{}
		iamConfiguration.UniformBucketLevelAccess = ubla
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *bucketsServer) PatchBucket(ctx context.Context, req *pb.PatchBucketRequest) (*pb.Bucket, error) {
	fqn := req.GetName()

	obj := &pb.Bucket{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if patch := req.Bucket; patch != nil {
		if patch.Labels != nil {
			obj.Labels = patch.Labels
		}
		if patch.Lifecycle != nil {
			obj.Lifecycle = patch.Lifecycle
		}
		if patch.Versioning != nil {
			obj.Versioning = patch.Versioning
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *bucketsServer) DeleteBucket(ctx context.Context, req *pb.DeleteBucketRequest) (*empty.Empty, error) {
	fqn := req.GetName()

	deletedObj := &pb.Bucket{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
