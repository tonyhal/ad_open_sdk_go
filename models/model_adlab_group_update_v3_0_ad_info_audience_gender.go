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

// AdlabGroupUpdateV30AdInfoAudienceGender
type AdlabGroupUpdateV30AdInfoAudienceGender string

// List of adlab_group_update_v3.0_ad_info_audience_gender
const (
	GENDER_FEMALE_AdlabGroupUpdateV30AdInfoAudienceGender AdlabGroupUpdateV30AdInfoAudienceGender = "GENDER_FEMALE"
	GENDER_MALE_AdlabGroupUpdateV30AdInfoAudienceGender   AdlabGroupUpdateV30AdInfoAudienceGender = "GENDER_MALE"
	NONE_AdlabGroupUpdateV30AdInfoAudienceGender          AdlabGroupUpdateV30AdInfoAudienceGender = "NONE"
)

// All allowed values of AdlabGroupUpdateV30AdInfoAudienceGender enum
var AllowedAdlabGroupUpdateV30AdInfoAudienceGenderEnumValues = []AdlabGroupUpdateV30AdInfoAudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
	"NONE",
}

// NewAdlabGroupUpdateV30AdInfoAudienceGenderFromValue returns a pointer to a valid AdlabGroupUpdateV30AdInfoAudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupUpdateV30AdInfoAudienceGenderFromValue(v string) (*AdlabGroupUpdateV30AdInfoAudienceGender, error) {
	ev := AdlabGroupUpdateV30AdInfoAudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupUpdateV30AdInfoAudienceGender: valid values are %v", v, AllowedAdlabGroupUpdateV30AdInfoAudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupUpdateV30AdInfoAudienceGender) IsValid() bool {
	for _, existing := range AllowedAdlabGroupUpdateV30AdInfoAudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_update_v3.0_ad_info_audience_gender value
func (v AdlabGroupUpdateV30AdInfoAudienceGender) Ptr() *AdlabGroupUpdateV30AdInfoAudienceGender {
	return &v
}
