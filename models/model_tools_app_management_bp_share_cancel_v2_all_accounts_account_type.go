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

// ToolsAppManagementBpShareCancelV2AllAccountsAccountType
type ToolsAppManagementBpShareCancelV2AllAccountsAccountType string

// List of tools_app_management_bp_share_cancel_v2_all_accounts_account_type
const (
	BP_ToolsAppManagementBpShareCancelV2AllAccountsAccountType   ToolsAppManagementBpShareCancelV2AllAccountsAccountType = "BP"
	AD_ToolsAppManagementBpShareCancelV2AllAccountsAccountType   ToolsAppManagementBpShareCancelV2AllAccountsAccountType = "AD"
	STAR_ToolsAppManagementBpShareCancelV2AllAccountsAccountType ToolsAppManagementBpShareCancelV2AllAccountsAccountType = "STAR"
)

// All allowed values of ToolsAppManagementBpShareCancelV2AllAccountsAccountType enum
var AllowedToolsAppManagementBpShareCancelV2AllAccountsAccountTypeEnumValues = []ToolsAppManagementBpShareCancelV2AllAccountsAccountType{
	"BP",
	"AD",
	"STAR",
}

// NewToolsAppManagementBpShareCancelV2AllAccountsAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareCancelV2AllAccountsAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareCancelV2AllAccountsAccountTypeFromValue(v string) (*ToolsAppManagementBpShareCancelV2AllAccountsAccountType, error) {
	ev := ToolsAppManagementBpShareCancelV2AllAccountsAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareCancelV2AllAccountsAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareCancelV2AllAccountsAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareCancelV2AllAccountsAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareCancelV2AllAccountsAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_cancel_v2_all_accounts_account_type value
func (v ToolsAppManagementBpShareCancelV2AllAccountsAccountType) Ptr() *ToolsAppManagementBpShareCancelV2AllAccountsAccountType {
	return &v
}
