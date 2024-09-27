/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCreativeGetV2FilteringPricingCategories
type ReportCreativeGetV2FilteringPricingCategories string

// List of report_creative_get_v2_filtering_pricing_categories
const (
	PRICING_CATEGORY_BRAND_AND_PRICING_ReportCreativeGetV2FilteringPricingCategories ReportCreativeGetV2FilteringPricingCategories = "PRICING_CATEGORY_BRAND_AND_PRICING"
	PRICING_CATEGORY_FREE_ReportCreativeGetV2FilteringPricingCategories              ReportCreativeGetV2FilteringPricingCategories = "PRICING_CATEGORY_FREE"
	PRICING_CATEGORY_BID_ReportCreativeGetV2FilteringPricingCategories               ReportCreativeGetV2FilteringPricingCategories = "PRICING_CATEGORY_BID"
	PRICING_CATEGORY_BRAND_ReportCreativeGetV2FilteringPricingCategories             ReportCreativeGetV2FilteringPricingCategories = "PRICING_CATEGORY_BRAND"
	PRICING_CATEGORY_NOC_ReportCreativeGetV2FilteringPricingCategories               ReportCreativeGetV2FilteringPricingCategories = "PRICING_CATEGORY_NOC"
)

// Ptr returns reference to report_creative_get_v2_filtering_pricing_categories value
func (v ReportCreativeGetV2FilteringPricingCategories) Ptr() *ReportCreativeGetV2FilteringPricingCategories {
	return &v
}
