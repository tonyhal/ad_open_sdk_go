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

// BrandAdGetV30DataAdsAudienceInfoExcludeDistrict
type BrandAdGetV30DataAdsAudienceInfoExcludeDistrict int64

// List of brand_ad_get_v3.0_data_ads_audience_info_exclude_district
const (
	Enum_0_BrandAdGetV30DataAdsAudienceInfoExcludeDistrict BrandAdGetV30DataAdsAudienceInfoExcludeDistrict = 0
	Enum_1_BrandAdGetV30DataAdsAudienceInfoExcludeDistrict BrandAdGetV30DataAdsAudienceInfoExcludeDistrict = 1
)

// All allowed values of BrandAdGetV30DataAdsAudienceInfoExcludeDistrict enum
var AllowedBrandAdGetV30DataAdsAudienceInfoExcludeDistrictEnumValues = []BrandAdGetV30DataAdsAudienceInfoExcludeDistrict{
	0,
	1,
}

// NewBrandAdGetV30DataAdsAudienceInfoExcludeDistrictFromValue returns a pointer to a valid BrandAdGetV30DataAdsAudienceInfoExcludeDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30DataAdsAudienceInfoExcludeDistrictFromValue(v int64) (*BrandAdGetV30DataAdsAudienceInfoExcludeDistrict, error) {
	ev := BrandAdGetV30DataAdsAudienceInfoExcludeDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30DataAdsAudienceInfoExcludeDistrict: valid values are %v", v, AllowedBrandAdGetV30DataAdsAudienceInfoExcludeDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30DataAdsAudienceInfoExcludeDistrict) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30DataAdsAudienceInfoExcludeDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_data_ads_audience_info_exclude_district value
func (v BrandAdGetV30DataAdsAudienceInfoExcludeDistrict) Ptr() *BrandAdGetV30DataAdsAudienceInfoExcludeDistrict {
	return &v
}
