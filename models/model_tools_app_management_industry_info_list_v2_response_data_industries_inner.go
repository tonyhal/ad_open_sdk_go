/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementIndustryInfoListV2ResponseDataIndustriesInner struct for ToolsAppManagementIndustryInfoListV2ResponseDataIndustriesInner
type ToolsAppManagementIndustryInfoListV2ResponseDataIndustriesInner struct {
	// 子分类
	Children []*ToolsAppManagementIndustryInfoListV2ResponseDataIndustriesInnerChildrenInner `json:"children,omitempty"`
	// 分类id
	Id *string `json:"id,omitempty"`
	// 级别
	Level *int64 `json:"level,omitempty"`
	// 分类名称
	Name *string `json:"name,omitempty"`
}
