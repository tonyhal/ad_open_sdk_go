/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportLongTransferOrderGetV10ResponseData
type QianchuanReportLongTransferOrderGetV10ResponseData struct {
	//
	List     []*QianchuanReportLongTransferOrderGetV10ResponseDataListInner `json:"list,omitempty"`
	PageInfo *QianchuanAwemeReportOrderGetV10ResponseDataPageInfo           `json:"page_info,omitempty"`
}
