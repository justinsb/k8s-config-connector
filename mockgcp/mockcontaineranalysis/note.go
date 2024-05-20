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

package mockcontaineranalysis

import (
	"context"
	"fmt"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/devtools/containeranalysis/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type ContainerAnalysisV1 struct {
	*MockService
	pb.UnimplementedProjectsNotesServerServer
}

func (s *ContainerAnalysisV1) GetProjectsNote(ctx context.Context, req *pb.GetProjectsNoteRequest) (*pb.Note, error) {
	name, err := s.parseProjectNoteName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Note{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContainerAnalysisV1) CreateProjectsNote(ctx context.Context, req *pb.CreateProjectsNoteRequest) (*pb.Note, error) {
	fmt.Println("________________________________________________________________________________")
	fmt.Println("CREATE")
	fmt.Println("________________________________________________________________________________")
	reqName := req.Parent + "/notes/" + req.NoteId
	name, err := s.parseProjectNoteName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.ProjectsNote).(*pb.Note)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContainerAnalysisV1) UpdateProjectsNote(ctx context.Context, req *pb.PatchProjectsNoteRequest) (*pb.Note, error) {
	reqName := req.GetName()

	name, err := s.parseProjectNoteName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Note{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {

		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

//func (s *ContainerAnalysisV1) DeleteProjectsNote(ctx context.Context, req *pb.DeleteProjectsNoteRequest) (*Empty, error) {
//	name, err := s.parseProjectNoteName(req.Name)
//	if err != nil {
//		return nil, err
//	}
//
//	fqn := name.String()
//
//	oldObj := &pb.Note{}
//	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
//		return nil, err
//	}
//
//	return &emptypb.Empty{}, nil
//}
