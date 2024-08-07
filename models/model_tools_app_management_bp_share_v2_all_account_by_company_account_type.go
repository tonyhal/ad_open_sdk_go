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

// ToolsAppManagementBpShareV2AllAccountByCompanyAccountType
type ToolsAppManagementBpShareV2AllAccountByCompanyAccountType string

// List of tools_app_management_bp_share_v2_all_account_by_company_account_type
const (
	AD_ToolsAppManagementBpShareV2AllAccountByCompanyAccountType   ToolsAppManagementBpShareV2AllAccountByCompanyAccountType = "AD"
	BP_ToolsAppManagementBpShareV2AllAccountByCompanyAccountType   ToolsAppManagementBpShareV2AllAccountByCompanyAccountType = "BP"
	STAR_ToolsAppManagementBpShareV2AllAccountByCompanyAccountType ToolsAppManagementBpShareV2AllAccountByCompanyAccountType = "STAR"
)

// All allowed values of ToolsAppManagementBpShareV2AllAccountByCompanyAccountType enum
var AllowedToolsAppManagementBpShareV2AllAccountByCompanyAccountTypeEnumValues = []ToolsAppManagementBpShareV2AllAccountByCompanyAccountType{
	"AD",
	"BP",
	"STAR",
}

// NewToolsAppManagementBpShareV2AllAccountByCompanyAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareV2AllAccountByCompanyAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareV2AllAccountByCompanyAccountTypeFromValue(v string) (*ToolsAppManagementBpShareV2AllAccountByCompanyAccountType, error) {
	ev := ToolsAppManagementBpShareV2AllAccountByCompanyAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareV2AllAccountByCompanyAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareV2AllAccountByCompanyAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareV2AllAccountByCompanyAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareV2AllAccountByCompanyAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_v2_all_account_by_company_account_type value
func (v ToolsAppManagementBpShareV2AllAccountByCompanyAccountType) Ptr() *ToolsAppManagementBpShareV2AllAccountByCompanyAccountType {
	return &v
}
