/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPlayableSaveV2DataStatus
type ToolsPlayableSaveV2DataStatus string

// List of tools_playable_save_v2_data_status
const (
	AUDIT_FAIL_ToolsPlayableSaveV2DataStatus       ToolsPlayableSaveV2DataStatus = "AUDIT_FAIL"
	AUDIT_SUCCESS_ToolsPlayableSaveV2DataStatus    ToolsPlayableSaveV2DataStatus = "AUDIT_SUCCESS"
	VALIDATE_FAIL_ToolsPlayableSaveV2DataStatus    ToolsPlayableSaveV2DataStatus = "VALIDATE_FAIL"
	VALIDATE_SUCCESS_ToolsPlayableSaveV2DataStatus ToolsPlayableSaveV2DataStatus = "VALIDATE_SUCCESS"
	VALIDATING_ToolsPlayableSaveV2DataStatus       ToolsPlayableSaveV2DataStatus = "VALIDATING"
)

// All allowed values of ToolsPlayableSaveV2DataStatus enum
var AllowedToolsPlayableSaveV2DataStatusEnumValues = []ToolsPlayableSaveV2DataStatus{
	"AUDIT_FAIL",
	"AUDIT_SUCCESS",
	"VALIDATE_FAIL",
	"VALIDATE_SUCCESS",
	"VALIDATING",
}

// NewToolsPlayableSaveV2DataStatusFromValue returns a pointer to a valid ToolsPlayableSaveV2DataStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPlayableSaveV2DataStatusFromValue(v string) (*ToolsPlayableSaveV2DataStatus, error) {
	ev := ToolsPlayableSaveV2DataStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPlayableSaveV2DataStatus: valid values are %v", v, AllowedToolsPlayableSaveV2DataStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPlayableSaveV2DataStatus) IsValid() bool {
	for _, existing := range AllowedToolsPlayableSaveV2DataStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_playable_save_v2_data_status value
func (v ToolsPlayableSaveV2DataStatus) Ptr() *ToolsPlayableSaveV2DataStatus {
	return &v
}