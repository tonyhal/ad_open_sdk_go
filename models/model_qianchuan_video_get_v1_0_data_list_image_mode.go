/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanVideoGetV10DataListImageMode
type QianchuanVideoGetV10DataListImageMode string

// List of qianchuan_video_get_v1.0_data_list_image_mode
const (
	LARGE_QianchuanVideoGetV10DataListImageMode          QianchuanVideoGetV10DataListImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanVideoGetV10DataListImageMode QianchuanVideoGetV10DataListImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanVideoGetV10DataListImageMode          QianchuanVideoGetV10DataListImageMode = "SMALL"
	UNION_SPLASH_QianchuanVideoGetV10DataListImageMode   QianchuanVideoGetV10DataListImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanVideoGetV10DataListImageMode    QianchuanVideoGetV10DataListImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanVideoGetV10DataListImageMode QianchuanVideoGetV10DataListImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanVideoGetV10DataListImageMode enum
var AllowedQianchuanVideoGetV10DataListImageModeEnumValues = []QianchuanVideoGetV10DataListImageMode{
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanVideoGetV10DataListImageModeFromValue returns a pointer to a valid QianchuanVideoGetV10DataListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanVideoGetV10DataListImageModeFromValue(v string) (*QianchuanVideoGetV10DataListImageMode, error) {
	ev := QianchuanVideoGetV10DataListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanVideoGetV10DataListImageMode: valid values are %v", v, AllowedQianchuanVideoGetV10DataListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanVideoGetV10DataListImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanVideoGetV10DataListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_video_get_v1.0_data_list_image_mode value
func (v QianchuanVideoGetV10DataListImageMode) Ptr() *QianchuanVideoGetV10DataListImageMode {
	return &v
}
