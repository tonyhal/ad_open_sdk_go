/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionAidGetV30ResponseDataPromotionMapDataInnerAdsInnerCreativesInnerImageMaterials
type PromotionAidGetV30ResponseDataPromotionMapDataInnerAdsInnerCreativesInnerImageMaterials struct {
	ImageMode *PromotionAidGetV30DataPromotionMapDataAdsCreativesImageMaterialsImageMode `json:"image_mode,omitempty"`
	//
	ImageUris []string `json:"image_uris,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
}
