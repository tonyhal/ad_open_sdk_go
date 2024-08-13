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

// QianchuanReportMaterialGetV10DataListMaterialMode
type QianchuanReportMaterialGetV10DataListMaterialMode string

// List of qianchuan_report_material_get_v1.0_data_list_material_mode
const (
	LARGE_QianchuanReportMaterialGetV10DataListMaterialMode          QianchuanReportMaterialGetV10DataListMaterialMode = "LARGE"
	LARGE_VERTICAL_QianchuanReportMaterialGetV10DataListMaterialMode QianchuanReportMaterialGetV10DataListMaterialMode = "LARGE_VERTICAL"
	SMALL_QianchuanReportMaterialGetV10DataListMaterialMode          QianchuanReportMaterialGetV10DataListMaterialMode = "SMALL"
	SQUARE_QianchuanReportMaterialGetV10DataListMaterialMode         QianchuanReportMaterialGetV10DataListMaterialMode = "SQUARE"
	UNION_SPLASH_QianchuanReportMaterialGetV10DataListMaterialMode   QianchuanReportMaterialGetV10DataListMaterialMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanReportMaterialGetV10DataListMaterialMode    QianchuanReportMaterialGetV10DataListMaterialMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanReportMaterialGetV10DataListMaterialMode QianchuanReportMaterialGetV10DataListMaterialMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanReportMaterialGetV10DataListMaterialMode enum
var AllowedQianchuanReportMaterialGetV10DataListMaterialModeEnumValues = []QianchuanReportMaterialGetV10DataListMaterialMode{
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanReportMaterialGetV10DataListMaterialModeFromValue returns a pointer to a valid QianchuanReportMaterialGetV10DataListMaterialMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportMaterialGetV10DataListMaterialModeFromValue(v string) (*QianchuanReportMaterialGetV10DataListMaterialMode, error) {
	ev := QianchuanReportMaterialGetV10DataListMaterialMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportMaterialGetV10DataListMaterialMode: valid values are %v", v, AllowedQianchuanReportMaterialGetV10DataListMaterialModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportMaterialGetV10DataListMaterialMode) IsValid() bool {
	for _, existing := range AllowedQianchuanReportMaterialGetV10DataListMaterialModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_material_get_v1.0_data_list_material_mode value
func (v QianchuanReportMaterialGetV10DataListMaterialMode) Ptr() *QianchuanReportMaterialGetV10DataListMaterialMode {
	return &v
}
