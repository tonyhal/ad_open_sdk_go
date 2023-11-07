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

// ToolsPlayableCloudGameListV2FilteringOrientation
type ToolsPlayableCloudGameListV2FilteringOrientation string

// List of tools_playable_cloud_game_list_v2_filtering_orientation
const (
	HORIZONTAL_ToolsPlayableCloudGameListV2FilteringOrientation ToolsPlayableCloudGameListV2FilteringOrientation = "HORIZONTAL"
	VERTICAL_ToolsPlayableCloudGameListV2FilteringOrientation   ToolsPlayableCloudGameListV2FilteringOrientation = "VERTICAL"
)

// All allowed values of ToolsPlayableCloudGameListV2FilteringOrientation enum
var AllowedToolsPlayableCloudGameListV2FilteringOrientationEnumValues = []ToolsPlayableCloudGameListV2FilteringOrientation{
	"HORIZONTAL",
	"VERTICAL",
}

// NewToolsPlayableCloudGameListV2FilteringOrientationFromValue returns a pointer to a valid ToolsPlayableCloudGameListV2FilteringOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPlayableCloudGameListV2FilteringOrientationFromValue(v string) (*ToolsPlayableCloudGameListV2FilteringOrientation, error) {
	ev := ToolsPlayableCloudGameListV2FilteringOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPlayableCloudGameListV2FilteringOrientation: valid values are %v", v, AllowedToolsPlayableCloudGameListV2FilteringOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPlayableCloudGameListV2FilteringOrientation) IsValid() bool {
	for _, existing := range AllowedToolsPlayableCloudGameListV2FilteringOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_playable_cloud_game_list_v2_filtering_orientation value
func (v ToolsPlayableCloudGameListV2FilteringOrientation) Ptr() *ToolsPlayableCloudGameListV2FilteringOrientation {
	return &v
}
