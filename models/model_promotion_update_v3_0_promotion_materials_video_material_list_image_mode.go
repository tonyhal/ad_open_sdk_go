/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode
type PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode string

// List of promotion_update_v3.0_promotion_materials_video_material_list_image_mode
const (
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode                  PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode         PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
)

// All allowed values of PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode enum
var AllowedPromotionUpdateV30PromotionMaterialsVideoMaterialListImageModeEnumValues = []PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode{
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
}

// NewPromotionUpdateV30PromotionMaterialsVideoMaterialListImageModeFromValue returns a pointer to a valid PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30PromotionMaterialsVideoMaterialListImageModeFromValue(v string) (*PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode, error) {
	ev := PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode: valid values are %v", v, AllowedPromotionUpdateV30PromotionMaterialsVideoMaterialListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30PromotionMaterialsVideoMaterialListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_promotion_materials_video_material_list_image_mode value
func (v PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode) Ptr() *PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode {
	return &v
}
