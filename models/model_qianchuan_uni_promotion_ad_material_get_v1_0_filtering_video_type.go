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

// QianchuanUniPromotionAdMaterialGetV10FilteringVideoType
type QianchuanUniPromotionAdMaterialGetV10FilteringVideoType string

// List of qianchuan_uni_promotion_ad_material_get_v1.0_filtering_video_type
const (
	ALL_QianchuanUniPromotionAdMaterialGetV10FilteringVideoType    QianchuanUniPromotionAdMaterialGetV10FilteringVideoType = "ALL"
	CUSTOM_QianchuanUniPromotionAdMaterialGetV10FilteringVideoType QianchuanUniPromotionAdMaterialGetV10FilteringVideoType = "CUSTOM"
	AUTO_QianchuanUniPromotionAdMaterialGetV10FilteringVideoType   QianchuanUniPromotionAdMaterialGetV10FilteringVideoType = "AUTO"
)

// All allowed values of QianchuanUniPromotionAdMaterialGetV10FilteringVideoType enum
var AllowedQianchuanUniPromotionAdMaterialGetV10FilteringVideoTypeEnumValues = []QianchuanUniPromotionAdMaterialGetV10FilteringVideoType{
	"ALL",
	"CUSTOM",
	"AUTO",
}

// NewQianchuanUniPromotionAdMaterialGetV10FilteringVideoTypeFromValue returns a pointer to a valid QianchuanUniPromotionAdMaterialGetV10FilteringVideoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionAdMaterialGetV10FilteringVideoTypeFromValue(v string) (*QianchuanUniPromotionAdMaterialGetV10FilteringVideoType, error) {
	ev := QianchuanUniPromotionAdMaterialGetV10FilteringVideoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionAdMaterialGetV10FilteringVideoType: valid values are %v", v, AllowedQianchuanUniPromotionAdMaterialGetV10FilteringVideoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionAdMaterialGetV10FilteringVideoType) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionAdMaterialGetV10FilteringVideoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_ad_material_get_v1.0_filtering_video_type value
func (v QianchuanUniPromotionAdMaterialGetV10FilteringVideoType) Ptr() *QianchuanUniPromotionAdMaterialGetV10FilteringVideoType {
	return &v
}