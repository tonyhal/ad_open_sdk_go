/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsKeywordsBidRatioDeleteV30Request struct for ToolsKeywordsBidRatioDeleteV30Request
type ToolsKeywordsBidRatioDeleteV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	// 要删除的优词计划id列表
	PromotionWordIds []string `json:"promotion_word_ids"`
}