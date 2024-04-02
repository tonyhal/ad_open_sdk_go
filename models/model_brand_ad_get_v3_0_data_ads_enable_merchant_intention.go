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

// BrandAdGetV30DataAdsEnableMerchantIntention
type BrandAdGetV30DataAdsEnableMerchantIntention int64

// List of brand_ad_get_v3.0_data_ads_enable_merchant_intention
const (
	Enum_0_BrandAdGetV30DataAdsEnableMerchantIntention BrandAdGetV30DataAdsEnableMerchantIntention = 0
	Enum_1_BrandAdGetV30DataAdsEnableMerchantIntention BrandAdGetV30DataAdsEnableMerchantIntention = 1
)

// All allowed values of BrandAdGetV30DataAdsEnableMerchantIntention enum
var AllowedBrandAdGetV30DataAdsEnableMerchantIntentionEnumValues = []BrandAdGetV30DataAdsEnableMerchantIntention{
	0,
	1,
}

// NewBrandAdGetV30DataAdsEnableMerchantIntentionFromValue returns a pointer to a valid BrandAdGetV30DataAdsEnableMerchantIntention
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30DataAdsEnableMerchantIntentionFromValue(v int64) (*BrandAdGetV30DataAdsEnableMerchantIntention, error) {
	ev := BrandAdGetV30DataAdsEnableMerchantIntention(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30DataAdsEnableMerchantIntention: valid values are %v", v, AllowedBrandAdGetV30DataAdsEnableMerchantIntentionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30DataAdsEnableMerchantIntention) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30DataAdsEnableMerchantIntentionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_data_ads_enable_merchant_intention value
func (v BrandAdGetV30DataAdsEnableMerchantIntention) Ptr() *BrandAdGetV30DataAdsEnableMerchantIntention {
	return &v
}
