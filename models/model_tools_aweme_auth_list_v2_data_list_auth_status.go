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

// ToolsAwemeAuthListV2DataListAuthStatus
type ToolsAwemeAuthListV2DataListAuthStatus string

// List of tools_aweme_auth_list_v2_data_list_auth_status
const (
	AUTHRIZED_ToolsAwemeAuthListV2DataListAuthStatus  ToolsAwemeAuthListV2DataListAuthStatus = "AUTHRIZED"
	AUTHRIZING_ToolsAwemeAuthListV2DataListAuthStatus ToolsAwemeAuthListV2DataListAuthStatus = "AUTHRIZING"
	INVALID_ToolsAwemeAuthListV2DataListAuthStatus    ToolsAwemeAuthListV2DataListAuthStatus = "INVALID"
)

// All allowed values of ToolsAwemeAuthListV2DataListAuthStatus enum
var AllowedToolsAwemeAuthListV2DataListAuthStatusEnumValues = []ToolsAwemeAuthListV2DataListAuthStatus{
	"AUTHRIZED",
	"AUTHRIZING",
	"INVALID",
}

// NewToolsAwemeAuthListV2DataListAuthStatusFromValue returns a pointer to a valid ToolsAwemeAuthListV2DataListAuthStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeAuthListV2DataListAuthStatusFromValue(v string) (*ToolsAwemeAuthListV2DataListAuthStatus, error) {
	ev := ToolsAwemeAuthListV2DataListAuthStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeAuthListV2DataListAuthStatus: valid values are %v", v, AllowedToolsAwemeAuthListV2DataListAuthStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeAuthListV2DataListAuthStatus) IsValid() bool {
	for _, existing := range AllowedToolsAwemeAuthListV2DataListAuthStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_auth_list_v2_data_list_auth_status value
func (v ToolsAwemeAuthListV2DataListAuthStatus) Ptr() *ToolsAwemeAuthListV2DataListAuthStatus {
	return &v
}
