/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType
type QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType string

// List of qianchuan_uni_promotion_ad_material_get_v1.0_filtering_analysis_type
const (
	FIRST_PUBLISH_MATERIAL_QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType = "FIRST_PUBLISH_MATERIAL"
	HIGH_QUALITY_MATERIAL_QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType  QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType = "HIGH_QUALITY_MATERIAL"
	LOW_QUALITY_MATERIAL_QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType   QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType = "LOW_QUALITY_MATERIAL"
	INEFFICIENT_MATERIAL_QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType   QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType = "INEFFICIENT_MATERIAL"
	CARRY_MATERIAL_QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType         QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType = "CARRY_MATERIAL"
	SIMILAR_MATERIAL_QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType       QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType = "SIMILAR_MATERIAL"
)

// All allowed values of QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType enum
var AllowedQianchuanUniPromotionAdMaterialGetV10FilteringAnalysisTypeEnumValues = []QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType{
	"FIRST_PUBLISH_MATERIAL",
	"HIGH_QUALITY_MATERIAL",
	"LOW_QUALITY_MATERIAL",
	"INEFFICIENT_MATERIAL",
	"CARRY_MATERIAL",
	"SIMILAR_MATERIAL",
}

// NewQianchuanUniPromotionAdMaterialGetV10FilteringAnalysisTypeFromValue returns a pointer to a valid QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionAdMaterialGetV10FilteringAnalysisTypeFromValue(v string) (*QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType, error) {
	ev := QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType: valid values are %v", v, AllowedQianchuanUniPromotionAdMaterialGetV10FilteringAnalysisTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionAdMaterialGetV10FilteringAnalysisTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_ad_material_get_v1.0_filtering_analysis_type value
func (v QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType) Ptr() *QianchuanUniPromotionAdMaterialGetV10FilteringAnalysisType {
	return &v
}