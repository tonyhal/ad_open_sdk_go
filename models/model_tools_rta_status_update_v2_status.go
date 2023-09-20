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

// ToolsRtaStatusUpdateV2Status
type ToolsRtaStatusUpdateV2Status string

// List of tools_rta_status_update_v2_status
const (
	ENABLE_ToolsRtaStatusUpdateV2Status  ToolsRtaStatusUpdateV2Status = "ENABLE"
	DISABLE_ToolsRtaStatusUpdateV2Status ToolsRtaStatusUpdateV2Status = "DISABLE"
)

// All allowed values of ToolsRtaStatusUpdateV2Status enum
var AllowedToolsRtaStatusUpdateV2StatusEnumValues = []ToolsRtaStatusUpdateV2Status{
	"ENABLE",
	"DISABLE",
}

// NewToolsRtaStatusUpdateV2StatusFromValue returns a pointer to a valid ToolsRtaStatusUpdateV2Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRtaStatusUpdateV2StatusFromValue(v string) (*ToolsRtaStatusUpdateV2Status, error) {
	ev := ToolsRtaStatusUpdateV2Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRtaStatusUpdateV2Status: valid values are %v", v, AllowedToolsRtaStatusUpdateV2StatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRtaStatusUpdateV2Status) IsValid() bool {
	for _, existing := range AllowedToolsRtaStatusUpdateV2StatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rta_status_update_v2_status value
func (v ToolsRtaStatusUpdateV2Status) Ptr() *ToolsRtaStatusUpdateV2Status {
	return &v
}
