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

// ToolsAppManagementBpShareCancelV2AccountInfosAccountType
type ToolsAppManagementBpShareCancelV2AccountInfosAccountType string

// List of tools_app_management_bp_share_cancel_v2_account_infos_account_type
const (
	STAR_ToolsAppManagementBpShareCancelV2AccountInfosAccountType ToolsAppManagementBpShareCancelV2AccountInfosAccountType = "STAR"
	AD_ToolsAppManagementBpShareCancelV2AccountInfosAccountType   ToolsAppManagementBpShareCancelV2AccountInfosAccountType = "AD"
	BP_ToolsAppManagementBpShareCancelV2AccountInfosAccountType   ToolsAppManagementBpShareCancelV2AccountInfosAccountType = "BP"
)

// All allowed values of ToolsAppManagementBpShareCancelV2AccountInfosAccountType enum
var AllowedToolsAppManagementBpShareCancelV2AccountInfosAccountTypeEnumValues = []ToolsAppManagementBpShareCancelV2AccountInfosAccountType{
	"STAR",
	"AD",
	"BP",
}

// NewToolsAppManagementBpShareCancelV2AccountInfosAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareCancelV2AccountInfosAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareCancelV2AccountInfosAccountTypeFromValue(v string) (*ToolsAppManagementBpShareCancelV2AccountInfosAccountType, error) {
	ev := ToolsAppManagementBpShareCancelV2AccountInfosAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareCancelV2AccountInfosAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareCancelV2AccountInfosAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareCancelV2AccountInfosAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareCancelV2AccountInfosAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_cancel_v2_account_infos_account_type value
func (v ToolsAppManagementBpShareCancelV2AccountInfosAccountType) Ptr() *ToolsAppManagementBpShareCancelV2AccountInfosAccountType {
	return &v
}
