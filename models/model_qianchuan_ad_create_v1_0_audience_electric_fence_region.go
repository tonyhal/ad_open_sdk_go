/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10AudienceElectricFenceRegion
type QianchuanAdCreateV10AudienceElectricFenceRegion int64

// List of qianchuan_ad_create_v1.0_audience_electric_fence_region
const (
	Enum_1_QianchuanAdCreateV10AudienceElectricFenceRegion QianchuanAdCreateV10AudienceElectricFenceRegion = 1
	Enum_2_QianchuanAdCreateV10AudienceElectricFenceRegion QianchuanAdCreateV10AudienceElectricFenceRegion = 2
)

// All allowed values of QianchuanAdCreateV10AudienceElectricFenceRegion enum
var AllowedQianchuanAdCreateV10AudienceElectricFenceRegionEnumValues = []QianchuanAdCreateV10AudienceElectricFenceRegion{
	1,
	2,
}

// NewQianchuanAdCreateV10AudienceElectricFenceRegionFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceElectricFenceRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceElectricFenceRegionFromValue(v int64) (*QianchuanAdCreateV10AudienceElectricFenceRegion, error) {
	ev := QianchuanAdCreateV10AudienceElectricFenceRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceElectricFenceRegion: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceElectricFenceRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceElectricFenceRegion) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceElectricFenceRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_electric_fence_region value
func (v QianchuanAdCreateV10AudienceElectricFenceRegion) Ptr() *QianchuanAdCreateV10AudienceElectricFenceRegion {
	return &v
}
