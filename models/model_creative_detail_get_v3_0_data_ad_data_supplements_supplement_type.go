/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeDetailGetV30DataAdDataSupplementsSupplementType
type CreativeDetailGetV30DataAdDataSupplementsSupplementType string

// List of creative_detail_get_v3.0_data_ad_data_supplements_supplement_type
const (
	CLOUD_GAME_CreativeDetailGetV30DataAdDataSupplementsSupplementType CreativeDetailGetV30DataAdDataSupplementsSupplementType = "CLOUD_GAME"
	NORMAL_CreativeDetailGetV30DataAdDataSupplementsSupplementType     CreativeDetailGetV30DataAdDataSupplementsSupplementType = "NORMAL"
)

// All allowed values of CreativeDetailGetV30DataAdDataSupplementsSupplementType enum
var AllowedCreativeDetailGetV30DataAdDataSupplementsSupplementTypeEnumValues = []CreativeDetailGetV30DataAdDataSupplementsSupplementType{
	"CLOUD_GAME",
	"NORMAL",
}

// NewCreativeDetailGetV30DataAdDataSupplementsSupplementTypeFromValue returns a pointer to a valid CreativeDetailGetV30DataAdDataSupplementsSupplementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeDetailGetV30DataAdDataSupplementsSupplementTypeFromValue(v string) (*CreativeDetailGetV30DataAdDataSupplementsSupplementType, error) {
	ev := CreativeDetailGetV30DataAdDataSupplementsSupplementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeDetailGetV30DataAdDataSupplementsSupplementType: valid values are %v", v, AllowedCreativeDetailGetV30DataAdDataSupplementsSupplementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeDetailGetV30DataAdDataSupplementsSupplementType) IsValid() bool {
	for _, existing := range AllowedCreativeDetailGetV30DataAdDataSupplementsSupplementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_detail_get_v3.0_data_ad_data_supplements_supplement_type value
func (v CreativeDetailGetV30DataAdDataSupplementsSupplementType) Ptr() *CreativeDetailGetV30DataAdDataSupplementsSupplementType {
	return &v
}
