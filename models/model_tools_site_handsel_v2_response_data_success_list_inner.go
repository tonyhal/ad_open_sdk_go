/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteHandselV2ResponseDataSuccessListInner struct for ToolsSiteHandselV2ResponseDataSuccessListInner
type ToolsSiteHandselV2ResponseDataSuccessListInner struct {
	// 返回转赠成功的原site_ids
	OriginSiteId *string `json:"origin_site_id,omitempty"`
	// 转赠成功后的生成的新站点id，失败的数据无此参数返回
	SiteId *string `json:"site_id,omitempty"`
	// 返回转赠成功的目标广告主id
	TargetAdvertiserId *string `json:"target_advertiser_id,omitempty"`
}