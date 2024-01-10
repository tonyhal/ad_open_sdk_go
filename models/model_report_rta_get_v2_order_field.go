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

// ReportRtaGetV2OrderField
type ReportRtaGetV2OrderField string

// List of report_rta_get_v2_order_field
const (
	STAT_COST_ReportRtaGetV2OrderField   ReportRtaGetV2OrderField = "stat_cost"
	SHOW_CNT_ReportRtaGetV2OrderField    ReportRtaGetV2OrderField = "show_cnt"
	CLICK_CNT_ReportRtaGetV2OrderField   ReportRtaGetV2OrderField = "click_cnt"
	CONVERT_CNT_ReportRtaGetV2OrderField ReportRtaGetV2OrderField = "convert_cnt"
)

// All allowed values of ReportRtaGetV2OrderField enum
var AllowedReportRtaGetV2OrderFieldEnumValues = []ReportRtaGetV2OrderField{
	"stat_cost",
	"show_cnt",
	"click_cnt",
	"convert_cnt",
}

// NewReportRtaGetV2OrderFieldFromValue returns a pointer to a valid ReportRtaGetV2OrderField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportRtaGetV2OrderFieldFromValue(v string) (*ReportRtaGetV2OrderField, error) {
	ev := ReportRtaGetV2OrderField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportRtaGetV2OrderField: valid values are %v", v, AllowedReportRtaGetV2OrderFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportRtaGetV2OrderField) IsValid() bool {
	for _, existing := range AllowedReportRtaGetV2OrderFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_rta_get_v2_order_field value
func (v ReportRtaGetV2OrderField) Ptr() *ReportRtaGetV2OrderField {
	return &v
}
