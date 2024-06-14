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

// AdGetV2DataFlowControlMode
type AdGetV2DataFlowControlMode string

// List of ad_get_v2_data_flow_control_mode
const (
	FLOW_CONTROL_MODE_BALANCE_AdGetV2DataFlowControlMode    AdGetV2DataFlowControlMode = "FLOW_CONTROL_MODE_BALANCE"
	FLOW_CONTROL_MODE_SMOOTH_AdGetV2DataFlowControlMode     AdGetV2DataFlowControlMode = "FLOW_CONTROL_MODE_SMOOTH"
	FLOW_CONTROL_MODE_FAST_AdGetV2DataFlowControlMode       AdGetV2DataFlowControlMode = "FLOW_CONTROL_MODE_FAST"
	FLOW_CONTROL_MODE_HOURLY_AdGetV2DataFlowControlMode     AdGetV2DataFlowControlMode = "FLOW_CONTROL_MODE_HOURLY"
	FLOW_CONTROL_MODE_TWO_PHASES_AdGetV2DataFlowControlMode AdGetV2DataFlowControlMode = "FLOW_CONTROL_MODE_TWO_PHASES"
)

// All allowed values of AdGetV2DataFlowControlMode enum
var AllowedAdGetV2DataFlowControlModeEnumValues = []AdGetV2DataFlowControlMode{
	"FLOW_CONTROL_MODE_BALANCE",
	"FLOW_CONTROL_MODE_SMOOTH",
	"FLOW_CONTROL_MODE_FAST",
	"FLOW_CONTROL_MODE_HOURLY",
	"FLOW_CONTROL_MODE_TWO_PHASES",
}

// NewAdGetV2DataFlowControlModeFromValue returns a pointer to a valid AdGetV2DataFlowControlMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataFlowControlModeFromValue(v string) (*AdGetV2DataFlowControlMode, error) {
	ev := AdGetV2DataFlowControlMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataFlowControlMode: valid values are %v", v, AllowedAdGetV2DataFlowControlModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataFlowControlMode) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataFlowControlModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_flow_control_mode value
func (v AdGetV2DataFlowControlMode) Ptr() *AdGetV2DataFlowControlMode {
	return &v
}
