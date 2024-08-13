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

// QianchuanAdCreateV10CreativeListImageMode
type QianchuanAdCreateV10CreativeListImageMode string

// List of qianchuan_ad_create_v1.0_creative_list_image_mode
const (
	AWEME_LIVE_ROOM_QianchuanAdCreateV10CreativeListImageMode QianchuanAdCreateV10CreativeListImageMode = "AWEME_LIVE_ROOM"
	CAROUSEL_QianchuanAdCreateV10CreativeListImageMode        QianchuanAdCreateV10CreativeListImageMode = "CAROUSEL"
	LARGE_QianchuanAdCreateV10CreativeListImageMode           QianchuanAdCreateV10CreativeListImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanAdCreateV10CreativeListImageMode  QianchuanAdCreateV10CreativeListImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanAdCreateV10CreativeListImageMode           QianchuanAdCreateV10CreativeListImageMode = "SMALL"
	SQUARE_QianchuanAdCreateV10CreativeListImageMode          QianchuanAdCreateV10CreativeListImageMode = "SQUARE"
	UNION_SPLASH_QianchuanAdCreateV10CreativeListImageMode    QianchuanAdCreateV10CreativeListImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanAdCreateV10CreativeListImageMode     QianchuanAdCreateV10CreativeListImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanAdCreateV10CreativeListImageMode  QianchuanAdCreateV10CreativeListImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanAdCreateV10CreativeListImageMode enum
var AllowedQianchuanAdCreateV10CreativeListImageModeEnumValues = []QianchuanAdCreateV10CreativeListImageMode{
	"AWEME_LIVE_ROOM",
	"CAROUSEL",
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanAdCreateV10CreativeListImageModeFromValue returns a pointer to a valid QianchuanAdCreateV10CreativeListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10CreativeListImageModeFromValue(v string) (*QianchuanAdCreateV10CreativeListImageMode, error) {
	ev := QianchuanAdCreateV10CreativeListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10CreativeListImageMode: valid values are %v", v, AllowedQianchuanAdCreateV10CreativeListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10CreativeListImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10CreativeListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_creative_list_image_mode value
func (v QianchuanAdCreateV10CreativeListImageMode) Ptr() *QianchuanAdCreateV10CreativeListImageMode {
	return &v
}
