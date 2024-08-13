/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementAndroidBasicPackagePublishV2AccountType
type ToolsAppManagementAndroidBasicPackagePublishV2AccountType string

// List of tools_app_management_android_basic_package_publish_v2_account_type
const (
	AD_ToolsAppManagementAndroidBasicPackagePublishV2AccountType ToolsAppManagementAndroidBasicPackagePublishV2AccountType = "AD"
	BP_ToolsAppManagementAndroidBasicPackagePublishV2AccountType ToolsAppManagementAndroidBasicPackagePublishV2AccountType = "BP"
)

// All allowed values of ToolsAppManagementAndroidBasicPackagePublishV2AccountType enum
var AllowedToolsAppManagementAndroidBasicPackagePublishV2AccountTypeEnumValues = []ToolsAppManagementAndroidBasicPackagePublishV2AccountType{
	"AD",
	"BP",
}

// NewToolsAppManagementAndroidBasicPackagePublishV2AccountTypeFromValue returns a pointer to a valid ToolsAppManagementAndroidBasicPackagePublishV2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementAndroidBasicPackagePublishV2AccountTypeFromValue(v string) (*ToolsAppManagementAndroidBasicPackagePublishV2AccountType, error) {
	ev := ToolsAppManagementAndroidBasicPackagePublishV2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementAndroidBasicPackagePublishV2AccountType: valid values are %v", v, AllowedToolsAppManagementAndroidBasicPackagePublishV2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementAndroidBasicPackagePublishV2AccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementAndroidBasicPackagePublishV2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_android_basic_package_publish_v2_account_type value
func (v ToolsAppManagementAndroidBasicPackagePublishV2AccountType) Ptr() *ToolsAppManagementAndroidBasicPackagePublishV2AccountType {
	return &v
}
