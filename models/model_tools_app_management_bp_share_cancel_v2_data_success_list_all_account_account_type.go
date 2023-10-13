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

// ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType
type ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType string

// List of tools_app_management_bp_share_cancel_v2_data_success_list_all_account_account_type
const (
	AD_ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType   ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType = "AD"
	BP_ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType   ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType = "BP"
	STAR_ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType = "STAR"
)

// All allowed values of ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType enum
var AllowedToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountTypeEnumValues = []ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType{
	"AD",
	"BP",
	"STAR",
}

// NewToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountTypeFromValue(v string) (*ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType, error) {
	ev := ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_cancel_v2_data_success_list_all_account_account_type value
func (v ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType) Ptr() *ToolsAppManagementBpShareCancelV2DataSuccessListAllAccountAccountType {
	return &v
}
