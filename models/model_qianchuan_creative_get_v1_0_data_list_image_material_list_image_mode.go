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

// QianchuanCreativeGetV10DataListImageMaterialListImageMode
type QianchuanCreativeGetV10DataListImageMaterialListImageMode string

// List of qianchuan_creative_get_v1.0_data_list_image_material_list_image_mode
const (
	AWEME_LIVE_ROOM_QianchuanCreativeGetV10DataListImageMaterialListImageMode QianchuanCreativeGetV10DataListImageMaterialListImageMode = "AWEME_LIVE_ROOM"
	LARGE_QianchuanCreativeGetV10DataListImageMaterialListImageMode           QianchuanCreativeGetV10DataListImageMaterialListImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanCreativeGetV10DataListImageMaterialListImageMode  QianchuanCreativeGetV10DataListImageMaterialListImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanCreativeGetV10DataListImageMaterialListImageMode           QianchuanCreativeGetV10DataListImageMaterialListImageMode = "SMALL"
	SQUARE_QianchuanCreativeGetV10DataListImageMaterialListImageMode          QianchuanCreativeGetV10DataListImageMaterialListImageMode = "SQUARE"
	UNION_SPLASH_QianchuanCreativeGetV10DataListImageMaterialListImageMode    QianchuanCreativeGetV10DataListImageMaterialListImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanCreativeGetV10DataListImageMaterialListImageMode     QianchuanCreativeGetV10DataListImageMaterialListImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanCreativeGetV10DataListImageMaterialListImageMode  QianchuanCreativeGetV10DataListImageMaterialListImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanCreativeGetV10DataListImageMaterialListImageMode enum
var AllowedQianchuanCreativeGetV10DataListImageMaterialListImageModeEnumValues = []QianchuanCreativeGetV10DataListImageMaterialListImageMode{
	"AWEME_LIVE_ROOM",
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanCreativeGetV10DataListImageMaterialListImageModeFromValue returns a pointer to a valid QianchuanCreativeGetV10DataListImageMaterialListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10DataListImageMaterialListImageModeFromValue(v string) (*QianchuanCreativeGetV10DataListImageMaterialListImageMode, error) {
	ev := QianchuanCreativeGetV10DataListImageMaterialListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10DataListImageMaterialListImageMode: valid values are %v", v, AllowedQianchuanCreativeGetV10DataListImageMaterialListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10DataListImageMaterialListImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10DataListImageMaterialListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_data_list_image_material_list_image_mode value
func (v QianchuanCreativeGetV10DataListImageMaterialListImageMode) Ptr() *QianchuanCreativeGetV10DataListImageMaterialListImageMode {
	return &v
}
