/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLandingGroupUpdateV2ResponseData
type ToolsLandingGroupUpdateV2ResponseData struct {
	// 流量分配方式。
	GroupFlowType *string `json:"group_flow_type,omitempty"`
	// 落地页组 ID
	GroupId *string `json:"group_id,omitempty"`
	// 落地页组状态。
	GroupStatus *string `json:"group_status,omitempty"`
	// 落地页组名称
	GroupTitle *string `json:"group_title,omitempty"`
	// 落地页组 URL
	GroupUrl *string `json:"group_url,omitempty"`
	// 站点列表
	Sites []*ToolsLandingGroupUpdateV2ResponseDataSitesInner `json:"sites,omitempty"`
}
