/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordUpdateV2V2KeywordsBidType
type KeywordUpdateV2V2KeywordsBidType string

// List of keyword_update_v2_v2_keywords_bid_type
const (
	WITH_AD_KeywordUpdateV2V2KeywordsBidType        KeywordUpdateV2V2KeywordsBidType = "WITH_AD"
	BRAND_AD_KeywordUpdateV2V2KeywordsBidType       KeywordUpdateV2V2KeywordsBidType = "BRAND_AD"
	SUGGEST_KeywordUpdateV2V2KeywordsBidType        KeywordUpdateV2V2KeywordsBidType = "SUGGEST"
	CUSTOM_KeywordUpdateV2V2KeywordsBidType         KeywordUpdateV2V2KeywordsBidType = "CUSTOM"
	FEED_TO_SEARCH_KeywordUpdateV2V2KeywordsBidType KeywordUpdateV2V2KeywordsBidType = "FEED_TO_SEARCH"
)

// Ptr returns reference to keyword_update_v2_v2_keywords_bid_type value
func (v KeywordUpdateV2V2KeywordsBidType) Ptr() *KeywordUpdateV2V2KeywordsBidType {
	return &v
}
