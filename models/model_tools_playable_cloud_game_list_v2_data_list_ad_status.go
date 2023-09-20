/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPlayableCloudGameListV2DataListAdStatus
type ToolsPlayableCloudGameListV2DataListAdStatus string

// List of tools_playable_cloud_game_list_v2_data_list_ad_status
const (
	DELETE_ToolsPlayableCloudGameListV2DataListAdStatus ToolsPlayableCloudGameListV2DataListAdStatus = "DELETE"
	UNUSED_ToolsPlayableCloudGameListV2DataListAdStatus ToolsPlayableCloudGameListV2DataListAdStatus = "UNUSED"
	INUSE_ToolsPlayableCloudGameListV2DataListAdStatus  ToolsPlayableCloudGameListV2DataListAdStatus = "INUSE"
)

// All allowed values of ToolsPlayableCloudGameListV2DataListAdStatus enum
var AllowedToolsPlayableCloudGameListV2DataListAdStatusEnumValues = []ToolsPlayableCloudGameListV2DataListAdStatus{
	"DELETE",
	"UNUSED",
	"INUSE",
}

// NewToolsPlayableCloudGameListV2DataListAdStatusFromValue returns a pointer to a valid ToolsPlayableCloudGameListV2DataListAdStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPlayableCloudGameListV2DataListAdStatusFromValue(v string) (*ToolsPlayableCloudGameListV2DataListAdStatus, error) {
	ev := ToolsPlayableCloudGameListV2DataListAdStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPlayableCloudGameListV2DataListAdStatus: valid values are %v", v, AllowedToolsPlayableCloudGameListV2DataListAdStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPlayableCloudGameListV2DataListAdStatus) IsValid() bool {
	for _, existing := range AllowedToolsPlayableCloudGameListV2DataListAdStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_playable_cloud_game_list_v2_data_list_ad_status value
func (v ToolsPlayableCloudGameListV2DataListAdStatus) Ptr() *ToolsPlayableCloudGameListV2DataListAdStatus {
	return &v
}
