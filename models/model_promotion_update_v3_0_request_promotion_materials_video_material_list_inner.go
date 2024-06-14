/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30RequestPromotionMaterialsVideoMaterialListInner struct for PromotionUpdateV30RequestPromotionMaterialsVideoMaterialListInner
type PromotionUpdateV30RequestPromotionMaterialsVideoMaterialListInner struct {
	// 引导视频id *不传表示删除引导视频。引导视频一旦修改，广告将重新进入审核，请谨慎操作
	GuideVideoId *string                                                        `json:"guide_video_id,omitempty"`
	ImageMode    PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode `json:"image_mode"`
	//
	ItemId *int64 `json:"item_id,omitempty"`
	//
	VideoCoverId      *string                                                                 `json:"video_cover_id,omitempty"`
	VideoHpVisibility *PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility `json:"video_hp_visibility,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
	//
	VideoTaskIds      []string                                                                `json:"video_task_ids,omitempty"`
	VideoTemplateType *PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType `json:"video_template_type,omitempty"`
}
