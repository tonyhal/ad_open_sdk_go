/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag
type ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag string

// List of tools_app_management_android_basic_package_update_v2_files_file_tag
const (
	AGE_REMINDER_ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag        ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag = "AGE_REMINDER"
	ANTI_ADDICTION_TIPS_ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag = "ANTI_ADDICTION_TIPS"
	DEFAULT_ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag             ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag = "DEFAULT"
	MATERIAL_SCREENSHOT_ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag = "MATERIAL_SCREENSHOT"
	REAL_NAME_VERIFIED_ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag  ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag = "REAL_NAME_VERIFIED"
)

// All allowed values of ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag enum
var AllowedToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTagEnumValues = []ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag{
	"AGE_REMINDER",
	"ANTI_ADDICTION_TIPS",
	"DEFAULT",
	"MATERIAL_SCREENSHOT",
	"REAL_NAME_VERIFIED",
}

// NewToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTagFromValue returns a pointer to a valid ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTagFromValue(v string) (*ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag, error) {
	ev := ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag: valid values are %v", v, AllowedToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTagEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTagEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_android_basic_package_update_v2_files_file_tag value
func (v ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag) Ptr() *ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag {
	return &v
}
