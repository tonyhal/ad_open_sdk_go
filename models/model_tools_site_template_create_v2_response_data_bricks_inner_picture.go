/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateCreateV2ResponseDataBricksInnerPicture 图片组件描述
type ToolsSiteTemplateCreateV2ResponseDataBricksInnerPicture struct {
	LinkDto *ToolsSiteTemplateCreateV2ResponseDataBricksInnerPictureLinkDto `json:"link_dto,omitempty"`
	// 标签，用户自定义标注
	Tag *string `json:"tag,omitempty"`
	// 用户自行上传的图片url，当`picture`不为空时，有值
	Url string `json:"url"`
}
