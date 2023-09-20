/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLandingGroupUpdateV2Request struct for ToolsLandingGroupUpdateV2Request
type ToolsLandingGroupUpdateV2Request struct {
	// 广告主id，范围：1 <= advertiser_id <= MAX_INT64
	AdvertiserId int64 `json:"advertiser_id"`
	// 新加站点列表
	AppendSites []string `json:"append_sites"`
	// 落地页组 ID
	GroupId string `json:"group_id"`
	// 落地页组名称
	GroupTitle string `json:"group_title"`
}
