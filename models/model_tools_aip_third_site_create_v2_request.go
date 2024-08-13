/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAipThirdSiteCreateV2Request struct for ToolsAipThirdSiteCreateV2Request
type ToolsAipThirdSiteCreateV2Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 站点名称，长度限制，1-50 字
	Name *string `json:"name,omitempty"`
	// 站点URL
	Url string `json:"url"`
}
