/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdRecommendKeywordsGetV10ResponseDataPageInfo 分页信息
type QianchuanAdRecommendKeywordsGetV10ResponseDataPageInfo struct {
	// 页码
	Page *int32 `json:"page,omitempty"`
	// 页面大小
	PageSize *int32 `json:"page_size,omitempty"`
	// 关键词总数
	TotalNumber *int32 `json:"total_number,omitempty"`
	// 页面总数
	TotalPage *int32 `json:"total_page,omitempty"`
}
