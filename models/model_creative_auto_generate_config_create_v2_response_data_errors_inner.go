/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeAutoGenerateConfigCreateV2ResponseDataErrorsInner struct for CreativeAutoGenerateConfigCreateV2ResponseDataErrorsInner
type CreativeAutoGenerateConfigCreateV2ResponseDataErrorsInner struct {
	// 错误信息列举： 1.template_id查不到：模板ID不存在 2.模板填充的内容校验不通过：模板填充内容不符合要求，请确认填充内容的类型、格式、大小、长度等要求 3.内部报错等其他原因：系统错误
	ErrorMessage *string `json:"error_message,omitempty"`
	// 模板ID
	TemplateId *int64 `json:"template_id,omitempty"`
}
