package mappings

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"unicode"

	"k8s.io/klog/v2"
)

type ValidationError struct {
	Message  string
	Proposal string
}

func (m *Mapping) Validate() []ValidationError {
	var errors []ValidationError
	for _, typeMapping := range m.Mappings {
		checkMissing := (os.Getenv("CHECK_COVERAGE") != "")
		errors = append(errors, typeMapping.Validate(checkMissing)...)
	}
	return errors
}

func (m *TypeMapping) Validate(checkMissing bool) []ValidationError {
	var errors []ValidationError

	inFields := make(map[string]bool)
	for _, mapping := range m.fields {
		if mapping.krmToCloud.InPath == nil {
			continue
		}

		inFields[strings.Join(mapping.krmToCloud.InPath, ".")] = true
		inField := m.inType.LookupField(mapping.krmToCloud.InPath)

		var outField *ReflectField
		if len(mapping.krmToCloud.OutPath) != 0 {
			outField = m.outType.LookupField(mapping.krmToCloud.OutPath)
		}

		if inField == nil {
			err := ValidationError{Message: fmt.Sprintf("field %s not found in input type %v", mapping.krmToCloud.InPath, m.inType)}
			if outField != nil {
				err.Proposal = buildGoField(outField)
			}

			errors = append(errors, err)
			continue
		}
		if len(mapping.krmToCloud.OutPath) != 0 {
			if outField == nil {
				proposal := buildGoField(inField)

				errors = append(errors, ValidationError{
					Message:  fmt.Sprintf("field %s not found in output type %v", mapping.krmToCloud.OutPath, m.outType),
					Proposal: proposal,
				})
				continue
			}
		}
	}

	specField := m.inType.LookupField([]string{"spec"})
	if specField == nil {
		for _, inField := range m.inType.Fields() {
			id := inField.ID()
			if id == "" {
				continue
			}

			if !inFields[id] {
				err := ValidationError{Message: fmt.Sprintf("field %s not mapped from %v to %v", id, m.inType, m.outType)}
				errors = append(errors, err)
				continue
			}
		}
	}

	if !checkMissing {
		return errors
	}

	if m.krmToCloud {
		specField := m.inType.LookupField([]string{"spec"})
		if specField == nil {
			errors = append(errors, ValidationError{
				Message: fmt.Sprintf("spec field not found in KRM type %s", m.inType.String()),
			})
			return errors
		}
		specType := specField.Type()
		statusField := m.inType.LookupField([]string{"status"})
		statusType := statusField.Type()
		// cloudType := m.ResourceCloud

		for _, cloudField := range m.outType.Fields() {
			id := cloudField.ID()

			specField := specType.LookupField([]string{id})
			statusField := statusType.LookupField([]string{id})

			if specField == nil && statusField == nil {
				proposal := buildGoField(cloudField)

				errors = append(errors, ValidationError{
					Message:  fmt.Sprintf("field %s not found in KRM spec %v nor status %v", id, specType, statusType),
					Proposal: proposal,
				})
				continue
			}
		}
	}

	if m.krmToCloud {
		for _, outField := range m.outType.Fields() {
			id := outField.ID()

			inField := m.inType.LookupField([]string{id})

			if inField == nil {
				proposal := buildGoField(outField)

				errors = append(errors, ValidationError{
					Message:  fmt.Sprintf("field %s not found in K%v", id, m.inType),
					Proposal: proposal,
				})
				continue
			}
		}
	}

	return errors
}

func buildGoField(f *ReflectField) string {
	id := f.ID()

	fieldName := jsonToGoFieldName(id)
	fieldType := convertToGoType(f.Type().t)
	jsonTag := id
	jsonTag += ",omitempty"

	requiredTag := ""
	if isMarkedAsRequired(f) {
		requiredTag = "true"
	}

	tags := []string{}
	if jsonTag != "" {
		tags = append(tags, fmt.Sprintf("json:%q", jsonTag))
	}
	if requiredTag != "" {
		tags = append(tags, fmt.Sprintf("required:%q", requiredTag))
	}

	fieldTags := ""
	if len(tags) != 0 {
		fieldTags = " `" + strings.Join(tags, " ") + "`"
	}

	proposal := fmt.Sprintf("%s %s%s", fieldName, fieldType, fieldTags)
	return proposal
}

func isMarkedAsRequired(f *ReflectField) bool {
	// AWS uses required
	requiredTag := f.f.Tag.Get("required")
	switch requiredTag {
	case "true":
		return true

	case "false":
		return false

	case "":
		return false

	default:
		klog.Fatalf("unexpected required value %q", requiredTag)
	}

	return false
}

func jsonToGoFieldName(jsonName string) string {
	var out []rune
	for i, r := range jsonName {
		if i == 0 {
			r = unicode.ToUpper(r)
		}
		out = append(out, r)
	}
	return string(out)
}

func goFieldToFieldID(fieldName string) string {
	var out []rune
	for i, r := range fieldName {
		if i == 0 {
			r = unicode.ToLower(r)
		}
		out = append(out, r)
	}
	s := string(out)
	return s
}

func convertToGoType(t reflect.Type) string {
	fieldGoType := t
	switch fieldGoType.Kind() {
	case reflect.Slice:
		return "[]" + convertToGoType(t.Elem())
	case reflect.Ptr:
		return "*" + convertToGoType(t.Elem())
	case reflect.Struct:
		return t.Name()
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Int64:
		return "int64"
	case reflect.Int:
		return "int" // TODO: Should we warn?
	case reflect.Map:
		return "map[todo]todo"
	default:
		klog.Fatalf("unsupported kind in convertToGoType %v", fieldGoType.Kind())
		return ""
	}
}
