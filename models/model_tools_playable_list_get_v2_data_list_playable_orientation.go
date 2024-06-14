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

// ToolsPlayableListGetV2DataListPlayableOrientation
type ToolsPlayableListGetV2DataListPlayableOrientation string

// List of tools_playable_list_get_v2_data_list_playable_orientation
const (
	PORTRAIT_ToolsPlayableListGetV2DataListPlayableOrientation  ToolsPlayableListGetV2DataListPlayableOrientation = "PORTRAIT"
	BOTH_ToolsPlayableListGetV2DataListPlayableOrientation      ToolsPlayableListGetV2DataListPlayableOrientation = "BOTH"
	LANDSCAPE_ToolsPlayableListGetV2DataListPlayableOrientation ToolsPlayableListGetV2DataListPlayableOrientation = "LANDSCAPE"
)

// All allowed values of ToolsPlayableListGetV2DataListPlayableOrientation enum
var AllowedToolsPlayableListGetV2DataListPlayableOrientationEnumValues = []ToolsPlayableListGetV2DataListPlayableOrientation{
	"PORTRAIT",
	"BOTH",
	"LANDSCAPE",
}

// NewToolsPlayableListGetV2DataListPlayableOrientationFromValue returns a pointer to a valid ToolsPlayableListGetV2DataListPlayableOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPlayableListGetV2DataListPlayableOrientationFromValue(v string) (*ToolsPlayableListGetV2DataListPlayableOrientation, error) {
	ev := ToolsPlayableListGetV2DataListPlayableOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPlayableListGetV2DataListPlayableOrientation: valid values are %v", v, AllowedToolsPlayableListGetV2DataListPlayableOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPlayableListGetV2DataListPlayableOrientation) IsValid() bool {
	for _, existing := range AllowedToolsPlayableListGetV2DataListPlayableOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_playable_list_get_v2_data_list_playable_orientation value
func (v ToolsPlayableListGetV2DataListPlayableOrientation) Ptr() *ToolsPlayableListGetV2DataListPlayableOrientation {
	return &v
}
