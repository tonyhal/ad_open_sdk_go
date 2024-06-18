/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarVasCreateBoostItemGroupV2CustomAudienceTagGender
type StarVasCreateBoostItemGroupV2CustomAudienceTagGender string

// List of star_vas_create_boost_item_group_v2_custom_audience_tag_gender
const (
	FEMALE_StarVasCreateBoostItemGroupV2CustomAudienceTagGender   StarVasCreateBoostItemGroupV2CustomAudienceTagGender = "FEMALE"
	INFINITE_StarVasCreateBoostItemGroupV2CustomAudienceTagGender StarVasCreateBoostItemGroupV2CustomAudienceTagGender = "INFINITE"
	MALE_StarVasCreateBoostItemGroupV2CustomAudienceTagGender     StarVasCreateBoostItemGroupV2CustomAudienceTagGender = "MALE"
)

// All allowed values of StarVasCreateBoostItemGroupV2CustomAudienceTagGender enum
var AllowedStarVasCreateBoostItemGroupV2CustomAudienceTagGenderEnumValues = []StarVasCreateBoostItemGroupV2CustomAudienceTagGender{
	"FEMALE",
	"INFINITE",
	"MALE",
}

// NewStarVasCreateBoostItemGroupV2CustomAudienceTagGenderFromValue returns a pointer to a valid StarVasCreateBoostItemGroupV2CustomAudienceTagGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarVasCreateBoostItemGroupV2CustomAudienceTagGenderFromValue(v string) (*StarVasCreateBoostItemGroupV2CustomAudienceTagGender, error) {
	ev := StarVasCreateBoostItemGroupV2CustomAudienceTagGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarVasCreateBoostItemGroupV2CustomAudienceTagGender: valid values are %v", v, AllowedStarVasCreateBoostItemGroupV2CustomAudienceTagGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarVasCreateBoostItemGroupV2CustomAudienceTagGender) IsValid() bool {
	for _, existing := range AllowedStarVasCreateBoostItemGroupV2CustomAudienceTagGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_vas_create_boost_item_group_v2_custom_audience_tag_gender value
func (v StarVasCreateBoostItemGroupV2CustomAudienceTagGender) Ptr() *StarVasCreateBoostItemGroupV2CustomAudienceTagGender {
	return &v
}