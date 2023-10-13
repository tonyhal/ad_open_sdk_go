/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionBidUpdateV30ResponseData
type PromotionBidUpdateV30ResponseData struct {
	// 删除失败的广告ID列表
	Errors []*PromotionBidUpdateV30ResponseDataErrorsInner `json:"errors,omitempty"`
	// 更新成功的广告ID列表
	PromotionIds []int64 `json:"promotion_ids,omitempty"`
}
