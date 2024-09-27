/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableCloudGameListV2DataListStatus
type ToolsPlayableCloudGameListV2DataListStatus string

// List of tools_playable_cloud_game_list_v2_data_list_status
const (
	OFF_SHELF_ToolsPlayableCloudGameListV2DataListStatus     ToolsPlayableCloudGameListV2DataListStatus = "OFF_SHELF"
	AUDIT_FAIL_ToolsPlayableCloudGameListV2DataListStatus    ToolsPlayableCloudGameListV2DataListStatus = "AUDIT_FAIL"
	UPLOAD_STATUS_ToolsPlayableCloudGameListV2DataListStatus ToolsPlayableCloudGameListV2DataListStatus = "UPLOAD_STATUS"
	ON_SHELF_ToolsPlayableCloudGameListV2DataListStatus      ToolsPlayableCloudGameListV2DataListStatus = "ON_SHELF"
	AUDIT_SUCCESS_ToolsPlayableCloudGameListV2DataListStatus ToolsPlayableCloudGameListV2DataListStatus = "AUDIT_SUCCESS"
)

// Ptr returns reference to tools_playable_cloud_game_list_v2_data_list_status value
func (v ToolsPlayableCloudGameListV2DataListStatus) Ptr() *ToolsPlayableCloudGameListV2DataListStatus {
	return &v
}
