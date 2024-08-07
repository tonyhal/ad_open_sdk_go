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

// ToolsAppIosListV2DataListStatus
type ToolsAppIosListV2DataListStatus string

// List of tools_app_ios_list_v2_data_list_status
const (
	AUDIT_REJECTED_ToolsAppIosListV2DataListStatus ToolsAppIosListV2DataListStatus = "AUDIT_REJECTED"
	ALL_ToolsAppIosListV2DataListStatus            ToolsAppIosListV2DataListStatus = "ALL"
	ENABLE_ToolsAppIosListV2DataListStatus         ToolsAppIosListV2DataListStatus = "ENABLE"
	AUDIT_DOING_ToolsAppIosListV2DataListStatus    ToolsAppIosListV2DataListStatus = "AUDIT_DOING"
	AUDIT_ACCEPTED_ToolsAppIosListV2DataListStatus ToolsAppIosListV2DataListStatus = "AUDIT_ACCEPTED"
)

// All allowed values of ToolsAppIosListV2DataListStatus enum
var AllowedToolsAppIosListV2DataListStatusEnumValues = []ToolsAppIosListV2DataListStatus{
	"AUDIT_REJECTED",
	"ALL",
	"ENABLE",
	"AUDIT_DOING",
	"AUDIT_ACCEPTED",
}

// NewToolsAppIosListV2DataListStatusFromValue returns a pointer to a valid ToolsAppIosListV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppIosListV2DataListStatusFromValue(v string) (*ToolsAppIosListV2DataListStatus, error) {
	ev := ToolsAppIosListV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppIosListV2DataListStatus: valid values are %v", v, AllowedToolsAppIosListV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppIosListV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedToolsAppIosListV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_ios_list_v2_data_list_status value
func (v ToolsAppIosListV2DataListStatus) Ptr() *ToolsAppIosListV2DataListStatus {
	return &v
}
