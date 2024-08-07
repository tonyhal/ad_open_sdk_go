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

// ToolsAppManagementUploadTaskListV2AccountType
type ToolsAppManagementUploadTaskListV2AccountType string

// List of tools_app_management_upload_task_list_v2_account_type
const (
	AD_ToolsAppManagementUploadTaskListV2AccountType ToolsAppManagementUploadTaskListV2AccountType = "AD"
	BP_ToolsAppManagementUploadTaskListV2AccountType ToolsAppManagementUploadTaskListV2AccountType = "BP"
)

// All allowed values of ToolsAppManagementUploadTaskListV2AccountType enum
var AllowedToolsAppManagementUploadTaskListV2AccountTypeEnumValues = []ToolsAppManagementUploadTaskListV2AccountType{
	"AD",
	"BP",
}

// NewToolsAppManagementUploadTaskListV2AccountTypeFromValue returns a pointer to a valid ToolsAppManagementUploadTaskListV2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementUploadTaskListV2AccountTypeFromValue(v string) (*ToolsAppManagementUploadTaskListV2AccountType, error) {
	ev := ToolsAppManagementUploadTaskListV2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementUploadTaskListV2AccountType: valid values are %v", v, AllowedToolsAppManagementUploadTaskListV2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementUploadTaskListV2AccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementUploadTaskListV2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_upload_task_list_v2_account_type value
func (v ToolsAppManagementUploadTaskListV2AccountType) Ptr() *ToolsAppManagementUploadTaskListV2AccountType {
	return &v
}
