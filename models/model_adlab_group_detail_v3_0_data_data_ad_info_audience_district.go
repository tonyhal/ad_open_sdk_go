/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataAdInfoAudienceDistrict
type AdlabGroupDetailV30DataDataAdInfoAudienceDistrict string

// List of adlab_group_detail_v3.0_data_data_ad_info_audience_district
const (
	BUSINESS_DISTRICT_AdlabGroupDetailV30DataDataAdInfoAudienceDistrict AdlabGroupDetailV30DataDataAdInfoAudienceDistrict = "BUSINESS_DISTRICT"
	CITY_AdlabGroupDetailV30DataDataAdInfoAudienceDistrict              AdlabGroupDetailV30DataDataAdInfoAudienceDistrict = "CITY"
	COUNTY_AdlabGroupDetailV30DataDataAdInfoAudienceDistrict            AdlabGroupDetailV30DataDataAdInfoAudienceDistrict = "COUNTY"
	NONE_AdlabGroupDetailV30DataDataAdInfoAudienceDistrict              AdlabGroupDetailV30DataDataAdInfoAudienceDistrict = "NONE"
	REGION_AdlabGroupDetailV30DataDataAdInfoAudienceDistrict            AdlabGroupDetailV30DataDataAdInfoAudienceDistrict = "REGION"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoAudienceDistrict enum
var AllowedAdlabGroupDetailV30DataDataAdInfoAudienceDistrictEnumValues = []AdlabGroupDetailV30DataDataAdInfoAudienceDistrict{
	"BUSINESS_DISTRICT",
	"CITY",
	"COUNTY",
	"NONE",
	"REGION",
}

// NewAdlabGroupDetailV30DataDataAdInfoAudienceDistrictFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoAudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoAudienceDistrictFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoAudienceDistrict, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoAudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoAudienceDistrict: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoAudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoAudienceDistrict) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoAudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_audience_district value
func (v AdlabGroupDetailV30DataDataAdInfoAudienceDistrict) Ptr() *AdlabGroupDetailV30DataDataAdInfoAudienceDistrict {
	return &v
}
