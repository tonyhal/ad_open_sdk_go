/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdGetV10ResponseData
type QianchuanAdGetV10ResponseData struct {
	//
	FailList []int64 `json:"fail_list,omitempty"`
	//
	List     []*QianchuanAdGetV10ResponseDataListInner `json:"list,omitempty"`
	PageInfo *QianchuanAdGetV10ResponseDataPageInfo    `json:"page_info,omitempty"`
}
