/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeAutoGenerateConfigCreateV2ResponseData
type CreativeAutoGenerateConfigCreateV2ResponseData struct {
	// 配置ID
	ConfigId *int64 `json:"config_id,omitempty"`
	// 保存失败的模板列表
	Errors []*CreativeAutoGenerateConfigCreateV2ResponseDataErrorsInner `json:"errors,omitempty"`
	// 保存成功的模板列表
	Templates []*CreativeAutoGenerateConfigCreateV2ResponseDataTemplatesInner `json:"templates,omitempty"`
}
