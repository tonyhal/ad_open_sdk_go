/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandMaterialListV30ResponseDataMaterialsInnerMaterialComponentCommonExternalMicroAppInfo 小程序链接
type BrandMaterialListV30ResponseDataMaterialsInnerMaterialComponentCommonExternalMicroAppInfo struct {
	// appid
	AppId *string `json:"app_id,omitempty"`
	// 启动参数
	StartParam *string `json:"start_param,omitempty"`
	// 启动路径
	StartPath *string `json:"start_path,omitempty"`
	// 小程序类型
	Type *int64 `json:"type,omitempty"`
	// 完整链接
	Url *string `json:"url,omitempty"`
}