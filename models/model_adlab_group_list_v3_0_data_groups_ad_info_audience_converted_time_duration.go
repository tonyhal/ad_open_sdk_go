/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration
type AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration string

// List of adlab_group_list_v3.0_data_groups_ad_info_audience_converted_time_duration
const (
	NONE_AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration         AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration = "NONE"
	ONE_MONTH_AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration    AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration = "ONE_MONTH"
	SEVEN_DAY_AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration    AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration = "SEVEN_DAY"
	SIX_MONTH_AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration    AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration = "SIX_MONTH"
	THREE_MONTH_AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration  AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration = "THREE_MONTH"
	TODAY_AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration        AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration = "TODAY"
	TWELVE_MONTH_AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration = "TWELVE_MONTH"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration enum
var AllowedAdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDurationEnumValues = []AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration{
	"NONE",
	"ONE_MONTH",
	"SEVEN_DAY",
	"SIX_MONTH",
	"THREE_MONTH",
	"TODAY",
	"TWELVE_MONTH",
}

// NewAdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDurationFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDurationFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_audience_converted_time_duration value
func (v AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration) Ptr() *AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration {
	return &v
}
