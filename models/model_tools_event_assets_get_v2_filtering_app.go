/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEventAssetsGetV2FilteringApp 应用数据
type ToolsEventAssetsGetV2FilteringApp struct {
	// 应用包名，精确搜索
	AppPackageName *string `json:"app_package_name,omitempty"`
	// 资产名，模糊搜索
	AssetName *string `json:"asset_name,omitempty"`
}
