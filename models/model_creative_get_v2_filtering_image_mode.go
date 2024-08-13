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

// CreativeGetV2FilteringImageMode
type CreativeGetV2FilteringImageMode string

// List of creative_get_v2_filtering_image_mode
const (
	CREATIVE_IMAGE_MODE_AWEME_LIVE_CreativeGetV2FilteringImageMode                    CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_CreativeGetV2FilteringImageMode CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_CreativeGetV2FilteringImageMode        CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_CreativeGetV2FilteringImageMode             CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_CreativeGetV2FilteringImageMode               CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_CreativeGetV2FilteringImageMode                CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_GIF_CreativeGetV2FilteringImageMode                           CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_GROUP_CreativeGetV2FilteringImageMode                         CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_LARGE_CreativeGetV2FilteringImageMode                         CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_CreativeGetV2FilteringImageMode                CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_CreativeGetV2FilteringImageMode           CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_CreativeGetV2FilteringImageMode             CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_SMALL_CreativeGetV2FilteringImageMode                         CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_CreativeGetV2FilteringImageMode                  CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_CreativeGetV2FilteringImageMode            CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_CreativeGetV2FilteringImageMode                         CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_CreativeGetV2FilteringImageMode                CreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	INVALID_DATA_CreativeGetV2FilteringImageMode                                      CreativeGetV2FilteringImageMode = "INVALID_DATA"
	SEARCH_AD_SMALL_IMAGE_CreativeGetV2FilteringImageMode                             CreativeGetV2FilteringImageMode = "SEARCH_AD_SMALL_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_CreativeGetV2FilteringImageMode                           CreativeGetV2FilteringImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of CreativeGetV2FilteringImageMode enum
var AllowedCreativeGetV2FilteringImageModeEnumValues = []CreativeGetV2FilteringImageMode{
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
	"INVALID_DATA",
	"SEARCH_AD_SMALL_IMAGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
}

// NewCreativeGetV2FilteringImageModeFromValue returns a pointer to a valid CreativeGetV2FilteringImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeGetV2FilteringImageModeFromValue(v string) (*CreativeGetV2FilteringImageMode, error) {
	ev := CreativeGetV2FilteringImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeGetV2FilteringImageMode: valid values are %v", v, AllowedCreativeGetV2FilteringImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeGetV2FilteringImageMode) IsValid() bool {
	for _, existing := range AllowedCreativeGetV2FilteringImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_get_v2_filtering_image_mode value
func (v CreativeGetV2FilteringImageMode) Ptr() *CreativeGetV2FilteringImageMode {
	return &v
}
