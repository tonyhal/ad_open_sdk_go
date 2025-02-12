/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignCreateV2LandingType
type CampaignCreateV2LandingType string

// List of campaign_create_v2_landing_type
const (
	LIVE_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "LIVE"
	STORE_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "STORE"
	SHOP_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "SHOP"
	ARTICLE_CampaignCreateV2LandingType        CampaignCreateV2LandingType = "ARTICLE"
	GOODS_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "GOODS"
	LINK_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "LINK"
	BRAND_EXTERNAL_CampaignCreateV2LandingType CampaignCreateV2LandingType = "BRAND_EXTERNAL"
	AWEME_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "AWEME"
	QUICK_APP_CampaignCreateV2LandingType      CampaignCreateV2LandingType = "QUICK_APP"
	APP_CampaignCreateV2LandingType            CampaignCreateV2LandingType = "APP"
	DPA_CampaignCreateV2LandingType            CampaignCreateV2LandingType = "DPA"
)

// Ptr returns reference to campaign_create_v2_landing_type value
func (v CampaignCreateV2LandingType) Ptr() *CampaignCreateV2LandingType {
	return &v
}
