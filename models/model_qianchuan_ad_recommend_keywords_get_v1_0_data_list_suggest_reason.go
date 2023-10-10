/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdRecommendKeywordsGetV10DataListSuggestReason
type QianchuanAdRecommendKeywordsGetV10DataListSuggestReason string

// List of qianchuan_ad_recommend_keywords_get_v1.0_data_list_suggest_reason
const (
	CLICK_QianchuanAdRecommendKeywordsGetV10DataListSuggestReason      QianchuanAdRecommendKeywordsGetV10DataListSuggestReason = "CLICK"
	DARKHORSE_QianchuanAdRecommendKeywordsGetV10DataListSuggestReason  QianchuanAdRecommendKeywordsGetV10DataListSuggestReason = "DARKHORSE"
	LOW_COST_QianchuanAdRecommendKeywordsGetV10DataListSuggestReason   QianchuanAdRecommendKeywordsGetV10DataListSuggestReason = "LOW_COST"
	POTENTIAL_QianchuanAdRecommendKeywordsGetV10DataListSuggestReason  QianchuanAdRecommendKeywordsGetV10DataListSuggestReason = "POTENTIAL"
	PREFERENCE_QianchuanAdRecommendKeywordsGetV10DataListSuggestReason QianchuanAdRecommendKeywordsGetV10DataListSuggestReason = "PREFERENCE"
)

// All allowed values of QianchuanAdRecommendKeywordsGetV10DataListSuggestReason enum
var AllowedQianchuanAdRecommendKeywordsGetV10DataListSuggestReasonEnumValues = []QianchuanAdRecommendKeywordsGetV10DataListSuggestReason{
	"CLICK",
	"DARKHORSE",
	"LOW_COST",
	"POTENTIAL",
	"PREFERENCE",
}

// NewQianchuanAdRecommendKeywordsGetV10DataListSuggestReasonFromValue returns a pointer to a valid QianchuanAdRecommendKeywordsGetV10DataListSuggestReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdRecommendKeywordsGetV10DataListSuggestReasonFromValue(v string) (*QianchuanAdRecommendKeywordsGetV10DataListSuggestReason, error) {
	ev := QianchuanAdRecommendKeywordsGetV10DataListSuggestReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdRecommendKeywordsGetV10DataListSuggestReason: valid values are %v", v, AllowedQianchuanAdRecommendKeywordsGetV10DataListSuggestReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdRecommendKeywordsGetV10DataListSuggestReason) IsValid() bool {
	for _, existing := range AllowedQianchuanAdRecommendKeywordsGetV10DataListSuggestReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_recommend_keywords_get_v1.0_data_list_suggest_reason value
func (v QianchuanAdRecommendKeywordsGetV10DataListSuggestReason) Ptr() *QianchuanAdRecommendKeywordsGetV10DataListSuggestReason {
	return &v
}
