package mappings

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

type refField struct {
	ID string
}

func Ref(id string) Field {
	return &refField{ID: id}
}

// func (f *refField) mapKRMToCloud(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
// 	return &fieldMapping{
// 		InPath:   ParseFieldPath(f.ID + "Ref"),
// 		OutPath:  ParseFieldPath(f.ID),
// 		SetValue: f.setKRMToCloud,
// 	}
// }

// func (f *refField) mapCloudToKRM(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
// 	return nil
// }

func (f *refField) buildFieldMapping(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
	cloudPath := ParseFieldPath(f.ID)
	krmPath := ParseFieldPath(f.ID + "Ref")

	m := &fieldMapping{}
	m.krmToCloud.InPath = krmPath
	m.krmToCloud.OutPath = cloudPath
	m.krmToCloud.SetValue = f.setKRMToCloud
	return m
}

func (f *refField) setKRMToCloud(op *mapOp, outPath []string, dest *ReflectValue, inValue *ReflectValue) error {
	outValue, err := f.resolveValue(op, inValue)
	if err != nil {
		return err
	}

	return dest.SetValue(op, outPath, outValue)
}

func (f *refField) resolveValue(op *mapOp, inValue *ReflectValue) (*ReflectValue, error) {
	switch v := inValue.v.Interface().(type) {
	case *v1alpha1.ResourceRef:
		if v.External != "" {
			return &ReflectValue{v: reflect.ValueOf(v.External)}, nil
		}

		var kccOp *KCCOperation
		if err := op.Operation.Get(&kccOp); err != nil {
			return nil, err
		}

		k8s := kccOp.Client
		u := &unstructured.Unstructured{}

		kind := v.Kind
		if kind == "" {
			// TODO: Default kind
			kind = "IAMServiceAccount"
		}
		switch kind {
		case "IAMServiceAccount":
			u.SetGroupVersionKind(schema.GroupVersionKind{Kind: kind, Group: "iam.cnrm.cloud.google.com", Version: "v1beta1"})
			id := types.NamespacedName{Name: v.Name, Namespace: kccOp.Namespace}

			// TODO: Allow specification of namespace
			// TODO: validate name is set

			if err := k8s.Get(kccOp.Ctx, id, u); err != nil {
				// TODO: Dependency not found
				return nil, fmt.Errorf("reading IAMServiceAccount %v: %w", id, err)
			}

			email, _, _ := unstructured.NestedString(u.Object, "status", "email")
			if email == "" {
				// TODO: Dependency not ready
				return nil, fmt.Errorf("IAMServiceAccount %v did not yet have status.email value", id)
			}

			return &ReflectValue{v: reflect.ValueOf(email)}, nil
		default:
			return nil, fmt.Errorf("unhandled kind %q for reference", kind)
		}
	default:
		return nil, fmt.Errorf("unhandled type in refField::transformKRMToCloud: %T", v)
	}
}
