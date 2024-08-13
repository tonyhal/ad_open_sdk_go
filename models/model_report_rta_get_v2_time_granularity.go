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

// ReportRtaGetV2TimeGranularity
type ReportRtaGetV2TimeGranularity string

// List of report_rta_get_v2_time_granularity
const (
	STAT_TIME_DAY_ReportRtaGetV2TimeGranularity ReportRtaGetV2TimeGranularity = "STAT_TIME_DAY"
)

// All allowed values of ReportRtaGetV2TimeGranularity enum
var AllowedReportRtaGetV2TimeGranularityEnumValues = []ReportRtaGetV2TimeGranularity{
	"STAT_TIME_DAY",
}

// NewReportRtaGetV2TimeGranularityFromValue returns a pointer to a valid ReportRtaGetV2TimeGranularity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportRtaGetV2TimeGranularityFromValue(v string) (*ReportRtaGetV2TimeGranularity, error) {
	ev := ReportRtaGetV2TimeGranularity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportRtaGetV2TimeGranularity: valid values are %v", v, AllowedReportRtaGetV2TimeGranularityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportRtaGetV2TimeGranularity) IsValid() bool {
	for _, existing := range AllowedReportRtaGetV2TimeGranularityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_rta_get_v2_time_granularity value
func (v ReportRtaGetV2TimeGranularity) Ptr() *ReportRtaGetV2TimeGranularity {
	return &v
}
