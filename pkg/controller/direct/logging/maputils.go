package logging

// import (
// 	"errors"
// 	"fmt"
// 	"strconv"
// 	"time"

// 	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
// 	"google.golang.org/protobuf/reflect/protoreflect"
// 	"google.golang.org/protobuf/types/known/durationpb"
// 	"sigs.k8s.io/controller-runtime/pkg/client"
// )

// type MapContext struct {
// 	kube client.Reader
// 	errs []error
// }

// func (c *MapContext) Errorf(msg string, args ...interface{}) {
// 	c.errs = append(c.errs, fmt.Errorf(msg, args...))
// }

// func (c *MapContext) Err() error {
// 	return errors.Join(c.errs...)
// }

// func (c *MapContext) ResolveRef(ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
// 	retval := *ref
// 	if ref.External != "" {
// 		return &retval, nil
// 	}

// 	return nil, fmt.Errorf("ResolveRef not implemented")
// 	// parentObj := &unstructured.Unstructured{}
// 	// parentObj.SetGroupVersionKind(krm.TagsTagKeyGVK)
// 	// key := types.NamespacedName{
// 	// 	Name:      obj.Spec.ParentRef.Name,
// 	// 	Namespace: obj.Spec.ParentRef.Namespace,
// 	// }
// 	// if key.Namespace == "" {
// 	// 	key.Namespace = obj.GetNamespace()
// 	// }
// 	// if err := client.Get(ctx, key, parentObj); err != nil {
// 	// 	return nil, fmt.Errorf("getting parent %v: %w", key, err)
// 	// }
// 	// name, _, err := unstructured.NestedString(parentObj.Object, "status", "name")
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("getting status.name: %w", err)
// 	// }
// 	// if name == "" {
// 	// 	// TODO: Return correct dependency-not-ready value
// 	// 	return nil, fmt.Errorf("not ready")
// 	// }
// 	// external := "tagKeys/" + name
// 	// obj.Spec.ParentRef = v1alpha1.ResourceRef{
// 	// 	External: external,
// 	// }
// }

// func Slice_ToProto[T, U any](ctx *MapContext, in []T, mapper func(ctx *MapContext, in *T) *U) []*U {
// 	if in == nil {
// 		return nil
// 	}

// 	outSlice := make([]*U, 0, len(in))
// 	for _, inItem := range in {
// 		outItem := mapper(ctx, &inItem)
// 		outSlice = append(outSlice, outItem)
// 	}
// 	return outSlice
// }

// func Slice_FromProto[T, U any](ctx *MapContext, in []*T, mapper func(ctx *MapContext, in *T) *U) []U {
// 	if in == nil {
// 		return nil
// 	}

// 	outSlice := make([]U, 0, len(in))
// 	for _, inItem := range in {
// 		outItem := mapper(ctx, inItem)
// 		outSlice = append(outSlice, *outItem)
// 	}
// 	return outSlice
// }

// type ProtoEnum interface {
// 	~int32
// 	Descriptor() protoreflect.EnumDescriptor
// }

// func Enum_ToProto[U ProtoEnum](ctx *MapContext, in *string) U {
// 	var defaultU U
// 	descriptor := defaultU.Descriptor()

// 	inValue := ValueOf(in)
// 	if inValue == "" {
// 		unspecifiedValue := U(int32(0)) // defaultU.New(protoreflect.EnumNumber(0)).(U)
// 		return unspecifiedValue
// 	}

// 	n := descriptor.Values().Len()
// 	for i := 0; i < n; i++ {
// 		value := descriptor.Values().Get(i)
// 		if string(value.Name()) == inValue {
// 			v := U(int32(value.Number()))
// 			return v
// 		}
// 	}

// 	ctx.Errorf("unknown enum value %q", in)
// 	return 0
// }

// func Enum_FromProto[U ProtoEnum](ctx *MapContext, v U) *string {
// 	// if in == nil {
// 	// 	return nil
// 	// }

// 	// v := *in
// 	descriptor := v.Descriptor()

// 	if v == 0 {
// 		return nil
// 	}

// 	val := descriptor.Values().ByNumber(protoreflect.EnumNumber(v))
// 	if val == nil {
// 		ctx.Errorf("unknown enum value %d", v)
// 		return nil
// 	}
// 	s := string(val.Name())
// 	return &s
// }

// func Duration_ToProto(ctx *MapContext, in *string) *durationpb.Duration {
// 	if in == nil {
// 		return nil
// 	}

// 	s := *in
// 	if s == "" {
// 		return nil
// 	}

// 	// TODO: Is this 1:1 with durationpb?
// 	d, err := time.ParseDuration(s)
// 	if err != nil {
// 		ctx.Errorf("parsing duration %q: %w", s, err)
// 		return nil
// 	}
// 	out := durationpb.New(d)
// 	return out
// }

// func Duration_FromProto(ctx *MapContext, in *durationpb.Duration) *string {
// 	if in == nil {
// 		return nil
// 	}

// 	s := in.String()
// 	return &s
// }

// func LazyPtr[V comparable](v V) *V {
// 	var defaultV V
// 	if v == defaultV {
// 		return nil
// 	}
// 	return &v
// }

// func SecondsString_FromProto(ctx *MapContext, in *durationpb.Duration) *string {
// 	if in == nil {
// 		return nil
// 	}
// 	seconds := in.GetSeconds()
// 	out := strconv.FormatInt(seconds, 10)
// 	return &out
// }

// func SecondsString_ToProto(ctx *MapContext, in *string, fieldName string) *durationpb.Duration {
// 	if in == nil {
// 		return nil
// 	}
// 	seconds, err := strconv.ParseInt(*in, 10, 64)
// 	if err != nil {
// 		ctx.Errorf("%s value %q is not valid", fieldName, *in)
// 		return nil
// 	}
// 	out := &durationpb.Duration{Seconds: seconds}
// 	return out
// }
