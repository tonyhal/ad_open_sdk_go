/*
API Title

巨量引擎开放平台

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionStatusUpdateV30RequestDataInner struct for PromotionStatusUpdateV30RequestDataInner
type PromotionStatusUpdateV30RequestDataInner struct {
	OptStatus PromotionStatusUpdateV30DataOptStatus `json:"opt_status"`
	// 广告ID
	PromotionId int64 `json:"promotion_id"`
}
