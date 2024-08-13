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

// StarReportOrderUserDistributionGetV2FanType
type StarReportOrderUserDistributionGetV2FanType string

// List of star_report_order_user_distribution_get_v2_fan_type
const (
	ALL_StarReportOrderUserDistributionGetV2FanType     StarReportOrderUserDistributionGetV2FanType = "ALL"
	FANS_StarReportOrderUserDistributionGetV2FanType    StarReportOrderUserDistributionGetV2FanType = "FANS"
	NO_FANS_StarReportOrderUserDistributionGetV2FanType StarReportOrderUserDistributionGetV2FanType = "NO_FANS"
)

// All allowed values of StarReportOrderUserDistributionGetV2FanType enum
var AllowedStarReportOrderUserDistributionGetV2FanTypeEnumValues = []StarReportOrderUserDistributionGetV2FanType{
	"ALL",
	"FANS",
	"NO_FANS",
}

// NewStarReportOrderUserDistributionGetV2FanTypeFromValue returns a pointer to a valid StarReportOrderUserDistributionGetV2FanType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarReportOrderUserDistributionGetV2FanTypeFromValue(v string) (*StarReportOrderUserDistributionGetV2FanType, error) {
	ev := StarReportOrderUserDistributionGetV2FanType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarReportOrderUserDistributionGetV2FanType: valid values are %v", v, AllowedStarReportOrderUserDistributionGetV2FanTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarReportOrderUserDistributionGetV2FanType) IsValid() bool {
	for _, existing := range AllowedStarReportOrderUserDistributionGetV2FanTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_report_order_user_distribution_get_v2_fan_type value
func (v StarReportOrderUserDistributionGetV2FanType) Ptr() *StarReportOrderUserDistributionGetV2FanType {
	return &v
}
