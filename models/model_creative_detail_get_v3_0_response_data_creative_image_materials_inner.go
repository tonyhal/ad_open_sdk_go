/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeDetailGetV30ResponseDataCreativeImageMaterialsInner struct for CreativeDetailGetV30ResponseDataCreativeImageMaterialsInner
type CreativeDetailGetV30ResponseDataCreativeImageMaterialsInner struct {
	// 图片素材信息
	ImageInfo     []*CreativeDetailGetV30ResponseDataCreativeImageMaterialsInnerImageInfoInner `json:"image_info,omitempty"`
	ImageMode     *CreativeDetailGetV30DataCreativeImageMaterialsImageMode                     `json:"image_mode,omitempty"`
	TemplateImage *CreativeDetailGetV30ResponseDataCreativeImageMaterialsInnerTemplateImage    `json:"template_image,omitempty"`
}
