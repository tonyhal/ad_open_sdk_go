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

// ToolsOrangeSiteGetV30FilteringSearchMode
type ToolsOrangeSiteGetV30FilteringSearchMode string

// List of tools_orange_site_get_v3.0_filtering_search_mode
const (
	ALL_ToolsOrangeSiteGetV30FilteringSearchMode   ToolsOrangeSiteGetV30FilteringSearchMode = "ALL"
	OTHER_ToolsOrangeSiteGetV30FilteringSearchMode ToolsOrangeSiteGetV30FilteringSearchMode = "OTHER"
	OUR_ToolsOrangeSiteGetV30FilteringSearchMode   ToolsOrangeSiteGetV30FilteringSearchMode = "OUR"
)

// All allowed values of ToolsOrangeSiteGetV30FilteringSearchMode enum
var AllowedToolsOrangeSiteGetV30FilteringSearchModeEnumValues = []ToolsOrangeSiteGetV30FilteringSearchMode{
	"ALL",
	"OTHER",
	"OUR",
}

// NewToolsOrangeSiteGetV30FilteringSearchModeFromValue returns a pointer to a valid ToolsOrangeSiteGetV30FilteringSearchMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsOrangeSiteGetV30FilteringSearchModeFromValue(v string) (*ToolsOrangeSiteGetV30FilteringSearchMode, error) {
	ev := ToolsOrangeSiteGetV30FilteringSearchMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsOrangeSiteGetV30FilteringSearchMode: valid values are %v", v, AllowedToolsOrangeSiteGetV30FilteringSearchModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsOrangeSiteGetV30FilteringSearchMode) IsValid() bool {
	for _, existing := range AllowedToolsOrangeSiteGetV30FilteringSearchModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_orange_site_get_v3.0_filtering_search_mode value
func (v ToolsOrangeSiteGetV30FilteringSearchMode) Ptr() *ToolsOrangeSiteGetV30FilteringSearchMode {
	return &v
}
