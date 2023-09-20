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

// QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode
type QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode string

// List of qianchuan_ad_update_v1.0_programmatic_creative_media_list_image_mode
const (
	AWEME_LIVE_ROOM_QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode = "AWEME_LIVE_ROOM"
	LARGE_QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode           QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode  QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode           QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode = "SMALL"
	SQUARE_QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode          QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode = "SQUARE"
	UNION_SPLASH_QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode    QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode     QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode  QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode enum
var AllowedQianchuanAdUpdateV10ProgrammaticCreativeMediaListImageModeEnumValues = []QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode{
	"AWEME_LIVE_ROOM",
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanAdUpdateV10ProgrammaticCreativeMediaListImageModeFromValue returns a pointer to a valid QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10ProgrammaticCreativeMediaListImageModeFromValue(v string) (*QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode, error) {
	ev := QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode: valid values are %v", v, AllowedQianchuanAdUpdateV10ProgrammaticCreativeMediaListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10ProgrammaticCreativeMediaListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_programmatic_creative_media_list_image_mode value
func (v QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode) Ptr() *QianchuanAdUpdateV10ProgrammaticCreativeMediaListImageMode {
	return &v
}
