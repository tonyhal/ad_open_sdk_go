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

// QianchuanToolsEstimateAudienceV10AudienceAge
type QianchuanToolsEstimateAudienceV10AudienceAge string

// List of qianchuan_tools_estimate_audience_v1.0_audience_age
const (
	AGE_ABOVE_50_QianchuanToolsEstimateAudienceV10AudienceAge      QianchuanToolsEstimateAudienceV10AudienceAge = "AGE_ABOVE_50"
	AGE_BETWEEN_18_23_QianchuanToolsEstimateAudienceV10AudienceAge QianchuanToolsEstimateAudienceV10AudienceAge = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_24_30_QianchuanToolsEstimateAudienceV10AudienceAge QianchuanToolsEstimateAudienceV10AudienceAge = "AGE_BETWEEN_24_30"
	AGE_BETWEEN_31_40_QianchuanToolsEstimateAudienceV10AudienceAge QianchuanToolsEstimateAudienceV10AudienceAge = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_41_49_QianchuanToolsEstimateAudienceV10AudienceAge QianchuanToolsEstimateAudienceV10AudienceAge = "AGE_BETWEEN_41_49"
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceAge enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceAgeEnumValues = []QianchuanToolsEstimateAudienceV10AudienceAge{
	"AGE_ABOVE_50",
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_24_30",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_41_49",
}

// NewQianchuanToolsEstimateAudienceV10AudienceAgeFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceAge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceAgeFromValue(v string) (*QianchuanToolsEstimateAudienceV10AudienceAge, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceAge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceAge: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceAgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceAge) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceAgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_age value
func (v QianchuanToolsEstimateAudienceV10AudienceAge) Ptr() *QianchuanToolsEstimateAudienceV10AudienceAge {
	return &v
}
