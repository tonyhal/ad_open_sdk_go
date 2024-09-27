/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableCloudGameListV2FilteringStatus
type ToolsPlayableCloudGameListV2FilteringStatus string

// List of tools_playable_cloud_game_list_v2_filtering_status
const (
	OFF_SHELF_ToolsPlayableCloudGameListV2FilteringStatus     ToolsPlayableCloudGameListV2FilteringStatus = "OFF_SHELF"
	AUDIT_FAIL_ToolsPlayableCloudGameListV2FilteringStatus    ToolsPlayableCloudGameListV2FilteringStatus = "AUDIT_FAIL"
	UPLOAD_STATUS_ToolsPlayableCloudGameListV2FilteringStatus ToolsPlayableCloudGameListV2FilteringStatus = "UPLOAD_STATUS"
	ON_SHELF_ToolsPlayableCloudGameListV2FilteringStatus      ToolsPlayableCloudGameListV2FilteringStatus = "ON_SHELF"
	AUDIT_SUCCESS_ToolsPlayableCloudGameListV2FilteringStatus ToolsPlayableCloudGameListV2FilteringStatus = "AUDIT_SUCCESS"
)

// Ptr returns reference to tools_playable_cloud_game_list_v2_filtering_status value
func (v ToolsPlayableCloudGameListV2FilteringStatus) Ptr() *ToolsPlayableCloudGameListV2FilteringStatus {
	return &v
}
