/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandAdGetV30DataAdsAudienceInfoRetargetingType
type BrandAdGetV30DataAdsAudienceInfoRetargetingType int64

// List of brand_ad_get_v3.0_data_ads_audience_info_retargeting_type
const (
	Enum_0_BrandAdGetV30DataAdsAudienceInfoRetargetingType BrandAdGetV30DataAdsAudienceInfoRetargetingType = 0
	Enum_1_BrandAdGetV30DataAdsAudienceInfoRetargetingType BrandAdGetV30DataAdsAudienceInfoRetargetingType = 1
	Enum_2_BrandAdGetV30DataAdsAudienceInfoRetargetingType BrandAdGetV30DataAdsAudienceInfoRetargetingType = 2
	Enum_3_BrandAdGetV30DataAdsAudienceInfoRetargetingType BrandAdGetV30DataAdsAudienceInfoRetargetingType = 3
	Enum_4_BrandAdGetV30DataAdsAudienceInfoRetargetingType BrandAdGetV30DataAdsAudienceInfoRetargetingType = 4
	Enum_5_BrandAdGetV30DataAdsAudienceInfoRetargetingType BrandAdGetV30DataAdsAudienceInfoRetargetingType = 5
)

// All allowed values of BrandAdGetV30DataAdsAudienceInfoRetargetingType enum
var AllowedBrandAdGetV30DataAdsAudienceInfoRetargetingTypeEnumValues = []BrandAdGetV30DataAdsAudienceInfoRetargetingType{
	0,
	1,
	2,
	3,
	4,
	5,
}

// NewBrandAdGetV30DataAdsAudienceInfoRetargetingTypeFromValue returns a pointer to a valid BrandAdGetV30DataAdsAudienceInfoRetargetingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30DataAdsAudienceInfoRetargetingTypeFromValue(v int64) (*BrandAdGetV30DataAdsAudienceInfoRetargetingType, error) {
	ev := BrandAdGetV30DataAdsAudienceInfoRetargetingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30DataAdsAudienceInfoRetargetingType: valid values are %v", v, AllowedBrandAdGetV30DataAdsAudienceInfoRetargetingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30DataAdsAudienceInfoRetargetingType) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30DataAdsAudienceInfoRetargetingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_data_ads_audience_info_retargeting_type value
func (v BrandAdGetV30DataAdsAudienceInfoRetargetingType) Ptr() *BrandAdGetV30DataAdsAudienceInfoRetargetingType {
	return &v
}
