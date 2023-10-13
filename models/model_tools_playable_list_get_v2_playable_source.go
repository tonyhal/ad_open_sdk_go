/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPlayableListGetV2PlayableSource
type ToolsPlayableListGetV2PlayableSource string

// List of tools_playable_list_get_v2_playable_source
const (
	EXCLUDE_DIRECT_ToolsPlayableListGetV2PlayableSource  ToolsPlayableListGetV2PlayableSource = "EXCLUDE_DIRECT"
	DIRECT_PLAYABLE_ToolsPlayableListGetV2PlayableSource ToolsPlayableListGetV2PlayableSource = "DIRECT_PLAYABLE"
)

// All allowed values of ToolsPlayableListGetV2PlayableSource enum
var AllowedToolsPlayableListGetV2PlayableSourceEnumValues = []ToolsPlayableListGetV2PlayableSource{
	"EXCLUDE_DIRECT",
	"DIRECT_PLAYABLE",
}

// NewToolsPlayableListGetV2PlayableSourceFromValue returns a pointer to a valid ToolsPlayableListGetV2PlayableSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPlayableListGetV2PlayableSourceFromValue(v string) (*ToolsPlayableListGetV2PlayableSource, error) {
	ev := ToolsPlayableListGetV2PlayableSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPlayableListGetV2PlayableSource: valid values are %v", v, AllowedToolsPlayableListGetV2PlayableSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPlayableListGetV2PlayableSource) IsValid() bool {
	for _, existing := range AllowedToolsPlayableListGetV2PlayableSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_playable_list_get_v2_playable_source value
func (v ToolsPlayableListGetV2PlayableSource) Ptr() *ToolsPlayableListGetV2PlayableSource {
	return &v
}
