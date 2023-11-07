/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10AudienceSmartInterestAction
type QianchuanAdCreateV10AudienceSmartInterestAction string

// List of qianchuan_ad_create_v1.0_audience_smart_interest_action
const (
	CUSTOM_QianchuanAdCreateV10AudienceSmartInterestAction    QianchuanAdCreateV10AudienceSmartInterestAction = "CUSTOM"
	RECOMMEND_QianchuanAdCreateV10AudienceSmartInterestAction QianchuanAdCreateV10AudienceSmartInterestAction = "RECOMMEND"
)

// All allowed values of QianchuanAdCreateV10AudienceSmartInterestAction enum
var AllowedQianchuanAdCreateV10AudienceSmartInterestActionEnumValues = []QianchuanAdCreateV10AudienceSmartInterestAction{
	"CUSTOM",
	"RECOMMEND",
}

// NewQianchuanAdCreateV10AudienceSmartInterestActionFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceSmartInterestAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceSmartInterestActionFromValue(v string) (*QianchuanAdCreateV10AudienceSmartInterestAction, error) {
	ev := QianchuanAdCreateV10AudienceSmartInterestAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceSmartInterestAction: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceSmartInterestActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceSmartInterestAction) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceSmartInterestActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_smart_interest_action value
func (v QianchuanAdCreateV10AudienceSmartInterestAction) Ptr() *QianchuanAdCreateV10AudienceSmartInterestAction {
	return &v
}
