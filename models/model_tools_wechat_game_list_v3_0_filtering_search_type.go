/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatGameListV30FilteringSearchType
type ToolsWechatGameListV30FilteringSearchType string

// List of tools_wechat_game_list_v3.0_filtering_search_type
const (
	CREATE_ONLY_ToolsWechatGameListV30FilteringSearchType ToolsWechatGameListV30FilteringSearchType = "CREATE_ONLY"
	SHARE_ONLY_ToolsWechatGameListV30FilteringSearchType  ToolsWechatGameListV30FilteringSearchType = "SHARE_ONLY"
)

// Ptr returns reference to tools_wechat_game_list_v3.0_filtering_search_type value
func (v ToolsWechatGameListV30FilteringSearchType) Ptr() *ToolsWechatGameListV30FilteringSearchType {
	return &v
}
