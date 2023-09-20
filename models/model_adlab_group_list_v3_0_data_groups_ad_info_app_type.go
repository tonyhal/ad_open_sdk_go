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

// AdlabGroupListV30DataGroupsAdInfoAppType
type AdlabGroupListV30DataGroupsAdInfoAppType string

// List of adlab_group_list_v3.0_data_groups_ad_info_app_type
const (
	APP_ANDROID_AdlabGroupListV30DataGroupsAdInfoAppType AdlabGroupListV30DataGroupsAdInfoAppType = "APP_ANDROID"
	APP_IOS_AdlabGroupListV30DataGroupsAdInfoAppType     AdlabGroupListV30DataGroupsAdInfoAppType = "APP_IOS"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoAppType enum
var AllowedAdlabGroupListV30DataGroupsAdInfoAppTypeEnumValues = []AdlabGroupListV30DataGroupsAdInfoAppType{
	"APP_ANDROID",
	"APP_IOS",
}

// NewAdlabGroupListV30DataGroupsAdInfoAppTypeFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoAppType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoAppTypeFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoAppType, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoAppType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoAppType: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoAppTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoAppType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoAppTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_app_type value
func (v AdlabGroupListV30DataGroupsAdInfoAppType) Ptr() *AdlabGroupListV30DataGroupsAdInfoAppType {
	return &v
}
