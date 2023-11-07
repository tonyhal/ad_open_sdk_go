/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization
type QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization int64

// List of qianchuan_creative_get_v1.0_data_list_promotion_card_material_button_smart_optimization
const (
	Enum_0_QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization = 0
	Enum_1_QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization = 1
)

// All allowed values of QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization enum
var AllowedQianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimizationEnumValues = []QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization{
	0,
	1,
}

// NewQianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimizationFromValue returns a pointer to a valid QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimizationFromValue(v int64) (*QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization, error) {
	ev := QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization: valid values are %v", v, AllowedQianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimizationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimizationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_data_list_promotion_card_material_button_smart_optimization value
func (v QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization) Ptr() *QianchuanCreativeGetV10DataListPromotionCardMaterialButtonSmartOptimization {
	return &v
}
