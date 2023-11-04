package mappings

import (
	"reflect"
	"strings"
)

type TypeMapping struct {
	inType  *ReflectType
	outType *ReflectType

	krmToCloud bool

	fields []*fieldMapping
}

type fieldMapping struct {
	krmToCloud fieldMeta
	cloudToKRM fieldMeta
}

type fieldMeta struct {
	InPath  []string
	OutPath []string

	SetValue func(op *mapOp, outPath []string, dest *ReflectValue, inValue *ReflectValue) error
}

func ParseFieldPath(s string) []string {
	return strings.Split(s, ".")
}

func (m *TypeMapping) Map(op *mapOp, in any, out any) error {
	// specField := m.ResourceKRM.LookupField("spec")
	// specType := specField.Type()

	// cloudType := m.ResourceCloud

	var inVal *ReflectValue
	if v, ok := in.(*ReflectValue); ok {
		inVal = v
	} else {
		inVal = &ReflectValue{reflect.ValueOf(in)}
	}

	var outVal *ReflectValue
	if v, ok := out.(*ReflectValue); ok {
		outVal = v
	} else {
		outVal = &ReflectValue{reflect.ValueOf(out)}
	}

	// outType := outVal.Type()

	// specField := inVal.Type().LookupField("spec")
	// v, err := inVal.GetField(specField)
	// if err != nil {
	// 	return err
	// }
	// inVal = v

	// outVal := &ReflectValue{reflect.ValueOf(out)}

	for _, mapping := range m.fields {
		// inField := inType.LookupField(mapping.ID)
		// if inField == nil {
		// 	return fmt.Errorf("field %s not found in type %v", mapping.ID, inType)
		// }

		// outField := outType.LookupField(mapping.ID)
		// if outField == nil {
		// 	return fmt.Errorf("field %s not found in type %v", mapping.ID, outType)
		// }

		meta := mapping.cloudToKRM
		if op.krmToCloud {
			meta = mapping.krmToCloud
		}
		v, err := inVal.GetValue(meta.InPath)
		if err != nil {
			return err
		}

		if meta.SetValue != nil {
			if err := meta.SetValue(op, meta.OutPath, outVal, v); err != nil {
				return err
			}
		} else {
			if err := outVal.SetValue(op, meta.OutPath, v); err != nil {
				return err
			}
		}

	}

	return nil
}
