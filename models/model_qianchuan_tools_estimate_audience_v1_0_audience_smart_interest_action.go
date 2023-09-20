/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction
type QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction string

// List of qianchuan_tools_estimate_audience_v1.0_audience_smart_interest_action
const (
	CUSTOM_QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction    QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction = "CUSTOM"
	RECOMMEND_QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction = "RECOMMEND"
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceSmartInterestActionEnumValues = []QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction{
	"CUSTOM",
	"RECOMMEND",
}

// NewQianchuanToolsEstimateAudienceV10AudienceSmartInterestActionFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceSmartInterestActionFromValue(v string) (*QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceSmartInterestActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceSmartInterestActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_smart_interest_action value
func (v QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction) Ptr() *QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction {
	return &v
}
