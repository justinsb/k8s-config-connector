package mappings

type specField struct {
	ID string
}

func Spec(id string) Field {
	return &specField{ID: id}
}

// func (f *specField) mapKRMToCloud(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
// 	return &fieldMapping{
// 		InPath:  ParseFieldPath("spec." + f.ID),
// 		OutPath: ParseFieldPath(f.ID),
// 	}
// }

// func (f *specField) mapCloudToKRM(cloudType *ReflectType, krmType *ReflectType) *fieldMapping {
// 	return &fieldMapping{
// 		InPath:  ParseFieldPath(f.ID),
// 		OutPath: ParseFieldPath("spec." + f.ID),
// 	}
// }

func (f *specField) buildFieldMapping(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
	cloudPath := ParseFieldPath(f.ID)
	krmPath := ParseFieldPath("spec." + f.ID)

	m := &fieldMapping{}
	m.krmToCloud.InPath = krmPath
	m.krmToCloud.OutPath = cloudPath
	return m
}
