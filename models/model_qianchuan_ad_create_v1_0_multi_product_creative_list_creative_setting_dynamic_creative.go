/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative
type QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative int64

// List of qianchuan_ad_create_v1.0_multi_product_creative_list_creative_setting_dynamic_creative
const (
	Enum_0_QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative = 0
	Enum_1_QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative = 1
)

// All allowed values of QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative enum
var AllowedQianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreativeEnumValues = []QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative{
	0,
	1,
}

// NewQianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreativeFromValue returns a pointer to a valid QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreativeFromValue(v int64) (*QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative, error) {
	ev := QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative: valid values are %v", v, AllowedQianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreativeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreativeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_multi_product_creative_list_creative_setting_dynamic_creative value
func (v QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative) Ptr() *QianchuanAdCreateV10MultiProductCreativeListCreativeSettingDynamicCreative {
	return &v
}
