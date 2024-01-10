/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsKeywordsBidRatioCreateV30Request struct for ToolsKeywordsBidRatioCreateV30Request
type ToolsKeywordsBidRatioCreateV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	//
	BidRatio  float64                                 `json:"bid_ratio"`
	Dimension ToolsKeywordsBidRatioCreateV30Dimension `json:"dimension"`
	//
	ProjectIds []int64 `json:"project_ids,omitempty"`
	// 优词列表
	WordList []string `json:"word_list"`
}
