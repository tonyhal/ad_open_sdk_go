/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableGrantV2Request struct for ToolsPlayableGrantV2Request
type ToolsPlayableGrantV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 推送的目标（广告主id）列表，一次最多同时推送给50个广告主
	GrantedIds []int64 `json:"granted_ids"`
	// 需要推送的试玩素材ID，与playable_url有且仅有一个
	PlayableId *int64 `json:"playable_id,omitempty"`
	// 需要推送的试玩素材url，与playable_id有且仅有一个
	PlayableUrl *string `json:"playable_url,omitempty"`
}
