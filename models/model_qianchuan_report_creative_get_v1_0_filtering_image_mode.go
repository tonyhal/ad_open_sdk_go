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

// QianchuanReportCreativeGetV10FilteringImageMode
type QianchuanReportCreativeGetV10FilteringImageMode string

// List of qianchuan_report_creative_get_v1.0_filtering_image_mode
const (
	LARGE_QianchuanReportCreativeGetV10FilteringImageMode          QianchuanReportCreativeGetV10FilteringImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanReportCreativeGetV10FilteringImageMode QianchuanReportCreativeGetV10FilteringImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanReportCreativeGetV10FilteringImageMode          QianchuanReportCreativeGetV10FilteringImageMode = "SMALL"
	SQUARE_QianchuanReportCreativeGetV10FilteringImageMode         QianchuanReportCreativeGetV10FilteringImageMode = "SQUARE"
	UNION_SPLASH_QianchuanReportCreativeGetV10FilteringImageMode   QianchuanReportCreativeGetV10FilteringImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanReportCreativeGetV10FilteringImageMode    QianchuanReportCreativeGetV10FilteringImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanReportCreativeGetV10FilteringImageMode QianchuanReportCreativeGetV10FilteringImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanReportCreativeGetV10FilteringImageMode enum
var AllowedQianchuanReportCreativeGetV10FilteringImageModeEnumValues = []QianchuanReportCreativeGetV10FilteringImageMode{
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanReportCreativeGetV10FilteringImageModeFromValue returns a pointer to a valid QianchuanReportCreativeGetV10FilteringImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportCreativeGetV10FilteringImageModeFromValue(v string) (*QianchuanReportCreativeGetV10FilteringImageMode, error) {
	ev := QianchuanReportCreativeGetV10FilteringImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportCreativeGetV10FilteringImageMode: valid values are %v", v, AllowedQianchuanReportCreativeGetV10FilteringImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportCreativeGetV10FilteringImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanReportCreativeGetV10FilteringImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_creative_get_v1.0_filtering_image_mode value
func (v QianchuanReportCreativeGetV10FilteringImageMode) Ptr() *QianchuanReportCreativeGetV10FilteringImageMode {
	return &v
}
