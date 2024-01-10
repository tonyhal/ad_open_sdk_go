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

// ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType
type ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType string

// List of tools_bp_asset_management_share_get_v3.0_data_sharedAccounts_account_info_account_type
const (
	AD_ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType = "AD"
	BP_ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType = "BP"
)

// All allowed values of ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType enum
var AllowedToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountTypeEnumValues = []ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType{
	"AD",
	"BP",
}

// NewToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountTypeFromValue returns a pointer to a valid ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountTypeFromValue(v string) (*ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType, error) {
	ev := ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType: valid values are %v", v, AllowedToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_get_v3.0_data_sharedAccounts_account_info_account_type value
func (v ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType) Ptr() *ToolsBpAssetManagementShareGetV30DataSharedAccountsAccountInfoAccountType {
	return &v
}
