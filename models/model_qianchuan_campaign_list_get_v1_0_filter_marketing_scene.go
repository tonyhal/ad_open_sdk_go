/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCampaignListGetV10FilterMarketingScene
type QianchuanCampaignListGetV10FilterMarketingScene string

// List of qianchuan_campaign_list_get_v1.0_filter_marketing_scene
const (
	FEED_QianchuanCampaignListGetV10FilterMarketingScene          QianchuanCampaignListGetV10FilterMarketingScene = "FEED"
	SEARCH_QianchuanCampaignListGetV10FilterMarketingScene        QianchuanCampaignListGetV10FilterMarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanCampaignListGetV10FilterMarketingScene QianchuanCampaignListGetV10FilterMarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanCampaignListGetV10FilterMarketingScene enum
var AllowedQianchuanCampaignListGetV10FilterMarketingSceneEnumValues = []QianchuanCampaignListGetV10FilterMarketingScene{
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanCampaignListGetV10FilterMarketingSceneFromValue returns a pointer to a valid QianchuanCampaignListGetV10FilterMarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCampaignListGetV10FilterMarketingSceneFromValue(v string) (*QianchuanCampaignListGetV10FilterMarketingScene, error) {
	ev := QianchuanCampaignListGetV10FilterMarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCampaignListGetV10FilterMarketingScene: valid values are %v", v, AllowedQianchuanCampaignListGetV10FilterMarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCampaignListGetV10FilterMarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanCampaignListGetV10FilterMarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_campaign_list_get_v1.0_filter_marketing_scene value
func (v QianchuanCampaignListGetV10FilterMarketingScene) Ptr() *QianchuanCampaignListGetV10FilterMarketingScene {
	return &v
}
