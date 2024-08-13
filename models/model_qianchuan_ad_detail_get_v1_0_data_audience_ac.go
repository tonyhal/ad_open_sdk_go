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

// QianchuanAdDetailGetV10DataAudienceAc
type QianchuanAdDetailGetV10DataAudienceAc string

// List of qianchuan_ad_detail_get_v1.0_data_audience_ac
const (
	Enum_2_G_QianchuanAdDetailGetV10DataAudienceAc QianchuanAdDetailGetV10DataAudienceAc = "2G"
	Enum_3_G_QianchuanAdDetailGetV10DataAudienceAc QianchuanAdDetailGetV10DataAudienceAc = "3G"
	Enum_4_G_QianchuanAdDetailGetV10DataAudienceAc QianchuanAdDetailGetV10DataAudienceAc = "4G"
	Enum_5_G_QianchuanAdDetailGetV10DataAudienceAc QianchuanAdDetailGetV10DataAudienceAc = "5G"
	WIFI_QianchuanAdDetailGetV10DataAudienceAc     QianchuanAdDetailGetV10DataAudienceAc = "WIFI"
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceAc enum
var AllowedQianchuanAdDetailGetV10DataAudienceAcEnumValues = []QianchuanAdDetailGetV10DataAudienceAc{
	"2G",
	"3G",
	"4G",
	"5G",
	"WIFI",
}

// NewQianchuanAdDetailGetV10DataAudienceAcFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceAcFromValue(v string) (*QianchuanAdDetailGetV10DataAudienceAc, error) {
	ev := QianchuanAdDetailGetV10DataAudienceAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceAc: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceAc) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_ac value
func (v QianchuanAdDetailGetV10DataAudienceAc) Ptr() *QianchuanAdDetailGetV10DataAudienceAc {
	return &v
}
