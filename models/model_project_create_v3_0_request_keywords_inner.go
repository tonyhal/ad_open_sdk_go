/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30RequestKeywordsInner struct for ProjectCreateV30RequestKeywordsInner
type ProjectCreateV30RequestKeywordsInner struct {
	BidType   ProjectCreateV30KeywordsBidType   `json:"bid_type"`
	MatchType ProjectCreateV30KeywordsMatchType `json:"match_type"`
	//
	Word string `json:"word"`
}
