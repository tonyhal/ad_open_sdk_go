/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueLifeGetV2DataListFollowStateName
type ToolsClueLifeGetV2DataListFollowStateName string

// List of tools_clue_life_get_v2_data_list_follow_state_name
const (
	NOT_CALLED_ToolsClueLifeGetV2DataListFollowStateName     ToolsClueLifeGetV2DataListFollowStateName = "NOT_CALLED"
	NOT_ANSWERED_ToolsClueLifeGetV2DataListFollowStateName   ToolsClueLifeGetV2DataListFollowStateName = "NOT_ANSWERED"
	SHORT_ANSWERED_ToolsClueLifeGetV2DataListFollowStateName ToolsClueLifeGetV2DataListFollowStateName = "SHORT_ANSWERED"
	ANSWERED_ToolsClueLifeGetV2DataListFollowStateName       ToolsClueLifeGetV2DataListFollowStateName = "ANSWERED"
	DEEP_ANSWERED_ToolsClueLifeGetV2DataListFollowStateName  ToolsClueLifeGetV2DataListFollowStateName = "DEEP_ANSWERED"
)

// Ptr returns reference to tools_clue_life_get_v2_data_list_follow_state_name value
func (v ToolsClueLifeGetV2DataListFollowStateName) Ptr() *ToolsClueLifeGetV2DataListFollowStateName {
	return &v
}
