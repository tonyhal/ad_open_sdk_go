/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerAssetsCreateV2RequestAppAsset 应用信息
type EventManagerAssetsCreateV2RequestAppAsset struct {
	AppCreateType *EventManagerAssetsCreateV2AppAssetAppCreateType `json:"app_create_type,omitempty"`
	// 应用ID，Android应用必填
	AppId   *int64                                    `json:"app_id,omitempty"`
	AppType EventManagerAssetsCreateV2AppAssetAppType `json:"app_type"`
	// 应用下载链接，IOS、Android应用必填
	DownloadUrl string `json:"download_url"`
	// 应用名称，长度限制为125，一个字符长度为1，IOS、Android应用必填
	Name string `json:"name"`
	// 母包ID，Android应用必填
	PackageId *string `json:"package_id,omitempty"`
	// 应用包名，长度限制为140，IOS、Android应用必填
	PackageName string `json:"package_name"`
}
