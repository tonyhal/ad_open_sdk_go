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

// ReportAdvertiserGetV2TimeGranularity
type ReportAdvertiserGetV2TimeGranularity string

// List of report_advertiser_get_v2_time_granularity
const (
	STAT_TIME_GRANULARITY_MONTH_ReportAdvertiserGetV2TimeGranularity  ReportAdvertiserGetV2TimeGranularity = "STAT_TIME_GRANULARITY_MONTH"
	STAT_TIME_GRANULARITY_WEEK_ReportAdvertiserGetV2TimeGranularity   ReportAdvertiserGetV2TimeGranularity = "STAT_TIME_GRANULARITY_WEEK"
	STAT_TIME_GRANULARITY_DAILY_ReportAdvertiserGetV2TimeGranularity  ReportAdvertiserGetV2TimeGranularity = "STAT_TIME_GRANULARITY_DAILY"
	STAT_TIME_GRANULARITY_HOURLY_ReportAdvertiserGetV2TimeGranularity ReportAdvertiserGetV2TimeGranularity = "STAT_TIME_GRANULARITY_HOURLY"
)

// All allowed values of ReportAdvertiserGetV2TimeGranularity enum
var AllowedReportAdvertiserGetV2TimeGranularityEnumValues = []ReportAdvertiserGetV2TimeGranularity{
	"STAT_TIME_GRANULARITY_MONTH",
	"STAT_TIME_GRANULARITY_WEEK",
	"STAT_TIME_GRANULARITY_DAILY",
	"STAT_TIME_GRANULARITY_HOURLY",
}

// NewReportAdvertiserGetV2TimeGranularityFromValue returns a pointer to a valid ReportAdvertiserGetV2TimeGranularity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2TimeGranularityFromValue(v string) (*ReportAdvertiserGetV2TimeGranularity, error) {
	ev := ReportAdvertiserGetV2TimeGranularity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2TimeGranularity: valid values are %v", v, AllowedReportAdvertiserGetV2TimeGranularityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2TimeGranularity) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2TimeGranularityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_time_granularity value
func (v ReportAdvertiserGetV2TimeGranularity) Ptr() *ReportAdvertiserGetV2TimeGranularity {
	return &v
}
