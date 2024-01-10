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

// ToolsAdConvertQueryV2DataListExternalActionsSource
type ToolsAdConvertQueryV2DataListExternalActionsSource string

// List of tools_ad_convert_query_v2_data_list_external_actions_source
const (
	SOURCE_TETRIS_ToolsAdConvertQueryV2DataListExternalActionsSource ToolsAdConvertQueryV2DataListExternalActionsSource = "SOURCE_TETRIS"
	SOURCE_CONFIG_ToolsAdConvertQueryV2DataListExternalActionsSource ToolsAdConvertQueryV2DataListExternalActionsSource = "SOURCE_CONFIG"
	SOURCE_METEOR_ToolsAdConvertQueryV2DataListExternalActionsSource ToolsAdConvertQueryV2DataListExternalActionsSource = "SOURCE_METEOR"
)

// All allowed values of ToolsAdConvertQueryV2DataListExternalActionsSource enum
var AllowedToolsAdConvertQueryV2DataListExternalActionsSourceEnumValues = []ToolsAdConvertQueryV2DataListExternalActionsSource{
	"SOURCE_TETRIS",
	"SOURCE_CONFIG",
	"SOURCE_METEOR",
}

// NewToolsAdConvertQueryV2DataListExternalActionsSourceFromValue returns a pointer to a valid ToolsAdConvertQueryV2DataListExternalActionsSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertQueryV2DataListExternalActionsSourceFromValue(v string) (*ToolsAdConvertQueryV2DataListExternalActionsSource, error) {
	ev := ToolsAdConvertQueryV2DataListExternalActionsSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertQueryV2DataListExternalActionsSource: valid values are %v", v, AllowedToolsAdConvertQueryV2DataListExternalActionsSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertQueryV2DataListExternalActionsSource) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertQueryV2DataListExternalActionsSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_query_v2_data_list_external_actions_source value
func (v ToolsAdConvertQueryV2DataListExternalActionsSource) Ptr() *ToolsAdConvertQueryV2DataListExternalActionsSource {
	return &v
}
