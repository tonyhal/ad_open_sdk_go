/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsRubeexPlayableListV2DataListPlayableOrientation
type ToolsRubeexPlayableListV2DataListPlayableOrientation string

// List of tools_rubeex_playable_list_v2_data_list_playable_orientation
const (
	BOTH_ToolsRubeexPlayableListV2DataListPlayableOrientation      ToolsRubeexPlayableListV2DataListPlayableOrientation = "BOTH"
	PORTRAIT_ToolsRubeexPlayableListV2DataListPlayableOrientation  ToolsRubeexPlayableListV2DataListPlayableOrientation = "PORTRAIT"
	LANDSCAPE_ToolsRubeexPlayableListV2DataListPlayableOrientation ToolsRubeexPlayableListV2DataListPlayableOrientation = "LANDSCAPE"
)

// All allowed values of ToolsRubeexPlayableListV2DataListPlayableOrientation enum
var AllowedToolsRubeexPlayableListV2DataListPlayableOrientationEnumValues = []ToolsRubeexPlayableListV2DataListPlayableOrientation{
	"BOTH",
	"PORTRAIT",
	"LANDSCAPE",
}

// NewToolsRubeexPlayableListV2DataListPlayableOrientationFromValue returns a pointer to a valid ToolsRubeexPlayableListV2DataListPlayableOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRubeexPlayableListV2DataListPlayableOrientationFromValue(v string) (*ToolsRubeexPlayableListV2DataListPlayableOrientation, error) {
	ev := ToolsRubeexPlayableListV2DataListPlayableOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRubeexPlayableListV2DataListPlayableOrientation: valid values are %v", v, AllowedToolsRubeexPlayableListV2DataListPlayableOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRubeexPlayableListV2DataListPlayableOrientation) IsValid() bool {
	for _, existing := range AllowedToolsRubeexPlayableListV2DataListPlayableOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rubeex_playable_list_v2_data_list_playable_orientation value
func (v ToolsRubeexPlayableListV2DataListPlayableOrientation) Ptr() *ToolsRubeexPlayableListV2DataListPlayableOrientation {
	return &v
}
