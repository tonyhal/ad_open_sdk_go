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

// QianchuanReportMaterialGetV10FilteringCarouselSource
type QianchuanReportMaterialGetV10FilteringCarouselSource string

// List of qianchuan_report_material_get_v1.0_filtering_carousel_source
const (
	AWEME_QianchuanReportMaterialGetV10FilteringCarouselSource      QianchuanReportMaterialGetV10FilteringCarouselSource = "AWEME"
	E_COMMERCE_QianchuanReportMaterialGetV10FilteringCarouselSource QianchuanReportMaterialGetV10FilteringCarouselSource = "E_COMMERCE"
	JI_CHUANG_QianchuanReportMaterialGetV10FilteringCarouselSource  QianchuanReportMaterialGetV10FilteringCarouselSource = "JI_CHUANG"
)

// All allowed values of QianchuanReportMaterialGetV10FilteringCarouselSource enum
var AllowedQianchuanReportMaterialGetV10FilteringCarouselSourceEnumValues = []QianchuanReportMaterialGetV10FilteringCarouselSource{
	"AWEME",
	"E_COMMERCE",
	"JI_CHUANG",
}

// NewQianchuanReportMaterialGetV10FilteringCarouselSourceFromValue returns a pointer to a valid QianchuanReportMaterialGetV10FilteringCarouselSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportMaterialGetV10FilteringCarouselSourceFromValue(v string) (*QianchuanReportMaterialGetV10FilteringCarouselSource, error) {
	ev := QianchuanReportMaterialGetV10FilteringCarouselSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportMaterialGetV10FilteringCarouselSource: valid values are %v", v, AllowedQianchuanReportMaterialGetV10FilteringCarouselSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportMaterialGetV10FilteringCarouselSource) IsValid() bool {
	for _, existing := range AllowedQianchuanReportMaterialGetV10FilteringCarouselSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_material_get_v1.0_filtering_carousel_source value
func (v QianchuanReportMaterialGetV10FilteringCarouselSource) Ptr() *QianchuanReportMaterialGetV10FilteringCarouselSource {
	return &v
}
