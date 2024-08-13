/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit
type NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit string

// List of native_anchor_create_v3.0_anchor_info_game_anchor_game_package_list_gift_gift_unit
const (
	INDIVIDUAL_NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit = "INDIVIDUAL"
	MYRIAD_NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit     NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit = "MYRIAD"
)

// All allowed values of NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit enum
var AllowedNativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnitEnumValues = []NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit{
	"INDIVIDUAL",
	"MYRIAD",
}

// NewNativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnitFromValue returns a pointer to a valid NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnitFromValue(v string) (*NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit, error) {
	ev := NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit: valid values are %v", v, AllowedNativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit) IsValid() bool {
	for _, existing := range AllowedNativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_create_v3.0_anchor_info_game_anchor_game_package_list_gift_gift_unit value
func (v NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit) Ptr() *NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit {
	return &v
}
