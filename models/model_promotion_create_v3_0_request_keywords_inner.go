/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30RequestKeywordsInner struct for PromotionCreateV30RequestKeywordsInner
type PromotionCreateV30RequestKeywordsInner struct {
	//
	Bid       *float64                            `json:"bid,omitempty"`
	BidType   *PromotionCreateV30KeywordsBidType  `json:"bid_type,omitempty"`
	MatchType PromotionCreateV30KeywordsMatchType `json:"match_type"`
	//
	Word string `json:"word"`
}
