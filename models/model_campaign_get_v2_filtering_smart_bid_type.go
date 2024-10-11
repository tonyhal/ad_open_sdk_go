/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignGetV2FilteringSmartBidType
type CampaignGetV2FilteringSmartBidType string

// List of campaign_get_v2_filtering_smart_bid_type
const (
	BAOSHOU_CampaignGetV2FilteringSmartBidType           CampaignGetV2FilteringSmartBidType = "BAOSHOU"
	SMART_BID_NO_BID_CampaignGetV2FilteringSmartBidType  CampaignGetV2FilteringSmartBidType = "SMART_BID_NO_BID"
	MAXCV_CampaignGetV2FilteringSmartBidType             CampaignGetV2FilteringSmartBidType = "MAXCV"
	JIJIN_CampaignGetV2FilteringSmartBidType             CampaignGetV2FilteringSmartBidType = "JIJIN"
	AWEME_LITE_PACING_CampaignGetV2FilteringSmartBidType CampaignGetV2FilteringSmartBidType = "AWEME_LITE_PACING"
	GD_PROMOTE_CampaignGetV2FilteringSmartBidType        CampaignGetV2FilteringSmartBidType = "GD_PROMOTE"
	LITE_PACING_CampaignGetV2FilteringSmartBidType       CampaignGetV2FilteringSmartBidType = "LITE_PACING"
	GUARANTEED_SHOW_CampaignGetV2FilteringSmartBidType   CampaignGetV2FilteringSmartBidType = "GUARANTEED_SHOW"
	CUSTOM_CampaignGetV2FilteringSmartBidType            CampaignGetV2FilteringSmartBidType = "CUSTOM"
	MAX_CONVERSION_CampaignGetV2FilteringSmartBidType    CampaignGetV2FilteringSmartBidType = "MAX_CONVERSION"
)

// Ptr returns reference to campaign_get_v2_filtering_smart_bid_type value
func (v CampaignGetV2FilteringSmartBidType) Ptr() *CampaignGetV2FilteringSmartBidType {
	return &v
}
