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

// AdGetV2DataAudienceDistrict
type AdGetV2DataAudienceDistrict string

// List of ad_get_v2_data_audience_district
const (
	BUSINESS_DISTRICT_AdGetV2DataAudienceDistrict AdGetV2DataAudienceDistrict = "BUSINESS_DISTRICT"
	COUNTY_AdGetV2DataAudienceDistrict            AdGetV2DataAudienceDistrict = "COUNTY"
	CITY_AdGetV2DataAudienceDistrict              AdGetV2DataAudienceDistrict = "CITY"
	OVERSEA_AdGetV2DataAudienceDistrict           AdGetV2DataAudienceDistrict = "OVERSEA"
	REGION_AdGetV2DataAudienceDistrict            AdGetV2DataAudienceDistrict = "REGION"
	NONE_AdGetV2DataAudienceDistrict              AdGetV2DataAudienceDistrict = "NONE"
)

// All allowed values of AdGetV2DataAudienceDistrict enum
var AllowedAdGetV2DataAudienceDistrictEnumValues = []AdGetV2DataAudienceDistrict{
	"BUSINESS_DISTRICT",
	"COUNTY",
	"CITY",
	"OVERSEA",
	"REGION",
	"NONE",
}

// NewAdGetV2DataAudienceDistrictFromValue returns a pointer to a valid AdGetV2DataAudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceDistrictFromValue(v string) (*AdGetV2DataAudienceDistrict, error) {
	ev := AdGetV2DataAudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceDistrict: valid values are %v", v, AllowedAdGetV2DataAudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceDistrict) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_district value
func (v AdGetV2DataAudienceDistrict) Ptr() *AdGetV2DataAudienceDistrict {
	return &v
}
