/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeTemplateListGetV2ResponseData
type CreativeTemplateListGetV2ResponseData struct {
	// 模版列表
	Lists    []*CreativeTemplateListGetV2ResponseDataListsInner `json:"lists,omitempty"`
	PageInfo *CreativeTemplateListGetV2ResponseDataPageInfo     `json:"page_info,omitempty"`
}
