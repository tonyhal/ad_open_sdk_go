/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportCreativeGetV10ResponseData
type QianchuanReportCreativeGetV10ResponseData struct {
	//
	List     []*QianchuanReportCreativeGetV10ResponseDataListInner `json:"list,omitempty"`
	PageInfo *QianchuanCreativeGetV10ResponseDataPageInfo          `json:"page_info,omitempty"`
}
