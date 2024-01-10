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

// AdlabGroupListV30DataGroupsAdInfoAudienceLocationType
type AdlabGroupListV30DataGroupsAdInfoAudienceLocationType string

// List of adlab_group_list_v3.0_data_groups_ad_info_audience_location_type
const (
	ALL_AdlabGroupListV30DataGroupsAdInfoAudienceLocationType     AdlabGroupListV30DataGroupsAdInfoAudienceLocationType = "ALL"
	CURRENT_AdlabGroupListV30DataGroupsAdInfoAudienceLocationType AdlabGroupListV30DataGroupsAdInfoAudienceLocationType = "CURRENT"
	HOME_AdlabGroupListV30DataGroupsAdInfoAudienceLocationType    AdlabGroupListV30DataGroupsAdInfoAudienceLocationType = "HOME"
	TRAVEL_AdlabGroupListV30DataGroupsAdInfoAudienceLocationType  AdlabGroupListV30DataGroupsAdInfoAudienceLocationType = "TRAVEL"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoAudienceLocationType enum
var AllowedAdlabGroupListV30DataGroupsAdInfoAudienceLocationTypeEnumValues = []AdlabGroupListV30DataGroupsAdInfoAudienceLocationType{
	"ALL",
	"CURRENT",
	"HOME",
	"TRAVEL",
}

// NewAdlabGroupListV30DataGroupsAdInfoAudienceLocationTypeFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoAudienceLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoAudienceLocationTypeFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoAudienceLocationType, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoAudienceLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoAudienceLocationType: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoAudienceLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoAudienceLocationType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoAudienceLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_audience_location_type value
func (v AdlabGroupListV30DataGroupsAdInfoAudienceLocationType) Ptr() *AdlabGroupListV30DataGroupsAdInfoAudienceLocationType {
	return &v
}
