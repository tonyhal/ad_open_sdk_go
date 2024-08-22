/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableValidateV2ResponseData
type ToolsPlayableValidateV2ResponseData struct {
	// 验证失败原因
	FailMessage *string `json:"fail_message,omitempty"`
	// 试玩素材ID
	PlayableId          *int64                                          `json:"playable_id,omitempty"`
	PlayableOrientation *ToolsPlayableValidateV2DataPlayableOrientation `json:"playable_orientation,omitempty"`
	// 试玩素材url
	PlayableUrl *string                            `json:"playable_url,omitempty"`
	Status      *ToolsPlayableValidateV2DataStatus `json:"status,omitempty"`
}
