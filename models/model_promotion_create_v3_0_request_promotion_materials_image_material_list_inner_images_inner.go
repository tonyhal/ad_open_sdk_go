/*
API Title

巨量引擎开放平台

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30RequestPromotionMaterialsImageMaterialListInnerImagesInner struct for PromotionCreateV30RequestPromotionMaterialsImageMaterialListInnerImagesInner
type PromotionCreateV30RequestPromotionMaterialsImageMaterialListInnerImagesInner struct {
	//
	ImageId *string `json:"image_id,omitempty"`
	//
	TemplateDataList []*PromotionCreateV30RequestPromotionMaterialsImageMaterialListInnerImagesInnerTemplateDataListInner `json:"template_data_list,omitempty"`
	//
	TemplateId *int64 `json:"template_id,omitempty"`
}
