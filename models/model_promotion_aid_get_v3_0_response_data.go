/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionAidGetV30ResponseData
type PromotionAidGetV30ResponseData struct {
	//
	ErrorList []*PromotionAidGetV30ResponseDataErrorListInner `json:"error_list,omitempty"`
	//
	PromotionMapData []*PromotionAidGetV30ResponseDataPromotionMapDataInner `json:"promotion_map_data,omitempty"`
}
