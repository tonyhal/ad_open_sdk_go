/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponGetV2ResponseData
type ClueCouponGetV2ResponseData struct {
	//
	List     []*ClueCouponGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ClueCouponGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
