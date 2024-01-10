/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCreativeGetV10DataListCreativeMaterialMode
type QianchuanCreativeGetV10DataListCreativeMaterialMode string

// List of qianchuan_creative_get_v1.0_data_list_creative_material_mode
const (
	CUSTOM_CREATIVE_QianchuanCreativeGetV10DataListCreativeMaterialMode       QianchuanCreativeGetV10DataListCreativeMaterialMode = "CUSTOM_CREATIVE"
	PROGRAMMATIC_CREATIVE_QianchuanCreativeGetV10DataListCreativeMaterialMode QianchuanCreativeGetV10DataListCreativeMaterialMode = "PROGRAMMATIC_CREATIVE"
)

// All allowed values of QianchuanCreativeGetV10DataListCreativeMaterialMode enum
var AllowedQianchuanCreativeGetV10DataListCreativeMaterialModeEnumValues = []QianchuanCreativeGetV10DataListCreativeMaterialMode{
	"CUSTOM_CREATIVE",
	"PROGRAMMATIC_CREATIVE",
}

// NewQianchuanCreativeGetV10DataListCreativeMaterialModeFromValue returns a pointer to a valid QianchuanCreativeGetV10DataListCreativeMaterialMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10DataListCreativeMaterialModeFromValue(v string) (*QianchuanCreativeGetV10DataListCreativeMaterialMode, error) {
	ev := QianchuanCreativeGetV10DataListCreativeMaterialMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10DataListCreativeMaterialMode: valid values are %v", v, AllowedQianchuanCreativeGetV10DataListCreativeMaterialModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10DataListCreativeMaterialMode) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10DataListCreativeMaterialModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_data_list_creative_material_mode value
func (v QianchuanCreativeGetV10DataListCreativeMaterialMode) Ptr() *QianchuanCreativeGetV10DataListCreativeMaterialMode {
	return &v
}
