/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderDetailGetV10DataAudienceAge
type QianchuanAwemeOrderDetailGetV10DataAudienceAge string

// List of qianchuan_aweme_order_detail_get_v1.0_data_audience_age
const (
	AGE_BETWEEN_18_23_QianchuanAwemeOrderDetailGetV10DataAudienceAge QianchuanAwemeOrderDetailGetV10DataAudienceAge = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_24_30_QianchuanAwemeOrderDetailGetV10DataAudienceAge QianchuanAwemeOrderDetailGetV10DataAudienceAge = "AGE_BETWEEN_24_30"
	AGE_BETWEEN_31_40_QianchuanAwemeOrderDetailGetV10DataAudienceAge QianchuanAwemeOrderDetailGetV10DataAudienceAge = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_41_50_QianchuanAwemeOrderDetailGetV10DataAudienceAge QianchuanAwemeOrderDetailGetV10DataAudienceAge = "AGE_BETWEEN_41_50"
)

// All allowed values of QianchuanAwemeOrderDetailGetV10DataAudienceAge enum
var AllowedQianchuanAwemeOrderDetailGetV10DataAudienceAgeEnumValues = []QianchuanAwemeOrderDetailGetV10DataAudienceAge{
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_24_30",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_41_50",
}

// NewQianchuanAwemeOrderDetailGetV10DataAudienceAgeFromValue returns a pointer to a valid QianchuanAwemeOrderDetailGetV10DataAudienceAge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderDetailGetV10DataAudienceAgeFromValue(v string) (*QianchuanAwemeOrderDetailGetV10DataAudienceAge, error) {
	ev := QianchuanAwemeOrderDetailGetV10DataAudienceAge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderDetailGetV10DataAudienceAge: valid values are %v", v, AllowedQianchuanAwemeOrderDetailGetV10DataAudienceAgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderDetailGetV10DataAudienceAge) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderDetailGetV10DataAudienceAgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_detail_get_v1.0_data_audience_age value
func (v QianchuanAwemeOrderDetailGetV10DataAudienceAge) Ptr() *QianchuanAwemeOrderDetailGetV10DataAudienceAge {
	return &v
}
