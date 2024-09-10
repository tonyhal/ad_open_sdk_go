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
	LIVE_CampaignGetV2FilteringLandingType           CampaignGetV2FilteringLandingType = "LIVE"
	APP_CampaignGetV2FilteringLandingType            CampaignGetV2FilteringLandingType = "APP"
	LINK_CampaignGetV2FilteringLandingType           CampaignGetV2FilteringLandingType = "LINK"
	QUICK_APP_CampaignGetV2FilteringLandingType      CampaignGetV2FilteringLandingType = "QUICK_APP"
	BRAND_EXTERNAL_CampaignGetV2FilteringLandingType CampaignGetV2FilteringLandingType = "BRAND_EXTERNAL"
	GOODS_CampaignGetV2FilteringLandingType          CampaignGetV2FilteringLandingType = "GOODS"
	ARTICLE_CampaignGetV2FilteringLandingType        CampaignGetV2FilteringLandingType = "ARTICLE"
	STORE_CampaignGetV2FilteringLandingType          CampaignGetV2FilteringLandingType = "STORE"
	SHOP_CampaignGetV2FilteringLandingType           CampaignGetV2FilteringLandingType = "SHOP"
	DPA_CampaignGetV2FilteringLandingType            CampaignGetV2FilteringLandingType = "DPA"
	AWEME_CampaignGetV2FilteringLandingType          CampaignGetV2FilteringLandingType = "AWEME"
)

// Ptr returns reference to campaign_get_v2_filtering_landing_type value
func (v CampaignGetV2FilteringLandingType) Ptr() *CampaignGetV2FilteringLandingType {
	return &v
}
