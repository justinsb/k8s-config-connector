package mockcomposer

import (
	"context"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/orchestration/airflow/service/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/util/uuid"
)

func (s *MockService) startLRO(ctx context.Context, obj *pb.Environment, operationType pb.OperationMetadata_Type) (*longrunningpb.Operation, error) {
	prefix := obj.Name
	if ix := strings.Index(prefix, "/environments/"); ix != -1 {
		prefix = prefix[:ix]
	}
	operationUUID := string(uuid.NewUUID())

	name := prefix + "/operations/" + operationUUID
	lro, err := s.operations.StartLRO(ctx, name, func() (proto.Message, error) {
		now := timestamppb.New(time.Now().Round(time.Microsecond))
		metadata := &pb.OperationMetadata{
			CreateTime:    now,
			OperationType: operationType,
			Resource:      obj.Name,
			ResourceUuid:  obj.Uuid,
			State:         pb.OperationMetadata_SUCCEEDED,
		}

		return metadata, nil
	})
	if err != nil {
		return nil, err
	}

	now := timestamppb.New(time.Now().Round(time.Microsecond))

	metadata := &pb.OperationMetadata{
		CreateTime:    now,
		OperationType: operationType,
		Resource:      obj.Name,
		ResourceUuid:  obj.Uuid,
		State:         pb.OperationMetadata_PENDING,
	}

	any, err := anypb.New(metadata)
	if err != nil {
		return nil, err
	}
	lro.Metadata = any
	return lro, nil
}
