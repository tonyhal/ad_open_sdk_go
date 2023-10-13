/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormDetailV2DataFormFormType
type ClueFormDetailV2DataFormFormType string

// List of clue_form_detail_v2_data_form_form_type
const (
	NATIVE_FORM_ClueFormDetailV2DataFormFormType            ClueFormDetailV2DataFormFormType = "NATIVE_FORM"
	NORMAL_FORM_ClueFormDetailV2DataFormFormType            ClueFormDetailV2DataFormFormType = "NORMAL_FORM"
	ADVANCED_CREATIVE_FORM_ClueFormDetailV2DataFormFormType ClueFormDetailV2DataFormFormType = "ADVANCED_CREATIVE_FORM"
)

// All allowed values of ClueFormDetailV2DataFormFormType enum
var AllowedClueFormDetailV2DataFormFormTypeEnumValues = []ClueFormDetailV2DataFormFormType{
	"NATIVE_FORM",
	"NORMAL_FORM",
	"ADVANCED_CREATIVE_FORM",
}

// NewClueFormDetailV2DataFormFormTypeFromValue returns a pointer to a valid ClueFormDetailV2DataFormFormType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormDetailV2DataFormFormTypeFromValue(v string) (*ClueFormDetailV2DataFormFormType, error) {
	ev := ClueFormDetailV2DataFormFormType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormDetailV2DataFormFormType: valid values are %v", v, AllowedClueFormDetailV2DataFormFormTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormDetailV2DataFormFormType) IsValid() bool {
	for _, existing := range AllowedClueFormDetailV2DataFormFormTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_detail_v2_data_form_form_type value
func (v ClueFormDetailV2DataFormFormType) Ptr() *ClueFormDetailV2DataFormFormType {
	return &v
}
