/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeAuthListV2Filtering
type ToolsAwemeAuthListV2Filtering struct {
	// 授权状态
	AuthStatus []*ToolsAwemeAuthListV2FilteringAuthStatus `json:"auth_status,omitempty"`
	// 授权类型
	AuthType []*ToolsAwemeAuthListV2FilteringAuthType `json:"auth_type,omitempty"`
}
