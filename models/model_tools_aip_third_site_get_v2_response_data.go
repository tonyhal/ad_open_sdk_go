/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAipThirdSiteGetV2ResponseData
type ToolsAipThirdSiteGetV2ResponseData struct {
	// 广告主id
	AdveriserId *int64 `json:"adveriser_id,omitempty"`
	// 站点审核信息
	AuditMessage *string                                `json:"audit_message,omitempty"`
	AuditStatus  *ToolsAipThirdSiteGetV2DataAuditStatus `json:"audit_status,omitempty"`
	// 站点创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// 站点名称
	Name *string `json:"name,omitempty"`
	// 站点id
	SiteId *int64 `json:"site_id,omitempty"`
	// 站点更新时间
	UpdateTime *string `json:"update_time,omitempty"`
	// 站点URL
	Url *string `json:"url,omitempty"`
	// 站点标准化校验失败信息
	ValidateMessage *string                                   `json:"validate_message,omitempty"`
	ValidateStatus  *ToolsAipThirdSiteGetV2DataValidateStatus `json:"validate_status,omitempty"`
}
