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

// ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType
type ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType string

// List of tools_app_management_share_account_list_v2_data_list_all_account_by_company_account_type
const (
	STAR_ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType = "STAR"
	BP_ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType   ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType = "BP"
	AD_ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType   ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType = "AD"
)

// All allowed values of ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType enum
var AllowedToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountTypeEnumValues = []ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType{
	"STAR",
	"BP",
	"AD",
}

// NewToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountTypeFromValue returns a pointer to a valid ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountTypeFromValue(v string) (*ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType, error) {
	ev := ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType: valid values are %v", v, AllowedToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_share_account_list_v2_data_list_all_account_by_company_account_type value
func (v ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType) Ptr() *ToolsAppManagementShareAccountListV2DataListAllAccountByCompanyAccountType {
	return &v
}
