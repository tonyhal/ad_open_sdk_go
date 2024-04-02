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

// CreativeDetailGetV30DataCreativeVideoMaterialsImageMode
type CreativeDetailGetV30DataCreativeVideoMaterialsImageMode string

// List of creative_detail_get_v3.0_data_creative_video_materials_image_mode
const (
	CREATIVE_IMAGE_MODE_AWEME_LIVE_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode                    CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode        CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode               CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode                CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_GIF_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode                           CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_GROUP_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode                         CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_LARGE_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode                         CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode                CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode           CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode             CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_SMALL_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode                         CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode                  CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode            CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode                         CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode                CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	SEARCH_AD_SMALL_IMAGE_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode                             CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "SEARCH_AD_SMALL_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_CreativeDetailGetV30DataCreativeVideoMaterialsImageMode                           CreativeDetailGetV30DataCreativeVideoMaterialsImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of CreativeDetailGetV30DataCreativeVideoMaterialsImageMode enum
var AllowedCreativeDetailGetV30DataCreativeVideoMaterialsImageModeEnumValues = []CreativeDetailGetV30DataCreativeVideoMaterialsImageMode{
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO",
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

// NewCreativeDetailGetV30DataCreativeVideoMaterialsImageModeFromValue returns a pointer to a valid CreativeDetailGetV30DataCreativeVideoMaterialsImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeDetailGetV30DataCreativeVideoMaterialsImageModeFromValue(v string) (*CreativeDetailGetV30DataCreativeVideoMaterialsImageMode, error) {
	ev := CreativeDetailGetV30DataCreativeVideoMaterialsImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeDetailGetV30DataCreativeVideoMaterialsImageMode: valid values are %v", v, AllowedCreativeDetailGetV30DataCreativeVideoMaterialsImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeDetailGetV30DataCreativeVideoMaterialsImageMode) IsValid() bool {
	for _, existing := range AllowedCreativeDetailGetV30DataCreativeVideoMaterialsImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_detail_get_v3.0_data_creative_video_materials_image_mode value
func (v CreativeDetailGetV30DataCreativeVideoMaterialsImageMode) Ptr() *CreativeDetailGetV30DataCreativeVideoMaterialsImageMode {
	return &v
}
