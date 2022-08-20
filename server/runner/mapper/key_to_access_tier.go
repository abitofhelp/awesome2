package mapper

import (
	"fmt"
	enumsv1 "github.com/abitofhelp/awesome/gen/go/awesome/v1/enums"
	enumsv21 "github.com/abitofhelp/awesome2/gen/go/awesome2/v1/enums"
)

//func Enumsv1AccessTierToEnums1V1AccessTier(pb enumsv1.AccessTier) (enumsv21.AccessTier, error) {
//
//	var value enumsv21.AccessTier
//	if _, found := enumsv1.AccessTier_name[int32(pb.Number())]; found {
//		value = enumsv21.AccessTier(int32(pb.Number()))
//	} else {
//		return enumsv21.AccessTier_ACCESS_TIER_UNSPECIFIED, fmt.Errorf("\nfailed to find an enumsv1.AccessTier named '%s'", pb.String())
//	}
//
//	return value, nil
//}

func KeyToAccessTier(key string) (enumsv21.AccessTier, error) {

	var value enumsv21.AccessTier
	if v, found := enumsv1.AccessTier_value[key]; found {
		value = enumsv21.AccessTier(v)
	} else {
		return enumsv21.AccessTier_ACCESS_TIER_UNSPECIFIED, fmt.Errorf("\nfailed to find an enumsv1.AccessTier named '%s'", key)
	}

	return value, nil
}
