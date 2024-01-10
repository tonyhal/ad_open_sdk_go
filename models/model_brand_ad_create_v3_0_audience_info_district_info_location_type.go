/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandAdCreateV30AudienceInfoDistrictInfoLocationType
type BrandAdCreateV30AudienceInfoDistrictInfoLocationType string

// List of brand_ad_create_v3.0_audience_info_district_info_location_type
const (
	ALL_LOCATION_BrandAdCreateV30AudienceInfoDistrictInfoLocationType BrandAdCreateV30AudienceInfoDistrictInfoLocationType = "ALL_LOCATION"
	TRAVEL_BrandAdCreateV30AudienceInfoDistrictInfoLocationType       BrandAdCreateV30AudienceInfoDistrictInfoLocationType = "TRAVEL"
)

// All allowed values of BrandAdCreateV30AudienceInfoDistrictInfoLocationType enum
var AllowedBrandAdCreateV30AudienceInfoDistrictInfoLocationTypeEnumValues = []BrandAdCreateV30AudienceInfoDistrictInfoLocationType{
	"ALL_LOCATION",
	"TRAVEL",
}

// NewBrandAdCreateV30AudienceInfoDistrictInfoLocationTypeFromValue returns a pointer to a valid BrandAdCreateV30AudienceInfoDistrictInfoLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdCreateV30AudienceInfoDistrictInfoLocationTypeFromValue(v string) (*BrandAdCreateV30AudienceInfoDistrictInfoLocationType, error) {
	ev := BrandAdCreateV30AudienceInfoDistrictInfoLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdCreateV30AudienceInfoDistrictInfoLocationType: valid values are %v", v, AllowedBrandAdCreateV30AudienceInfoDistrictInfoLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdCreateV30AudienceInfoDistrictInfoLocationType) IsValid() bool {
	for _, existing := range AllowedBrandAdCreateV30AudienceInfoDistrictInfoLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_create_v3.0_audience_info_district_info_location_type value
func (v BrandAdCreateV30AudienceInfoDistrictInfoLocationType) Ptr() *BrandAdCreateV30AudienceInfoDistrictInfoLocationType {
	return &v
}
