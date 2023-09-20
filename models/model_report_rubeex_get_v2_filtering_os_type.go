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

// ReportRubeexGetV2FilteringOsType
type ReportRubeexGetV2FilteringOsType string

// List of report_rubeex_get_v2_filtering_os_type
const (
	IOS_ReportRubeexGetV2FilteringOsType     ReportRubeexGetV2FilteringOsType = "ios"
	ANDROID_ReportRubeexGetV2FilteringOsType ReportRubeexGetV2FilteringOsType = "android"
)

// All allowed values of ReportRubeexGetV2FilteringOsType enum
var AllowedReportRubeexGetV2FilteringOsTypeEnumValues = []ReportRubeexGetV2FilteringOsType{
	"ios",
	"android",
}

// NewReportRubeexGetV2FilteringOsTypeFromValue returns a pointer to a valid ReportRubeexGetV2FilteringOsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportRubeexGetV2FilteringOsTypeFromValue(v string) (*ReportRubeexGetV2FilteringOsType, error) {
	ev := ReportRubeexGetV2FilteringOsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportRubeexGetV2FilteringOsType: valid values are %v", v, AllowedReportRubeexGetV2FilteringOsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportRubeexGetV2FilteringOsType) IsValid() bool {
	for _, existing := range AllowedReportRubeexGetV2FilteringOsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_rubeex_get_v2_filtering_os_type value
func (v ReportRubeexGetV2FilteringOsType) Ptr() *ReportRubeexGetV2FilteringOsType {
	return &v
}
