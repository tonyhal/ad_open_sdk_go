/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportAdvertiserGetV10ResponseData
type QianchuanReportAdvertiserGetV10ResponseData struct {
	//
	List     []*QianchuanReportAdvertiserGetV10ResponseDataListInner `json:"list,omitempty"`
	PageInfo *QianchuanCreativeGetV10ResponseDataPageInfo            `json:"page_info,omitempty"`
}
