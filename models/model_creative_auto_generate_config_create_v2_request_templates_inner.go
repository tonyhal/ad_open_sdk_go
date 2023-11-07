/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeAutoGenerateConfigCreateV2RequestTemplatesInner struct for CreativeAutoGenerateConfigCreateV2RequestTemplatesInner
type CreativeAutoGenerateConfigCreateV2RequestTemplatesInner struct {
	// 模板ID
	TemplateId int64 `json:"template_id"`
	// 模板填充的图片内容，不填写则表示使用默认填充值
	TemplateImgSchema []*CreativeAutoGenerateConfigCreateV2RequestTemplatesInnerTemplateImgSchemaInner `json:"template_img_schema,omitempty"`
	// 模板填充的文本内容，不填写则表示使用默认填充值
	TemplateTextSchema []*CreativeAutoGenerateConfigCreateV2RequestTemplatesInnerTemplateTextSchemaInner `json:"template_text_schema,omitempty"`
	TemplateType       CreativeAutoGenerateConfigCreateV2TemplatesTemplateType                           `json:"template_type"`
}
