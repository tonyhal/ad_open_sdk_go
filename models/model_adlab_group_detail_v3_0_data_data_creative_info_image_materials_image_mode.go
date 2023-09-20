/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode
type AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode string

// List of adlab_group_detail_v3.0_data_data_creative_info_image_materials_image_mode
const (
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode     AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_LARGE_AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode              AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode     AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_SMALL_AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode              AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode       AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode              AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode     AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	SEARCH_AD_SMALL_IMAGE_AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode                  AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode = "SEARCH_AD_SMALL_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode                AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode enum
var AllowedAdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageModeEnumValues = []AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode{
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"SEARCH_AD_SMALL_IMAGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
}

// NewAdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageModeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageModeFromValue(v string) (*AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode, error) {
	ev := AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_creative_info_image_materials_image_mode value
func (v AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode) Ptr() *AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode {
	return &v
}
