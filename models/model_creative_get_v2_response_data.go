/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeGetV2ResponseData
type CreativeGetV2ResponseData struct {
	CursorInfo *CreativeGetV2ResponseDataCursorInfo `json:"cursor_info,omitempty"`
	//
	List     []*CreativeGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *CreativeGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
