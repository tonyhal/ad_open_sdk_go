/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdKeywordsGetV10ResponseData
type QianchuanAdKeywordsGetV10ResponseData struct {
	// 关键词列表
	List     []*QianchuanAdKeywordsGetV10ResponseDataListInner `json:"list,omitempty"`
	PageInfo *QianchuanAdKeywordsGetV10ResponseDataPageInfo    `json:"page_info,omitempty"`
}
