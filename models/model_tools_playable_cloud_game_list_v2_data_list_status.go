/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPlayableCloudGameListV2DataListStatus
type ToolsPlayableCloudGameListV2DataListStatus string

// List of tools_playable_cloud_game_list_v2_data_list_status
const (
	ON_SHELF_ToolsPlayableCloudGameListV2DataListStatus      ToolsPlayableCloudGameListV2DataListStatus = "ON_SHELF"
	AUDIT_FAIL_ToolsPlayableCloudGameListV2DataListStatus    ToolsPlayableCloudGameListV2DataListStatus = "AUDIT_FAIL"
	AUDIT_SUCCESS_ToolsPlayableCloudGameListV2DataListStatus ToolsPlayableCloudGameListV2DataListStatus = "AUDIT_SUCCESS"
	OFF_SHELF_ToolsPlayableCloudGameListV2DataListStatus     ToolsPlayableCloudGameListV2DataListStatus = "OFF_SHELF"
	UPLOAD_STATUS_ToolsPlayableCloudGameListV2DataListStatus ToolsPlayableCloudGameListV2DataListStatus = "UPLOAD_STATUS"
)

// All allowed values of ToolsPlayableCloudGameListV2DataListStatus enum
var AllowedToolsPlayableCloudGameListV2DataListStatusEnumValues = []ToolsPlayableCloudGameListV2DataListStatus{
	"ON_SHELF",
	"AUDIT_FAIL",
	"AUDIT_SUCCESS",
	"OFF_SHELF",
	"UPLOAD_STATUS",
}

// NewToolsPlayableCloudGameListV2DataListStatusFromValue returns a pointer to a valid ToolsPlayableCloudGameListV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPlayableCloudGameListV2DataListStatusFromValue(v string) (*ToolsPlayableCloudGameListV2DataListStatus, error) {
	ev := ToolsPlayableCloudGameListV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPlayableCloudGameListV2DataListStatus: valid values are %v", v, AllowedToolsPlayableCloudGameListV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPlayableCloudGameListV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedToolsPlayableCloudGameListV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_playable_cloud_game_list_v2_data_list_status value
func (v ToolsPlayableCloudGameListV2DataListStatus) Ptr() *ToolsPlayableCloudGameListV2DataListStatus {
	return &v
}
