/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanVideoGetV10ResponseData
type QianchuanVideoGetV10ResponseData struct {
	//
	List     []*QianchuanVideoGetV10ResponseDataListInner `json:"list"`
	PageInfo *QianchuanImageGetV10ResponseDataPageInfo    `json:"page_info,omitempty"`
}
