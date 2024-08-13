/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel
type ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel string

// List of tools_admin_info_v2_data_districts_sub_districts_sub_districts_sub_districts_sub_districts_level
const (
	FOUR_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel  ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel = "FOUR_LEVEL"
	NONE_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel        ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel = "NONE"
	ONE_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel   ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel = "ONE_LEVEL"
	THREE_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel = "THREE_LEVEL"
	TWO_LEVEL_ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel   ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel = "TWO_LEVEL"
)

// All allowed values of ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel enum
var AllowedToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevelEnumValues = []ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel{
	"FOUR_LEVEL",
	"NONE",
	"ONE_LEVEL",
	"THREE_LEVEL",
	"TWO_LEVEL",
}

// NewToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevelFromValue returns a pointer to a valid ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevelFromValue(v string) (*ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel, error) {
	ev := ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel: valid values are %v", v, AllowedToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel) IsValid() bool {
	for _, existing := range AllowedToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_admin_info_v2_data_districts_sub_districts_sub_districts_sub_districts_sub_districts_level value
func (v ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel) Ptr() *ToolsAdminInfoV2DataDistrictsSubDistrictsSubDistrictsSubDistrictsSubDistrictsLevel {
	return &v
}
