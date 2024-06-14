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

// ToolsBidSuggestV2FilterAwemeAbnormalActive
type ToolsBidSuggestV2FilterAwemeAbnormalActive int64

// List of tools_bid_suggest_v2_filter_aweme_abnormal_active
const (
	Enum_0_ToolsBidSuggestV2FilterAwemeAbnormalActive ToolsBidSuggestV2FilterAwemeAbnormalActive = 0
	Enum_1_ToolsBidSuggestV2FilterAwemeAbnormalActive ToolsBidSuggestV2FilterAwemeAbnormalActive = 1
)

// All allowed values of ToolsBidSuggestV2FilterAwemeAbnormalActive enum
var AllowedToolsBidSuggestV2FilterAwemeAbnormalActiveEnumValues = []ToolsBidSuggestV2FilterAwemeAbnormalActive{
	0,
	1,
}

// NewToolsBidSuggestV2FilterAwemeAbnormalActiveFromValue returns a pointer to a valid ToolsBidSuggestV2FilterAwemeAbnormalActive
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2FilterAwemeAbnormalActiveFromValue(v int64) (*ToolsBidSuggestV2FilterAwemeAbnormalActive, error) {
	ev := ToolsBidSuggestV2FilterAwemeAbnormalActive(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2FilterAwemeAbnormalActive: valid values are %v", v, AllowedToolsBidSuggestV2FilterAwemeAbnormalActiveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2FilterAwemeAbnormalActive) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2FilterAwemeAbnormalActiveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_filter_aweme_abnormal_active value
func (v ToolsBidSuggestV2FilterAwemeAbnormalActive) Ptr() *ToolsBidSuggestV2FilterAwemeAbnormalActive {
	return &v
}
