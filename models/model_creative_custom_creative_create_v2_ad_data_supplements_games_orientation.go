/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation
type CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation string

// List of creative_custom_creative_create_v2_ad_data_supplements_games_orientation
const (
	HORIZONTAL_CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation = "HORIZONTAL"
	VERTICAL_CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation   CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation = "VERTICAL"
)

// All allowed values of CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation enum
var AllowedCreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientationEnumValues = []CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation{
	"HORIZONTAL",
	"VERTICAL",
}

// NewCreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientationFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientationFromValue(v string) (*CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation, error) {
	ev := CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_ad_data_supplements_games_orientation value
func (v CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation) Ptr() *CreativeCustomCreativeCreateV2AdDataSupplementsGamesOrientation {
	return &v
}
