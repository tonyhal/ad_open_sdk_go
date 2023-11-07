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

// ToolsLiveAuthorizeListV2DataListStatus
type ToolsLiveAuthorizeListV2DataListStatus string

// List of tools_live_authorize_list_v2_data_list_status
const (
	APPLYING_ToolsLiveAuthorizeListV2DataListStatus          ToolsLiveAuthorizeListV2DataListStatus = "APPLYING"
	AUTHORIZE_OVERDUE_ToolsLiveAuthorizeListV2DataListStatus ToolsLiveAuthorizeListV2DataListStatus = "AUTHORIZE_OVERDUE"
	APPLY_OVERDUE_ToolsLiveAuthorizeListV2DataListStatus     ToolsLiveAuthorizeListV2DataListStatus = "APPLY_OVERDUE"
	AUTHORIZING_ToolsLiveAuthorizeListV2DataListStatus       ToolsLiveAuthorizeListV2DataListStatus = "AUTHORIZING"
)

// All allowed values of ToolsLiveAuthorizeListV2DataListStatus enum
var AllowedToolsLiveAuthorizeListV2DataListStatusEnumValues = []ToolsLiveAuthorizeListV2DataListStatus{
	"APPLYING",
	"AUTHORIZE_OVERDUE",
	"APPLY_OVERDUE",
	"AUTHORIZING",
}

// NewToolsLiveAuthorizeListV2DataListStatusFromValue returns a pointer to a valid ToolsLiveAuthorizeListV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsLiveAuthorizeListV2DataListStatusFromValue(v string) (*ToolsLiveAuthorizeListV2DataListStatus, error) {
	ev := ToolsLiveAuthorizeListV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsLiveAuthorizeListV2DataListStatus: valid values are %v", v, AllowedToolsLiveAuthorizeListV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsLiveAuthorizeListV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedToolsLiveAuthorizeListV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_live_authorize_list_v2_data_list_status value
func (v ToolsLiveAuthorizeListV2DataListStatus) Ptr() *ToolsLiveAuthorizeListV2DataListStatus {
	return &v
}
