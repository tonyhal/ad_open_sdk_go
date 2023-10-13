/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCampaignCreateV10MarketingScene
type QianchuanCampaignCreateV10MarketingScene string

// List of qianchuan_campaign_create_v1.0_marketing_scene
const (
	FEED_QianchuanCampaignCreateV10MarketingScene          QianchuanCampaignCreateV10MarketingScene = "FEED"
	SEARCH_QianchuanCampaignCreateV10MarketingScene        QianchuanCampaignCreateV10MarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanCampaignCreateV10MarketingScene QianchuanCampaignCreateV10MarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanCampaignCreateV10MarketingScene enum
var AllowedQianchuanCampaignCreateV10MarketingSceneEnumValues = []QianchuanCampaignCreateV10MarketingScene{
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanCampaignCreateV10MarketingSceneFromValue returns a pointer to a valid QianchuanCampaignCreateV10MarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCampaignCreateV10MarketingSceneFromValue(v string) (*QianchuanCampaignCreateV10MarketingScene, error) {
	ev := QianchuanCampaignCreateV10MarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCampaignCreateV10MarketingScene: valid values are %v", v, AllowedQianchuanCampaignCreateV10MarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCampaignCreateV10MarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanCampaignCreateV10MarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_campaign_create_v1.0_marketing_scene value
func (v QianchuanCampaignCreateV10MarketingScene) Ptr() *QianchuanCampaignCreateV10MarketingScene {
	return &v
}
