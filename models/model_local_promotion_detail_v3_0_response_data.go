/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalPromotionDetailV30ResponseData
type LocalPromotionDetailV30ResponseData struct {
	// 抖音号
	AwemeId *string `json:"aweme_id,omitempty"`
	// 自定义素材组合
	CustomerMaterialList []*LocalPromotionDetailV30ResponseDataCustomerMaterialListInner `json:"customer_material_list,omitempty"`
	// 是否开启团购卡
	EnableGraphicDelivery *bool                                        `json:"enable_graphic_delivery,omitempty"`
	LiveMaterialType      *LocalPromotionDetailV30DataLiveMaterialType `json:"live_material_type,omitempty"`
	// 广告ID
	PromotionId       *int64                                        `json:"promotion_id,omitempty"`
	VideoHpVisibility *LocalPromotionDetailV30DataVideoHpVisibility `json:"video_hp_visibility,omitempty"`
}
