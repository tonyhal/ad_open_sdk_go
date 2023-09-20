/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10AudienceAutoExtendTargets
type QianchuanAdUpdateV10AudienceAutoExtendTargets string

// List of qianchuan_ad_update_v1.0_audience_auto_extend_targets
const (
	AGE_QianchuanAdUpdateV10AudienceAutoExtendTargets             QianchuanAdUpdateV10AudienceAutoExtendTargets = "AGE"
	CUSTOM_AUDIENCE_QianchuanAdUpdateV10AudienceAutoExtendTargets QianchuanAdUpdateV10AudienceAutoExtendTargets = "CUSTOM_AUDIENCE"
	GENDER_QianchuanAdUpdateV10AudienceAutoExtendTargets          QianchuanAdUpdateV10AudienceAutoExtendTargets = "GENDER"
	INTEREST_ACTION_QianchuanAdUpdateV10AudienceAutoExtendTargets QianchuanAdUpdateV10AudienceAutoExtendTargets = "INTEREST_ACTION"
	REGION_QianchuanAdUpdateV10AudienceAutoExtendTargets          QianchuanAdUpdateV10AudienceAutoExtendTargets = "REGION"
)

// All allowed values of QianchuanAdUpdateV10AudienceAutoExtendTargets enum
var AllowedQianchuanAdUpdateV10AudienceAutoExtendTargetsEnumValues = []QianchuanAdUpdateV10AudienceAutoExtendTargets{
	"AGE",
	"CUSTOM_AUDIENCE",
	"GENDER",
	"INTEREST_ACTION",
	"REGION",
}

// NewQianchuanAdUpdateV10AudienceAutoExtendTargetsFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceAutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceAutoExtendTargetsFromValue(v string) (*QianchuanAdUpdateV10AudienceAutoExtendTargets, error) {
	ev := QianchuanAdUpdateV10AudienceAutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceAutoExtendTargets: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceAutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceAutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceAutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_auto_extend_targets value
func (v QianchuanAdUpdateV10AudienceAutoExtendTargets) Ptr() *QianchuanAdUpdateV10AudienceAutoExtendTargets {
	return &v
}
