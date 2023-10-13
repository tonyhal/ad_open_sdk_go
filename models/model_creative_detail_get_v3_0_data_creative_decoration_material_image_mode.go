/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeDetailGetV30DataCreativeDecorationMaterialImageMode
type CreativeDetailGetV30DataCreativeDecorationMaterialImageMode string

// List of creative_detail_get_v3.0_data_creative_decoration_material_image_mode
const (
	CREATIVE_IMAGE_MODE_AWEME_LIVE_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode                    CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode        CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode             CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode               CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode                CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_GIF_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode                           CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_GROUP_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode                         CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_LARGE_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode                         CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode                CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode           CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode             CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_SMALL_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode                         CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode                  CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode            CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode                         CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode                CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	SEARCH_AD_SMALL_IMAGE_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode                             CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "SEARCH_AD_SMALL_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_CreativeDetailGetV30DataCreativeDecorationMaterialImageMode                           CreativeDetailGetV30DataCreativeDecorationMaterialImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of CreativeDetailGetV30DataCreativeDecorationMaterialImageMode enum
var AllowedCreativeDetailGetV30DataCreativeDecorationMaterialImageModeEnumValues = []CreativeDetailGetV30DataCreativeDecorationMaterialImageMode{
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO",
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
	"CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_GROUP",
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"SEARCH_AD_SMALL_IMAGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
}

// NewCreativeDetailGetV30DataCreativeDecorationMaterialImageModeFromValue returns a pointer to a valid CreativeDetailGetV30DataCreativeDecorationMaterialImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeDetailGetV30DataCreativeDecorationMaterialImageModeFromValue(v string) (*CreativeDetailGetV30DataCreativeDecorationMaterialImageMode, error) {
	ev := CreativeDetailGetV30DataCreativeDecorationMaterialImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeDetailGetV30DataCreativeDecorationMaterialImageMode: valid values are %v", v, AllowedCreativeDetailGetV30DataCreativeDecorationMaterialImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeDetailGetV30DataCreativeDecorationMaterialImageMode) IsValid() bool {
	for _, existing := range AllowedCreativeDetailGetV30DataCreativeDecorationMaterialImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_detail_get_v3.0_data_creative_decoration_material_image_mode value
func (v CreativeDetailGetV30DataCreativeDecorationMaterialImageMode) Ptr() *CreativeDetailGetV30DataCreativeDecorationMaterialImageMode {
	return &v
}
