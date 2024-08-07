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

// QianchuanAdRegionUpdateV10ElectricFenceRegion
type QianchuanAdRegionUpdateV10ElectricFenceRegion int64

// List of qianchuan_ad_region_update_v1.0_electric_fence_region
const (
	Enum_1_QianchuanAdRegionUpdateV10ElectricFenceRegion QianchuanAdRegionUpdateV10ElectricFenceRegion = 1
	Enum_2_QianchuanAdRegionUpdateV10ElectricFenceRegion QianchuanAdRegionUpdateV10ElectricFenceRegion = 2
)

// All allowed values of QianchuanAdRegionUpdateV10ElectricFenceRegion enum
var AllowedQianchuanAdRegionUpdateV10ElectricFenceRegionEnumValues = []QianchuanAdRegionUpdateV10ElectricFenceRegion{
	1,
	2,
}

// NewQianchuanAdRegionUpdateV10ElectricFenceRegionFromValue returns a pointer to a valid QianchuanAdRegionUpdateV10ElectricFenceRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdRegionUpdateV10ElectricFenceRegionFromValue(v int64) (*QianchuanAdRegionUpdateV10ElectricFenceRegion, error) {
	ev := QianchuanAdRegionUpdateV10ElectricFenceRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdRegionUpdateV10ElectricFenceRegion: valid values are %v", v, AllowedQianchuanAdRegionUpdateV10ElectricFenceRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdRegionUpdateV10ElectricFenceRegion) IsValid() bool {
	for _, existing := range AllowedQianchuanAdRegionUpdateV10ElectricFenceRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_region_update_v1.0_electric_fence_region value
func (v QianchuanAdRegionUpdateV10ElectricFenceRegion) Ptr() *QianchuanAdRegionUpdateV10ElectricFenceRegion {
	return &v
}
