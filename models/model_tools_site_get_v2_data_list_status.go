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

// ToolsSiteGetV2DataListStatus
type ToolsSiteGetV2DataListStatus string

// List of tools_site_get_v2_data_list_status
const (
	AUDIT_ACCEPTED_ToolsSiteGetV2DataListStatus ToolsSiteGetV2DataListStatus = "AUDIT_ACCEPTED"
	AUDIT_BANNED_ToolsSiteGetV2DataListStatus   ToolsSiteGetV2DataListStatus = "AUDIT_BANNED"
	AUDIT_DOING_ToolsSiteGetV2DataListStatus    ToolsSiteGetV2DataListStatus = "AUDIT_DOING"
	AUDIT_REJECTED_ToolsSiteGetV2DataListStatus ToolsSiteGetV2DataListStatus = "AUDIT_REJECTED"
	DELETED_ToolsSiteGetV2DataListStatus        ToolsSiteGetV2DataListStatus = "DELETED"
	DISABLE_ToolsSiteGetV2DataListStatus        ToolsSiteGetV2DataListStatus = "DISABLE"
	EDIT_ToolsSiteGetV2DataListStatus           ToolsSiteGetV2DataListStatus = "EDIT"
	ENABLE_ToolsSiteGetV2DataListStatus         ToolsSiteGetV2DataListStatus = "ENABLE"
	FORBIDDEN_ToolsSiteGetV2DataListStatus      ToolsSiteGetV2DataListStatus = "FORBIDDEN"
)

// All allowed values of ToolsSiteGetV2DataListStatus enum
var AllowedToolsSiteGetV2DataListStatusEnumValues = []ToolsSiteGetV2DataListStatus{
	"AUDIT_ACCEPTED",
	"AUDIT_BANNED",
	"AUDIT_DOING",
	"AUDIT_REJECTED",
	"DELETED",
	"DISABLE",
	"EDIT",
	"ENABLE",
	"FORBIDDEN",
}

// NewToolsSiteGetV2DataListStatusFromValue returns a pointer to a valid ToolsSiteGetV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteGetV2DataListStatusFromValue(v string) (*ToolsSiteGetV2DataListStatus, error) {
	ev := ToolsSiteGetV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteGetV2DataListStatus: valid values are %v", v, AllowedToolsSiteGetV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteGetV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedToolsSiteGetV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_get_v2_data_list_status value
func (v ToolsSiteGetV2DataListStatus) Ptr() *ToolsSiteGetV2DataListStatus {
	return &v
}
