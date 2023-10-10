/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanOrientationPackageGetV10DataListAge
type QianchuanOrientationPackageGetV10DataListAge string

// List of qianchuan_orientation_package_get_v1.0_data_list_age
const (
	AGE_ABOVE_50_QianchuanOrientationPackageGetV10DataListAge      QianchuanOrientationPackageGetV10DataListAge = "AGE_ABOVE_50"
	AGE_BETWEEN_18_23_QianchuanOrientationPackageGetV10DataListAge QianchuanOrientationPackageGetV10DataListAge = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_24_30_QianchuanOrientationPackageGetV10DataListAge QianchuanOrientationPackageGetV10DataListAge = "AGE_BETWEEN_24_30"
	AGE_BETWEEN_31_40_QianchuanOrientationPackageGetV10DataListAge QianchuanOrientationPackageGetV10DataListAge = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_41_49_QianchuanOrientationPackageGetV10DataListAge QianchuanOrientationPackageGetV10DataListAge = "AGE_BETWEEN_41_49"
)

// All allowed values of QianchuanOrientationPackageGetV10DataListAge enum
var AllowedQianchuanOrientationPackageGetV10DataListAgeEnumValues = []QianchuanOrientationPackageGetV10DataListAge{
	"AGE_ABOVE_50",
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_24_30",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_41_49",
}

// NewQianchuanOrientationPackageGetV10DataListAgeFromValue returns a pointer to a valid QianchuanOrientationPackageGetV10DataListAge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanOrientationPackageGetV10DataListAgeFromValue(v string) (*QianchuanOrientationPackageGetV10DataListAge, error) {
	ev := QianchuanOrientationPackageGetV10DataListAge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanOrientationPackageGetV10DataListAge: valid values are %v", v, AllowedQianchuanOrientationPackageGetV10DataListAgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanOrientationPackageGetV10DataListAge) IsValid() bool {
	for _, existing := range AllowedQianchuanOrientationPackageGetV10DataListAgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_orientation_package_get_v1.0_data_list_age value
func (v QianchuanOrientationPackageGetV10DataListAge) Ptr() *QianchuanOrientationPackageGetV10DataListAge {
	return &v
}
