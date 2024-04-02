/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPreAuditGetV2ResponseDataPageInfo
type ToolsPreAuditGetV2ResponseDataPageInfo struct {
	// 页码
	Page *int64 `json:"page,omitempty"`
	// 分页大小
	PageSize *int64 `json:"page_size,omitempty"`
	// 总数
	TotalNumber *int64 `json:"total_number,omitempty"`
	// 总分页数
	TotalPage *int64 `json:"total_page,omitempty"`
}
