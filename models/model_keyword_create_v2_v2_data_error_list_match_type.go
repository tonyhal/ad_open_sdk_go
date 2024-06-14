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

// KeywordCreateV2V2DataErrorListMatchType
type KeywordCreateV2V2DataErrorListMatchType string

// List of keyword_create_v2_v2_data_error_list_match_type
const (
	PRECISION_KeywordCreateV2V2DataErrorListMatchType KeywordCreateV2V2DataErrorListMatchType = "PRECISION"
	PHRASE_KeywordCreateV2V2DataErrorListMatchType    KeywordCreateV2V2DataErrorListMatchType = "PHRASE"
	EXTENSIVE_KeywordCreateV2V2DataErrorListMatchType KeywordCreateV2V2DataErrorListMatchType = "EXTENSIVE"
)

// All allowed values of KeywordCreateV2V2DataErrorListMatchType enum
var AllowedKeywordCreateV2V2DataErrorListMatchTypeEnumValues = []KeywordCreateV2V2DataErrorListMatchType{
	"PRECISION",
	"PHRASE",
	"EXTENSIVE",
}

// NewKeywordCreateV2V2DataErrorListMatchTypeFromValue returns a pointer to a valid KeywordCreateV2V2DataErrorListMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordCreateV2V2DataErrorListMatchTypeFromValue(v string) (*KeywordCreateV2V2DataErrorListMatchType, error) {
	ev := KeywordCreateV2V2DataErrorListMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordCreateV2V2DataErrorListMatchType: valid values are %v", v, AllowedKeywordCreateV2V2DataErrorListMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordCreateV2V2DataErrorListMatchType) IsValid() bool {
	for _, existing := range AllowedKeywordCreateV2V2DataErrorListMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_create_v2_v2_data_error_list_match_type value
func (v KeywordCreateV2V2DataErrorListMatchType) Ptr() *KeywordCreateV2V2DataErrorListMatchType {
	return &v
}
