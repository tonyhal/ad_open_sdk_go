/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateGetV2ResponseData
type ToolsSiteTemplateGetV2ResponseData struct {
	// 模板列表
	List     []*ToolsSiteTemplateGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ToolsSiteTemplateGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
