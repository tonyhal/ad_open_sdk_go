/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormListV2ResponseData
type ClueFormListV2ResponseData struct {
	//
	List     []*ClueFormListV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ClueFormListV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
