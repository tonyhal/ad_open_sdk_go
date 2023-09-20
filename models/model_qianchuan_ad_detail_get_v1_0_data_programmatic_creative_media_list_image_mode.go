/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode
type QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode string

// List of qianchuan_ad_detail_get_v1.0_data_programmatic_creative_media_list_image_mode
const (
	AWEME_LIVE_ROOM_QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode = "AWEME_LIVE_ROOM"
	LARGE_QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode           QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode  QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode           QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode = "SMALL"
	SQUARE_QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode          QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode = "SQUARE"
	UNION_SPLASH_QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode    QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode     QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode  QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode enum
var AllowedQianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageModeEnumValues = []QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode{
	"AWEME_LIVE_ROOM",
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageModeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageModeFromValue(v string) (*QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode, error) {
	ev := QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_programmatic_creative_media_list_image_mode value
func (v QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode) Ptr() *QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode {
	return &v
}
