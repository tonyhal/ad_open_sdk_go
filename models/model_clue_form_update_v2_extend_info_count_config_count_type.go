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

// ClueFormUpdateV2ExtendInfoCountConfigCountType
type ClueFormUpdateV2ExtendInfoCountConfigCountType string

// List of clue_form_update_v2_extend_info_count_config_count_type
const (
	COUNT_TYPE_INCREMENT_ClueFormUpdateV2ExtendInfoCountConfigCountType ClueFormUpdateV2ExtendInfoCountConfigCountType = "COUNT_TYPE_INCREMENT"
	COUNT_TYPE_DECREMENT_ClueFormUpdateV2ExtendInfoCountConfigCountType ClueFormUpdateV2ExtendInfoCountConfigCountType = "COUNT_TYPE_DECREMENT"
)

// All allowed values of ClueFormUpdateV2ExtendInfoCountConfigCountType enum
var AllowedClueFormUpdateV2ExtendInfoCountConfigCountTypeEnumValues = []ClueFormUpdateV2ExtendInfoCountConfigCountType{
	"COUNT_TYPE_INCREMENT",
	"COUNT_TYPE_DECREMENT",
}

// NewClueFormUpdateV2ExtendInfoCountConfigCountTypeFromValue returns a pointer to a valid ClueFormUpdateV2ExtendInfoCountConfigCountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormUpdateV2ExtendInfoCountConfigCountTypeFromValue(v string) (*ClueFormUpdateV2ExtendInfoCountConfigCountType, error) {
	ev := ClueFormUpdateV2ExtendInfoCountConfigCountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormUpdateV2ExtendInfoCountConfigCountType: valid values are %v", v, AllowedClueFormUpdateV2ExtendInfoCountConfigCountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormUpdateV2ExtendInfoCountConfigCountType) IsValid() bool {
	for _, existing := range AllowedClueFormUpdateV2ExtendInfoCountConfigCountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_update_v2_extend_info_count_config_count_type value
func (v ClueFormUpdateV2ExtendInfoCountConfigCountType) Ptr() *ClueFormUpdateV2ExtendInfoCountConfigCountType {
	return &v
}
