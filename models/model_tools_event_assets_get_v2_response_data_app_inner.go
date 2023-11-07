/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEventAssetsGetV2ResponseDataAppInner struct for ToolsEventAssetsGetV2ResponseDataAppInner
type ToolsEventAssetsGetV2ResponseDataAppInner struct {
	// 应用ID
	AppId *int64 `json:"app_id,omitempty"`
	// 应用名
	AppName *string                              `json:"app_name,omitempty"`
	AppType *ToolsEventAssetsGetV2DataAppAppType `json:"app_type,omitempty"`
	// 应用资产ID
	AssetId *int64 `json:"asset_id,omitempty"`
	// 资产名
	AssetName *string `json:"asset_name,omitempty"`
	// 应用下载链接
	DownloadUrl *string `json:"download_url,omitempty"`
	// 母包ID
	PackageId *string `json:"package_id,omitempty"`
	// 应用包名
	PackageName *string                           `json:"package_name,omitempty"`
	Role        *ToolsEventAssetsGetV2DataAppRole `json:"role,omitempty"`
}
