/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignGetV2FilteringLandingType
type CampaignGetV2FilteringLandingType string

// List of campaign_get_v2_filtering_landing_type
const (
	ARTICLE_CampaignGetV2FilteringLandingType        CampaignGetV2FilteringLandingType = "ARTICLE"
	BRAND_EXTERNAL_CampaignGetV2FilteringLandingType CampaignGetV2FilteringLandingType = "BRAND_EXTERNAL"
	STORE_CampaignGetV2FilteringLandingType          CampaignGetV2FilteringLandingType = "STORE"
	DPA_CampaignGetV2FilteringLandingType            CampaignGetV2FilteringLandingType = "DPA"
	LIVE_CampaignGetV2FilteringLandingType           CampaignGetV2FilteringLandingType = "LIVE"
	QUICK_APP_CampaignGetV2FilteringLandingType      CampaignGetV2FilteringLandingType = "QUICK_APP"
	SHOP_CampaignGetV2FilteringLandingType           CampaignGetV2FilteringLandingType = "SHOP"
	APP_CampaignGetV2FilteringLandingType            CampaignGetV2FilteringLandingType = "APP"
	GOODS_CampaignGetV2FilteringLandingType          CampaignGetV2FilteringLandingType = "GOODS"
	LINK_CampaignGetV2FilteringLandingType           CampaignGetV2FilteringLandingType = "LINK"
	AWEME_CampaignGetV2FilteringLandingType          CampaignGetV2FilteringLandingType = "AWEME"
)

// Ptr returns reference to campaign_get_v2_filtering_landing_type value
func (v CampaignGetV2FilteringLandingType) Ptr() *CampaignGetV2FilteringLandingType {
	return &v
}
