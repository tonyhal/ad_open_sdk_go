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

// QianchuanReportCreativeGetV10FilteringMarketingGoal
type QianchuanReportCreativeGetV10FilteringMarketingGoal string

// List of qianchuan_report_creative_get_v1.0_filtering_marketing_goal
const (
	ALL_QianchuanReportCreativeGetV10FilteringMarketingGoal              QianchuanReportCreativeGetV10FilteringMarketingGoal = "ALL"
	LIVE_PROM_GOODS_QianchuanReportCreativeGetV10FilteringMarketingGoal  QianchuanReportCreativeGetV10FilteringMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanReportCreativeGetV10FilteringMarketingGoal QianchuanReportCreativeGetV10FilteringMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanReportCreativeGetV10FilteringMarketingGoal enum
var AllowedQianchuanReportCreativeGetV10FilteringMarketingGoalEnumValues = []QianchuanReportCreativeGetV10FilteringMarketingGoal{
	"ALL",
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanReportCreativeGetV10FilteringMarketingGoalFromValue returns a pointer to a valid QianchuanReportCreativeGetV10FilteringMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportCreativeGetV10FilteringMarketingGoalFromValue(v string) (*QianchuanReportCreativeGetV10FilteringMarketingGoal, error) {
	ev := QianchuanReportCreativeGetV10FilteringMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportCreativeGetV10FilteringMarketingGoal: valid values are %v", v, AllowedQianchuanReportCreativeGetV10FilteringMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportCreativeGetV10FilteringMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanReportCreativeGetV10FilteringMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_creative_get_v1.0_filtering_marketing_goal value
func (v QianchuanReportCreativeGetV10FilteringMarketingGoal) Ptr() *QianchuanReportCreativeGetV10FilteringMarketingGoal {
	return &v
}
