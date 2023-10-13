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

// QianchuanCampaignListGetV10DataListMarketingScene
type QianchuanCampaignListGetV10DataListMarketingScene string

// List of qianchuan_campaign_list_get_v1.0_data_list_marketing_scene
const (
	ALL_QianchuanCampaignListGetV10DataListMarketingScene           QianchuanCampaignListGetV10DataListMarketingScene = "ALL"
	FEED_QianchuanCampaignListGetV10DataListMarketingScene          QianchuanCampaignListGetV10DataListMarketingScene = "FEED"
	SEARCH_QianchuanCampaignListGetV10DataListMarketingScene        QianchuanCampaignListGetV10DataListMarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanCampaignListGetV10DataListMarketingScene QianchuanCampaignListGetV10DataListMarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanCampaignListGetV10DataListMarketingScene enum
var AllowedQianchuanCampaignListGetV10DataListMarketingSceneEnumValues = []QianchuanCampaignListGetV10DataListMarketingScene{
	"ALL",
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanCampaignListGetV10DataListMarketingSceneFromValue returns a pointer to a valid QianchuanCampaignListGetV10DataListMarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCampaignListGetV10DataListMarketingSceneFromValue(v string) (*QianchuanCampaignListGetV10DataListMarketingScene, error) {
	ev := QianchuanCampaignListGetV10DataListMarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCampaignListGetV10DataListMarketingScene: valid values are %v", v, AllowedQianchuanCampaignListGetV10DataListMarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCampaignListGetV10DataListMarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanCampaignListGetV10DataListMarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_campaign_list_get_v1.0_data_list_marketing_scene value
func (v QianchuanCampaignListGetV10DataListMarketingScene) Ptr() *QianchuanCampaignListGetV10DataListMarketingScene {
	return &v
}
