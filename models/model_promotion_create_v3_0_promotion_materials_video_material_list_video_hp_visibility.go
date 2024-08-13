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

// PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility
type PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility string

// List of promotion_create_v3.0_promotion_materials_video_material_list_video_hp_visibility
const (
	ALWAYS_VISIBLE_PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility   PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility = "ALWAYS_VISIBLE"
	HIDE_VIDEO_ON_HP_PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility = "HIDE_VIDEO_ON_HP"
)

// All allowed values of PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility enum
var AllowedPromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibilityEnumValues = []PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility{
	"ALWAYS_VISIBLE",
	"HIDE_VIDEO_ON_HP",
}

// NewPromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibilityFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibilityFromValue(v string) (*PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility, error) {
	ev := PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_video_material_list_video_hp_visibility value
func (v PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility) Ptr() *PromotionCreateV30PromotionMaterialsVideoMaterialListVideoHpVisibility {
	return &v
}
