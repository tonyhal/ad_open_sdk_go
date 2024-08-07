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

// PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility
type PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility string

// List of promotion_update_v3.0_promotion_materials_video_material_list_video_hp_visibility
const (
	ALWAYS_VISIBLE_PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility   PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility = "ALWAYS_VISIBLE"
	HIDE_VIDEO_ON_HP_PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility = "HIDE_VIDEO_ON_HP"
)

// All allowed values of PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility enum
var AllowedPromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibilityEnumValues = []PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility{
	"ALWAYS_VISIBLE",
	"HIDE_VIDEO_ON_HP",
}

// NewPromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibilityFromValue returns a pointer to a valid PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibilityFromValue(v string) (*PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility, error) {
	ev := PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility: valid values are %v", v, AllowedPromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_promotion_materials_video_material_list_video_hp_visibility value
func (v PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility) Ptr() *PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoHpVisibility {
	return &v
}
