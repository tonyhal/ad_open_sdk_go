/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCampaignCreateV10MarketingScene
type QianchuanCampaignCreateV10MarketingScene string

// List of qianchuan_campaign_create_v1.0_marketing_scene
const (
	FEED_QianchuanCampaignCreateV10MarketingScene          QianchuanCampaignCreateV10MarketingScene = "FEED"
	SEARCH_QianchuanCampaignCreateV10MarketingScene        QianchuanCampaignCreateV10MarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanCampaignCreateV10MarketingScene QianchuanCampaignCreateV10MarketingScene = "SHOPPING_MALL"
)

// Ptr returns reference to qianchuan_campaign_create_v1.0_marketing_scene value
func (v QianchuanCampaignCreateV10MarketingScene) Ptr() *QianchuanCampaignCreateV10MarketingScene {
	return &v
}
