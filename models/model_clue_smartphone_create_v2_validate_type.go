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

// ClueSmartphoneCreateV2ValidateType
type ClueSmartphoneCreateV2ValidateType string

// List of clue_smartphone_create_v2_validate_type
const (
	NONE_VERIFICATION_ClueSmartphoneCreateV2ValidateType ClueSmartphoneCreateV2ValidateType = "NONE_VERIFICATION"
	AUTO_VERIFICATION_ClueSmartphoneCreateV2ValidateType ClueSmartphoneCreateV2ValidateType = "AUTO_VERIFICATION"
	ALL_VERIFICATION_ClueSmartphoneCreateV2ValidateType  ClueSmartphoneCreateV2ValidateType = "ALL_VERIFICATION"
	CLUE_PRIORITY_ClueSmartphoneCreateV2ValidateType     ClueSmartphoneCreateV2ValidateType = "CLUE_PRIORITY"
	VALIDITY_PRIORITY_ClueSmartphoneCreateV2ValidateType ClueSmartphoneCreateV2ValidateType = "VALIDITY_PRIORITY"
)

// All allowed values of ClueSmartphoneCreateV2ValidateType enum
var AllowedClueSmartphoneCreateV2ValidateTypeEnumValues = []ClueSmartphoneCreateV2ValidateType{
	"NONE_VERIFICATION",
	"AUTO_VERIFICATION",
	"ALL_VERIFICATION",
	"CLUE_PRIORITY",
	"VALIDITY_PRIORITY",
}

// NewClueSmartphoneCreateV2ValidateTypeFromValue returns a pointer to a valid ClueSmartphoneCreateV2ValidateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueSmartphoneCreateV2ValidateTypeFromValue(v string) (*ClueSmartphoneCreateV2ValidateType, error) {
	ev := ClueSmartphoneCreateV2ValidateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueSmartphoneCreateV2ValidateType: valid values are %v", v, AllowedClueSmartphoneCreateV2ValidateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueSmartphoneCreateV2ValidateType) IsValid() bool {
	for _, existing := range AllowedClueSmartphoneCreateV2ValidateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_smartphone_create_v2_validate_type value
func (v ClueSmartphoneCreateV2ValidateType) Ptr() *ClueSmartphoneCreateV2ValidateType {
	return &v
}
