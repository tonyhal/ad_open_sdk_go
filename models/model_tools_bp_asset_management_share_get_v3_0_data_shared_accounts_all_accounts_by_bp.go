/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBp
type ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBp string

// List of tools_bp_asset_management_share_get_v3.0_data_sharedAccounts_all_accounts_by_bp
const (
	AD_ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBp ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBp = "AD"
)

// All allowed values of ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBp enum
var AllowedToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBpEnumValues = []ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBp{
	"AD",
}

// NewToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBpFromValue returns a pointer to a valid ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBpFromValue(v string) (*ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBp, error) {
	ev := ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBp: valid values are %v", v, AllowedToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBp) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_get_v3.0_data_sharedAccounts_all_accounts_by_bp value
func (v ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBp) Ptr() *ToolsBpAssetManagementShareGetV30DataSharedAccountsAllAccountsByBp {
	return &v
}
