/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCampaignGetV2DataListPricingCategory
type ReportCampaignGetV2DataListPricingCategory string

// List of report_campaign_get_v2_data_list_pricing_category
const (
	PRICING_CATEGORY_BRAND_AND_PRICING_ReportCampaignGetV2DataListPricingCategory ReportCampaignGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND_AND_PRICING"
	PRICING_CATEGORY_BID_ReportCampaignGetV2DataListPricingCategory               ReportCampaignGetV2DataListPricingCategory = "PRICING_CATEGORY_BID"
	PRICING_CATEGORY_FREE_ReportCampaignGetV2DataListPricingCategory              ReportCampaignGetV2DataListPricingCategory = "PRICING_CATEGORY_FREE"
	PRICING_CATEGORY_BRAND_ReportCampaignGetV2DataListPricingCategory             ReportCampaignGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND"
	PRICING_CATEGORY_NOC_ReportCampaignGetV2DataListPricingCategory               ReportCampaignGetV2DataListPricingCategory = "PRICING_CATEGORY_NOC"
)

// Ptr returns reference to report_campaign_get_v2_data_list_pricing_category value
func (v ReportCampaignGetV2DataListPricingCategory) Ptr() *ReportCampaignGetV2DataListPricingCategory {
	return &v
}
