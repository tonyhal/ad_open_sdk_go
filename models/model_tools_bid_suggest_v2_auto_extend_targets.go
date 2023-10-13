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

// ToolsBidSuggestV2AutoExtendTargets
type ToolsBidSuggestV2AutoExtendTargets string

// List of tools_bid_suggest_v2_auto_extend_targets
const (
	CUSTOM_AUDIENCE_ToolsBidSuggestV2AutoExtendTargets ToolsBidSuggestV2AutoExtendTargets = "CUSTOM_AUDIENCE"
	INTEREST_TAG_ToolsBidSuggestV2AutoExtendTargets    ToolsBidSuggestV2AutoExtendTargets = "INTEREST_TAG"
	REGION_ToolsBidSuggestV2AutoExtendTargets          ToolsBidSuggestV2AutoExtendTargets = "REGION"
	GENDER_ToolsBidSuggestV2AutoExtendTargets          ToolsBidSuggestV2AutoExtendTargets = "GENDER"
	AD_TAG_ToolsBidSuggestV2AutoExtendTargets          ToolsBidSuggestV2AutoExtendTargets = "AD_TAG"
	AGE_ToolsBidSuggestV2AutoExtendTargets             ToolsBidSuggestV2AutoExtendTargets = "AGE"
)

// All allowed values of ToolsBidSuggestV2AutoExtendTargets enum
var AllowedToolsBidSuggestV2AutoExtendTargetsEnumValues = []ToolsBidSuggestV2AutoExtendTargets{
	"CUSTOM_AUDIENCE",
	"INTEREST_TAG",
	"REGION",
	"GENDER",
	"AD_TAG",
	"AGE",
}

// NewToolsBidSuggestV2AutoExtendTargetsFromValue returns a pointer to a valid ToolsBidSuggestV2AutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2AutoExtendTargetsFromValue(v string) (*ToolsBidSuggestV2AutoExtendTargets, error) {
	ev := ToolsBidSuggestV2AutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2AutoExtendTargets: valid values are %v", v, AllowedToolsBidSuggestV2AutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2AutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2AutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_auto_extend_targets value
func (v ToolsBidSuggestV2AutoExtendTargets) Ptr() *ToolsBidSuggestV2AutoExtendTargets {
	return &v
}
