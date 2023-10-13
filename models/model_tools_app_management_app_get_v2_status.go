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

// ToolsAppManagementAppGetV2Status
type ToolsAppManagementAppGetV2Status string

// List of tools_app_management_app_get_v2_status
const (
	ALL_ToolsAppManagementAppGetV2Status            ToolsAppManagementAppGetV2Status = "ALL"
	AUDIT_ACCEPTED_ToolsAppManagementAppGetV2Status ToolsAppManagementAppGetV2Status = "AUDIT_ACCEPTED"
	AUDIT_DOING_ToolsAppManagementAppGetV2Status    ToolsAppManagementAppGetV2Status = "AUDIT_DOING"
	AUDIT_REJECTED_ToolsAppManagementAppGetV2Status ToolsAppManagementAppGetV2Status = "AUDIT_REJECTED"
	ENABLE_ToolsAppManagementAppGetV2Status         ToolsAppManagementAppGetV2Status = "ENABLE"
)

// All allowed values of ToolsAppManagementAppGetV2Status enum
var AllowedToolsAppManagementAppGetV2StatusEnumValues = []ToolsAppManagementAppGetV2Status{
	"ALL",
	"AUDIT_ACCEPTED",
	"AUDIT_DOING",
	"AUDIT_REJECTED",
	"ENABLE",
}

// NewToolsAppManagementAppGetV2StatusFromValue returns a pointer to a valid ToolsAppManagementAppGetV2Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementAppGetV2StatusFromValue(v string) (*ToolsAppManagementAppGetV2Status, error) {
	ev := ToolsAppManagementAppGetV2Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementAppGetV2Status: valid values are %v", v, AllowedToolsAppManagementAppGetV2StatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementAppGetV2Status) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementAppGetV2StatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_app_get_v2_status value
func (v ToolsAppManagementAppGetV2Status) Ptr() *ToolsAppManagementAppGetV2Status {
	return &v
}
