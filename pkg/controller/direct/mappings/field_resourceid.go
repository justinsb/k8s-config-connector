package mappings

type resourceID struct {
	ID string
}

func ResourceID(id string) Field {
	return &resourceID{ID: id}
}

// func (f *resourceID) mapKRMToCloud(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
// 	return &fieldMapping{
// 		InPath:  ParseFieldPath("spec.resourceID"),
// 		OutPath: ParseFieldPath(f.ID),
// 	}
// }

// func (f *resourceID) mapCloudToKRM(cloudType *ReflectType, krmType *ReflectType) *fieldMapping {
// 	return &fieldMapping{
// 		InPath:  ParseFieldPath(f.ID),
// 		OutPath: ParseFieldPath("spec.resourceID"),
// 	}
// }

func (f *resourceID) buildFieldMapping(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
	cloudPath := ParseFieldPath(f.ID)
	krmPath := ParseFieldPath("spec.resourceID")

	m := &fieldMapping{}
	m.krmToCloud.InPath = krmPath
	m.krmToCloud.OutPath = cloudPath
	return m
}
