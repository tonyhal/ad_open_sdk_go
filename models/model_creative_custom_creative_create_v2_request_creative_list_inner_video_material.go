/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeCustomCreativeCreateV2RequestCreativeListInnerVideoMaterial
type CreativeCustomCreativeCreateV2RequestCreativeListInnerVideoMaterial struct {
	//
	AwemeItemId *int64 `json:"aweme_item_id,omitempty"`
	//
	DpaVideoTaskIds      []string                                                                      `json:"dpa_video_task_ids,omitempty"`
	DpaVideoTemplateType *CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType  `json:"dpa_video_template_type,omitempty"`
	ImageInfo            *CreativeCustomCreativeCreateV2RequestCreativeListInnerVideoMaterialImageInfo `json:"image_info,omitempty"`
	VideoInfo            *CreativeCustomCreativeCreateV2RequestCreativeListInnerVideoMaterialVideoInfo `json:"video_info,omitempty"`
}