/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2ActivateType
type ToolsBidSuggestV2ActivateType string

// List of tools_bid_suggest_v2_activate_type
const (
	WITH_IN_A_MONTH_ToolsBidSuggestV2ActivateType         ToolsBidSuggestV2ActivateType = "WITH_IN_A_MONTH"
	ONE_MONTH_2_THREE_MONTH_ToolsBidSuggestV2ActivateType ToolsBidSuggestV2ActivateType = "ONE_MONTH_2_THREE_MONTH"
	UNLIMITED_ToolsBidSuggestV2ActivateType               ToolsBidSuggestV2ActivateType = "UNLIMITED"
	THREE_MONTH_EAILIER_ToolsBidSuggestV2ActivateType     ToolsBidSuggestV2ActivateType = "THREE_MONTH_EAILIER"
)

// All allowed values of ToolsBidSuggestV2ActivateType enum
var AllowedToolsBidSuggestV2ActivateTypeEnumValues = []ToolsBidSuggestV2ActivateType{
	"WITH_IN_A_MONTH",
	"ONE_MONTH_2_THREE_MONTH",
	"UNLIMITED",
	"THREE_MONTH_EAILIER",
}

// NewToolsBidSuggestV2ActivateTypeFromValue returns a pointer to a valid ToolsBidSuggestV2ActivateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2ActivateTypeFromValue(v string) (*ToolsBidSuggestV2ActivateType, error) {
	ev := ToolsBidSuggestV2ActivateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2ActivateType: valid values are %v", v, AllowedToolsBidSuggestV2ActivateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2ActivateType) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2ActivateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_activate_type value
func (v ToolsBidSuggestV2ActivateType) Ptr() *ToolsBidSuggestV2ActivateType {
	return &v
}
