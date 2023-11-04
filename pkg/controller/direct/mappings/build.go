package mappings

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/assert"
	"k8s.io/klog/v2"
)

type Mapping struct {
	ResourceCloudType *ReflectType
	ResourceKRMType   *ReflectType

	Mappings []*TypeMapping
}

func (m *Mapping) MapKRMToCloud(operation Operation, in any, out any) error {
	return m.mapRoot(operation, true, in, out)
}

func (m *Mapping) MapCloudToKRM(operation Operation, in any, out any) error {
	return m.mapRoot(operation, true, in, out)
}

func (m *Mapping) mapRoot(operation Operation, krmToCloud bool, in any, out any) error {
	inTypeKey := typeOf(in).String()
	outTypeKey := typeOf(out).String()
	for _, typeMapping := range m.Mappings {
		if typeMapping.inType.String() != inTypeKey {
			continue
		}

		if typeMapping.outType.String() != outTypeKey {
			continue
		}

		op := &mapOp{
			Operation:  operation,
			mapping:    m,
			krmToCloud: krmToCloud,
			rootIn:     &ReflectValue{reflect.ValueOf(in)},
			rootOut:    &ReflectValue{reflect.ValueOf(out)},
		}

		return typeMapping.Map(op, op.rootIn, op.rootOut)
	}

	for _, typeMapping := range m.Mappings {
		klog.Infof("mapping from %q -> %q", typeMapping.inType.String(), typeMapping.outType.String())
	}

	return fmt.Errorf("type mapping not found for %q -> %q", inTypeKey, outTypeKey)
}

func (m *Mapping) nestedMap(op *mapOp, in *ReflectValue, out *ReflectValue) error {

	inTypeKey := typeOf(in).String()
	outTypeKey := typeOf(out).String()
	for _, typeMapping := range m.Mappings {
		if typeMapping.inType.String() != inTypeKey {
			continue
		}

		if typeMapping.outType.String() != outTypeKey {
			continue
		}

		return typeMapping.Map(op, op.rootIn, op.rootOut)
	}

	return fmt.Errorf("type mapping not found for %q -> %q", inTypeKey, outTypeKey)
}

type Field interface {
	buildFieldMapping(krmType *ReflectType, cloudType *ReflectType) *fieldMapping
	// 	mapCloudToKRM(cloudType *ReflectType, krmType *ReflectType) *fieldMapping
	// mapKRMToCloud(krmType *ReflectType, cloudType *ReflectType) *fieldMapping
}

type MappingBuilder struct {
	mapping *Mapping
	errors  []error
}

func (b *MappingBuilder) Build() (*Mapping, error) {
	if len(b.errors) != 0 {
		return nil, errors.Join(b.errors...)
	}

	if assert.Enabled {
		errs := b.mapping.Validate()
		if len(errs) != 0 {
			for _, err := range errs {
				fmt.Fprintf(os.Stderr, "%v\n", err.Message)
				if err.Proposal != "" {
					fmt.Fprintf(os.Stderr, "    %v\n", err.Proposal)
				}
			}
			assert.Fail()
		}
	}

	return b.mapping, nil
}

func (b *MappingBuilder) MustBuild() *Mapping {
	m, err := b.Build()
	if err != nil {
		klog.Fatalf("error building mapping: %v", err)
	}
	return m
}

func NewMapping(krmObj any, cloudObj any, fields ...any) *MappingBuilder {
	resourceCloudType := typeOf(cloudObj)
	resourceKRMType := typeOf(krmObj)

	m := &Mapping{
		ResourceCloudType: resourceCloudType,
		ResourceKRMType:   resourceKRMType,
	}

	b := &MappingBuilder{
		mapping: m,
	}

	b = b.mapKRMToCloud(resourceKRMType, resourceCloudType, fields...)

	return b

}

func (b *MappingBuilder) WithCreate(out any, fields ...any) *MappingBuilder {
	return b.mapKRMToCloud(b.mapping.ResourceKRMType, typeOf(out), fields...)
}

func (b *MappingBuilder) WithDelete(out any, fields ...any) *MappingBuilder {
	return b.mapKRMToCloud(b.mapping.ResourceKRMType, typeOf(out), fields...)
}

func (b *MappingBuilder) MapType(krmObj any, cloudObj any, fields ...any) *MappingBuilder {
	cloudType := typeOf(cloudObj)
	krmType := typeOf(krmObj)

	b = b.mapKRMToCloud(krmType, cloudType, fields...)
	// b = b.mapCloudToKRM(cloudType, krmType, fields...)
	return b
}

func (b *MappingBuilder) mapKRMToCloud(krmType *ReflectType, cloudType *ReflectType, fields ...any) *MappingBuilder {

	// specType := b.mapping.ResourceKRMType.LookupField("spec").Type()
	// statusType := b.mapping.ResourceKRMType.LookupField("status").Type()

	createMapping := &TypeMapping{
		inType:     krmType,
		outType:    cloudType,
		krmToCloud: true,
	}
	b.mapping.Mappings = append(b.mapping.Mappings, createMapping)

	for _, field := range fields {
		switch field := field.(type) {
		case string:
			fieldPath := ParseFieldPath(field)

			m := &fieldMapping{}
			m.krmToCloud.InPath = fieldPath
			m.krmToCloud.OutPath = fieldPath
			m.cloudToKRM.InPath = fieldPath
			m.cloudToKRM.OutPath = fieldPath

			createMapping.fields = append(createMapping.fields, m)

		case Field:
			m := field.buildFieldMapping(krmType, cloudType)
			createMapping.fields = append(createMapping.fields, m)

		default:
			klog.Fatalf("unhandled field type %T", field)
		}
	}

	return b
}

// func (b *MappingBuilder) mapCloudToKRM(cloudType *ReflectType, krmType *ReflectType, fields ...any) *MappingBuilder {

// 	// specType := b.mapping.ResourceKRMType.LookupField("spec").Type()
// 	// statusType := b.mapping.ResourceKRMType.LookupField("status").Type()

// 	createMapping := &TypeMapping{
// 		parent:  b.mapping,
// 		inType:  cloudType,
// 		outType: krmType,
// 	}
// 	b.mapping.Mappings = append(b.mapping.Mappings, createMapping)

// 	for _, field := range fields {
// 		switch field := field.(type) {
// 		case string:
// 			createMapping.fields = append(createMapping.fields, &fieldMapping{
// 				InPath:  ParseFieldPath(field),
// 				OutPath: ParseFieldPath(field),
// 			})
// 		case Field:
// 			m := field.mapCloudToKRM(krmType, cloudType)
// 			if m != nil {
// 				createMapping.fields = append(createMapping.fields, m)
// 			}

// 		default:
// 			klog.Fatalf("unhandled field type %T", field)
// 		}
// 	}

// 	return b
// }
