/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2BidMode
type ToolsBidSuggestV2BidMode string

// List of tools_bid_suggest_v2_bid_mode
const (
	SUGGEST_ToolsBidSuggestV2BidMode  ToolsBidSuggestV2BidMode = "SUGGEST"
	AUTO_BID_ToolsBidSuggestV2BidMode ToolsBidSuggestV2BidMode = "AUTO_BID"
)

// All allowed values of ToolsBidSuggestV2BidMode enum
var AllowedToolsBidSuggestV2BidModeEnumValues = []ToolsBidSuggestV2BidMode{
	"SUGGEST",
	"AUTO_BID",
}

// NewToolsBidSuggestV2BidModeFromValue returns a pointer to a valid ToolsBidSuggestV2BidMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2BidModeFromValue(v string) (*ToolsBidSuggestV2BidMode, error) {
	ev := ToolsBidSuggestV2BidMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2BidMode: valid values are %v", v, AllowedToolsBidSuggestV2BidModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2BidMode) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2BidModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_bid_mode value
func (v ToolsBidSuggestV2BidMode) Ptr() *ToolsBidSuggestV2BidMode {
	return &v
}
