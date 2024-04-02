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

// ReportAdGetV2FilteringImageModes
type ReportAdGetV2FilteringImageModes string

// List of report_ad_get_v2_filtering_image_modes
const (
	CREATIVE_IMAGE_MODE_LARGE_ReportAdGetV2FilteringImageModes              ReportAdGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_LARGE"
	TOUTIAO_SEARCH_AD_IMAGE_ReportAdGetV2FilteringImageModes                ReportAdGetV2FilteringImageModes = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_ReportAdGetV2FilteringImageModes     ReportAdGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_ReportAdGetV2FilteringImageModes     ReportAdGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_ReportAdGetV2FilteringImageModes ReportAdGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_ReportAdGetV2FilteringImageModes  ReportAdGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_GROUP_ReportAdGetV2FilteringImageModes              ReportAdGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_GIF_ReportAdGetV2FilteringImageModes                ReportAdGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_ReportAdGetV2FilteringImageModes       ReportAdGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_ReportAdGetV2FilteringImageModes         ReportAdGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_ReportAdGetV2FilteringImageModes     ReportAdGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_VIDEO_ReportAdGetV2FilteringImageModes              ReportAdGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_SMALL_ReportAdGetV2FilteringImageModes              ReportAdGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_SMALL"
	SEARCH_AD_SMALL_IMAGE_ReportAdGetV2FilteringImageModes                  ReportAdGetV2FilteringImageModes = "SEARCH_AD_SMALL_IMAGE"
)

// All allowed values of ReportAdGetV2FilteringImageModes enum
var AllowedReportAdGetV2FilteringImageModesEnumValues = []ReportAdGetV2FilteringImageModes{
	"CREATIVE_IMAGE_MODE_LARGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
	"CREATIVE_IMAGE_MODE_GROUP",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_SMALL",
	"SEARCH_AD_SMALL_IMAGE",
}

// NewReportAdGetV2FilteringImageModesFromValue returns a pointer to a valid ReportAdGetV2FilteringImageModes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2FilteringImageModesFromValue(v string) (*ReportAdGetV2FilteringImageModes, error) {
	ev := ReportAdGetV2FilteringImageModes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2FilteringImageModes: valid values are %v", v, AllowedReportAdGetV2FilteringImageModesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2FilteringImageModes) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2FilteringImageModesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_filtering_image_modes value
func (v ReportAdGetV2FilteringImageModes) Ptr() *ReportAdGetV2FilteringImageModes {
	return &v
}
