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

// QianchuanReportSearchWordGetV10TimeGranularity
type QianchuanReportSearchWordGetV10TimeGranularity string

// List of qianchuan_report_search_word_get_v1.0_time_granularity
const (
	TIME_GRANULARITY_DAILY_QianchuanReportSearchWordGetV10TimeGranularity QianchuanReportSearchWordGetV10TimeGranularity = "TIME_GRANULARITY_DAILY"
	TIME_GRANULARITY_WEEK_QianchuanReportSearchWordGetV10TimeGranularity  QianchuanReportSearchWordGetV10TimeGranularity = "TIME_GRANULARITY_WEEK"
)

// All allowed values of QianchuanReportSearchWordGetV10TimeGranularity enum
var AllowedQianchuanReportSearchWordGetV10TimeGranularityEnumValues = []QianchuanReportSearchWordGetV10TimeGranularity{
	"TIME_GRANULARITY_DAILY",
	"TIME_GRANULARITY_WEEK",
}

// NewQianchuanReportSearchWordGetV10TimeGranularityFromValue returns a pointer to a valid QianchuanReportSearchWordGetV10TimeGranularity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportSearchWordGetV10TimeGranularityFromValue(v string) (*QianchuanReportSearchWordGetV10TimeGranularity, error) {
	ev := QianchuanReportSearchWordGetV10TimeGranularity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportSearchWordGetV10TimeGranularity: valid values are %v", v, AllowedQianchuanReportSearchWordGetV10TimeGranularityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportSearchWordGetV10TimeGranularity) IsValid() bool {
	for _, existing := range AllowedQianchuanReportSearchWordGetV10TimeGranularityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_search_word_get_v1.0_time_granularity value
func (v QianchuanReportSearchWordGetV10TimeGranularity) Ptr() *QianchuanReportSearchWordGetV10TimeGranularity {
	return &v
}
