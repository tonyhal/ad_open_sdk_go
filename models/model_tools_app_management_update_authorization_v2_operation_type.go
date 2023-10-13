/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementUpdateAuthorizationV2OperationType
type ToolsAppManagementUpdateAuthorizationV2OperationType string

// List of tools_app_management_update_authorization_v2_operation_type
const (
	ADD_ToolsAppManagementUpdateAuthorizationV2OperationType ToolsAppManagementUpdateAuthorizationV2OperationType = "ADD"
	DEL_ToolsAppManagementUpdateAuthorizationV2OperationType ToolsAppManagementUpdateAuthorizationV2OperationType = "DEL"
)

// All allowed values of ToolsAppManagementUpdateAuthorizationV2OperationType enum
var AllowedToolsAppManagementUpdateAuthorizationV2OperationTypeEnumValues = []ToolsAppManagementUpdateAuthorizationV2OperationType{
	"ADD",
	"DEL",
}

// NewToolsAppManagementUpdateAuthorizationV2OperationTypeFromValue returns a pointer to a valid ToolsAppManagementUpdateAuthorizationV2OperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementUpdateAuthorizationV2OperationTypeFromValue(v string) (*ToolsAppManagementUpdateAuthorizationV2OperationType, error) {
	ev := ToolsAppManagementUpdateAuthorizationV2OperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementUpdateAuthorizationV2OperationType: valid values are %v", v, AllowedToolsAppManagementUpdateAuthorizationV2OperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementUpdateAuthorizationV2OperationType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementUpdateAuthorizationV2OperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_update_authorization_v2_operation_type value
func (v ToolsAppManagementUpdateAuthorizationV2OperationType) Ptr() *ToolsAppManagementUpdateAuthorizationV2OperationType {
	return &v
}
