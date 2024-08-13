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

// QianchuanAdMaterialGetV10FilteringVideoSource
type QianchuanAdMaterialGetV10FilteringVideoSource string

// List of qianchuan_ad_material_get_v1.0_filtering_video_source
const (
	ARTHUR_QianchuanAdMaterialGetV10FilteringVideoSource          QianchuanAdMaterialGetV10FilteringVideoSource = "ARTHUR"
	AWEME_QianchuanAdMaterialGetV10FilteringVideoSource           QianchuanAdMaterialGetV10FilteringVideoSource = "AWEME"
	BP_QianchuanAdMaterialGetV10FilteringVideoSource              QianchuanAdMaterialGetV10FilteringVideoSource = "BP"
	CREATIVE_CENTER_QianchuanAdMaterialGetV10FilteringVideoSource QianchuanAdMaterialGetV10FilteringVideoSource = "CREATIVE_CENTER"
	E_COMMERCE_QianchuanAdMaterialGetV10FilteringVideoSource      QianchuanAdMaterialGetV10FilteringVideoSource = "E_COMMERCE"
	JIANYING_QianchuanAdMaterialGetV10FilteringVideoSource        QianchuanAdMaterialGetV10FilteringVideoSource = "JIANYING"
	JI_CHUANG_QianchuanAdMaterialGetV10FilteringVideoSource       QianchuanAdMaterialGetV10FilteringVideoSource = "JI_CHUANG"
	LIVE_HIGHLIGHT_QianchuanAdMaterialGetV10FilteringVideoSource  QianchuanAdMaterialGetV10FilteringVideoSource = "LIVE_HIGHLIGHT"
	STAR_QianchuanAdMaterialGetV10FilteringVideoSource            QianchuanAdMaterialGetV10FilteringVideoSource = "STAR"
	TADA_QianchuanAdMaterialGetV10FilteringVideoSource            QianchuanAdMaterialGetV10FilteringVideoSource = "TADA"
	VIDEO_CAPTURE_QianchuanAdMaterialGetV10FilteringVideoSource   QianchuanAdMaterialGetV10FilteringVideoSource = "VIDEO_CAPTURE"
)

// All allowed values of QianchuanAdMaterialGetV10FilteringVideoSource enum
var AllowedQianchuanAdMaterialGetV10FilteringVideoSourceEnumValues = []QianchuanAdMaterialGetV10FilteringVideoSource{
	"ARTHUR",
	"AWEME",
	"BP",
	"CREATIVE_CENTER",
	"E_COMMERCE",
	"JIANYING",
	"JI_CHUANG",
	"LIVE_HIGHLIGHT",
	"STAR",
	"TADA",
	"VIDEO_CAPTURE",
}

// NewQianchuanAdMaterialGetV10FilteringVideoSourceFromValue returns a pointer to a valid QianchuanAdMaterialGetV10FilteringVideoSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdMaterialGetV10FilteringVideoSourceFromValue(v string) (*QianchuanAdMaterialGetV10FilteringVideoSource, error) {
	ev := QianchuanAdMaterialGetV10FilteringVideoSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdMaterialGetV10FilteringVideoSource: valid values are %v", v, AllowedQianchuanAdMaterialGetV10FilteringVideoSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdMaterialGetV10FilteringVideoSource) IsValid() bool {
	for _, existing := range AllowedQianchuanAdMaterialGetV10FilteringVideoSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_material_get_v1.0_filtering_video_source value
func (v QianchuanAdMaterialGetV10FilteringVideoSource) Ptr() *QianchuanAdMaterialGetV10FilteringVideoSource {
	return &v
}
