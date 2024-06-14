/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableSaveV2ResponseData
type ToolsPlayableSaveV2ResponseData struct {
	// 试玩素材ID
	PlayableId *int64 `json:"playable_id,omitempty"`
	// 试玩素材名
	PlayableName        *string                                     `json:"playable_name,omitempty"`
	PlayableOrientation *ToolsPlayableSaveV2DataPlayableOrientation `json:"playable_orientation,omitempty"`
	// 试玩素材url
	PlayableUrl *string                        `json:"playable_url,omitempty"`
	Status      *ToolsPlayableSaveV2DataStatus `json:"status,omitempty"`
}
