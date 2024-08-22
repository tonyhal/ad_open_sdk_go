/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30KeywordsMatchType
type PromotionCreateV30KeywordsMatchType string

// List of promotion_create_v3.0_keywords_match_type
const (
	EXTENSIVE_PromotionCreateV30KeywordsMatchType PromotionCreateV30KeywordsMatchType = "EXTENSIVE"
	PHRASE_PromotionCreateV30KeywordsMatchType    PromotionCreateV30KeywordsMatchType = "PHRASE"
	PRECISION_PromotionCreateV30KeywordsMatchType PromotionCreateV30KeywordsMatchType = "PRECISION"
)

// Ptr returns reference to promotion_create_v3.0_keywords_match_type value
func (v PromotionCreateV30KeywordsMatchType) Ptr() *PromotionCreateV30KeywordsMatchType {
	return &v
}
