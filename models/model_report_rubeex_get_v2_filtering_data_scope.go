/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportRubeexGetV2FilteringDataScope
type ReportRubeexGetV2FilteringDataScope string

// List of report_rubeex_get_v2_filtering_data_scope
const (
	ALL_ReportRubeexGetV2FilteringDataScope ReportRubeexGetV2FilteringDataScope = "ALL"
	ADV_ReportRubeexGetV2FilteringDataScope ReportRubeexGetV2FilteringDataScope = "ADV"
)

// All allowed values of ReportRubeexGetV2FilteringDataScope enum
var AllowedReportRubeexGetV2FilteringDataScopeEnumValues = []ReportRubeexGetV2FilteringDataScope{
	"ALL",
	"ADV",
}

// NewReportRubeexGetV2FilteringDataScopeFromValue returns a pointer to a valid ReportRubeexGetV2FilteringDataScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportRubeexGetV2FilteringDataScopeFromValue(v string) (*ReportRubeexGetV2FilteringDataScope, error) {
	ev := ReportRubeexGetV2FilteringDataScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportRubeexGetV2FilteringDataScope: valid values are %v", v, AllowedReportRubeexGetV2FilteringDataScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportRubeexGetV2FilteringDataScope) IsValid() bool {
	for _, existing := range AllowedReportRubeexGetV2FilteringDataScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_rubeex_get_v2_filtering_data_scope value
func (v ReportRubeexGetV2FilteringDataScope) Ptr() *ReportRubeexGetV2FilteringDataScope {
	return &v
}
