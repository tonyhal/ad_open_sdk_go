/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCreativeGetV10FilteringCreativeMaterialMode
type QianchuanCreativeGetV10FilteringCreativeMaterialMode string

// List of qianchuan_creative_get_v1.0_filtering_creative_material_mode
const (
	CUSTOM_CREATIVE_QianchuanCreativeGetV10FilteringCreativeMaterialMode       QianchuanCreativeGetV10FilteringCreativeMaterialMode = "CUSTOM_CREATIVE"
	PROGRAMMATIC_CREATIVE_QianchuanCreativeGetV10FilteringCreativeMaterialMode QianchuanCreativeGetV10FilteringCreativeMaterialMode = "PROGRAMMATIC_CREATIVE"
)

// All allowed values of QianchuanCreativeGetV10FilteringCreativeMaterialMode enum
var AllowedQianchuanCreativeGetV10FilteringCreativeMaterialModeEnumValues = []QianchuanCreativeGetV10FilteringCreativeMaterialMode{
	"CUSTOM_CREATIVE",
	"PROGRAMMATIC_CREATIVE",
}

// NewQianchuanCreativeGetV10FilteringCreativeMaterialModeFromValue returns a pointer to a valid QianchuanCreativeGetV10FilteringCreativeMaterialMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10FilteringCreativeMaterialModeFromValue(v string) (*QianchuanCreativeGetV10FilteringCreativeMaterialMode, error) {
	ev := QianchuanCreativeGetV10FilteringCreativeMaterialMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10FilteringCreativeMaterialMode: valid values are %v", v, AllowedQianchuanCreativeGetV10FilteringCreativeMaterialModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10FilteringCreativeMaterialMode) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10FilteringCreativeMaterialModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_filtering_creative_material_mode value
func (v QianchuanCreativeGetV10FilteringCreativeMaterialMode) Ptr() *QianchuanCreativeGetV10FilteringCreativeMaterialMode {
	return &v
}
