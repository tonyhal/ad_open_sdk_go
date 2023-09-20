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

// ToolsPlayableCloudGameListV2FilteringAdStatus
type ToolsPlayableCloudGameListV2FilteringAdStatus string

// List of tools_playable_cloud_game_list_v2_filtering_ad_status
const (
	DELETE_ToolsPlayableCloudGameListV2FilteringAdStatus ToolsPlayableCloudGameListV2FilteringAdStatus = "DELETE"
	UNUSED_ToolsPlayableCloudGameListV2FilteringAdStatus ToolsPlayableCloudGameListV2FilteringAdStatus = "UNUSED"
	INUSE_ToolsPlayableCloudGameListV2FilteringAdStatus  ToolsPlayableCloudGameListV2FilteringAdStatus = "INUSE"
)

// All allowed values of ToolsPlayableCloudGameListV2FilteringAdStatus enum
var AllowedToolsPlayableCloudGameListV2FilteringAdStatusEnumValues = []ToolsPlayableCloudGameListV2FilteringAdStatus{
	"DELETE",
	"UNUSED",
	"INUSE",
}

// NewToolsPlayableCloudGameListV2FilteringAdStatusFromValue returns a pointer to a valid ToolsPlayableCloudGameListV2FilteringAdStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPlayableCloudGameListV2FilteringAdStatusFromValue(v string) (*ToolsPlayableCloudGameListV2FilteringAdStatus, error) {
	ev := ToolsPlayableCloudGameListV2FilteringAdStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPlayableCloudGameListV2FilteringAdStatus: valid values are %v", v, AllowedToolsPlayableCloudGameListV2FilteringAdStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPlayableCloudGameListV2FilteringAdStatus) IsValid() bool {
	for _, existing := range AllowedToolsPlayableCloudGameListV2FilteringAdStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_playable_cloud_game_list_v2_filtering_ad_status value
func (v ToolsPlayableCloudGameListV2FilteringAdStatus) Ptr() *ToolsPlayableCloudGameListV2FilteringAdStatus {
	return &v
}
