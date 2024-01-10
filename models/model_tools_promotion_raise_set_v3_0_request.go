/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionRaiseSetV30Request struct for ToolsPromotionRaiseSetV30Request
type ToolsPromotionRaiseSetV30Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 广告ID
	PromotionId int64 `json:"promotion_id"`
	// 起量信息，如果需要删除请传[]
	RaiseInfo []*ToolsPromotionRaiseSetV30RequestRaiseInfoInner `json:"raise_info"`
}
