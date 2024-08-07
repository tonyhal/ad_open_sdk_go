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

// ToolsAipThirdSiteGetV2DataAuditStatus
type ToolsAipThirdSiteGetV2DataAuditStatus string

// List of tools_aip_third_site_get_v2_data_audit_status
const (
	AUDIT_UNKNOW_ToolsAipThirdSiteGetV2DataAuditStatus   ToolsAipThirdSiteGetV2DataAuditStatus = "AUDIT_UNKNOW"
	AUDIT_ACCEPTED_ToolsAipThirdSiteGetV2DataAuditStatus ToolsAipThirdSiteGetV2DataAuditStatus = "AUDIT_ACCEPTED"
	AUDIT_REJECTED_ToolsAipThirdSiteGetV2DataAuditStatus ToolsAipThirdSiteGetV2DataAuditStatus = "AUDIT_REJECTED"
	AUDIT_BANNED_ToolsAipThirdSiteGetV2DataAuditStatus   ToolsAipThirdSiteGetV2DataAuditStatus = "AUDIT_BANNED"
	AUDITING_ToolsAipThirdSiteGetV2DataAuditStatus       ToolsAipThirdSiteGetV2DataAuditStatus = "AUDITING"
	AWAIT_AUDIT_ToolsAipThirdSiteGetV2DataAuditStatus    ToolsAipThirdSiteGetV2DataAuditStatus = "AWAIT_AUDIT"
)

// All allowed values of ToolsAipThirdSiteGetV2DataAuditStatus enum
var AllowedToolsAipThirdSiteGetV2DataAuditStatusEnumValues = []ToolsAipThirdSiteGetV2DataAuditStatus{
	"AUDIT_UNKNOW",
	"AUDIT_ACCEPTED",
	"AUDIT_REJECTED",
	"AUDIT_BANNED",
	"AUDITING",
	"AWAIT_AUDIT",
}

// NewToolsAipThirdSiteGetV2DataAuditStatusFromValue returns a pointer to a valid ToolsAipThirdSiteGetV2DataAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAipThirdSiteGetV2DataAuditStatusFromValue(v string) (*ToolsAipThirdSiteGetV2DataAuditStatus, error) {
	ev := ToolsAipThirdSiteGetV2DataAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAipThirdSiteGetV2DataAuditStatus: valid values are %v", v, AllowedToolsAipThirdSiteGetV2DataAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAipThirdSiteGetV2DataAuditStatus) IsValid() bool {
	for _, existing := range AllowedToolsAipThirdSiteGetV2DataAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aip_third_site_get_v2_data_audit_status value
func (v ToolsAipThirdSiteGetV2DataAuditStatus) Ptr() *ToolsAipThirdSiteGetV2DataAuditStatus {
	return &v
}
