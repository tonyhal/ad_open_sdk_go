/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionListV10ResponseDataPageInfo
type QianchuanUniPromotionListV10ResponseDataPageInfo struct {
	//
	Page int64 `json:"page"`
	//
	PageSize int64 `json:"page_size"`
	//
	TotalNum *int64 `json:"total_num,omitempty"`
	//
	TotalPage *int64 `json:"total_page,omitempty"`
}
