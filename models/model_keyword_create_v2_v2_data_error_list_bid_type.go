/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordCreateV2V2DataErrorListBidType
type KeywordCreateV2V2DataErrorListBidType string

// List of keyword_create_v2_v2_data_error_list_bid_type
const (
	CUSTOM_KeywordCreateV2V2DataErrorListBidType         KeywordCreateV2V2DataErrorListBidType = "CUSTOM"
	BRAND_AD_KeywordCreateV2V2DataErrorListBidType       KeywordCreateV2V2DataErrorListBidType = "BRAND_AD"
	SUGGEST_KeywordCreateV2V2DataErrorListBidType        KeywordCreateV2V2DataErrorListBidType = "SUGGEST"
	FEED_TO_SEARCH_KeywordCreateV2V2DataErrorListBidType KeywordCreateV2V2DataErrorListBidType = "FEED_TO_SEARCH"
	WITH_AD_KeywordCreateV2V2DataErrorListBidType        KeywordCreateV2V2DataErrorListBidType = "WITH_AD"
)

// Ptr returns reference to keyword_create_v2_v2_data_error_list_bid_type value
func (v KeywordCreateV2V2DataErrorListBidType) Ptr() *KeywordCreateV2V2DataErrorListBidType {
	return &v
}
