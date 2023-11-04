package mappings

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"k8s.io/klog/v2"
)

type ReflectType struct {
	t reflect.Type
}

func typeOf(o any) *ReflectType {
	rv, ok := o.(reflect.Value)
	if !ok {
		rv = reflect.ValueOf(o)
	}

	return &ReflectType{rv.Type()}
}

func (t *ReflectType) String() string {
	prefix := ""
	reflectType := t.t

	for {
		if reflectType.Kind() == reflect.Ptr {
			prefix = "*" + prefix
			reflectType = reflectType.Elem()
		} else if reflectType.Kind() == reflect.Slice {
			prefix = "[]" + prefix
			reflectType = reflectType.Elem()
		} else {
			break
		}
	}

	return prefix + reflectType.PkgPath() + "." + reflectType.Name()
}

func getFieldID(f *reflect.StructField) string {
	jsonTag := getJSONFieldTag(f)

	if strings.HasSuffix(jsonTag, "Ids") {
		jsonTag = strings.TrimSuffix(jsonTag, "Ids") + "IDs"
	}

	if strings.HasSuffix(jsonTag, "Arn") {
		jsonTag = strings.TrimSuffix(jsonTag, "Arn") + "ARN"
	}

	var words []string
	var sb strings.Builder
	for _, r := range jsonTag {
		if unicode.IsUpper(r) {
			words = append(words, sb.String())
			sb.Reset()
		}
		sb.WriteRune(r)
	}
	words = append(words, sb.String())

	for i, s := range words {
		switch s {
		case "Id":
			words[i] = "ID"
		case "Vpc":
			words[i] = "VPC"
		case "Cidr":
			words[i] = "CIDR"
		}
	}
	return strings.Join(words, "")
}

func getJSONFieldTag(f *reflect.StructField) string {
	jsonTag := f.Tag.Get("json")
	switch jsonTag {
	case "-":
		// Not marshaled
		return ""

	case "":
		break

	default:
		parts := strings.Split(jsonTag, ",")
		name := parts[0]
		if name == "" {
			name = f.Name
		}
		return name
	}

	// locationName is used by AWS types
	// TODO: Make this conditions?
	locationName := f.Tag.Get("locationName")
	if locationName != "" {
		return locationName
	}
	name := f.Name
	return goFieldToFieldID(name)
}

func (t *ReflectType) findField(id string) *ReflectField {
	reflectType := t.t
	if reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	n := reflectType.NumField()
	for i := 0; i < n; i++ {
		f := reflectType.Field(i)
		fieldID := getFieldID(&f)
		if fieldID == id {
			return &ReflectField{&f}
		}
	}
	return nil
}

func (t *ReflectType) LookupField(fieldPath []string) *ReflectField {
	f := t.findField(fieldPath[0])
	if f == nil {
		return nil
	}
	if len(fieldPath) == 1 {
		return f
	}
	return f.Type().LookupField(fieldPath[1:])
}

func (t *ReflectType) Fields() []*ReflectField {
	var out []*ReflectField
	reflectType := t.t
	if reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	n := reflectType.NumField()
	for i := 0; i < n; i++ {
		f := reflectType.Field(i)
		out = append(out, &ReflectField{&f})
	}
	return out
}

type ReflectField struct {
	f *reflect.StructField
}

func (f *ReflectField) ID() string {
	fieldID := getFieldID(f.f)
	return fieldID
}

func (f *ReflectField) Type() *ReflectType {
	return &ReflectType{t: f.f.Type}
}

// func (v *ReflectValue) GetField(f *ReflectField) (*ReflectValue, error) {
// 	sv := v.v
// 	if sv.Kind() == reflect.Ptr {
// 		sv = sv.Elem()
// 	}
// 	rv := sv.FieldByIndex(f.f.Index)
// 	return &ReflectValue{v: rv}, nil
// }

// func (v *ReflectValue) SetField(f *ReflectField, arg *ReflectValue) error {
// 	sv := v.v
// 	if sv.Kind() == reflect.Ptr {
// 		sv = sv.Elem()
// 	}
// 	field := sv.FieldByIndex(f.f.Index)
// 	field.Set(arg.v)
// 	return nil
// }

func (v *ReflectValue) Type() *ReflectType {
	t := v.v.Type()
	return &ReflectType{t: t}
}

type ReflectValue struct {
	v reflect.Value
}

func (v *ReflectValue) GetValue(fieldPath []string) (*ReflectValue, error) {
	sv := v.v
	if sv.Kind() == reflect.Ptr {
		sv = sv.Elem()
	}
	f := v.Type().findField(fieldPath[0])
	if f == nil {
		return nil, fmt.Errorf("unable to find field %q in %v", fieldPath[0], v.Type())
	}
	rv := sv.FieldByIndex(f.f.Index)
	out := &ReflectValue{v: rv}
	if len(fieldPath) > 1 {
		return out.GetValue(fieldPath[1:])
	}
	return out, nil
}

func (v *ReflectValue) SetValue(op *mapOp, fieldPath []string, arg *ReflectValue) error {
	sv := v.v
	if sv.Kind() == reflect.Ptr {
		sv = sv.Elem()
	}
	f := v.Type().findField(fieldPath[0])
	if f == nil {
		return fmt.Errorf("unable to find field %q in %v", fieldPath[0], v.Type())
	}
	rv := sv.FieldByIndex(f.f.Index)
	if len(fieldPath) == 1 {
		destVal, err := op.asTargetValue(rv.Type(), arg.v)
		if err != nil {
			return err
		}
		if !destVal.IsValid() {
			return nil
		}
		rv.Set(destVal)
		return nil
	}

	out := &ReflectValue{v: rv}
	return out.SetValue(op, fieldPath[1:], arg)
}

func (op *mapOp) asTargetValue(destType reflect.Type, src reflect.Value) (reflect.Value, error) {
	if src.Kind() == reflect.Pointer {
		if src.IsNil() {
			// Nothing to set
			return reflect.Value{}, nil
		}
		src = src.Elem()
	}

	if src.Kind() == reflect.Slice {
		if src.IsNil() {
			// Nothing to set
			return reflect.Value{}, nil
		}
		dest := reflect.New(destType).Elem()
		n := src.Len()
		for i := 0; i < n; i++ {
			srcElem := src.Index(i)
			destElem, err := op.asTargetValue(destType.Elem(), srcElem)
			if err != nil {
				return reflect.Value{}, err
			}
			// TODO: What if destElem is not valid
			dest = reflect.Append(dest, destElem)
		}
		klog.Infof("copied slice %v -> %v", src.Interface(), dest.Interface())
		return dest, nil
	}

	srcType := src.Type()

	switch srcType.String() {
	case "string":
		v := src.String()
		switch destType.String() {
		case "string":
			return reflect.ValueOf(v), nil
		case "*string":
			return reflect.ValueOf(&v), nil
		}
	}

	if src.CanInterface() {
		srcVal := src
		if srcVal.Kind() == reflect.Struct {
			srcVal = srcVal.Addr()
		}
		destVal := reflect.New(destType.Elem())
		if err := op.mapping.nestedMap(op, &ReflectValue{srcVal}, &ReflectValue{destVal}); err != nil {
			return reflect.Value{}, err
		}
		return destVal, nil
	}

	// if dest.Kind() == reflect.Ptr {
	// 	destType := dest.Type().Elem()
	// 	srcType := src.Type()

	// 	switch srcType.String() {
	// 	// case reflect.Ptr:
	// 	// 	dest.Set(src)
	// 	// 	return nil
	// 	case "*string":
	// 		if src.IsNil() {
	// 			return nil
	// 		}
	// 		s := src.Elem().String()
	// 		switch destType.String() {
	// 			case
	// 		dest.Set(reflect.ValueOf(&s))
	// 		return nil
	// 	default:
	// 		return fmt.Errorf("conversion from %v to %v not implemented", srcType.String(), destType.String())
	// 	}
	// }
	// if dest.Kind() == reflect.String {
	// 	if src.Kind() == reflect.Ptr {
	// 		src = src.Elem()
	// 	}
	// 	switch src.Kind() {
	// 	case reflect.String:
	// 		dest.Set(src)
	// 		return nil
	// 	default:
	// 		return fmt.Errorf("conversion from %v to %v not implemented", dest.Type(), src.Type())
	// 	}
	// }
	return reflect.Value{}, fmt.Errorf("conversion from %v to %v not implemented", srcType.String(), destType.String())
}
