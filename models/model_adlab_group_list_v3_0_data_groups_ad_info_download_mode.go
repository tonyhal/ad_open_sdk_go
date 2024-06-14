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

// AdlabGroupListV30DataGroupsAdInfoDownloadMode
type AdlabGroupListV30DataGroupsAdInfoDownloadMode string

// List of adlab_group_list_v3.0_data_groups_ad_info_download_mode
const (
	APP_STORE_DELIVERY_AdlabGroupListV30DataGroupsAdInfoDownloadMode AdlabGroupListV30DataGroupsAdInfoDownloadMode = "APP_STORE_DELIVERY"
	DEFAULT_AdlabGroupListV30DataGroupsAdInfoDownloadMode            AdlabGroupListV30DataGroupsAdInfoDownloadMode = "DEFAULT"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoDownloadMode enum
var AllowedAdlabGroupListV30DataGroupsAdInfoDownloadModeEnumValues = []AdlabGroupListV30DataGroupsAdInfoDownloadMode{
	"APP_STORE_DELIVERY",
	"DEFAULT",
}

// NewAdlabGroupListV30DataGroupsAdInfoDownloadModeFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoDownloadMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoDownloadModeFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoDownloadMode, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoDownloadMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoDownloadMode: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoDownloadModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoDownloadMode) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoDownloadModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_download_mode value
func (v AdlabGroupListV30DataGroupsAdInfoDownloadMode) Ptr() *AdlabGroupListV30DataGroupsAdInfoDownloadMode {
	return &v
}
