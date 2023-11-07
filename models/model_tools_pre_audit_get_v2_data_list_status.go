/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPreAuditGetV2DataListStatus
type ToolsPreAuditGetV2DataListStatus string

// List of tools_pre_audit_get_v2_data_list_status
const (
	AUDITING_ToolsPreAuditGetV2DataListStatus       ToolsPreAuditGetV2DataListStatus = "AUDITING"
	AUDIT_ACCEPTED_ToolsPreAuditGetV2DataListStatus ToolsPreAuditGetV2DataListStatus = "AUDIT_ACCEPTED"
	AUDIT_FAILED_ToolsPreAuditGetV2DataListStatus   ToolsPreAuditGetV2DataListStatus = "AUDIT_FAILED"
	AUDIT_REJECTED_ToolsPreAuditGetV2DataListStatus ToolsPreAuditGetV2DataListStatus = "AUDIT_REJECTED"
)

// All allowed values of ToolsPreAuditGetV2DataListStatus enum
var AllowedToolsPreAuditGetV2DataListStatusEnumValues = []ToolsPreAuditGetV2DataListStatus{
	"AUDITING",
	"AUDIT_ACCEPTED",
	"AUDIT_FAILED",
	"AUDIT_REJECTED",
}

// NewToolsPreAuditGetV2DataListStatusFromValue returns a pointer to a valid ToolsPreAuditGetV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPreAuditGetV2DataListStatusFromValue(v string) (*ToolsPreAuditGetV2DataListStatus, error) {
	ev := ToolsPreAuditGetV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPreAuditGetV2DataListStatus: valid values are %v", v, AllowedToolsPreAuditGetV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPreAuditGetV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedToolsPreAuditGetV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_pre_audit_get_v2_data_list_status value
func (v ToolsPreAuditGetV2DataListStatus) Ptr() *ToolsPreAuditGetV2DataListStatus {
	return &v
}
