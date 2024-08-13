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

// QianchuanAdDetailGetV10DataLabAdType
type QianchuanAdDetailGetV10DataLabAdType string

// List of qianchuan_ad_detail_get_v1.0_data_lab_ad_type
const (
	LAB_AD_QianchuanAdDetailGetV10DataLabAdType     QianchuanAdDetailGetV10DataLabAdType = "LAB_AD"
	NOT_LAB_AD_QianchuanAdDetailGetV10DataLabAdType QianchuanAdDetailGetV10DataLabAdType = "NOT_LAB_AD"
)

// All allowed values of QianchuanAdDetailGetV10DataLabAdType enum
var AllowedQianchuanAdDetailGetV10DataLabAdTypeEnumValues = []QianchuanAdDetailGetV10DataLabAdType{
	"LAB_AD",
	"NOT_LAB_AD",
}

// NewQianchuanAdDetailGetV10DataLabAdTypeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataLabAdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataLabAdTypeFromValue(v string) (*QianchuanAdDetailGetV10DataLabAdType, error) {
	ev := QianchuanAdDetailGetV10DataLabAdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataLabAdType: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataLabAdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataLabAdType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataLabAdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_lab_ad_type value
func (v QianchuanAdDetailGetV10DataLabAdType) Ptr() *QianchuanAdDetailGetV10DataLabAdType {
	return &v
}
