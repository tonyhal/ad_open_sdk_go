/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerMaterialScoreInfo
type PromotionListV30ResponseDataListInnerMaterialScoreInfo struct {
	LowQualityMaterialList *PromotionListV30ResponseDataListInnerMaterialScoreInfoLowQualityMaterialList `json:"low_quality_material_list,omitempty"`
	//
	MaterialAdvice       []string                                                       `json:"material_advice,omitempty"`
	ScoreNumOfMaterial   *PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial   `json:"score_num_of_material,omitempty"`
	ScoreTypeOfMaterial  *PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial  `json:"score_type_of_material,omitempty"`
	ScoreValueOfMaterial *PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial `json:"score_value_of_material,omitempty"`
}
