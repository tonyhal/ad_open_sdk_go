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

// KeywordGetV2DataListBidType
type KeywordGetV2DataListBidType string

// List of keyword_get_v2_data_list_bid_type
const (
	SUGGEST_KeywordGetV2DataListBidType        KeywordGetV2DataListBidType = "SUGGEST"
	CUSTOM_KeywordGetV2DataListBidType         KeywordGetV2DataListBidType = "CUSTOM"
	BRAND_AD_KeywordGetV2DataListBidType       KeywordGetV2DataListBidType = "BRAND_AD"
	FEED_TO_SEARCH_KeywordGetV2DataListBidType KeywordGetV2DataListBidType = "FEED_TO_SEARCH"
	WITH_AD_KeywordGetV2DataListBidType        KeywordGetV2DataListBidType = "WITH_AD"
)

// All allowed values of KeywordGetV2DataListBidType enum
var AllowedKeywordGetV2DataListBidTypeEnumValues = []KeywordGetV2DataListBidType{
	"SUGGEST",
	"CUSTOM",
	"BRAND_AD",
	"FEED_TO_SEARCH",
	"WITH_AD",
}

// NewKeywordGetV2DataListBidTypeFromValue returns a pointer to a valid KeywordGetV2DataListBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordGetV2DataListBidTypeFromValue(v string) (*KeywordGetV2DataListBidType, error) {
	ev := KeywordGetV2DataListBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordGetV2DataListBidType: valid values are %v", v, AllowedKeywordGetV2DataListBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordGetV2DataListBidType) IsValid() bool {
	for _, existing := range AllowedKeywordGetV2DataListBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_get_v2_data_list_bid_type value
func (v KeywordGetV2DataListBidType) Ptr() *KeywordGetV2DataListBidType {
	return &v
}
