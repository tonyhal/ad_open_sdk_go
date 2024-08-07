/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementExtendPackageUpdateV2Mode
type ToolsAppManagementExtendPackageUpdateV2Mode string

// List of tools_app_management_extend_package_update_v2_mode
const (
	ALL_ToolsAppManagementExtendPackageUpdateV2Mode       ToolsAppManagementExtendPackageUpdateV2Mode = "All"
	CUSTOMIZE_ToolsAppManagementExtendPackageUpdateV2Mode ToolsAppManagementExtendPackageUpdateV2Mode = "Customize"
	LIST_ToolsAppManagementExtendPackageUpdateV2Mode      ToolsAppManagementExtendPackageUpdateV2Mode = "List"
)

// All allowed values of ToolsAppManagementExtendPackageUpdateV2Mode enum
var AllowedToolsAppManagementExtendPackageUpdateV2ModeEnumValues = []ToolsAppManagementExtendPackageUpdateV2Mode{
	"All",
	"Customize",
	"List",
}

// NewToolsAppManagementExtendPackageUpdateV2ModeFromValue returns a pointer to a valid ToolsAppManagementExtendPackageUpdateV2Mode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementExtendPackageUpdateV2ModeFromValue(v string) (*ToolsAppManagementExtendPackageUpdateV2Mode, error) {
	ev := ToolsAppManagementExtendPackageUpdateV2Mode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementExtendPackageUpdateV2Mode: valid values are %v", v, AllowedToolsAppManagementExtendPackageUpdateV2ModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementExtendPackageUpdateV2Mode) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementExtendPackageUpdateV2ModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_extend_package_update_v2_mode value
func (v ToolsAppManagementExtendPackageUpdateV2Mode) Ptr() *ToolsAppManagementExtendPackageUpdateV2Mode {
	return &v
}
