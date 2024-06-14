/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueInfoUpdateV2RequestUpdateInfoInner struct for ToolsClueInfoUpdateV2RequestUpdateInfoInner
type ToolsClueInfoUpdateV2RequestUpdateInfoInner struct {
	// 线索id
	ClueId int64 `json:"clue_id"`
	// 具体更新的 key, value。更新name时，不能修改为空
	Fields *map[string]string `json:"fields,omitempty"`
}
