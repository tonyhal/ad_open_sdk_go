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

// QianchuanAdCreateV10AudienceAutoExtendTargets
type QianchuanAdCreateV10AudienceAutoExtendTargets string

// List of qianchuan_ad_create_v1.0_audience_auto_extend_targets
const (
	AGE_QianchuanAdCreateV10AudienceAutoExtendTargets             QianchuanAdCreateV10AudienceAutoExtendTargets = "AGE"
	CUSTOM_AUDIENCE_QianchuanAdCreateV10AudienceAutoExtendTargets QianchuanAdCreateV10AudienceAutoExtendTargets = "CUSTOM_AUDIENCE"
	GENDER_QianchuanAdCreateV10AudienceAutoExtendTargets          QianchuanAdCreateV10AudienceAutoExtendTargets = "GENDER"
	INTEREST_ACTION_QianchuanAdCreateV10AudienceAutoExtendTargets QianchuanAdCreateV10AudienceAutoExtendTargets = "INTEREST_ACTION"
	REGION_QianchuanAdCreateV10AudienceAutoExtendTargets          QianchuanAdCreateV10AudienceAutoExtendTargets = "REGION"
)

// All allowed values of QianchuanAdCreateV10AudienceAutoExtendTargets enum
var AllowedQianchuanAdCreateV10AudienceAutoExtendTargetsEnumValues = []QianchuanAdCreateV10AudienceAutoExtendTargets{
	"AGE",
	"CUSTOM_AUDIENCE",
	"GENDER",
	"INTEREST_ACTION",
	"REGION",
}

// NewQianchuanAdCreateV10AudienceAutoExtendTargetsFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceAutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceAutoExtendTargetsFromValue(v string) (*QianchuanAdCreateV10AudienceAutoExtendTargets, error) {
	ev := QianchuanAdCreateV10AudienceAutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceAutoExtendTargets: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceAutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceAutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceAutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_auto_extend_targets value
func (v QianchuanAdCreateV10AudienceAutoExtendTargets) Ptr() *QianchuanAdCreateV10AudienceAutoExtendTargets {
	return &v
}
