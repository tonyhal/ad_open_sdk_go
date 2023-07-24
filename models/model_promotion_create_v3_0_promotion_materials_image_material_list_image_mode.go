/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30PromotionMaterialsImageMaterialListImageMode
type PromotionCreateV30PromotionMaterialsImageMaterialListImageMode string

// List of promotion_create_v3.0_promotion_materials_image_material_list_image_mode
const (
	CREATIVE_IMAGE_MODE_LARGE_PromotionCreateV30PromotionMaterialsImageMaterialListImageMode          PromotionCreateV30PromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_PromotionCreateV30PromotionMaterialsImageMaterialListImageMode PromotionCreateV30PromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_SMALL_PromotionCreateV30PromotionMaterialsImageMaterialListImageMode          PromotionCreateV30PromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_PromotionCreateV30PromotionMaterialsImageMaterialListImageMode   PromotionCreateV30PromotionMaterialsImageMaterialListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	SEARCH_AD_SMALL_IMAGE_PromotionCreateV30PromotionMaterialsImageMaterialListImageMode              PromotionCreateV30PromotionMaterialsImageMaterialListImageMode = "SEARCH_AD_SMALL_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_PromotionCreateV30PromotionMaterialsImageMaterialListImageMode            PromotionCreateV30PromotionMaterialsImageMaterialListImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of PromotionCreateV30PromotionMaterialsImageMaterialListImageMode enum
var AllowedPromotionCreateV30PromotionMaterialsImageMaterialListImageModeEnumValues = []PromotionCreateV30PromotionMaterialsImageMaterialListImageMode{
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"SEARCH_AD_SMALL_IMAGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
}

// NewPromotionCreateV30PromotionMaterialsImageMaterialListImageModeFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsImageMaterialListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsImageMaterialListImageModeFromValue(v string) (*PromotionCreateV30PromotionMaterialsImageMaterialListImageMode, error) {
	ev := PromotionCreateV30PromotionMaterialsImageMaterialListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsImageMaterialListImageMode: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsImageMaterialListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsImageMaterialListImageMode) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsImageMaterialListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_image_material_list_image_mode value
func (v PromotionCreateV30PromotionMaterialsImageMaterialListImageMode) Ptr() *PromotionCreateV30PromotionMaterialsImageMaterialListImageMode {
	return &v
}
