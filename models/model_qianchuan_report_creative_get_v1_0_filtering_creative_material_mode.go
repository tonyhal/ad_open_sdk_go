/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportCreativeGetV10FilteringCreativeMaterialMode
type QianchuanReportCreativeGetV10FilteringCreativeMaterialMode string

// List of qianchuan_report_creative_get_v1.0_filtering_creative_material_mode
const (
	CUSTOM_CREATIVE_QianchuanReportCreativeGetV10FilteringCreativeMaterialMode       QianchuanReportCreativeGetV10FilteringCreativeMaterialMode = "CUSTOM_CREATIVE"
	PROGRAMMATIC_CREATIVE_QianchuanReportCreativeGetV10FilteringCreativeMaterialMode QianchuanReportCreativeGetV10FilteringCreativeMaterialMode = "PROGRAMMATIC_CREATIVE"
)

// All allowed values of QianchuanReportCreativeGetV10FilteringCreativeMaterialMode enum
var AllowedQianchuanReportCreativeGetV10FilteringCreativeMaterialModeEnumValues = []QianchuanReportCreativeGetV10FilteringCreativeMaterialMode{
	"CUSTOM_CREATIVE",
	"PROGRAMMATIC_CREATIVE",
}

// NewQianchuanReportCreativeGetV10FilteringCreativeMaterialModeFromValue returns a pointer to a valid QianchuanReportCreativeGetV10FilteringCreativeMaterialMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportCreativeGetV10FilteringCreativeMaterialModeFromValue(v string) (*QianchuanReportCreativeGetV10FilteringCreativeMaterialMode, error) {
	ev := QianchuanReportCreativeGetV10FilteringCreativeMaterialMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportCreativeGetV10FilteringCreativeMaterialMode: valid values are %v", v, AllowedQianchuanReportCreativeGetV10FilteringCreativeMaterialModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportCreativeGetV10FilteringCreativeMaterialMode) IsValid() bool {
	for _, existing := range AllowedQianchuanReportCreativeGetV10FilteringCreativeMaterialModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_creative_get_v1.0_filtering_creative_material_mode value
func (v QianchuanReportCreativeGetV10FilteringCreativeMaterialMode) Ptr() *QianchuanReportCreativeGetV10FilteringCreativeMaterialMode {
	return &v
}
