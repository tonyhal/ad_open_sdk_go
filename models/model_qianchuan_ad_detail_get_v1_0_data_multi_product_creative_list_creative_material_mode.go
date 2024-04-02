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

// QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode
type QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode string

// List of qianchuan_ad_detail_get_v1.0_data_multi_product_creative_list_creative_material_mode
const (
	CUSTOM_CREATIVE_QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode       QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode = "CUSTOM_CREATIVE"
	PROGRAMMATIC_CREATIVE_QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode = "PROGRAMMATIC_CREATIVE"
)

// All allowed values of QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode enum
var AllowedQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialModeEnumValues = []QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode{
	"CUSTOM_CREATIVE",
	"PROGRAMMATIC_CREATIVE",
}

// NewQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialModeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialModeFromValue(v string) (*QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode, error) {
	ev := QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_multi_product_creative_list_creative_material_mode value
func (v QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode) Ptr() *QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode {
	return &v
}
