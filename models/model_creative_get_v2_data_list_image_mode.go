/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeGetV2DataListImageMode
type CreativeGetV2DataListImageMode string

// List of creative_get_v2_data_list_image_mode
const (
	CREATIVE_IMAGE_MODE_AWEME_LIVE_CreativeGetV2DataListImageMode                    CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_CreativeGetV2DataListImageMode CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_CreativeGetV2DataListImageMode        CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_CreativeGetV2DataListImageMode             CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_CreativeGetV2DataListImageMode               CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_CreativeGetV2DataListImageMode                CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_GIF_CreativeGetV2DataListImageMode                           CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_GROUP_CreativeGetV2DataListImageMode                         CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_LARGE_CreativeGetV2DataListImageMode                         CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_CreativeGetV2DataListImageMode                CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_CreativeGetV2DataListImageMode           CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_CreativeGetV2DataListImageMode             CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_SMALL_CreativeGetV2DataListImageMode                         CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_CreativeGetV2DataListImageMode                  CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_CreativeGetV2DataListImageMode            CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_CreativeGetV2DataListImageMode                         CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_CreativeGetV2DataListImageMode                CreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	INVALID_DATA_CreativeGetV2DataListImageMode                                      CreativeGetV2DataListImageMode = "INVALID_DATA"
	SEARCH_AD_SMALL_IMAGE_CreativeGetV2DataListImageMode                             CreativeGetV2DataListImageMode = "SEARCH_AD_SMALL_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_CreativeGetV2DataListImageMode                           CreativeGetV2DataListImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of CreativeGetV2DataListImageMode enum
var AllowedCreativeGetV2DataListImageModeEnumValues = []CreativeGetV2DataListImageMode{
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

// NewCreativeGetV2DataListImageModeFromValue returns a pointer to a valid CreativeGetV2DataListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeGetV2DataListImageModeFromValue(v string) (*CreativeGetV2DataListImageMode, error) {
	ev := CreativeGetV2DataListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeGetV2DataListImageMode: valid values are %v", v, AllowedCreativeGetV2DataListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeGetV2DataListImageMode) IsValid() bool {
	for _, existing := range AllowedCreativeGetV2DataListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_get_v2_data_list_image_mode value
func (v CreativeGetV2DataListImageMode) Ptr() *CreativeGetV2DataListImageMode {
	return &v
}
