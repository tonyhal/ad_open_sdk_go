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

// ReportAdGetV2FilteringPricingCategories
type ReportAdGetV2FilteringPricingCategories string

// List of report_ad_get_v2_filtering_pricing_categories
const (
	PRICING_CATEGORY_FREE_ReportAdGetV2FilteringPricingCategories              ReportAdGetV2FilteringPricingCategories = "PRICING_CATEGORY_FREE"
	PRICING_CATEGORY_BRAND_AND_PRICING_ReportAdGetV2FilteringPricingCategories ReportAdGetV2FilteringPricingCategories = "PRICING_CATEGORY_BRAND_AND_PRICING"
	PRICING_CATEGORY_BRAND_ReportAdGetV2FilteringPricingCategories             ReportAdGetV2FilteringPricingCategories = "PRICING_CATEGORY_BRAND"
	PRICING_CATEGORY_BID_ReportAdGetV2FilteringPricingCategories               ReportAdGetV2FilteringPricingCategories = "PRICING_CATEGORY_BID"
	PRICING_CATEGORY_NOC_ReportAdGetV2FilteringPricingCategories               ReportAdGetV2FilteringPricingCategories = "PRICING_CATEGORY_NOC"
)

// All allowed values of ReportAdGetV2FilteringPricingCategories enum
var AllowedReportAdGetV2FilteringPricingCategoriesEnumValues = []ReportAdGetV2FilteringPricingCategories{
	"PRICING_CATEGORY_FREE",
	"PRICING_CATEGORY_BRAND_AND_PRICING",
	"PRICING_CATEGORY_BRAND",
	"PRICING_CATEGORY_BID",
	"PRICING_CATEGORY_NOC",
}

// NewReportAdGetV2FilteringPricingCategoriesFromValue returns a pointer to a valid ReportAdGetV2FilteringPricingCategories
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2FilteringPricingCategoriesFromValue(v string) (*ReportAdGetV2FilteringPricingCategories, error) {
	ev := ReportAdGetV2FilteringPricingCategories(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2FilteringPricingCategories: valid values are %v", v, AllowedReportAdGetV2FilteringPricingCategoriesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2FilteringPricingCategories) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2FilteringPricingCategoriesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_filtering_pricing_categories value
func (v ReportAdGetV2FilteringPricingCategories) Ptr() *ReportAdGetV2FilteringPricingCategories {
	return &v
}
