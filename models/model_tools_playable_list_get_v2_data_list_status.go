/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPlayableListGetV2DataListStatus
type ToolsPlayableListGetV2DataListStatus string

// List of tools_playable_list_get_v2_data_list_status
const (
	VALIDATING_ToolsPlayableListGetV2DataListStatus       ToolsPlayableListGetV2DataListStatus = "VALIDATING"
	AUDIT_FAIL_ToolsPlayableListGetV2DataListStatus       ToolsPlayableListGetV2DataListStatus = "AUDIT_FAIL"
	AUDIT_SUCCESS_ToolsPlayableListGetV2DataListStatus    ToolsPlayableListGetV2DataListStatus = "AUDIT_SUCCESS"
	VALIDATE_SUCCESS_ToolsPlayableListGetV2DataListStatus ToolsPlayableListGetV2DataListStatus = "VALIDATE_SUCCESS"
	VALIDATE_FAIL_ToolsPlayableListGetV2DataListStatus    ToolsPlayableListGetV2DataListStatus = "VALIDATE_FAIL"
)

// All allowed values of ToolsPlayableListGetV2DataListStatus enum
var AllowedToolsPlayableListGetV2DataListStatusEnumValues = []ToolsPlayableListGetV2DataListStatus{
	"VALIDATING",
	"AUDIT_FAIL",
	"AUDIT_SUCCESS",
	"VALIDATE_SUCCESS",
	"VALIDATE_FAIL",
}

// NewToolsPlayableListGetV2DataListStatusFromValue returns a pointer to a valid ToolsPlayableListGetV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPlayableListGetV2DataListStatusFromValue(v string) (*ToolsPlayableListGetV2DataListStatus, error) {
	ev := ToolsPlayableListGetV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPlayableListGetV2DataListStatus: valid values are %v", v, AllowedToolsPlayableListGetV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPlayableListGetV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedToolsPlayableListGetV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_playable_list_get_v2_data_list_status value
func (v ToolsPlayableListGetV2DataListStatus) Ptr() *ToolsPlayableListGetV2DataListStatus {
	return &v
}
