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

// ReportCampaignGetV2FilteringImageModes
type ReportCampaignGetV2FilteringImageModes string

// List of report_campaign_get_v2_filtering_image_modes
const (
	CREATIVE_IMAGE_MODE_LARGE_ReportCampaignGetV2FilteringImageModes              ReportCampaignGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_LARGE"
	TOUTIAO_SEARCH_AD_IMAGE_ReportCampaignGetV2FilteringImageModes                ReportCampaignGetV2FilteringImageModes = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_ReportCampaignGetV2FilteringImageModes     ReportCampaignGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_ReportCampaignGetV2FilteringImageModes     ReportCampaignGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_ReportCampaignGetV2FilteringImageModes ReportCampaignGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_ReportCampaignGetV2FilteringImageModes  ReportCampaignGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_GROUP_ReportCampaignGetV2FilteringImageModes              ReportCampaignGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_GIF_ReportCampaignGetV2FilteringImageModes                ReportCampaignGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_ReportCampaignGetV2FilteringImageModes       ReportCampaignGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_ReportCampaignGetV2FilteringImageModes         ReportCampaignGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_ReportCampaignGetV2FilteringImageModes     ReportCampaignGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_VIDEO_ReportCampaignGetV2FilteringImageModes              ReportCampaignGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_SMALL_ReportCampaignGetV2FilteringImageModes              ReportCampaignGetV2FilteringImageModes = "CREATIVE_IMAGE_MODE_SMALL"
	SEARCH_AD_SMALL_IMAGE_ReportCampaignGetV2FilteringImageModes                  ReportCampaignGetV2FilteringImageModes = "SEARCH_AD_SMALL_IMAGE"
)

// All allowed values of ReportCampaignGetV2FilteringImageModes enum
var AllowedReportCampaignGetV2FilteringImageModesEnumValues = []ReportCampaignGetV2FilteringImageModes{
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

// NewReportCampaignGetV2FilteringImageModesFromValue returns a pointer to a valid ReportCampaignGetV2FilteringImageModes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2FilteringImageModesFromValue(v string) (*ReportCampaignGetV2FilteringImageModes, error) {
	ev := ReportCampaignGetV2FilteringImageModes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2FilteringImageModes: valid values are %v", v, AllowedReportCampaignGetV2FilteringImageModesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2FilteringImageModes) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2FilteringImageModesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_filtering_image_modes value
func (v ReportCampaignGetV2FilteringImageModes) Ptr() *ReportCampaignGetV2FilteringImageModes {
	return &v
}
