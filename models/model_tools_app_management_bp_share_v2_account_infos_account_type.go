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

// ToolsAppManagementBpShareV2AccountInfosAccountType
type ToolsAppManagementBpShareV2AccountInfosAccountType string

// List of tools_app_management_bp_share_v2_account_infos_account_type
const (
	STAR_ToolsAppManagementBpShareV2AccountInfosAccountType ToolsAppManagementBpShareV2AccountInfosAccountType = "STAR"
	BP_ToolsAppManagementBpShareV2AccountInfosAccountType   ToolsAppManagementBpShareV2AccountInfosAccountType = "BP"
	AD_ToolsAppManagementBpShareV2AccountInfosAccountType   ToolsAppManagementBpShareV2AccountInfosAccountType = "AD"
)

// All allowed values of ToolsAppManagementBpShareV2AccountInfosAccountType enum
var AllowedToolsAppManagementBpShareV2AccountInfosAccountTypeEnumValues = []ToolsAppManagementBpShareV2AccountInfosAccountType{
	"STAR",
	"BP",
	"AD",
}

// NewToolsAppManagementBpShareV2AccountInfosAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareV2AccountInfosAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareV2AccountInfosAccountTypeFromValue(v string) (*ToolsAppManagementBpShareV2AccountInfosAccountType, error) {
	ev := ToolsAppManagementBpShareV2AccountInfosAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareV2AccountInfosAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareV2AccountInfosAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareV2AccountInfosAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareV2AccountInfosAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_v2_account_infos_account_type value
func (v ToolsAppManagementBpShareV2AccountInfosAccountType) Ptr() *ToolsAppManagementBpShareV2AccountInfosAccountType {
	return &v
}
