/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeUpdateV2RequestCreativeVideoMaterialsInner struct for CreativeProceduralCreativeUpdateV2RequestCreativeVideoMaterialsInner
type CreativeProceduralCreativeUpdateV2RequestCreativeVideoMaterialsInner struct {
	//
	DpaVideoTaskIds      []string                                                                       `json:"dpa_video_task_ids,omitempty"`
	DpaVideoTemplateType *CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType  `json:"dpa_video_template_type,omitempty"`
	ImageInfo            *CreativeProceduralCreativeUpdateV2RequestCreativeVideoMaterialsInnerImageInfo `json:"image_info,omitempty"`
	ImageMode            *CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsImageMode             `json:"image_mode,omitempty"`
	//
	MaterialId *int64                                                                         `json:"material_id,omitempty"`
	VideoInfo  *CreativeProceduralCreativeUpdateV2RequestCreativeVideoMaterialsInnerVideoInfo `json:"video_info,omitempty"`
}
