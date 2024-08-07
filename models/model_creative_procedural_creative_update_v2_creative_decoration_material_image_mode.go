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

// CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode
type CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode string

// List of creative_procedural_creative_update_v2_creative_decoration_material_image_mode
const (
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode                CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode                CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode               CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode            CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_SMALL_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode                         CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode             CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode                  CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode           CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_GROUP_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode                         CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	TOUTIAO_SEARCH_AD_IMAGE_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode                           CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode        CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_GIF_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode                           CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode                CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_LARGE_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode                         CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_VIDEO_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode                         CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode             CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode                    CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	SEARCH_AD_SMALL_IMAGE_CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode                             CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode = "SEARCH_AD_SMALL_IMAGE"
)

// All allowed values of CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode enum
var AllowedCreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageModeEnumValues = []CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode{
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
	"CREATIVE_IMAGE_MODE_GROUP",
	"TOUTIAO_SEARCH_AD_IMAGE",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"SEARCH_AD_SMALL_IMAGE",
}

// NewCreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageModeFromValue returns a pointer to a valid CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageModeFromValue(v string) (*CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode, error) {
	ev := CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode: valid values are %v", v, AllowedCreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_update_v2_creative_decoration_material_image_mode value
func (v CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode) Ptr() *CreativeProceduralCreativeUpdateV2CreativeDecorationMaterialImageMode {
	return &v
}
