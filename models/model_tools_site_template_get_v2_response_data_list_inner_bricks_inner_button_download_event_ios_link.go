/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerButtonDownloadEventIosLink ios链接信息
type ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerButtonDownloadEventIosLink struct {
	// 应用描述，为了展示效果，推荐12个中文字符长度
	Description *string                                                                `json:"description,omitempty"`
	LinkType    ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType `json:"link_type"`
	// 下载地址，当`link_type`等于`URL`时，有值
	Url *string `json:"url,omitempty"`
}
