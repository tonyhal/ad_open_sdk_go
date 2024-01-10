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

// QianchuanToolsAllowCouponV10MarketingGoal
type QianchuanToolsAllowCouponV10MarketingGoal string

// List of qianchuan_tools_allow_coupon_v1.0_marketing_goal
const (
	ALL_QianchuanToolsAllowCouponV10MarketingGoal              QianchuanToolsAllowCouponV10MarketingGoal = "ALL"
	LIVE_PROM_GOODS_QianchuanToolsAllowCouponV10MarketingGoal  QianchuanToolsAllowCouponV10MarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanToolsAllowCouponV10MarketingGoal QianchuanToolsAllowCouponV10MarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanToolsAllowCouponV10MarketingGoal enum
var AllowedQianchuanToolsAllowCouponV10MarketingGoalEnumValues = []QianchuanToolsAllowCouponV10MarketingGoal{
	"ALL",
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanToolsAllowCouponV10MarketingGoalFromValue returns a pointer to a valid QianchuanToolsAllowCouponV10MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsAllowCouponV10MarketingGoalFromValue(v string) (*QianchuanToolsAllowCouponV10MarketingGoal, error) {
	ev := QianchuanToolsAllowCouponV10MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsAllowCouponV10MarketingGoal: valid values are %v", v, AllowedQianchuanToolsAllowCouponV10MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsAllowCouponV10MarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsAllowCouponV10MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_allow_coupon_v1.0_marketing_goal value
func (v QianchuanToolsAllowCouponV10MarketingGoal) Ptr() *QianchuanToolsAllowCouponV10MarketingGoal {
	return &v
}
