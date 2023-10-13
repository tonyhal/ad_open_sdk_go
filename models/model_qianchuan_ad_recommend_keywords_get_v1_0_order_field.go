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

// QianchuanAdRecommendKeywordsGetV10OrderField
type QianchuanAdRecommendKeywordsGetV10OrderField string

// List of qianchuan_ad_recommend_keywords_get_v1.0_order_field
const (
	COMPETITION_QianchuanAdRecommendKeywordsGetV10OrderField QianchuanAdRecommendKeywordsGetV10OrderField = "COMPETITION"
	DEFAULT_QianchuanAdRecommendKeywordsGetV10OrderField     QianchuanAdRecommendKeywordsGetV10OrderField = "DEFAULT"
	PV_QianchuanAdRecommendKeywordsGetV10OrderField          QianchuanAdRecommendKeywordsGetV10OrderField = "PV"
	RELEVANCE_QianchuanAdRecommendKeywordsGetV10OrderField   QianchuanAdRecommendKeywordsGetV10OrderField = "RELEVANCE"
)

// All allowed values of QianchuanAdRecommendKeywordsGetV10OrderField enum
var AllowedQianchuanAdRecommendKeywordsGetV10OrderFieldEnumValues = []QianchuanAdRecommendKeywordsGetV10OrderField{
	"COMPETITION",
	"DEFAULT",
	"PV",
	"RELEVANCE",
}

// NewQianchuanAdRecommendKeywordsGetV10OrderFieldFromValue returns a pointer to a valid QianchuanAdRecommendKeywordsGetV10OrderField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdRecommendKeywordsGetV10OrderFieldFromValue(v string) (*QianchuanAdRecommendKeywordsGetV10OrderField, error) {
	ev := QianchuanAdRecommendKeywordsGetV10OrderField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdRecommendKeywordsGetV10OrderField: valid values are %v", v, AllowedQianchuanAdRecommendKeywordsGetV10OrderFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdRecommendKeywordsGetV10OrderField) IsValid() bool {
	for _, existing := range AllowedQianchuanAdRecommendKeywordsGetV10OrderFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_recommend_keywords_get_v1.0_order_field value
func (v QianchuanAdRecommendKeywordsGetV10OrderField) Ptr() *QianchuanAdRecommendKeywordsGetV10OrderField {
	return &v
}
