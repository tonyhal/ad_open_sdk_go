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

// StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType
type StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType string

// List of star_report_custom_data_topic_daily_report_v2_data_stats_data_metrics_type
const (
	FLOAT64_StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType = "FLOAT64"
	INT64_StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType   StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType = "INT64"
	JSON_StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType    StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType = "JSON"
	STRING_StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType  StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType = "STRING"
)

// All allowed values of StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType enum
var AllowedStarReportCustomDataTopicDailyReportV2DataStatsDataMetricsTypeEnumValues = []StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType{
	"FLOAT64",
	"INT64",
	"JSON",
	"STRING",
}

// NewStarReportCustomDataTopicDailyReportV2DataStatsDataMetricsTypeFromValue returns a pointer to a valid StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarReportCustomDataTopicDailyReportV2DataStatsDataMetricsTypeFromValue(v string) (*StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType, error) {
	ev := StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType: valid values are %v", v, AllowedStarReportCustomDataTopicDailyReportV2DataStatsDataMetricsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType) IsValid() bool {
	for _, existing := range AllowedStarReportCustomDataTopicDailyReportV2DataStatsDataMetricsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_report_custom_data_topic_daily_report_v2_data_stats_data_metrics_type value
func (v StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType) Ptr() *StarReportCustomDataTopicDailyReportV2DataStatsDataMetricsType {
	return &v
}
