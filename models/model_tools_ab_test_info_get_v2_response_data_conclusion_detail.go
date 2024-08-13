/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestInfoGetV2ResponseDataConclusionDetail 结论详情
type ToolsAbTestInfoGetV2ResponseDataConclusionDetail struct {
	// 普通对象ID列表
	Common []int64 `json:"common,omitempty"`
	// 失败对象ID列表
	Failed []int64 `json:"failed,omitempty"`
	// 获胜对象ID列表
	Winner []int64 `json:"winner,omitempty"`
}
