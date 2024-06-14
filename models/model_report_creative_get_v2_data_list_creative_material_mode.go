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

// ReportCreativeGetV2DataListCreativeMaterialMode
type ReportCreativeGetV2DataListCreativeMaterialMode string

// List of report_creative_get_v2_data_list_creative_material_mode
const (
	STATIC_ASSEMBLE_ReportCreativeGetV2DataListCreativeMaterialMode ReportCreativeGetV2DataListCreativeMaterialMode = "STATIC_ASSEMBLE"
	INTERVENE_ReportCreativeGetV2DataListCreativeMaterialMode       ReportCreativeGetV2DataListCreativeMaterialMode = "INTERVENE"
	CTR_ReportCreativeGetV2DataListCreativeMaterialMode             ReportCreativeGetV2DataListCreativeMaterialMode = "CTR"
)

// All allowed values of ReportCreativeGetV2DataListCreativeMaterialMode enum
var AllowedReportCreativeGetV2DataListCreativeMaterialModeEnumValues = []ReportCreativeGetV2DataListCreativeMaterialMode{
	"STATIC_ASSEMBLE",
	"INTERVENE",
	"CTR",
}

// NewReportCreativeGetV2DataListCreativeMaterialModeFromValue returns a pointer to a valid ReportCreativeGetV2DataListCreativeMaterialMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2DataListCreativeMaterialModeFromValue(v string) (*ReportCreativeGetV2DataListCreativeMaterialMode, error) {
	ev := ReportCreativeGetV2DataListCreativeMaterialMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2DataListCreativeMaterialMode: valid values are %v", v, AllowedReportCreativeGetV2DataListCreativeMaterialModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2DataListCreativeMaterialMode) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2DataListCreativeMaterialModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_data_list_creative_material_mode value
func (v ReportCreativeGetV2DataListCreativeMaterialMode) Ptr() *ReportCreativeGetV2DataListCreativeMaterialMode {
	return &v
}
