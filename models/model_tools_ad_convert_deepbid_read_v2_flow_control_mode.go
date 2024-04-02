/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdConvertDeepbidReadV2FlowControlMode
type ToolsAdConvertDeepbidReadV2FlowControlMode string

// List of tools_ad_convert_deepbid_read_v2_flow_control_mode
const (
	FLOW_CONTROL_MODE_HOURLY_ToolsAdConvertDeepbidReadV2FlowControlMode     ToolsAdConvertDeepbidReadV2FlowControlMode = "FLOW_CONTROL_MODE_HOURLY"
	FLOW_CONTROL_MODE_BALANCE_ToolsAdConvertDeepbidReadV2FlowControlMode    ToolsAdConvertDeepbidReadV2FlowControlMode = "FLOW_CONTROL_MODE_BALANCE"
	FLOW_CONTROL_MODE_FAST_ToolsAdConvertDeepbidReadV2FlowControlMode       ToolsAdConvertDeepbidReadV2FlowControlMode = "FLOW_CONTROL_MODE_FAST"
	FLOW_CONTROL_MODE_SMOOTH_ToolsAdConvertDeepbidReadV2FlowControlMode     ToolsAdConvertDeepbidReadV2FlowControlMode = "FLOW_CONTROL_MODE_SMOOTH"
	FLOW_CONTROL_MODE_TWO_PHASES_ToolsAdConvertDeepbidReadV2FlowControlMode ToolsAdConvertDeepbidReadV2FlowControlMode = "FLOW_CONTROL_MODE_TWO_PHASES"
)

// All allowed values of ToolsAdConvertDeepbidReadV2FlowControlMode enum
var AllowedToolsAdConvertDeepbidReadV2FlowControlModeEnumValues = []ToolsAdConvertDeepbidReadV2FlowControlMode{
	"FLOW_CONTROL_MODE_HOURLY",
	"FLOW_CONTROL_MODE_BALANCE",
	"FLOW_CONTROL_MODE_FAST",
	"FLOW_CONTROL_MODE_SMOOTH",
	"FLOW_CONTROL_MODE_TWO_PHASES",
}

// NewToolsAdConvertDeepbidReadV2FlowControlModeFromValue returns a pointer to a valid ToolsAdConvertDeepbidReadV2FlowControlMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertDeepbidReadV2FlowControlModeFromValue(v string) (*ToolsAdConvertDeepbidReadV2FlowControlMode, error) {
	ev := ToolsAdConvertDeepbidReadV2FlowControlMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertDeepbidReadV2FlowControlMode: valid values are %v", v, AllowedToolsAdConvertDeepbidReadV2FlowControlModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertDeepbidReadV2FlowControlMode) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertDeepbidReadV2FlowControlModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_deepbid_read_v2_flow_control_mode value
func (v ToolsAdConvertDeepbidReadV2FlowControlMode) Ptr() *ToolsAdConvertDeepbidReadV2FlowControlMode {
	return &v
}
