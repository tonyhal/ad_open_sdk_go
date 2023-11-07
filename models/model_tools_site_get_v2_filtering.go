/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteGetV2Filtering
type ToolsSiteGetV2Filtering struct {
	// 搜索字段，按照建站id和建站name进行模糊匹配，范围：1 <= 长度 <= 50
	Search *string `json:"search,omitempty"`
}
