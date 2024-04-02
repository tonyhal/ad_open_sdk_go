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

// CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation
type CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation string

// List of creative_procedural_creative_create_v2_ad_data_supplements_games_orientation
const (
	HORIZONTAL_CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation = "HORIZONTAL"
	VERTICAL_CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation   CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation = "VERTICAL"
)

// All allowed values of CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation enum
var AllowedCreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientationEnumValues = []CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation{
	"HORIZONTAL",
	"VERTICAL",
}

// NewCreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientationFromValue returns a pointer to a valid CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientationFromValue(v string) (*CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation, error) {
	ev := CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation: valid values are %v", v, AllowedCreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_create_v2_ad_data_supplements_games_orientation value
func (v CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation) Ptr() *CreativeProceduralCreativeCreateV2AdDataSupplementsGamesOrientation {
	return &v
}
