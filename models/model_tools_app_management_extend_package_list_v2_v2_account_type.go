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

// ToolsAppManagementExtendPackageListV2V2AccountType
type ToolsAppManagementExtendPackageListV2V2AccountType string

// List of tools_app_management_extend_package_list_v2_v2_account_type
const (
	AD_ToolsAppManagementExtendPackageListV2V2AccountType   ToolsAppManagementExtendPackageListV2V2AccountType = "AD"
	BP_ToolsAppManagementExtendPackageListV2V2AccountType   ToolsAppManagementExtendPackageListV2V2AccountType = "BP"
	STAR_ToolsAppManagementExtendPackageListV2V2AccountType ToolsAppManagementExtendPackageListV2V2AccountType = "STAR"
)

// All allowed values of ToolsAppManagementExtendPackageListV2V2AccountType enum
var AllowedToolsAppManagementExtendPackageListV2V2AccountTypeEnumValues = []ToolsAppManagementExtendPackageListV2V2AccountType{
	"AD",
	"BP",
	"STAR",
}

// NewToolsAppManagementExtendPackageListV2V2AccountTypeFromValue returns a pointer to a valid ToolsAppManagementExtendPackageListV2V2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementExtendPackageListV2V2AccountTypeFromValue(v string) (*ToolsAppManagementExtendPackageListV2V2AccountType, error) {
	ev := ToolsAppManagementExtendPackageListV2V2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementExtendPackageListV2V2AccountType: valid values are %v", v, AllowedToolsAppManagementExtendPackageListV2V2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementExtendPackageListV2V2AccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementExtendPackageListV2V2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_extend_package_list_v2_v2_account_type value
func (v ToolsAppManagementExtendPackageListV2V2AccountType) Ptr() *ToolsAppManagementExtendPackageListV2V2AccountType {
	return &v
}
