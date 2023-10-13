/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsWechatGameListV30FilteringAuditStatus
type ToolsWechatGameListV30FilteringAuditStatus string

// List of tools_wechat_game_list_v3.0_filtering_audit_status
const (
	AUDIT_ACCEPTED_ToolsWechatGameListV30FilteringAuditStatus ToolsWechatGameListV30FilteringAuditStatus = "AUDIT_ACCEPTED"
	AUDITING_ToolsWechatGameListV30FilteringAuditStatus       ToolsWechatGameListV30FilteringAuditStatus = "AUDITING"
	AUDIT_REJECTED_ToolsWechatGameListV30FilteringAuditStatus ToolsWechatGameListV30FilteringAuditStatus = "AUDIT_REJECTED"
	ALL_ToolsWechatGameListV30FilteringAuditStatus            ToolsWechatGameListV30FilteringAuditStatus = "ALL"
)

// All allowed values of ToolsWechatGameListV30FilteringAuditStatus enum
var AllowedToolsWechatGameListV30FilteringAuditStatusEnumValues = []ToolsWechatGameListV30FilteringAuditStatus{
	"AUDIT_ACCEPTED",
	"AUDITING",
	"AUDIT_REJECTED",
	"ALL",
}

// NewToolsWechatGameListV30FilteringAuditStatusFromValue returns a pointer to a valid ToolsWechatGameListV30FilteringAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsWechatGameListV30FilteringAuditStatusFromValue(v string) (*ToolsWechatGameListV30FilteringAuditStatus, error) {
	ev := ToolsWechatGameListV30FilteringAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsWechatGameListV30FilteringAuditStatus: valid values are %v", v, AllowedToolsWechatGameListV30FilteringAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsWechatGameListV30FilteringAuditStatus) IsValid() bool {
	for _, existing := range AllowedToolsWechatGameListV30FilteringAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_wechat_game_list_v3.0_filtering_audit_status value
func (v ToolsWechatGameListV30FilteringAuditStatus) Ptr() *ToolsWechatGameListV30FilteringAuditStatus {
	return &v
}
