/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteHandselV2ResponseDataErrorListInner struct for ToolsSiteHandselV2ResponseDataErrorListInner
type ToolsSiteHandselV2ResponseDataErrorListInner struct {
	// 失败原因，当对origin_site_id的操作失败时，返回error_reason字段，成功无
	ErrorReason *string `json:"error_reason,omitempty"`
	// 返回转赠失败的原site_ids
	OriginSiteId *string `json:"origin_site_id,omitempty"`
	// 返回转赠失败的目标广告主id
	TargetAdvertiserId *string `json:"target_advertiser_id,omitempty"`
}
