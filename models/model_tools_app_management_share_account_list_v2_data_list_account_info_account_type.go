/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType
type ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType string

// List of tools_app_management_share_account_list_v2_data_list_account_info_account_type
const (
	STAR_ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType = "STAR"
	BP_ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType   ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType = "BP"
	AD_ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType   ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType = "AD"
)

// All allowed values of ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType enum
var AllowedToolsAppManagementShareAccountListV2DataListAccountInfoAccountTypeEnumValues = []ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType{
	"STAR",
	"BP",
	"AD",
}

// NewToolsAppManagementShareAccountListV2DataListAccountInfoAccountTypeFromValue returns a pointer to a valid ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementShareAccountListV2DataListAccountInfoAccountTypeFromValue(v string) (*ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType, error) {
	ev := ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType: valid values are %v", v, AllowedToolsAppManagementShareAccountListV2DataListAccountInfoAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementShareAccountListV2DataListAccountInfoAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_share_account_list_v2_data_list_account_info_account_type value
func (v ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType) Ptr() *ToolsAppManagementShareAccountListV2DataListAccountInfoAccountType {
	return &v
}
