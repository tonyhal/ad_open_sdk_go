/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordPromotionAddV30Request struct for ToolsPrivativeWordPromotionAddV30Request
type ToolsPrivativeWordPromotionAddV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 项目列表
	PromotionList []*ToolsPrivativeWordPromotionAddV30RequestPromotionListInner `json:"promotion_list"`
}
