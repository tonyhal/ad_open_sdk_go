/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementBpShareV2ShareMode
type ToolsAppManagementBpShareV2ShareMode string

// List of tools_app_management_bp_share_v2_share_mode
const (
	ALL_ToolsAppManagementBpShareV2ShareMode     ToolsAppManagementBpShareV2ShareMode = "ALL"
	COMPANY_ToolsAppManagementBpShareV2ShareMode ToolsAppManagementBpShareV2ShareMode = "COMPANY"
	PART_ToolsAppManagementBpShareV2ShareMode    ToolsAppManagementBpShareV2ShareMode = "PART"
)

// All allowed values of ToolsAppManagementBpShareV2ShareMode enum
var AllowedToolsAppManagementBpShareV2ShareModeEnumValues = []ToolsAppManagementBpShareV2ShareMode{
	"ALL",
	"COMPANY",
	"PART",
}

// NewToolsAppManagementBpShareV2ShareModeFromValue returns a pointer to a valid ToolsAppManagementBpShareV2ShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareV2ShareModeFromValue(v string) (*ToolsAppManagementBpShareV2ShareMode, error) {
	ev := ToolsAppManagementBpShareV2ShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareV2ShareMode: valid values are %v", v, AllowedToolsAppManagementBpShareV2ShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareV2ShareMode) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareV2ShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_v2_share_mode value
func (v ToolsAppManagementBpShareV2ShareMode) Ptr() *ToolsAppManagementBpShareV2ShareMode {
	return &v
}
