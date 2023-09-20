/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteReadV2ResponseData
type ToolsSiteReadV2ResponseData struct {
	// 具体见返回示例业务数据（新建或更新时传递的数据）
	Bricks []map[string]interface{} `json:"bricks,omitempty"`
	// 站点ID
	Id *string `json:"id,omitempty"`
	// 建站类型
	SiteType *string `json:"site_type,omitempty"`
	// 站点状态
	Status *string `json:"status,omitempty"`
	// 缩略图（只有发布的站点才会生成）
	Thumbnail *string `json:"thumbnail,omitempty"`
}
