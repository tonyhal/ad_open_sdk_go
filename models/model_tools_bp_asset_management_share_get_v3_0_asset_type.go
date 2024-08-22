/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBpAssetManagementShareGetV30AssetType
type ToolsBpAssetManagementShareGetV30AssetType string

// List of tools_bp_asset_management_share_get_v3.0_asset_type
const (
	APPLETS_ToolsBpAssetManagementShareGetV30AssetType       ToolsBpAssetManagementShareGetV30AssetType = "APPLETS"
	BYTED_APPLETS_ToolsBpAssetManagementShareGetV30AssetType ToolsBpAssetManagementShareGetV30AssetType = "BYTED_APPLETS"
	BYTED_GAME_ToolsBpAssetManagementShareGetV30AssetType    ToolsBpAssetManagementShareGetV30AssetType = "BYTED_GAME"
	WECHAT_GAME_ToolsBpAssetManagementShareGetV30AssetType   ToolsBpAssetManagementShareGetV30AssetType = "WECHAT_GAME"
)

// Ptr returns reference to tools_bp_asset_management_share_get_v3.0_asset_type value
func (v ToolsBpAssetManagementShareGetV30AssetType) Ptr() *ToolsBpAssetManagementShareGetV30AssetType {
	return &v
}
