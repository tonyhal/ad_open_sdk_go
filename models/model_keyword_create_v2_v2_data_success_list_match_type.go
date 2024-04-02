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

// KeywordCreateV2V2DataSuccessListMatchType
type KeywordCreateV2V2DataSuccessListMatchType string

// List of keyword_create_v2_v2_data_success_list_match_type
const (
	PRECISION_KeywordCreateV2V2DataSuccessListMatchType KeywordCreateV2V2DataSuccessListMatchType = "PRECISION"
	EXTENSIVE_KeywordCreateV2V2DataSuccessListMatchType KeywordCreateV2V2DataSuccessListMatchType = "EXTENSIVE"
	PHRASE_KeywordCreateV2V2DataSuccessListMatchType    KeywordCreateV2V2DataSuccessListMatchType = "PHRASE"
)

// All allowed values of KeywordCreateV2V2DataSuccessListMatchType enum
var AllowedKeywordCreateV2V2DataSuccessListMatchTypeEnumValues = []KeywordCreateV2V2DataSuccessListMatchType{
	"PRECISION",
	"EXTENSIVE",
	"PHRASE",
}

// NewKeywordCreateV2V2DataSuccessListMatchTypeFromValue returns a pointer to a valid KeywordCreateV2V2DataSuccessListMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordCreateV2V2DataSuccessListMatchTypeFromValue(v string) (*KeywordCreateV2V2DataSuccessListMatchType, error) {
	ev := KeywordCreateV2V2DataSuccessListMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordCreateV2V2DataSuccessListMatchType: valid values are %v", v, AllowedKeywordCreateV2V2DataSuccessListMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordCreateV2V2DataSuccessListMatchType) IsValid() bool {
	for _, existing := range AllowedKeywordCreateV2V2DataSuccessListMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_create_v2_v2_data_success_list_match_type value
func (v KeywordCreateV2V2DataSuccessListMatchType) Ptr() *KeywordCreateV2V2DataSuccessListMatchType {
	return &v
}
