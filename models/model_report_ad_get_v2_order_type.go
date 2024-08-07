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

// ReportAdGetV2OrderType
type ReportAdGetV2OrderType string

// List of report_ad_get_v2_order_type
const (
	ASC_ReportAdGetV2OrderType  ReportAdGetV2OrderType = "ASC"
	DESC_ReportAdGetV2OrderType ReportAdGetV2OrderType = "DESC"
)

// All allowed values of ReportAdGetV2OrderType enum
var AllowedReportAdGetV2OrderTypeEnumValues = []ReportAdGetV2OrderType{
	"ASC",
	"DESC",
}

// NewReportAdGetV2OrderTypeFromValue returns a pointer to a valid ReportAdGetV2OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2OrderTypeFromValue(v string) (*ReportAdGetV2OrderType, error) {
	ev := ReportAdGetV2OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2OrderType: valid values are %v", v, AllowedReportAdGetV2OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2OrderType) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_order_type value
func (v ReportAdGetV2OrderType) Ptr() *ReportAdGetV2OrderType {
	return &v
}
