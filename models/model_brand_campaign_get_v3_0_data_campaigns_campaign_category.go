/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandCampaignGetV30DataCampaignsCampaignCategory
type BrandCampaignGetV30DataCampaignsCampaignCategory int64

// List of brand_campaign_get_v3.0_data_campaigns_campaign_category
const (
	Enum_1_BrandCampaignGetV30DataCampaignsCampaignCategory BrandCampaignGetV30DataCampaignsCampaignCategory = 1
	Enum_2_BrandCampaignGetV30DataCampaignsCampaignCategory BrandCampaignGetV30DataCampaignsCampaignCategory = 2
	Enum_3_BrandCampaignGetV30DataCampaignsCampaignCategory BrandCampaignGetV30DataCampaignsCampaignCategory = 3
	Enum_4_BrandCampaignGetV30DataCampaignsCampaignCategory BrandCampaignGetV30DataCampaignsCampaignCategory = 4
)

// All allowed values of BrandCampaignGetV30DataCampaignsCampaignCategory enum
var AllowedBrandCampaignGetV30DataCampaignsCampaignCategoryEnumValues = []BrandCampaignGetV30DataCampaignsCampaignCategory{
	1,
	2,
	3,
	4,
}

// NewBrandCampaignGetV30DataCampaignsCampaignCategoryFromValue returns a pointer to a valid BrandCampaignGetV30DataCampaignsCampaignCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandCampaignGetV30DataCampaignsCampaignCategoryFromValue(v int64) (*BrandCampaignGetV30DataCampaignsCampaignCategory, error) {
	ev := BrandCampaignGetV30DataCampaignsCampaignCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandCampaignGetV30DataCampaignsCampaignCategory: valid values are %v", v, AllowedBrandCampaignGetV30DataCampaignsCampaignCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandCampaignGetV30DataCampaignsCampaignCategory) IsValid() bool {
	for _, existing := range AllowedBrandCampaignGetV30DataCampaignsCampaignCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_campaign_get_v3.0_data_campaigns_campaign_category value
func (v BrandCampaignGetV30DataCampaignsCampaignCategory) Ptr() *BrandCampaignGetV30DataCampaignsCampaignCategory {
	return &v
}
