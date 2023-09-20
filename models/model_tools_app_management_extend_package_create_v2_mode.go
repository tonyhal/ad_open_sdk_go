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

// ToolsAppManagementExtendPackageCreateV2Mode
type ToolsAppManagementExtendPackageCreateV2Mode string

// List of tools_app_management_extend_package_create_v2_mode
const (
	MANUAL_ToolsAppManagementExtendPackageCreateV2Mode    ToolsAppManagementExtendPackageCreateV2Mode = "Manual"
	CUSTOMIZE_ToolsAppManagementExtendPackageCreateV2Mode ToolsAppManagementExtendPackageCreateV2Mode = "Customize"
	AUTO_ToolsAppManagementExtendPackageCreateV2Mode      ToolsAppManagementExtendPackageCreateV2Mode = "Auto"
)

// All allowed values of ToolsAppManagementExtendPackageCreateV2Mode enum
var AllowedToolsAppManagementExtendPackageCreateV2ModeEnumValues = []ToolsAppManagementExtendPackageCreateV2Mode{
	"Manual",
	"Customize",
	"Auto",
}

// NewToolsAppManagementExtendPackageCreateV2ModeFromValue returns a pointer to a valid ToolsAppManagementExtendPackageCreateV2Mode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementExtendPackageCreateV2ModeFromValue(v string) (*ToolsAppManagementExtendPackageCreateV2Mode, error) {
	ev := ToolsAppManagementExtendPackageCreateV2Mode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementExtendPackageCreateV2Mode: valid values are %v", v, AllowedToolsAppManagementExtendPackageCreateV2ModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementExtendPackageCreateV2Mode) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementExtendPackageCreateV2ModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_extend_package_create_v2_mode value
func (v ToolsAppManagementExtendPackageCreateV2Mode) Ptr() *ToolsAppManagementExtendPackageCreateV2Mode {
	return &v
}
