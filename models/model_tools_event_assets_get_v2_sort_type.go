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

// ToolsEventAssetsGetV2SortType
type ToolsEventAssetsGetV2SortType string

// List of tools_event_assets_get_v2_sort_type
const (
	DESC_ToolsEventAssetsGetV2SortType ToolsEventAssetsGetV2SortType = "DESC"
	ASC_ToolsEventAssetsGetV2SortType  ToolsEventAssetsGetV2SortType = "ASC"
)

// All allowed values of ToolsEventAssetsGetV2SortType enum
var AllowedToolsEventAssetsGetV2SortTypeEnumValues = []ToolsEventAssetsGetV2SortType{
	"DESC",
	"ASC",
}

// NewToolsEventAssetsGetV2SortTypeFromValue returns a pointer to a valid ToolsEventAssetsGetV2SortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventAssetsGetV2SortTypeFromValue(v string) (*ToolsEventAssetsGetV2SortType, error) {
	ev := ToolsEventAssetsGetV2SortType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventAssetsGetV2SortType: valid values are %v", v, AllowedToolsEventAssetsGetV2SortTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventAssetsGetV2SortType) IsValid() bool {
	for _, existing := range AllowedToolsEventAssetsGetV2SortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_assets_get_v2_sort_type value
func (v ToolsEventAssetsGetV2SortType) Ptr() *ToolsEventAssetsGetV2SortType {
	return &v
}
