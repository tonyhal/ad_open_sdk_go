/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsNoBidSuggestBidV2ResponseData
type ToolsNoBidSuggestBidV2ResponseData struct {
	//
	EstimatedConvertNumMax *int64 `json:"estimated_convert_num_max,omitempty"`
	//
	EstimatedConvertNumMin *int64 `json:"estimated_convert_num_min,omitempty"`
	//
	SuggestBidMax *float64 `json:"suggest_bid_max,omitempty"`
	//
	SuggestBidMin *float64 `json:"suggest_bid_min,omitempty"`
}
