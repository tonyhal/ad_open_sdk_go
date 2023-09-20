/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsThirdSiteGetV2DataListAuditStatus
type ToolsThirdSiteGetV2DataListAuditStatus string

// List of tools_third_site_get_v2_data_list_audit_status
const (
	AUDIT_UNKNOW_ToolsThirdSiteGetV2DataListAuditStatus   ToolsThirdSiteGetV2DataListAuditStatus = "AUDIT_UNKNOW"
	AUDIT_ACCEPTED_ToolsThirdSiteGetV2DataListAuditStatus ToolsThirdSiteGetV2DataListAuditStatus = "AUDIT_ACCEPTED"
	AUDIT_REJECTED_ToolsThirdSiteGetV2DataListAuditStatus ToolsThirdSiteGetV2DataListAuditStatus = "AUDIT_REJECTED"
	AUDIT_BANNED_ToolsThirdSiteGetV2DataListAuditStatus   ToolsThirdSiteGetV2DataListAuditStatus = "AUDIT_BANNED"
	AUDITING_ToolsThirdSiteGetV2DataListAuditStatus       ToolsThirdSiteGetV2DataListAuditStatus = "AUDITING"
	AWAIT_AUDIT_ToolsThirdSiteGetV2DataListAuditStatus    ToolsThirdSiteGetV2DataListAuditStatus = "AWAIT_AUDIT"
)

// All allowed values of ToolsThirdSiteGetV2DataListAuditStatus enum
var AllowedToolsThirdSiteGetV2DataListAuditStatusEnumValues = []ToolsThirdSiteGetV2DataListAuditStatus{
	"AUDIT_UNKNOW",
	"AUDIT_ACCEPTED",
	"AUDIT_REJECTED",
	"AUDIT_BANNED",
	"AUDITING",
	"AWAIT_AUDIT",
}

// NewToolsThirdSiteGetV2DataListAuditStatusFromValue returns a pointer to a valid ToolsThirdSiteGetV2DataListAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsThirdSiteGetV2DataListAuditStatusFromValue(v string) (*ToolsThirdSiteGetV2DataListAuditStatus, error) {
	ev := ToolsThirdSiteGetV2DataListAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsThirdSiteGetV2DataListAuditStatus: valid values are %v", v, AllowedToolsThirdSiteGetV2DataListAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsThirdSiteGetV2DataListAuditStatus) IsValid() bool {
	for _, existing := range AllowedToolsThirdSiteGetV2DataListAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_third_site_get_v2_data_list_audit_status value
func (v ToolsThirdSiteGetV2DataListAuditStatus) Ptr() *ToolsThirdSiteGetV2DataListAuditStatus {
	return &v
}
