/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteCopyV2ResponseDataSuccessListInner struct for ToolsSiteCopyV2ResponseDataSuccessListInner
type ToolsSiteCopyV2ResponseDataSuccessListInner struct {
	// 返回复制成功的原站点id
	OriginSiteId *string `json:"origin_site_id,omitempty"`
	// 返回复制成功后生成的新站点id
	SiteId *string `json:"site_id,omitempty"`
}
