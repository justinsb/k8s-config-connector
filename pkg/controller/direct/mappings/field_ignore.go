package mappings

type ignoreField struct {
	ID string
}

func Ignore(id string) Field {
	return &ignoreField{ID: id}
}

func TODO(id string) Field {
	return &ignoreField{ID: id}
}

// func (f *ignoreField) mapKRMToCloud(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
// 	return &fieldMapping{
// 		InPath:   ParseFieldPath(f.ID),
// 		SetValue: ignoreValue,
// 	}
// }

// func (f *ignoreField) mapCloudToKRM(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
// 	return &fieldMapping{
// 		InPath:   ParseFieldPath(f.ID),
// 		SetValue: ignoreValue,
// 	}
// }

func (f *ignoreField) buildFieldMapping(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
	p := ParseFieldPath(f.ID)

	m := &fieldMapping{}
	m.krmToCloud.InPath = p
	m.krmToCloud.SetValue = ignoreValue
	return m
}

func ignoreValue(op *mapOp, outPath []string, dest *ReflectValue, inValue *ReflectValue) error {
	return nil
}
