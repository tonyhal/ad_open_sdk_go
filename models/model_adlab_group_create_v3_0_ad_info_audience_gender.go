/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupCreateV30AdInfoAudienceGender
type AdlabGroupCreateV30AdInfoAudienceGender string

// List of adlab_group_create_v3.0_ad_info_audience_gender
const (
	GENDER_FEMALE_AdlabGroupCreateV30AdInfoAudienceGender AdlabGroupCreateV30AdInfoAudienceGender = "GENDER_FEMALE"
	GENDER_MALE_AdlabGroupCreateV30AdInfoAudienceGender   AdlabGroupCreateV30AdInfoAudienceGender = "GENDER_MALE"
	NONE_AdlabGroupCreateV30AdInfoAudienceGender          AdlabGroupCreateV30AdInfoAudienceGender = "NONE"
)

// All allowed values of AdlabGroupCreateV30AdInfoAudienceGender enum
var AllowedAdlabGroupCreateV30AdInfoAudienceGenderEnumValues = []AdlabGroupCreateV30AdInfoAudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
	"NONE",
}

// NewAdlabGroupCreateV30AdInfoAudienceGenderFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoAudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoAudienceGenderFromValue(v string) (*AdlabGroupCreateV30AdInfoAudienceGender, error) {
	ev := AdlabGroupCreateV30AdInfoAudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoAudienceGender: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoAudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoAudienceGender) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoAudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_audience_gender value
func (v AdlabGroupCreateV30AdInfoAudienceGender) Ptr() *AdlabGroupCreateV30AdInfoAudienceGender {
	return &v
}
