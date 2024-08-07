/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeAuthListV2Filtering
type ToolsAwemeAuthListV2Filtering struct {
	// 授权状态
	AuthStatus []*ToolsAwemeAuthListV2FilteringAuthStatus `json:"auth_status,omitempty"`
	// 授权类型
	AuthType []*ToolsAwemeAuthListV2FilteringAuthType `json:"auth_type,omitempty"`
	// 抖音号
	AwemeIds []string `json:"aweme_ids,omitempty"`
	// 抖音短视频ID
	ItemIds []int64 `json:"item_ids,omitempty"`
}
