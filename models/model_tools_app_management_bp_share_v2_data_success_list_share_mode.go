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

// ToolsAppManagementBpShareV2DataSuccessListShareMode
type ToolsAppManagementBpShareV2DataSuccessListShareMode string

// List of tools_app_management_bp_share_v2_data_success_list_share_mode
const (
	ALL_ToolsAppManagementBpShareV2DataSuccessListShareMode     ToolsAppManagementBpShareV2DataSuccessListShareMode = "ALL"
	PART_ToolsAppManagementBpShareV2DataSuccessListShareMode    ToolsAppManagementBpShareV2DataSuccessListShareMode = "PART"
	COMPANY_ToolsAppManagementBpShareV2DataSuccessListShareMode ToolsAppManagementBpShareV2DataSuccessListShareMode = "COMPANY"
)

// All allowed values of ToolsAppManagementBpShareV2DataSuccessListShareMode enum
var AllowedToolsAppManagementBpShareV2DataSuccessListShareModeEnumValues = []ToolsAppManagementBpShareV2DataSuccessListShareMode{
	"ALL",
	"PART",
	"COMPANY",
}

// NewToolsAppManagementBpShareV2DataSuccessListShareModeFromValue returns a pointer to a valid ToolsAppManagementBpShareV2DataSuccessListShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareV2DataSuccessListShareModeFromValue(v string) (*ToolsAppManagementBpShareV2DataSuccessListShareMode, error) {
	ev := ToolsAppManagementBpShareV2DataSuccessListShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareV2DataSuccessListShareMode: valid values are %v", v, AllowedToolsAppManagementBpShareV2DataSuccessListShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareV2DataSuccessListShareMode) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareV2DataSuccessListShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_v2_data_success_list_share_mode value
func (v ToolsAppManagementBpShareV2DataSuccessListShareMode) Ptr() *ToolsAppManagementBpShareV2DataSuccessListShareMode {
	return &v
}
