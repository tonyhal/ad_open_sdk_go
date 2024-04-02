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

// QianchuanAdUpdateV10AudienceGender
type QianchuanAdUpdateV10AudienceGender string

// List of qianchuan_ad_update_v1.0_audience_gender
const (
	GENDER_FEMALE_QianchuanAdUpdateV10AudienceGender QianchuanAdUpdateV10AudienceGender = "GENDER_FEMALE"
	GENDER_MALE_QianchuanAdUpdateV10AudienceGender   QianchuanAdUpdateV10AudienceGender = "GENDER_MALE"
	NONE_QianchuanAdUpdateV10AudienceGender          QianchuanAdUpdateV10AudienceGender = "NONE"
)

// All allowed values of QianchuanAdUpdateV10AudienceGender enum
var AllowedQianchuanAdUpdateV10AudienceGenderEnumValues = []QianchuanAdUpdateV10AudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
	"NONE",
}

// NewQianchuanAdUpdateV10AudienceGenderFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceGenderFromValue(v string) (*QianchuanAdUpdateV10AudienceGender, error) {
	ev := QianchuanAdUpdateV10AudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceGender: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceGender) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_gender value
func (v QianchuanAdUpdateV10AudienceGender) Ptr() *QianchuanAdUpdateV10AudienceGender {
	return &v
}
