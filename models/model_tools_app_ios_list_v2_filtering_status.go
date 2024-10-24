/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppIosListV2FilteringStatus
type ToolsAppIosListV2FilteringStatus string

// List of tools_app_ios_list_v2_filtering_status
const (
	AUDIT_REJECTED_ToolsAppIosListV2FilteringStatus ToolsAppIosListV2FilteringStatus = "AUDIT_REJECTED"
	ENABLE_ToolsAppIosListV2FilteringStatus         ToolsAppIosListV2FilteringStatus = "ENABLE"
	ALL_ToolsAppIosListV2FilteringStatus            ToolsAppIosListV2FilteringStatus = "ALL"
	AUDIT_ACCEPTED_ToolsAppIosListV2FilteringStatus ToolsAppIosListV2FilteringStatus = "AUDIT_ACCEPTED"
	AUDIT_DOING_ToolsAppIosListV2FilteringStatus    ToolsAppIosListV2FilteringStatus = "AUDIT_DOING"
)

// Ptr returns reference to tools_app_ios_list_v2_filtering_status value
func (v ToolsAppIosListV2FilteringStatus) Ptr() *ToolsAppIosListV2FilteringStatus {
	return &v
}
