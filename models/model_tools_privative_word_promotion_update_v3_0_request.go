/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordPromotionUpdateV30Request struct for ToolsPrivativeWordPromotionUpdateV30Request
type ToolsPrivativeWordPromotionUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	PromotionList []*ToolsPrivativeWordPromotionUpdateV30RequestPromotionListInner `json:"promotion_list"`
}
