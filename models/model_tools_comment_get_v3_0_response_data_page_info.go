/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentGetV30ResponseDataPageInfo 分页信息
type ToolsCommentGetV30ResponseDataPageInfo struct {
	// 页码
	Page int64 `json:"page"`
	// 页数
	PageSize *int64 `json:"page_size,omitempty"`
	// 总数
	TotalNumber int64 `json:"total_number"`
	// 总页数
	TotalPage *int64 `json:"total_page,omitempty"`
}
