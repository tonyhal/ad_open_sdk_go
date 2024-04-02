/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportLongTransferOrderGetV10MarketingGoal
type QianchuanReportLongTransferOrderGetV10MarketingGoal string

// List of qianchuan_report_long_transfer_order_get_v1.0_marketing_goal
const (
	ALL_QianchuanReportLongTransferOrderGetV10MarketingGoal              QianchuanReportLongTransferOrderGetV10MarketingGoal = "ALL"
	LIVE_PROM_GOODS_QianchuanReportLongTransferOrderGetV10MarketingGoal  QianchuanReportLongTransferOrderGetV10MarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanReportLongTransferOrderGetV10MarketingGoal QianchuanReportLongTransferOrderGetV10MarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanReportLongTransferOrderGetV10MarketingGoal enum
var AllowedQianchuanReportLongTransferOrderGetV10MarketingGoalEnumValues = []QianchuanReportLongTransferOrderGetV10MarketingGoal{
	"ALL",
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanReportLongTransferOrderGetV10MarketingGoalFromValue returns a pointer to a valid QianchuanReportLongTransferOrderGetV10MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportLongTransferOrderGetV10MarketingGoalFromValue(v string) (*QianchuanReportLongTransferOrderGetV10MarketingGoal, error) {
	ev := QianchuanReportLongTransferOrderGetV10MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportLongTransferOrderGetV10MarketingGoal: valid values are %v", v, AllowedQianchuanReportLongTransferOrderGetV10MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportLongTransferOrderGetV10MarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanReportLongTransferOrderGetV10MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_long_transfer_order_get_v1.0_marketing_goal value
func (v QianchuanReportLongTransferOrderGetV10MarketingGoal) Ptr() *QianchuanReportLongTransferOrderGetV10MarketingGoal {
	return &v
}
