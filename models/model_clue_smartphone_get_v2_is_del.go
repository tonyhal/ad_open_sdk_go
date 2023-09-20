/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueSmartphoneGetV2IsDel
type ClueSmartphoneGetV2IsDel string

// List of clue_smartphone_get_v2_is_del
const (
	Enum_0_ClueSmartphoneGetV2IsDel ClueSmartphoneGetV2IsDel = "0"
	Enum_1_ClueSmartphoneGetV2IsDel ClueSmartphoneGetV2IsDel = "1"
)

// All allowed values of ClueSmartphoneGetV2IsDel enum
var AllowedClueSmartphoneGetV2IsDelEnumValues = []ClueSmartphoneGetV2IsDel{
	"0",
	"1",
}

// NewClueSmartphoneGetV2IsDelFromValue returns a pointer to a valid ClueSmartphoneGetV2IsDel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueSmartphoneGetV2IsDelFromValue(v string) (*ClueSmartphoneGetV2IsDel, error) {
	ev := ClueSmartphoneGetV2IsDel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueSmartphoneGetV2IsDel: valid values are %v", v, AllowedClueSmartphoneGetV2IsDelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueSmartphoneGetV2IsDel) IsValid() bool {
	for _, existing := range AllowedClueSmartphoneGetV2IsDelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_smartphone_get_v2_is_del value
func (v ClueSmartphoneGetV2IsDel) Ptr() *ClueSmartphoneGetV2IsDel {
	return &v
}
