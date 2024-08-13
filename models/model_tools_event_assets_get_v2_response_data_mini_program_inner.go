/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEventAssetsGetV2ResponseDataMiniProgramInner struct for ToolsEventAssetsGetV2ResponseDataMiniProgramInner
type ToolsEventAssetsGetV2ResponseDataMiniProgramInner struct {
	// 字节小程序资产id
	AssetId *int64 `json:"asset_id,omitempty"`
	// 字节小程序资产名称
	AssetName *string `json:"asset_name,omitempty"`
	// 小程序instance_id
	InstanceId *int64 `json:"instance_id,omitempty"`
	// 字节小程序id
	MiniProgramId *string `json:"mini_program_id,omitempty"`
}
