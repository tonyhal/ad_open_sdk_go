/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarComponentQueryIndustryAnchorV2ResponseData
type StarComponentQueryIndustryAnchorV2ResponseData struct {
	// 组件列表
	Components []*StarComponentQueryIndustryAnchorV2ResponseDataComponentsInner `json:"components,omitempty"`
	PageInfo   *StarComponentQueryIndustryAnchorV2ResponseDataPageInfo          `json:"page_info,omitempty"`
}
