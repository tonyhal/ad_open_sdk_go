/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanImageGetV10DataListSource
type QianchuanImageGetV10DataListSource string

// List of qianchuan_image_get_v1.0_data_list_source
const (
	CREATIVE_CENTER_QianchuanImageGetV10DataListSource QianchuanImageGetV10DataListSource = "CREATIVE_CENTER"
	E_COMMERCE_QianchuanImageGetV10DataListSource      QianchuanImageGetV10DataListSource = "E_COMMERCE"
	JI_CHUANG_QianchuanImageGetV10DataListSource       QianchuanImageGetV10DataListSource = "JI_CHUANG"
)

// All allowed values of QianchuanImageGetV10DataListSource enum
var AllowedQianchuanImageGetV10DataListSourceEnumValues = []QianchuanImageGetV10DataListSource{
	"CREATIVE_CENTER",
	"E_COMMERCE",
	"JI_CHUANG",
}

// NewQianchuanImageGetV10DataListSourceFromValue returns a pointer to a valid QianchuanImageGetV10DataListSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanImageGetV10DataListSourceFromValue(v string) (*QianchuanImageGetV10DataListSource, error) {
	ev := QianchuanImageGetV10DataListSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanImageGetV10DataListSource: valid values are %v", v, AllowedQianchuanImageGetV10DataListSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanImageGetV10DataListSource) IsValid() bool {
	for _, existing := range AllowedQianchuanImageGetV10DataListSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_image_get_v1.0_data_list_source value
func (v QianchuanImageGetV10DataListSource) Ptr() *QianchuanImageGetV10DataListSource {
	return &v
}
