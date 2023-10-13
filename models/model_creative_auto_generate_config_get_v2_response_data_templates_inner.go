/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeAutoGenerateConfigGetV2ResponseDataTemplatesInner struct for CreativeAutoGenerateConfigGetV2ResponseDataTemplatesInner
type CreativeAutoGenerateConfigGetV2ResponseDataTemplatesInner struct {
	// 模板ID
	TemplateId int64 `json:"template_id"`
	// 模板填充的图片内容
	TemplateImgSchema []*CreativeAutoGenerateConfigGetV2ResponseDataTemplatesInnerTemplateImgSchemaInner `json:"template_img_schema,omitempty"`
	// 模板填充的文本内容
	TemplateTextSchema []*CreativeAutoGenerateConfigGetV2ResponseDataTemplatesInnerTemplateTextSchemaInner `json:"template_text_schema,omitempty"`
	TemplateType       CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType                            `json:"template_type"`
}
