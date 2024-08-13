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

// PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility
type PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility string

// List of promotion_create_v3.0_promotion_materials_carousel_material_list_video_hp_visibility
const (
	ALWAYS_VISIBLE_PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility   PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility = "ALWAYS_VISIBLE"
	HIDE_VIDEO_ON_HP_PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility = "HIDE_VIDEO_ON_HP"
)

// All allowed values of PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility enum
var AllowedPromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibilityEnumValues = []PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility{
	"ALWAYS_VISIBLE",
	"HIDE_VIDEO_ON_HP",
}

// NewPromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibilityFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibilityFromValue(v string) (*PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility, error) {
	ev := PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_carousel_material_list_video_hp_visibility value
func (v PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility) Ptr() *PromotionCreateV30PromotionMaterialsCarouselMaterialListVideoHpVisibility {
	return &v
}
