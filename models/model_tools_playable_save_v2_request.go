/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableSaveV2Request struct for ToolsPlayableSaveV2Request
type ToolsPlayableSaveV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 试玩素材ID
	PlayableId int64 `json:"playable_id"`
	// 试玩素材名称
	PlayableName string `json:"playable_name"`
}
