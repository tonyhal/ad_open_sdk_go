/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType
type ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType string

// List of tools_app_management_bp_share_cancel_v2_all_account_by_company_account_type
const (
	BP_ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType   ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType = "BP"
	AD_ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType   ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType = "AD"
	STAR_ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType = "STAR"
)

// All allowed values of ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType enum
var AllowedToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountTypeEnumValues = []ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType{
	"BP",
	"AD",
	"STAR",
}

// NewToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountTypeFromValue(v string) (*ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType, error) {
	ev := ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_cancel_v2_all_account_by_company_account_type value
func (v ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType) Ptr() *ToolsAppManagementBpShareCancelV2AllAccountByCompanyAccountType {
	return &v
}
