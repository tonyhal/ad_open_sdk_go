/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeTemplateListGetV2ResponseDataPageInfo 分页信息
type CreativeTemplateListGetV2ResponseDataPageInfo struct {
	// 是否还有下一页
	HasMore *bool `json:"has_more,omitempty"`
	// 页数
	Page *int32 `json:"page,omitempty"`
	// 页面大小
	PageSize *int32 `json:"page_size,omitempty"`
}
