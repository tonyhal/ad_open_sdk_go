/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2InterestActionMode
type ToolsBidSuggestV2InterestActionMode string

// List of tools_bid_suggest_v2_interest_action_mode
const (
	UNLIMITED_ToolsBidSuggestV2InterestActionMode ToolsBidSuggestV2InterestActionMode = "UNLIMITED"
	RECOMMEND_ToolsBidSuggestV2InterestActionMode ToolsBidSuggestV2InterestActionMode = "RECOMMEND"
	CUSTOM_ToolsBidSuggestV2InterestActionMode    ToolsBidSuggestV2InterestActionMode = "CUSTOM"
)

// All allowed values of ToolsBidSuggestV2InterestActionMode enum
var AllowedToolsBidSuggestV2InterestActionModeEnumValues = []ToolsBidSuggestV2InterestActionMode{
	"UNLIMITED",
	"RECOMMEND",
	"CUSTOM",
}

// NewToolsBidSuggestV2InterestActionModeFromValue returns a pointer to a valid ToolsBidSuggestV2InterestActionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2InterestActionModeFromValue(v string) (*ToolsBidSuggestV2InterestActionMode, error) {
	ev := ToolsBidSuggestV2InterestActionMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2InterestActionMode: valid values are %v", v, AllowedToolsBidSuggestV2InterestActionModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2InterestActionMode) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2InterestActionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_interest_action_mode value
func (v ToolsBidSuggestV2InterestActionMode) Ptr() *ToolsBidSuggestV2InterestActionMode {
	return &v
}
