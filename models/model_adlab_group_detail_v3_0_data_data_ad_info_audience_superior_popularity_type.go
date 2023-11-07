/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType
type AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType string

// List of adlab_group_detail_v3.0_data_data_ad_info_audience_superior_popularity_type
const (
	GAME_AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType = "GAME"
	NONE_AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType = "NONE"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType enum
var AllowedAdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityTypeEnumValues = []AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType{
	"GAME",
	"NONE",
}

// NewAdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityTypeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityTypeFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_audience_superior_popularity_type value
func (v AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType) Ptr() *AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType {
	return &v
}
