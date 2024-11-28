/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatGameCreateV30AgeLimit
type ToolsWechatGameCreateV30AgeLimit string

// List of tools_wechat_game_create_v3.0_age_limit
const (
	EIGHTEEN_PLUS_ToolsWechatGameCreateV30AgeLimit ToolsWechatGameCreateV30AgeLimit = "EIGHTEEN_PLUS"
	EIGHT_PLUS_ToolsWechatGameCreateV30AgeLimit    ToolsWechatGameCreateV30AgeLimit = "EIGHT_PLUS"
	FOUR_PLUS_ToolsWechatGameCreateV30AgeLimit     ToolsWechatGameCreateV30AgeLimit = "FOUR_PLUS"
	SIXTEEN_PLUS_ToolsWechatGameCreateV30AgeLimit  ToolsWechatGameCreateV30AgeLimit = "SIXTEEN_PLUS"
	TWELVE_PLUS_ToolsWechatGameCreateV30AgeLimit   ToolsWechatGameCreateV30AgeLimit = "TWELVE_PLUS"
)

// Ptr returns reference to tools_wechat_game_create_v3.0_age_limit value
func (v ToolsWechatGameCreateV30AgeLimit) Ptr() *ToolsWechatGameCreateV30AgeLimit {
	return &v
}