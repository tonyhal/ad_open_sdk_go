/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2FlowControlMode
type ToolsBidSuggestV2FlowControlMode string

// List of tools_bid_suggest_v2_flow_control_mode
const (
	FLOW_CONTROL_MODE_SMOOTH_ToolsBidSuggestV2FlowControlMode     ToolsBidSuggestV2FlowControlMode = "FLOW_CONTROL_MODE_SMOOTH"
	FLOW_CONTROL_MODE_TWO_PHASES_ToolsBidSuggestV2FlowControlMode ToolsBidSuggestV2FlowControlMode = "FLOW_CONTROL_MODE_TWO_PHASES"
	FLOW_CONTROL_MODE_HOURLY_ToolsBidSuggestV2FlowControlMode     ToolsBidSuggestV2FlowControlMode = "FLOW_CONTROL_MODE_HOURLY"
	FLOW_CONTROL_MODE_FAST_ToolsBidSuggestV2FlowControlMode       ToolsBidSuggestV2FlowControlMode = "FLOW_CONTROL_MODE_FAST"
	FLOW_CONTROL_MODE_BALANCE_ToolsBidSuggestV2FlowControlMode    ToolsBidSuggestV2FlowControlMode = "FLOW_CONTROL_MODE_BALANCE"
)

// All allowed values of ToolsBidSuggestV2FlowControlMode enum
var AllowedToolsBidSuggestV2FlowControlModeEnumValues = []ToolsBidSuggestV2FlowControlMode{
	"FLOW_CONTROL_MODE_SMOOTH",
	"FLOW_CONTROL_MODE_TWO_PHASES",
	"FLOW_CONTROL_MODE_HOURLY",
	"FLOW_CONTROL_MODE_FAST",
	"FLOW_CONTROL_MODE_BALANCE",
}

// NewToolsBidSuggestV2FlowControlModeFromValue returns a pointer to a valid ToolsBidSuggestV2FlowControlMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2FlowControlModeFromValue(v string) (*ToolsBidSuggestV2FlowControlMode, error) {
	ev := ToolsBidSuggestV2FlowControlMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2FlowControlMode: valid values are %v", v, AllowedToolsBidSuggestV2FlowControlModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2FlowControlMode) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2FlowControlModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_flow_control_mode value
func (v ToolsBidSuggestV2FlowControlMode) Ptr() *ToolsBidSuggestV2FlowControlMode {
	return &v
}
