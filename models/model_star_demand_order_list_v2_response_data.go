/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOrderListV2ResponseData
type StarDemandOrderListV2ResponseData struct {
	//
	List     []*StarDemandOrderListV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *StarDemandOrderListV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
