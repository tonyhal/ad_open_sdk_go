/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaProductDetailGetV2ResponseData
type DpaProductDetailGetV2ResponseData struct {
	//
	List     []*DpaProductDetailGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *DpaProductDetailGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
