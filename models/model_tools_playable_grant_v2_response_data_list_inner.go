/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableGrantV2ResponseDataListInner struct for ToolsPlayableGrantV2ResponseDataListInner
type ToolsPlayableGrantV2ResponseDataListInner struct {
	// 推送目标（广告主ID）
	GrantedId *int64 `json:"granted_id,omitempty"`
	// 推送任务id，可通过【查询推送结果】接口获取任务推送结果
	TaskId *int64 `json:"task_id,omitempty"`
}
