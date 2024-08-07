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

// BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType
type BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType string

// List of brand_ad_create_v3.0_audience_info_district_info_city_select_type
const (
	EXCLUDE_BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType = "EXCLUDE"
	SELECT_BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType  BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType = "SELECT"
)

// All allowed values of BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType enum
var AllowedBrandAdCreateV30AudienceInfoDistrictInfoCitySelectTypeEnumValues = []BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType{
	"EXCLUDE",
	"SELECT",
}

// NewBrandAdCreateV30AudienceInfoDistrictInfoCitySelectTypeFromValue returns a pointer to a valid BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdCreateV30AudienceInfoDistrictInfoCitySelectTypeFromValue(v string) (*BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType, error) {
	ev := BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType: valid values are %v", v, AllowedBrandAdCreateV30AudienceInfoDistrictInfoCitySelectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType) IsValid() bool {
	for _, existing := range AllowedBrandAdCreateV30AudienceInfoDistrictInfoCitySelectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_create_v3.0_audience_info_district_info_city_select_type value
func (v BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType) Ptr() *BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType {
	return &v
}
