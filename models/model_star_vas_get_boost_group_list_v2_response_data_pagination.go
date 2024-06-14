/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasGetBoostGroupListV2ResponseDataPagination 分页
type StarVasGetBoostGroupListV2ResponseDataPagination struct {
	// 是否还有下一页
	HasMore *bool `json:"has_more,omitempty"`
	// 页内数量
	Limit int64 `json:"limit"`
	// 页码
	Page int64 `json:"page"`
	// 总数
	TotalCount int64 `json:"total_count"`
}
