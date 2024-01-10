/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementExtendPackageUpdateV2V2Mode
type ToolsAppManagementExtendPackageUpdateV2V2Mode string

// List of tools_app_management_extend_package_update_v2_v2_mode
const (
	ALL_ToolsAppManagementExtendPackageUpdateV2V2Mode       ToolsAppManagementExtendPackageUpdateV2V2Mode = "All"
	LIST_ToolsAppManagementExtendPackageUpdateV2V2Mode      ToolsAppManagementExtendPackageUpdateV2V2Mode = "List"
	CUSTOMIZE_ToolsAppManagementExtendPackageUpdateV2V2Mode ToolsAppManagementExtendPackageUpdateV2V2Mode = "Customize"
)

// All allowed values of ToolsAppManagementExtendPackageUpdateV2V2Mode enum
var AllowedToolsAppManagementExtendPackageUpdateV2V2ModeEnumValues = []ToolsAppManagementExtendPackageUpdateV2V2Mode{
	"All",
	"List",
	"Customize",
}

// NewToolsAppManagementExtendPackageUpdateV2V2ModeFromValue returns a pointer to a valid ToolsAppManagementExtendPackageUpdateV2V2Mode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementExtendPackageUpdateV2V2ModeFromValue(v string) (*ToolsAppManagementExtendPackageUpdateV2V2Mode, error) {
	ev := ToolsAppManagementExtendPackageUpdateV2V2Mode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementExtendPackageUpdateV2V2Mode: valid values are %v", v, AllowedToolsAppManagementExtendPackageUpdateV2V2ModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementExtendPackageUpdateV2V2Mode) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementExtendPackageUpdateV2V2ModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_extend_package_update_v2_v2_mode value
func (v ToolsAppManagementExtendPackageUpdateV2V2Mode) Ptr() *ToolsAppManagementExtendPackageUpdateV2V2Mode {
	return &v
}
