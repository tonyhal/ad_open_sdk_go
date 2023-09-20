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

// ToolsAppManagementExtendPackageListV2V2DataListStatus
type ToolsAppManagementExtendPackageListV2V2DataListStatus string

// List of tools_app_management_extend_package_list_v2_v2_data_list_status
const (
	ALL_ToolsAppManagementExtendPackageListV2V2DataListStatus             ToolsAppManagementExtendPackageListV2V2DataListStatus = "ALL"
	NOT_UPDATE_ToolsAppManagementExtendPackageListV2V2DataListStatus      ToolsAppManagementExtendPackageListV2V2DataListStatus = "NOT_UPDATE"
	CREATING_ToolsAppManagementExtendPackageListV2V2DataListStatus        ToolsAppManagementExtendPackageListV2V2DataListStatus = "CREATING"
	UPDATING_ToolsAppManagementExtendPackageListV2V2DataListStatus        ToolsAppManagementExtendPackageListV2V2DataListStatus = "UPDATING"
	ENABLE_ToolsAppManagementExtendPackageListV2V2DataListStatus          ToolsAppManagementExtendPackageListV2V2DataListStatus = "ENABLE"
	CREATION_FAILED_ToolsAppManagementExtendPackageListV2V2DataListStatus ToolsAppManagementExtendPackageListV2V2DataListStatus = "CREATION_FAILED"
	UPDATE_FAILED_ToolsAppManagementExtendPackageListV2V2DataListStatus   ToolsAppManagementExtendPackageListV2V2DataListStatus = "UPDATE_FAILED"
	OFF_SHELF_ToolsAppManagementExtendPackageListV2V2DataListStatus       ToolsAppManagementExtendPackageListV2V2DataListStatus = "OFF_SHELF"
)

// All allowed values of ToolsAppManagementExtendPackageListV2V2DataListStatus enum
var AllowedToolsAppManagementExtendPackageListV2V2DataListStatusEnumValues = []ToolsAppManagementExtendPackageListV2V2DataListStatus{
	"ALL",
	"NOT_UPDATE",
	"CREATING",
	"UPDATING",
	"ENABLE",
	"CREATION_FAILED",
	"UPDATE_FAILED",
	"OFF_SHELF",
}

// NewToolsAppManagementExtendPackageListV2V2DataListStatusFromValue returns a pointer to a valid ToolsAppManagementExtendPackageListV2V2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementExtendPackageListV2V2DataListStatusFromValue(v string) (*ToolsAppManagementExtendPackageListV2V2DataListStatus, error) {
	ev := ToolsAppManagementExtendPackageListV2V2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementExtendPackageListV2V2DataListStatus: valid values are %v", v, AllowedToolsAppManagementExtendPackageListV2V2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementExtendPackageListV2V2DataListStatus) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementExtendPackageListV2V2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_extend_package_list_v2_v2_data_list_status value
func (v ToolsAppManagementExtendPackageListV2V2DataListStatus) Ptr() *ToolsAppManagementExtendPackageListV2V2DataListStatus {
	return &v
}