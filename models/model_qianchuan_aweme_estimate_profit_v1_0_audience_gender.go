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

// QianchuanAwemeEstimateProfitV10AudienceGender
type QianchuanAwemeEstimateProfitV10AudienceGender string

// List of qianchuan_aweme_estimate_profit_v1.0_audience_gender
const (
	GENDER_FEMALE_QianchuanAwemeEstimateProfitV10AudienceGender QianchuanAwemeEstimateProfitV10AudienceGender = "GENDER_FEMALE"
	GENDER_MALE_QianchuanAwemeEstimateProfitV10AudienceGender   QianchuanAwemeEstimateProfitV10AudienceGender = "GENDER_MALE"
)

// All allowed values of QianchuanAwemeEstimateProfitV10AudienceGender enum
var AllowedQianchuanAwemeEstimateProfitV10AudienceGenderEnumValues = []QianchuanAwemeEstimateProfitV10AudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
}

// NewQianchuanAwemeEstimateProfitV10AudienceGenderFromValue returns a pointer to a valid QianchuanAwemeEstimateProfitV10AudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeEstimateProfitV10AudienceGenderFromValue(v string) (*QianchuanAwemeEstimateProfitV10AudienceGender, error) {
	ev := QianchuanAwemeEstimateProfitV10AudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeEstimateProfitV10AudienceGender: valid values are %v", v, AllowedQianchuanAwemeEstimateProfitV10AudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeEstimateProfitV10AudienceGender) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeEstimateProfitV10AudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_estimate_profit_v1.0_audience_gender value
func (v QianchuanAwemeEstimateProfitV10AudienceGender) Ptr() *QianchuanAwemeEstimateProfitV10AudienceGender {
	return &v
}
