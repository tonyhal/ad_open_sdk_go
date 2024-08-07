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

// ToolsPlayableValidateV2DataStatus
type ToolsPlayableValidateV2DataStatus string

// List of tools_playable_validate_v2_data_status
const (
	AUDIT_FAIL_ToolsPlayableValidateV2DataStatus       ToolsPlayableValidateV2DataStatus = "AUDIT_FAIL"
	AUDIT_SUCCESS_ToolsPlayableValidateV2DataStatus    ToolsPlayableValidateV2DataStatus = "AUDIT_SUCCESS"
	VALIDATE_FAIL_ToolsPlayableValidateV2DataStatus    ToolsPlayableValidateV2DataStatus = "VALIDATE_FAIL"
	VALIDATE_SUCCESS_ToolsPlayableValidateV2DataStatus ToolsPlayableValidateV2DataStatus = "VALIDATE_SUCCESS"
	VALIDATING_ToolsPlayableValidateV2DataStatus       ToolsPlayableValidateV2DataStatus = "VALIDATING"
)

// All allowed values of ToolsPlayableValidateV2DataStatus enum
var AllowedToolsPlayableValidateV2DataStatusEnumValues = []ToolsPlayableValidateV2DataStatus{
	"AUDIT_FAIL",
	"AUDIT_SUCCESS",
	"VALIDATE_FAIL",
	"VALIDATE_SUCCESS",
	"VALIDATING",
}

// NewToolsPlayableValidateV2DataStatusFromValue returns a pointer to a valid ToolsPlayableValidateV2DataStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPlayableValidateV2DataStatusFromValue(v string) (*ToolsPlayableValidateV2DataStatus, error) {
	ev := ToolsPlayableValidateV2DataStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPlayableValidateV2DataStatus: valid values are %v", v, AllowedToolsPlayableValidateV2DataStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPlayableValidateV2DataStatus) IsValid() bool {
	for _, existing := range AllowedToolsPlayableValidateV2DataStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_playable_validate_v2_data_status value
func (v ToolsPlayableValidateV2DataStatus) Ptr() *ToolsPlayableValidateV2DataStatus {
	return &v
}
