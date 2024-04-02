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

// ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType
type ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType string

// List of tools_app_management_bp_share_v2_data_success_list_all_account_by_company_account_type
const (
	STAR_ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType = "STAR"
	AD_ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType   ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType = "AD"
	BP_ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType   ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType = "BP"
)

// All allowed values of ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType enum
var AllowedToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountTypeEnumValues = []ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType{
	"STAR",
	"AD",
	"BP",
}

// NewToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountTypeFromValue(v string) (*ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType, error) {
	ev := ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_v2_data_success_list_all_account_by_company_account_type value
func (v ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType) Ptr() *ToolsAppManagementBpShareV2DataSuccessListAllAccountByCompanyAccountType {
	return &v
}
