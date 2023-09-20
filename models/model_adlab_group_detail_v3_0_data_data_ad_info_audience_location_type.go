/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataAdInfoAudienceLocationType
type AdlabGroupDetailV30DataDataAdInfoAudienceLocationType string

// List of adlab_group_detail_v3.0_data_data_ad_info_audience_location_type
const (
	ALL_AdlabGroupDetailV30DataDataAdInfoAudienceLocationType     AdlabGroupDetailV30DataDataAdInfoAudienceLocationType = "ALL"
	CURRENT_AdlabGroupDetailV30DataDataAdInfoAudienceLocationType AdlabGroupDetailV30DataDataAdInfoAudienceLocationType = "CURRENT"
	HOME_AdlabGroupDetailV30DataDataAdInfoAudienceLocationType    AdlabGroupDetailV30DataDataAdInfoAudienceLocationType = "HOME"
	TRAVEL_AdlabGroupDetailV30DataDataAdInfoAudienceLocationType  AdlabGroupDetailV30DataDataAdInfoAudienceLocationType = "TRAVEL"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoAudienceLocationType enum
var AllowedAdlabGroupDetailV30DataDataAdInfoAudienceLocationTypeEnumValues = []AdlabGroupDetailV30DataDataAdInfoAudienceLocationType{
	"ALL",
	"CURRENT",
	"HOME",
	"TRAVEL",
}

// NewAdlabGroupDetailV30DataDataAdInfoAudienceLocationTypeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoAudienceLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoAudienceLocationTypeFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoAudienceLocationType, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoAudienceLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoAudienceLocationType: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoAudienceLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoAudienceLocationType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoAudienceLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_audience_location_type value
func (v AdlabGroupDetailV30DataDataAdInfoAudienceLocationType) Ptr() *AdlabGroupDetailV30DataDataAdInfoAudienceLocationType {
	return &v
}
