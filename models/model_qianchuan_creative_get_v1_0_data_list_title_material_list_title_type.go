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

// QianchuanCreativeGetV10DataListTitleMaterialListTitleType
type QianchuanCreativeGetV10DataListTitleMaterialListTitleType string

// List of qianchuan_creative_get_v1.0_data_list_title_material_list_title_type
const (
	AWEME_CAROUSEL_QianchuanCreativeGetV10DataListTitleMaterialListTitleType QianchuanCreativeGetV10DataListTitleMaterialListTitleType = "AWEME_CAROUSEL"
	COMMODITY_CARD_QianchuanCreativeGetV10DataListTitleMaterialListTitleType QianchuanCreativeGetV10DataListTitleMaterialListTitleType = "COMMODITY_CARD"
	CUSTOM_QianchuanCreativeGetV10DataListTitleMaterialListTitleType         QianchuanCreativeGetV10DataListTitleMaterialListTitleType = "CUSTOM"
)

// All allowed values of QianchuanCreativeGetV10DataListTitleMaterialListTitleType enum
var AllowedQianchuanCreativeGetV10DataListTitleMaterialListTitleTypeEnumValues = []QianchuanCreativeGetV10DataListTitleMaterialListTitleType{
	"AWEME_CAROUSEL",
	"COMMODITY_CARD",
	"CUSTOM",
}

// NewQianchuanCreativeGetV10DataListTitleMaterialListTitleTypeFromValue returns a pointer to a valid QianchuanCreativeGetV10DataListTitleMaterialListTitleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10DataListTitleMaterialListTitleTypeFromValue(v string) (*QianchuanCreativeGetV10DataListTitleMaterialListTitleType, error) {
	ev := QianchuanCreativeGetV10DataListTitleMaterialListTitleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10DataListTitleMaterialListTitleType: valid values are %v", v, AllowedQianchuanCreativeGetV10DataListTitleMaterialListTitleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10DataListTitleMaterialListTitleType) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10DataListTitleMaterialListTitleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_data_list_title_material_list_title_type value
func (v QianchuanCreativeGetV10DataListTitleMaterialListTitleType) Ptr() *QianchuanCreativeGetV10DataListTitleMaterialListTitleType {
	return &v
}
