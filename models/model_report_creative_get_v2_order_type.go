/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCreativeGetV2OrderType
type ReportCreativeGetV2OrderType string

// List of report_creative_get_v2_order_type
const (
	DESC_ReportCreativeGetV2OrderType ReportCreativeGetV2OrderType = "DESC"
	ASC_ReportCreativeGetV2OrderType  ReportCreativeGetV2OrderType = "ASC"
)

// All allowed values of ReportCreativeGetV2OrderType enum
var AllowedReportCreativeGetV2OrderTypeEnumValues = []ReportCreativeGetV2OrderType{
	"DESC",
	"ASC",
}

// NewReportCreativeGetV2OrderTypeFromValue returns a pointer to a valid ReportCreativeGetV2OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2OrderTypeFromValue(v string) (*ReportCreativeGetV2OrderType, error) {
	ev := ReportCreativeGetV2OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2OrderType: valid values are %v", v, AllowedReportCreativeGetV2OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2OrderType) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_order_type value
func (v ReportCreativeGetV2OrderType) Ptr() *ReportCreativeGetV2OrderType {
	return &v
}
