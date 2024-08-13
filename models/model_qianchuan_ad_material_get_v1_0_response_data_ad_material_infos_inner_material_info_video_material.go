/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdMaterialGetV10ResponseDataAdMaterialInfosInnerMaterialInfoVideoMaterial 视频素材
type QianchuanAdMaterialGetV10ResponseDataAdMaterialInfosInnerMaterialInfoVideoMaterial struct {
	CoverImage *QianchuanAdMaterialGetV10ResponseDataAdMaterialInfosInnerMaterialInfoVideoMaterialCoverImage `json:"cover_image,omitempty"`
	ImageMode  *QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode               `json:"image_mode,omitempty"`
	// 更新时需要传，创建时不需要传
	MaterialId *int64                                                                       `json:"material_id,omitempty"`
	Source     *QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource `json:"source,omitempty"`
	// 视频标题
	Title *string `json:"title,omitempty"`
	// 视频时长
	VideoDuration *int64 `json:"video_duration,omitempty"`
	// 视频 vid
	VideoId *string `json:"video_id,omitempty"`
}
