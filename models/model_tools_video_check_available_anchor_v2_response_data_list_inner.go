/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsVideoCheckAvailableAnchorV2ResponseDataListInner struct for ToolsVideoCheckAvailableAnchorV2ResponseDataListInner
type ToolsVideoCheckAvailableAnchorV2ResponseDataListInner struct {
	// 抖音视频ID
	ItemId int64 `json:"item_id"`
	// 是否支持双库存
	Valid bool `json:"valid"`
}
