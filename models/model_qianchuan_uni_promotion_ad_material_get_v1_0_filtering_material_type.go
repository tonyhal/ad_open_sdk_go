/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType
type QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType string

// List of qianchuan_uni_promotion_ad_material_get_v1.0_filtering_material_type
const (
	IMAGE_QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType     QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType = "IMAGE"
	LIVE_ROOM_QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType = "LIVE_ROOM"
	TITLE_QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType     QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType = "TITLE"
	VIDEO_QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType     QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType = "VIDEO"
)

// All allowed values of QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType enum
var AllowedQianchuanUniPromotionAdMaterialGetV10FilteringMaterialTypeEnumValues = []QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType{
	"IMAGE",
	"LIVE_ROOM",
	"TITLE",
	"VIDEO",
}

// NewQianchuanUniPromotionAdMaterialGetV10FilteringMaterialTypeFromValue returns a pointer to a valid QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionAdMaterialGetV10FilteringMaterialTypeFromValue(v string) (*QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType, error) {
	ev := QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType: valid values are %v", v, AllowedQianchuanUniPromotionAdMaterialGetV10FilteringMaterialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionAdMaterialGetV10FilteringMaterialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_ad_material_get_v1.0_filtering_material_type value
func (v QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType) Ptr() *QianchuanUniPromotionAdMaterialGetV10FilteringMaterialType {
	return &v
}
