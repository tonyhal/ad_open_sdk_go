/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportMaterialGetV10ResponseData
type QianchuanReportMaterialGetV10ResponseData struct {
	//
	List     []*QianchuanReportMaterialGetV10ResponseDataListInner `json:"list,omitempty"`
	PageInfo QianchuanReportMaterialGetV10ResponseDataPageInfo     `json:"page_info"`
}
