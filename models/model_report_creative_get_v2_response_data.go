/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCreativeGetV2ResponseData
type ReportCreativeGetV2ResponseData struct {
	//
	List     []*ReportCreativeGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ReportCreativeGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
