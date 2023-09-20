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

// ToolsBidSuggestV2SuperiorPopularityType
type ToolsBidSuggestV2SuperiorPopularityType string

// List of tools_bid_suggest_v2_superior_popularity_type
const (
	GAME_ToolsBidSuggestV2SuperiorPopularityType ToolsBidSuggestV2SuperiorPopularityType = "GAME"
	NONE_ToolsBidSuggestV2SuperiorPopularityType ToolsBidSuggestV2SuperiorPopularityType = "NONE"
)

// All allowed values of ToolsBidSuggestV2SuperiorPopularityType enum
var AllowedToolsBidSuggestV2SuperiorPopularityTypeEnumValues = []ToolsBidSuggestV2SuperiorPopularityType{
	"GAME",
	"NONE",
}

// NewToolsBidSuggestV2SuperiorPopularityTypeFromValue returns a pointer to a valid ToolsBidSuggestV2SuperiorPopularityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2SuperiorPopularityTypeFromValue(v string) (*ToolsBidSuggestV2SuperiorPopularityType, error) {
	ev := ToolsBidSuggestV2SuperiorPopularityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2SuperiorPopularityType: valid values are %v", v, AllowedToolsBidSuggestV2SuperiorPopularityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2SuperiorPopularityType) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2SuperiorPopularityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_superior_popularity_type value
func (v ToolsBidSuggestV2SuperiorPopularityType) Ptr() *ToolsBidSuggestV2SuperiorPopularityType {
	return &v
}