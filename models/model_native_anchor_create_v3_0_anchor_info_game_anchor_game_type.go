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

// NativeAnchorCreateV30AnchorInfoGameAnchorGameType
type NativeAnchorCreateV30AnchorInfoGameAnchorGameType string

// List of native_anchor_create_v3.0_anchor_info_game_anchor_game_type
const (
	GENERAL_NativeAnchorCreateV30AnchorInfoGameAnchorGameType    NativeAnchorCreateV30AnchorInfoGameAnchorGameType = "GENERAL"
	MICRO_GAME_NativeAnchorCreateV30AnchorInfoGameAnchorGameType NativeAnchorCreateV30AnchorInfoGameAnchorGameType = "MICRO_GAME"
)

// All allowed values of NativeAnchorCreateV30AnchorInfoGameAnchorGameType enum
var AllowedNativeAnchorCreateV30AnchorInfoGameAnchorGameTypeEnumValues = []NativeAnchorCreateV30AnchorInfoGameAnchorGameType{
	"GENERAL",
	"MICRO_GAME",
}

// NewNativeAnchorCreateV30AnchorInfoGameAnchorGameTypeFromValue returns a pointer to a valid NativeAnchorCreateV30AnchorInfoGameAnchorGameType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorCreateV30AnchorInfoGameAnchorGameTypeFromValue(v string) (*NativeAnchorCreateV30AnchorInfoGameAnchorGameType, error) {
	ev := NativeAnchorCreateV30AnchorInfoGameAnchorGameType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorCreateV30AnchorInfoGameAnchorGameType: valid values are %v", v, AllowedNativeAnchorCreateV30AnchorInfoGameAnchorGameTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorCreateV30AnchorInfoGameAnchorGameType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorCreateV30AnchorInfoGameAnchorGameTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_create_v3.0_anchor_info_game_anchor_game_type value
func (v NativeAnchorCreateV30AnchorInfoGameAnchorGameType) Ptr() *NativeAnchorCreateV30AnchorInfoGameAnchorGameType {
	return &v
}
