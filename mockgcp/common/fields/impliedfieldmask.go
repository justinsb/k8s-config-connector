package fields

import (
	"context"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ComputeImpliedFieldMask computes the implied field mask when it is omitted.
// According to AIP 134: this is "equivalent to all fields that are populated (have a non-empty value)."
// https://google.aip.dev/134
func ComputeImpliedFieldMask(ctx context.Context, req proto.Message) []string {
	var fieldMask []string
	req.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		fieldName := string(fd.Name())
		fieldMask = append(fieldMask, fieldName)
		return true
	})
	sort.Strings(fieldMask)
	return fieldMask
}
