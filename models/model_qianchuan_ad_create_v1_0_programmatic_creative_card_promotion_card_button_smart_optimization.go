/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization
type QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization int64

// List of qianchuan_ad_create_v1.0_programmatic_creative_card_promotion_card_button_smart_optimization
const (
	Enum_0_QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization = 0
	Enum_1_QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization = 1
)

// All allowed values of QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization enum
var AllowedQianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimizationEnumValues = []QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization{
	0,
	1,
}

// NewQianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimizationFromValue returns a pointer to a valid QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimizationFromValue(v int64) (*QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization, error) {
	ev := QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization: valid values are %v", v, AllowedQianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimizationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimizationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_programmatic_creative_card_promotion_card_button_smart_optimization value
func (v QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization) Ptr() *QianchuanAdCreateV10ProgrammaticCreativeCardPromotionCardButtonSmartOptimization {
	return &v
}
