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

// QianchuanAdDetailGetV10DataIsIntelligent
type QianchuanAdDetailGetV10DataIsIntelligent int64

// List of qianchuan_ad_detail_get_v1.0_data_is_intelligent
const (
	Enum_0_QianchuanAdDetailGetV10DataIsIntelligent QianchuanAdDetailGetV10DataIsIntelligent = 0
	Enum_1_QianchuanAdDetailGetV10DataIsIntelligent QianchuanAdDetailGetV10DataIsIntelligent = 1
)

// All allowed values of QianchuanAdDetailGetV10DataIsIntelligent enum
var AllowedQianchuanAdDetailGetV10DataIsIntelligentEnumValues = []QianchuanAdDetailGetV10DataIsIntelligent{
	0,
	1,
}

// NewQianchuanAdDetailGetV10DataIsIntelligentFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataIsIntelligent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataIsIntelligentFromValue(v int64) (*QianchuanAdDetailGetV10DataIsIntelligent, error) {
	ev := QianchuanAdDetailGetV10DataIsIntelligent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataIsIntelligent: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataIsIntelligentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataIsIntelligent) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataIsIntelligentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_is_intelligent value
func (v QianchuanAdDetailGetV10DataIsIntelligent) Ptr() *QianchuanAdDetailGetV10DataIsIntelligent {
	return &v
}
