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

// DpaVideoGetV2DataListVideoInfoImageMode
type DpaVideoGetV2DataListVideoInfoImageMode string

// List of dpa_video_get_v2_data_list_video_info_image_mode
const (
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_DpaVideoGetV2DataListVideoInfoImageMode                DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_DpaVideoGetV2DataListVideoInfoImageMode                DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_DpaVideoGetV2DataListVideoInfoImageMode               DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_DpaVideoGetV2DataListVideoInfoImageMode DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_DpaVideoGetV2DataListVideoInfoImageMode            DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_SMALL_DpaVideoGetV2DataListVideoInfoImageMode                         DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_DpaVideoGetV2DataListVideoInfoImageMode             DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_DpaVideoGetV2DataListVideoInfoImageMode                  DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_DpaVideoGetV2DataListVideoInfoImageMode           DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_GROUP_DpaVideoGetV2DataListVideoInfoImageMode                         DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	TOUTIAO_SEARCH_AD_IMAGE_DpaVideoGetV2DataListVideoInfoImageMode                           DpaVideoGetV2DataListVideoInfoImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_DpaVideoGetV2DataListVideoInfoImageMode        DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_GIF_DpaVideoGetV2DataListVideoInfoImageMode                           DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_DpaVideoGetV2DataListVideoInfoImageMode                DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_LARGE_DpaVideoGetV2DataListVideoInfoImageMode                         DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_VIDEO_DpaVideoGetV2DataListVideoInfoImageMode                         DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_DpaVideoGetV2DataListVideoInfoImageMode             DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_DpaVideoGetV2DataListVideoInfoImageMode                    DpaVideoGetV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	SEARCH_AD_SMALL_IMAGE_DpaVideoGetV2DataListVideoInfoImageMode                             DpaVideoGetV2DataListVideoInfoImageMode = "SEARCH_AD_SMALL_IMAGE"
)

// All allowed values of DpaVideoGetV2DataListVideoInfoImageMode enum
var AllowedDpaVideoGetV2DataListVideoInfoImageModeEnumValues = []DpaVideoGetV2DataListVideoInfoImageMode{
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

// NewDpaVideoGetV2DataListVideoInfoImageModeFromValue returns a pointer to a valid DpaVideoGetV2DataListVideoInfoImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaVideoGetV2DataListVideoInfoImageModeFromValue(v string) (*DpaVideoGetV2DataListVideoInfoImageMode, error) {
	ev := DpaVideoGetV2DataListVideoInfoImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaVideoGetV2DataListVideoInfoImageMode: valid values are %v", v, AllowedDpaVideoGetV2DataListVideoInfoImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaVideoGetV2DataListVideoInfoImageMode) IsValid() bool {
	for _, existing := range AllowedDpaVideoGetV2DataListVideoInfoImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_video_get_v2_data_list_video_info_image_mode value
func (v DpaVideoGetV2DataListVideoInfoImageMode) Ptr() *DpaVideoGetV2DataListVideoInfoImageMode {
	return &v
}
