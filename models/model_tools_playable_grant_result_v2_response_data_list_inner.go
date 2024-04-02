/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableGrantResultV2ResponseDataListInner struct for ToolsPlayableGrantResultV2ResponseDataListInner
type ToolsPlayableGrantResultV2ResponseDataListInner struct {
	// 推送任务创建的时间，格式：2020-06-03 16:08:47
	CreateTime *string `json:"create_time,omitempty"`
	// 推送目标（广告主id）
	GrantedId *int64 `json:"granted_id,omitempty"`
	// 推送成功后新生成的试玩素材ID
	NewPlayableId *int64 `json:"new_playable_id,omitempty"`
	// 推送成功后新生成的试玩素材url
	NewPlayableUrl *string `json:"new_playable_url,omitempty"`
	// 被推送的试玩素材id
	PlayableId *int64 `json:"playable_id,omitempty"`
	// 被推送的试玩素材url
	PlayableUrl *string                                   `json:"playable_url,omitempty"`
	Status      *ToolsPlayableGrantResultV2DataListStatus `json:"status,omitempty"`
	// 任务id
	TaskId *int64 `json:"task_id,omitempty"`
}
