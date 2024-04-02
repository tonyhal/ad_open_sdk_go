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

// ToolsBidSuggestV2Gender
type ToolsBidSuggestV2Gender string

// List of tools_bid_suggest_v2_gender
const (
	GENDER_UNLIMITED_ToolsBidSuggestV2Gender ToolsBidSuggestV2Gender = "GENDER_UNLIMITED"
	GENDER_FEMALE_ToolsBidSuggestV2Gender    ToolsBidSuggestV2Gender = "GENDER_FEMALE"
	GENDER_MALE_ToolsBidSuggestV2Gender      ToolsBidSuggestV2Gender = "GENDER_MALE"
	NONE_ToolsBidSuggestV2Gender             ToolsBidSuggestV2Gender = "NONE"
)

// All allowed values of ToolsBidSuggestV2Gender enum
var AllowedToolsBidSuggestV2GenderEnumValues = []ToolsBidSuggestV2Gender{
	"GENDER_UNLIMITED",
	"GENDER_FEMALE",
	"GENDER_MALE",
	"NONE",
}

// NewToolsBidSuggestV2GenderFromValue returns a pointer to a valid ToolsBidSuggestV2Gender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2GenderFromValue(v string) (*ToolsBidSuggestV2Gender, error) {
	ev := ToolsBidSuggestV2Gender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2Gender: valid values are %v", v, AllowedToolsBidSuggestV2GenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2Gender) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2GenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_gender value
func (v ToolsBidSuggestV2Gender) Ptr() *ToolsBidSuggestV2Gender {
	return &v
}
