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

// ReportRubeexGetV2FilteringType
type ReportRubeexGetV2FilteringType int64

// List of report_rubeex_get_v2_filtering_type
const (
	Enum_0_ReportRubeexGetV2FilteringType ReportRubeexGetV2FilteringType = 0
	Enum_1_ReportRubeexGetV2FilteringType ReportRubeexGetV2FilteringType = 1
)

// All allowed values of ReportRubeexGetV2FilteringType enum
var AllowedReportRubeexGetV2FilteringTypeEnumValues = []ReportRubeexGetV2FilteringType{
	0,
	1,
}

// NewReportRubeexGetV2FilteringTypeFromValue returns a pointer to a valid ReportRubeexGetV2FilteringType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportRubeexGetV2FilteringTypeFromValue(v int64) (*ReportRubeexGetV2FilteringType, error) {
	ev := ReportRubeexGetV2FilteringType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportRubeexGetV2FilteringType: valid values are %v", v, AllowedReportRubeexGetV2FilteringTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportRubeexGetV2FilteringType) IsValid() bool {
	for _, existing := range AllowedReportRubeexGetV2FilteringTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_rubeex_get_v2_filtering_type value
func (v ReportRubeexGetV2FilteringType) Ptr() *ReportRubeexGetV2FilteringType {
	return &v
}
