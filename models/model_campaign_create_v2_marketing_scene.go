/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignCreateV2MarketingScene
type CampaignCreateV2MarketingScene string

// List of campaign_create_v2_marketing_scene
const (
	EDUCATION_CampaignCreateV2MarketingScene         CampaignCreateV2MarketingScene = "EDUCATION"
	PROMOTION_PURPOSE_CampaignCreateV2MarketingScene CampaignCreateV2MarketingScene = "PROMOTION_PURPOSE"
	CAR_CampaignCreateV2MarketingScene               CampaignCreateV2MarketingScene = "CAR"
	MERCHANTS_CampaignCreateV2MarketingScene         CampaignCreateV2MarketingScene = "MERCHANTS"
	SOCIAL_CampaignCreateV2MarketingScene            CampaignCreateV2MarketingScene = "SOCIAL"
	GAME_PROMOTION_CampaignCreateV2MarketingScene    CampaignCreateV2MarketingScene = "GAME_PROMOTION"
	NOVEL_CampaignCreateV2MarketingScene             CampaignCreateV2MarketingScene = "NOVEL"
	VIDEO_INFO_CampaignCreateV2MarketingScene        CampaignCreateV2MarketingScene = "VIDEO_INFO"
	GAME_SUBSCRIBE_CampaignCreateV2MarketingScene    CampaignCreateV2MarketingScene = "GAME_SUBSCRIBE"
	ECOMMERCE_CampaignCreateV2MarketingScene         CampaignCreateV2MarketingScene = "ECOMMERCE"
)

// Ptr returns reference to campaign_create_v2_marketing_scene value
func (v CampaignCreateV2MarketingScene) Ptr() *CampaignCreateV2MarketingScene {
	return &v
}
