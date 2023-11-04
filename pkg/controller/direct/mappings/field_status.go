package mappings

type statusField struct {
	ID string
}

func Status(id string) Field {
	return &statusField{ID: id}
}

// func (f *statusField) mapKRMToCloud(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
// 	return &fieldMapping{
// 		InPath:  ParseFieldPath("status." + f.ID),
// 		OutPath: ParseFieldPath(f.ID),
// 	}
// }

//	func (f *statusField) mapCloudToKRM(cloudType *ReflectType, krmType *ReflectType) *fieldMapping {
//		return &fieldMapping{
//			InPath:  ParseFieldPath(f.ID),
//			OutPath: ParseFieldPath("status." + f.ID),
//		}
//	}
func (f *statusField) buildFieldMapping(krmType *ReflectType, cloudType *ReflectType) *fieldMapping {
	cloudPath := ParseFieldPath(f.ID)
	krmPath := ParseFieldPath("status." + f.ID)

	m := &fieldMapping{}
	m.cloudToKRM.InPath = cloudPath
	m.cloudToKRM.OutPath = krmPath
	return m
}
