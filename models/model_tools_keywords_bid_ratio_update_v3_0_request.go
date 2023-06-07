/*
API Title

巨量引擎开放平台

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsKeywordsBidRatioUpdateV30Request struct for ToolsKeywordsBidRatioUpdateV30Request
type ToolsKeywordsBidRatioUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	BidRatio *float32 `json:"bid_ratio,omitempty"`
	//
	ProjectIds []int64 `json:"project_ids,omitempty"`
	//
	PromotionWordId string `json:"promotion_word_id"`
}
