/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderGetV10ResponseData
type QianchuanAwemeOrderGetV10ResponseData struct {
	// 获取失败的订单id列表
	FailList []int64 `json:"fail_list,omitempty"`
	//
	List     []*QianchuanAwemeOrderGetV10ResponseDataListInner `json:"list,omitempty"`
	PageInfo *QianchuanAwemeOrderGetV10ResponseDataPageInfo    `json:"page_info,omitempty"`
}
