/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionBidUpdateV30RequestDataInner struct for PromotionBidUpdateV30RequestDataInner
type PromotionBidUpdateV30RequestDataInner struct {
	// 出价，单位为元
	Bid float64 `json:"bid"`
	// 广告 ID
	PromotionId int64 `json:"promotion_id"`
}
