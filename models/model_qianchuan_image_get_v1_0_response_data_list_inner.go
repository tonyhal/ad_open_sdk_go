/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanImageGetV10ResponseDataListInner struct for QianchuanImageGetV10ResponseDataListInner
type QianchuanImageGetV10ResponseDataListInner struct {
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	Filename *string `json:"filename,omitempty"`
	//
	Height *int64 `json:"height,omitempty"`
	//
	Id        *string                                `json:"id,omitempty"`
	ImageMode *QianchuanImageGetV10DataListImageMode `json:"image_mode,omitempty"`
	//
	IsAiCreate *bool `json:"is_ai_create,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	Signature *string `json:"signature,omitempty"`
	//
	Size   *int64                              `json:"size,omitempty"`
	Source *QianchuanImageGetV10DataListSource `json:"source,omitempty"`
	//
	Tag []string `json:"tag,omitempty"`
	//
	Url *string `json:"url,omitempty"`
	//
	Width *int64 `json:"width,omitempty"`
}