/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// KeywordListV30DataListBidType
type KeywordListV30DataListBidType string

// List of keyword_list_v3.0_data_list_bid_type
const (
	CUSTOM_KeywordListV30DataListBidType         KeywordListV30DataListBidType = "CUSTOM"
	FEED_TO_SEARCH_KeywordListV30DataListBidType KeywordListV30DataListBidType = "FEED_TO_SEARCH"
	WITH_PROMOTION_KeywordListV30DataListBidType KeywordListV30DataListBidType = "WITH_PROMOTION"
)

// All allowed values of KeywordListV30DataListBidType enum
var AllowedKeywordListV30DataListBidTypeEnumValues = []KeywordListV30DataListBidType{
	"CUSTOM",
	"FEED_TO_SEARCH",
	"WITH_PROMOTION",
}

// NewKeywordListV30DataListBidTypeFromValue returns a pointer to a valid KeywordListV30DataListBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordListV30DataListBidTypeFromValue(v string) (*KeywordListV30DataListBidType, error) {
	ev := KeywordListV30DataListBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordListV30DataListBidType: valid values are %v", v, AllowedKeywordListV30DataListBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordListV30DataListBidType) IsValid() bool {
	for _, existing := range AllowedKeywordListV30DataListBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_list_v3.0_data_list_bid_type value
func (v KeywordListV30DataListBidType) Ptr() *KeywordListV30DataListBidType {
	return &v
}
