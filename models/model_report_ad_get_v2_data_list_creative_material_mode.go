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

// ReportAdGetV2DataListCreativeMaterialMode
type ReportAdGetV2DataListCreativeMaterialMode string

// List of report_ad_get_v2_data_list_creative_material_mode
const (
	STATIC_ASSEMBLE_ReportAdGetV2DataListCreativeMaterialMode ReportAdGetV2DataListCreativeMaterialMode = "STATIC_ASSEMBLE"
	INTERVENE_ReportAdGetV2DataListCreativeMaterialMode       ReportAdGetV2DataListCreativeMaterialMode = "INTERVENE"
	CTR_ReportAdGetV2DataListCreativeMaterialMode             ReportAdGetV2DataListCreativeMaterialMode = "CTR"
)

// All allowed values of ReportAdGetV2DataListCreativeMaterialMode enum
var AllowedReportAdGetV2DataListCreativeMaterialModeEnumValues = []ReportAdGetV2DataListCreativeMaterialMode{
	"STATIC_ASSEMBLE",
	"INTERVENE",
	"CTR",
}

// NewReportAdGetV2DataListCreativeMaterialModeFromValue returns a pointer to a valid ReportAdGetV2DataListCreativeMaterialMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2DataListCreativeMaterialModeFromValue(v string) (*ReportAdGetV2DataListCreativeMaterialMode, error) {
	ev := ReportAdGetV2DataListCreativeMaterialMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2DataListCreativeMaterialMode: valid values are %v", v, AllowedReportAdGetV2DataListCreativeMaterialModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2DataListCreativeMaterialMode) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2DataListCreativeMaterialModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_data_list_creative_material_mode value
func (v ReportAdGetV2DataListCreativeMaterialMode) Ptr() *ReportAdGetV2DataListCreativeMaterialMode {
	return &v
}
