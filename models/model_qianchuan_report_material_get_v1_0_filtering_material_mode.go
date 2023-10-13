/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportMaterialGetV10FilteringMaterialMode
type QianchuanReportMaterialGetV10FilteringMaterialMode string

// List of qianchuan_report_material_get_v1.0_filtering_material_mode
const (
	LARGE_QianchuanReportMaterialGetV10FilteringMaterialMode          QianchuanReportMaterialGetV10FilteringMaterialMode = "LARGE"
	LARGE_VERTICAL_QianchuanReportMaterialGetV10FilteringMaterialMode QianchuanReportMaterialGetV10FilteringMaterialMode = "LARGE_VERTICAL"
	SMALL_QianchuanReportMaterialGetV10FilteringMaterialMode          QianchuanReportMaterialGetV10FilteringMaterialMode = "SMALL"
	SQUARE_QianchuanReportMaterialGetV10FilteringMaterialMode         QianchuanReportMaterialGetV10FilteringMaterialMode = "SQUARE"
	VIDEO_LARGE_QianchuanReportMaterialGetV10FilteringMaterialMode    QianchuanReportMaterialGetV10FilteringMaterialMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanReportMaterialGetV10FilteringMaterialMode QianchuanReportMaterialGetV10FilteringMaterialMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanReportMaterialGetV10FilteringMaterialMode enum
var AllowedQianchuanReportMaterialGetV10FilteringMaterialModeEnumValues = []QianchuanReportMaterialGetV10FilteringMaterialMode{
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanReportMaterialGetV10FilteringMaterialModeFromValue returns a pointer to a valid QianchuanReportMaterialGetV10FilteringMaterialMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportMaterialGetV10FilteringMaterialModeFromValue(v string) (*QianchuanReportMaterialGetV10FilteringMaterialMode, error) {
	ev := QianchuanReportMaterialGetV10FilteringMaterialMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportMaterialGetV10FilteringMaterialMode: valid values are %v", v, AllowedQianchuanReportMaterialGetV10FilteringMaterialModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportMaterialGetV10FilteringMaterialMode) IsValid() bool {
	for _, existing := range AllowedQianchuanReportMaterialGetV10FilteringMaterialModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_material_get_v1.0_filtering_material_mode value
func (v QianchuanReportMaterialGetV10FilteringMaterialMode) Ptr() *QianchuanReportMaterialGetV10FilteringMaterialMode {
	return &v
}
