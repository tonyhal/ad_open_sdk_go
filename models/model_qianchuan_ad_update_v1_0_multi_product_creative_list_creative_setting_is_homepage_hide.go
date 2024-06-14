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

// QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide
type QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide int64

// List of qianchuan_ad_update_v1.0_multi_product_creative_list_creative_setting_is_homepage_hide
const (
	Enum_0_QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide = 0
	Enum_1_QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide = 1
)

// All allowed values of QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide enum
var AllowedQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHideEnumValues = []QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide{
	0,
	1,
}

// NewQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHideFromValue returns a pointer to a valid QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHideFromValue(v int64) (*QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide, error) {
	ev := QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide: valid values are %v", v, AllowedQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHideEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_multi_product_creative_list_creative_setting_is_homepage_hide value
func (v QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide) Ptr() *QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide {
	return &v
}
