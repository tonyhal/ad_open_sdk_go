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

// ToolsAdminInfoV2DataDistrictsSubDistrictsLevel
type ToolsAdminInfoV2DataDistrictsSubDistrictsLevel string

// List of tools_admin_info_v2_data_districts_sub_districts_level
const (
	FOUR_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsLevel  ToolsAdminInfoV2DataDistrictsSubDistrictsLevel = "FOUR_LEVEL"
	NONE_ToolsAdminInfoV2DataDistrictsSubDistrictsLevel        ToolsAdminInfoV2DataDistrictsSubDistrictsLevel = "NONE"
	ONE_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsLevel   ToolsAdminInfoV2DataDistrictsSubDistrictsLevel = "ONE_LEVEL"
	THREE_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsLevel ToolsAdminInfoV2DataDistrictsSubDistrictsLevel = "THREE_LEVEL"
	TWO_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsLevel   ToolsAdminInfoV2DataDistrictsSubDistrictsLevel = "TWO_LEVEL"
)

// All allowed values of ToolsAdminInfoV2DataDistrictsSubDistrictsLevel enum
var AllowedToolsAdminInfoV2DataDistrictsSubDistrictsLevelEnumValues = []ToolsAdminInfoV2DataDistrictsSubDistrictsLevel{
	"FOUR_LEVEL",
	"NONE",
	"ONE_LEVEL",
	"THREE_LEVEL",
	"TWO_LEVEL",
}

// NewToolsAdminInfoV2DataDistrictsSubDistrictsLevelFromValue returns a pointer to a valid ToolsAdminInfoV2DataDistrictsSubDistrictsLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdminInfoV2DataDistrictsSubDistrictsLevelFromValue(v string) (*ToolsAdminInfoV2DataDistrictsSubDistrictsLevel, error) {
	ev := ToolsAdminInfoV2DataDistrictsSubDistrictsLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdminInfoV2DataDistrictsSubDistrictsLevel: valid values are %v", v, AllowedToolsAdminInfoV2DataDistrictsSubDistrictsLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdminInfoV2DataDistrictsSubDistrictsLevel) IsValid() bool {
	for _, existing := range AllowedToolsAdminInfoV2DataDistrictsSubDistrictsLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_admin_info_v2_data_districts_sub_districts_level value
func (v ToolsAdminInfoV2DataDistrictsSubDistrictsLevel) Ptr() *ToolsAdminInfoV2DataDistrictsSubDistrictsLevel {
	return &v
}
