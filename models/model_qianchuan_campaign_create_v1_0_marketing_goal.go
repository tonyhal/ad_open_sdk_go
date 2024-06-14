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

// QianchuanCampaignCreateV10MarketingGoal
type QianchuanCampaignCreateV10MarketingGoal string

// List of qianchuan_campaign_create_v1.0_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanCampaignCreateV10MarketingGoal  QianchuanCampaignCreateV10MarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanCampaignCreateV10MarketingGoal QianchuanCampaignCreateV10MarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanCampaignCreateV10MarketingGoal enum
var AllowedQianchuanCampaignCreateV10MarketingGoalEnumValues = []QianchuanCampaignCreateV10MarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanCampaignCreateV10MarketingGoalFromValue returns a pointer to a valid QianchuanCampaignCreateV10MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCampaignCreateV10MarketingGoalFromValue(v string) (*QianchuanCampaignCreateV10MarketingGoal, error) {
	ev := QianchuanCampaignCreateV10MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCampaignCreateV10MarketingGoal: valid values are %v", v, AllowedQianchuanCampaignCreateV10MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCampaignCreateV10MarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanCampaignCreateV10MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_campaign_create_v1.0_marketing_goal value
func (v QianchuanCampaignCreateV10MarketingGoal) Ptr() *QianchuanCampaignCreateV10MarketingGoal {
	return &v
}
