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

// QianchuanReportAdGetV10FilteringMarketingGoal
type QianchuanReportAdGetV10FilteringMarketingGoal string

// List of qianchuan_report_ad_get_v1.0_filtering_marketing_goal
const (
	ALL_QianchuanReportAdGetV10FilteringMarketingGoal              QianchuanReportAdGetV10FilteringMarketingGoal = "ALL"
	LIVE_PROM_GOODS_QianchuanReportAdGetV10FilteringMarketingGoal  QianchuanReportAdGetV10FilteringMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanReportAdGetV10FilteringMarketingGoal QianchuanReportAdGetV10FilteringMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanReportAdGetV10FilteringMarketingGoal enum
var AllowedQianchuanReportAdGetV10FilteringMarketingGoalEnumValues = []QianchuanReportAdGetV10FilteringMarketingGoal{
	"ALL",
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanReportAdGetV10FilteringMarketingGoalFromValue returns a pointer to a valid QianchuanReportAdGetV10FilteringMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdGetV10FilteringMarketingGoalFromValue(v string) (*QianchuanReportAdGetV10FilteringMarketingGoal, error) {
	ev := QianchuanReportAdGetV10FilteringMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdGetV10FilteringMarketingGoal: valid values are %v", v, AllowedQianchuanReportAdGetV10FilteringMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdGetV10FilteringMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdGetV10FilteringMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_ad_get_v1.0_filtering_marketing_goal value
func (v QianchuanReportAdGetV10FilteringMarketingGoal) Ptr() *QianchuanReportAdGetV10FilteringMarketingGoal {
	return &v
}