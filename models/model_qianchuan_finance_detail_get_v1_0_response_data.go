/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanFinanceDetailGetV10ResponseData
type QianchuanFinanceDetailGetV10ResponseData struct {
	// 流水列表
	List     []*QianchuanFinanceDetailGetV10ResponseDataListInner `json:"list"`
	PageInfo *QianchuanFinanceDetailGetV10ResponseDataPageInfo    `json:"page_info,omitempty"`
}
