/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdvertiserGetV2GroupBy
type ReportAdvertiserGetV2GroupBy string

// List of report_advertiser_get_v2_group_by
const (
	STAT_GROUP_BY_INVENTORY_ReportAdvertiserGetV2GroupBy              ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_INVENTORY"
	STAT_GROUP_BY_AC_ReportAdvertiserGetV2GroupBy                     ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_AC"
	STAT_GROUP_BY_PLATFORM_ReportAdvertiserGetV2GroupBy               ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_PLATFORM"
	STAT_GROUP_BY_FIELD_STAT_TIME_ReportAdvertiserGetV2GroupBy        ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_FIELD_STAT_TIME"
	STAT_GROUP_BY_LANDING_TYPE_ReportAdvertiserGetV2GroupBy           ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_LANDING_TYPE"
	STAT_GROUP_BY_PROVINCE_NAME_ReportAdvertiserGetV2GroupBy          ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_PROVINCE_NAME"
	STAT_GROUP_BY_CREATIVE_MATERIAL_MODE_ReportAdvertiserGetV2GroupBy ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_CREATIVE_MATERIAL_MODE"
	STAT_GROUP_BY_PRICING_ReportAdvertiserGetV2GroupBy                ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_PRICING"
	STAT_GROUP_BY_MATERIAL_ID_ReportAdvertiserGetV2GroupBy            ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_MATERIAL_ID"
	STAT_GROUP_BY_GENDER_ReportAdvertiserGetV2GroupBy                 ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_GENDER"
	STAT_GROUP_BY_AGE_ReportAdvertiserGetV2GroupBy                    ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_AGE"
	STAT_GROUP_BY_CREATIVE_COMPONENT_ID_ReportAdvertiserGetV2GroupBy  ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_CREATIVE_COMPONENT_ID"
	STAT_GROUP_BY_CAMPAIGN_TYPE_ReportAdvertiserGetV2GroupBy          ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_CAMPAIGN_TYPE"
	STAT_GROUP_BY_FIELD_ID_ReportAdvertiserGetV2GroupBy               ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_FIELD_ID"
	STAT_GROUP_BY_CITY_NAME_ReportAdvertiserGetV2GroupBy              ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_CITY_NAME"
	STAT_GROUP_BY_PRICING_CATEGORY_ReportAdvertiserGetV2GroupBy       ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_PRICING_CATEGORY"
	STAT_GROUP_BY_IMAGE_MODE_ReportAdvertiserGetV2GroupBy             ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_IMAGE_MODE"
	STAT_GROUP_BY_EXTERNAL_ACTION_ReportAdvertiserGetV2GroupBy        ReportAdvertiserGetV2GroupBy = "STAT_GROUP_BY_EXTERNAL_ACTION"
)

// All allowed values of ReportAdvertiserGetV2GroupBy enum
var AllowedReportAdvertiserGetV2GroupByEnumValues = []ReportAdvertiserGetV2GroupBy{
	"STAT_GROUP_BY_INVENTORY",
	"STAT_GROUP_BY_AC",
	"STAT_GROUP_BY_PLATFORM",
	"STAT_GROUP_BY_FIELD_STAT_TIME",
	"STAT_GROUP_BY_LANDING_TYPE",
	"STAT_GROUP_BY_PROVINCE_NAME",
	"STAT_GROUP_BY_CREATIVE_MATERIAL_MODE",
	"STAT_GROUP_BY_PRICING",
	"STAT_GROUP_BY_MATERIAL_ID",
	"STAT_GROUP_BY_GENDER",
	"STAT_GROUP_BY_AGE",
	"STAT_GROUP_BY_CREATIVE_COMPONENT_ID",
	"STAT_GROUP_BY_CAMPAIGN_TYPE",
	"STAT_GROUP_BY_FIELD_ID",
	"STAT_GROUP_BY_CITY_NAME",
	"STAT_GROUP_BY_PRICING_CATEGORY",
	"STAT_GROUP_BY_IMAGE_MODE",
	"STAT_GROUP_BY_EXTERNAL_ACTION",
}

// NewReportAdvertiserGetV2GroupByFromValue returns a pointer to a valid ReportAdvertiserGetV2GroupBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2GroupByFromValue(v string) (*ReportAdvertiserGetV2GroupBy, error) {
	ev := ReportAdvertiserGetV2GroupBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2GroupBy: valid values are %v", v, AllowedReportAdvertiserGetV2GroupByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2GroupBy) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2GroupByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_group_by value
func (v ReportAdvertiserGetV2GroupBy) Ptr() *ReportAdvertiserGetV2GroupBy {
	return &v
}
