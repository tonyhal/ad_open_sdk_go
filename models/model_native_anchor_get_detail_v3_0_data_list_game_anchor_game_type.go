/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorGetDetailV30DataListGameAnchorGameType
type NativeAnchorGetDetailV30DataListGameAnchorGameType string

// List of native_anchor_get_detail_v3.0_data_list_game_anchor_game_type
const (
	GENERAL_NativeAnchorGetDetailV30DataListGameAnchorGameType    NativeAnchorGetDetailV30DataListGameAnchorGameType = "GENERAL"
	MICRO_GAME_NativeAnchorGetDetailV30DataListGameAnchorGameType NativeAnchorGetDetailV30DataListGameAnchorGameType = "MICRO_GAME"
)

// All allowed values of NativeAnchorGetDetailV30DataListGameAnchorGameType enum
var AllowedNativeAnchorGetDetailV30DataListGameAnchorGameTypeEnumValues = []NativeAnchorGetDetailV30DataListGameAnchorGameType{
	"GENERAL",
	"MICRO_GAME",
}

// NewNativeAnchorGetDetailV30DataListGameAnchorGameTypeFromValue returns a pointer to a valid NativeAnchorGetDetailV30DataListGameAnchorGameType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetDetailV30DataListGameAnchorGameTypeFromValue(v string) (*NativeAnchorGetDetailV30DataListGameAnchorGameType, error) {
	ev := NativeAnchorGetDetailV30DataListGameAnchorGameType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetDetailV30DataListGameAnchorGameType: valid values are %v", v, AllowedNativeAnchorGetDetailV30DataListGameAnchorGameTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetDetailV30DataListGameAnchorGameType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetDetailV30DataListGameAnchorGameTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_detail_v3.0_data_list_game_anchor_game_type value
func (v NativeAnchorGetDetailV30DataListGameAnchorGameType) Ptr() *NativeAnchorGetDetailV30DataListGameAnchorGameType {
	return &v
}
