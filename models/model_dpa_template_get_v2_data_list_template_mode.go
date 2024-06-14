/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaTemplateGetV2DataListTemplateMode
type DpaTemplateGetV2DataListTemplateMode string

// List of dpa_template_get_v2_data_list_template_mode
const (
	CREATIVE_IMAGE_MODE_SMALL_DpaTemplateGetV2DataListTemplateMode               DpaTemplateGetV2DataListTemplateMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_LARGE_DpaTemplateGetV2DataListTemplateMode               DpaTemplateGetV2DataListTemplateMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_GROUP_DpaTemplateGetV2DataListTemplateMode               DpaTemplateGetV2DataListTemplateMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_VIDEO_DpaTemplateGetV2DataListTemplateMode               DpaTemplateGetV2DataListTemplateMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_GIF_DpaTemplateGetV2DataListTemplateMode                 DpaTemplateGetV2DataListTemplateMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_DpaTemplateGetV2DataListTemplateMode      DpaTemplateGetV2DataListTemplateMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_DpaTemplateGetV2DataListTemplateMode      DpaTemplateGetV2DataListTemplateMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	TOUTIAO_SEARCH_AD_IMAGE_DpaTemplateGetV2DataListTemplateMode                 DpaTemplateGetV2DataListTemplateMode = "TOUTIAO_SEARCH_AD_IMAGE"
	SEARCH_AD_SMALL_IMAGE_DpaTemplateGetV2DataListTemplateMode                   DpaTemplateGetV2DataListTemplateMode = "SEARCH_AD_SMALL_IMAGE"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_DpaTemplateGetV2DataListTemplateMode        DpaTemplateGetV2DataListTemplateMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_DpaTemplateGetV2DataListTemplateMode  DpaTemplateGetV2DataListTemplateMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_DpaTemplateGetV2DataListTemplateMode          DpaTemplateGetV2DataListTemplateMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_DpaTemplateGetV2DataListTemplateMode      DpaTemplateGetV2DataListTemplateMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_DpaTemplateGetV2DataListTemplateMode DpaTemplateGetV2DataListTemplateMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_DpaTemplateGetV2DataListTemplateMode   DpaTemplateGetV2DataListTemplateMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	ALL_DpaTemplateGetV2DataListTemplateMode                                     DpaTemplateGetV2DataListTemplateMode = "ALL"
	NOT_ALL_DpaTemplateGetV2DataListTemplateMode                                 DpaTemplateGetV2DataListTemplateMode = "NOT_ALL"
)

// All allowed values of DpaTemplateGetV2DataListTemplateMode enum
var AllowedDpaTemplateGetV2DataListTemplateModeEnumValues = []DpaTemplateGetV2DataListTemplateMode{
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_GROUP",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"TOUTIAO_SEARCH_AD_IMAGE",
	"SEARCH_AD_SMALL_IMAGE",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
	"ALL",
	"NOT_ALL",
}

// NewDpaTemplateGetV2DataListTemplateModeFromValue returns a pointer to a valid DpaTemplateGetV2DataListTemplateMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaTemplateGetV2DataListTemplateModeFromValue(v string) (*DpaTemplateGetV2DataListTemplateMode, error) {
	ev := DpaTemplateGetV2DataListTemplateMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaTemplateGetV2DataListTemplateMode: valid values are %v", v, AllowedDpaTemplateGetV2DataListTemplateModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaTemplateGetV2DataListTemplateMode) IsValid() bool {
	for _, existing := range AllowedDpaTemplateGetV2DataListTemplateModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_template_get_v2_data_list_template_mode value
func (v DpaTemplateGetV2DataListTemplateMode) Ptr() *DpaTemplateGetV2DataListTemplateMode {
	return &v
}
