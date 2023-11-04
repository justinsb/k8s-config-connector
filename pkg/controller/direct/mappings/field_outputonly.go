package mappings

type outputOnlyField struct {
	ID string
}

func OutputOnly(id string) Field {
	return &outputOnlyField{ID: id}
}

// func (f *outputOnlyField) mapKRMToCloud(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
// 	return nil
// }

// func (f *outputOnlyField) mapCloudToKRM(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
// 	return &fieldMapping{
// 		InPath:   ParseFieldPath(f.ID),
// 		OutPath:  ParseFieldPath("status." + f.ID),
// 		SetValue: f.setValue,
// 	}
// }

func (f *outputOnlyField) buildFieldMapping(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
	cloudPath := ParseFieldPath(f.ID)
	krmPath := ParseFieldPath("status." + f.ID)

	m := &fieldMapping{}
	m.cloudToKRM.InPath = cloudPath
	m.cloudToKRM.OutPath = krmPath
	m.cloudToKRM.SetValue = f.setValue
	return m
}

func (f *outputOnlyField) setValue(op *mapOp, outPath []string, dest *ReflectValue, inValue *ReflectValue) error {
	root := op.rootOut
	return root.SetValue(op, outPath, inValue)
}
