/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarReportDataTopicConfigV2DataStatMetricsType
type StarReportDataTopicConfigV2DataStatMetricsType string

// List of star_report_data_topic_config_v2_data_stat_metrics_type
const (
	FLOAT64_StarReportDataTopicConfigV2DataStatMetricsType StarReportDataTopicConfigV2DataStatMetricsType = "FLOAT64"
	INT64_StarReportDataTopicConfigV2DataStatMetricsType   StarReportDataTopicConfigV2DataStatMetricsType = "INT64"
	JSON_StarReportDataTopicConfigV2DataStatMetricsType    StarReportDataTopicConfigV2DataStatMetricsType = "JSON"
	STRING_StarReportDataTopicConfigV2DataStatMetricsType  StarReportDataTopicConfigV2DataStatMetricsType = "STRING"
)

// All allowed values of StarReportDataTopicConfigV2DataStatMetricsType enum
var AllowedStarReportDataTopicConfigV2DataStatMetricsTypeEnumValues = []StarReportDataTopicConfigV2DataStatMetricsType{
	"FLOAT64",
	"INT64",
	"JSON",
	"STRING",
}

// NewStarReportDataTopicConfigV2DataStatMetricsTypeFromValue returns a pointer to a valid StarReportDataTopicConfigV2DataStatMetricsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarReportDataTopicConfigV2DataStatMetricsTypeFromValue(v string) (*StarReportDataTopicConfigV2DataStatMetricsType, error) {
	ev := StarReportDataTopicConfigV2DataStatMetricsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarReportDataTopicConfigV2DataStatMetricsType: valid values are %v", v, AllowedStarReportDataTopicConfigV2DataStatMetricsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarReportDataTopicConfigV2DataStatMetricsType) IsValid() bool {
	for _, existing := range AllowedStarReportDataTopicConfigV2DataStatMetricsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_report_data_topic_config_v2_data_stat_metrics_type value
func (v StarReportDataTopicConfigV2DataStatMetricsType) Ptr() *StarReportDataTopicConfigV2DataStatMetricsType {
	return &v
}
