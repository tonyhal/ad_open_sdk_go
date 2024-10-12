/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignGetV2DataListMarketingPurpose
type CampaignGetV2DataListMarketingPurpose string

// List of campaign_get_v2_data_list_marketing_purpose
const (
	UNLIMITED_CampaignGetV2DataListMarketingPurpose   CampaignGetV2DataListMarketingPurpose = "UNLIMITED"
	INTENTION_CampaignGetV2DataListMarketingPurpose   CampaignGetV2DataListMarketingPurpose = "INTENTION"
	CONVERSION_CampaignGetV2DataListMarketingPurpose  CampaignGetV2DataListMarketingPurpose = "CONVERSION"
	ACKNOWLEDGE_CampaignGetV2DataListMarketingPurpose CampaignGetV2DataListMarketingPurpose = "ACKNOWLEDGE"
)

// Ptr returns reference to campaign_get_v2_data_list_marketing_purpose value
func (v CampaignGetV2DataListMarketingPurpose) Ptr() *CampaignGetV2DataListMarketingPurpose {
	return &v
}
