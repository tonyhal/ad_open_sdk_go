/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanEstimateEffectV10DeepExternalAction
type QianchuanEstimateEffectV10DeepExternalAction string

// List of qianchuan_estimate_effect_v1.0_deep_external_action
const (
	AD_CONVERT_TYPE_LIVE_PAY_ROI_QianchuanEstimateEffectV10DeepExternalAction QianchuanEstimateEffectV10DeepExternalAction = "AD_CONVERT_TYPE_LIVE_PAY_ROI"
)

// All allowed values of QianchuanEstimateEffectV10DeepExternalAction enum
var AllowedQianchuanEstimateEffectV10DeepExternalActionEnumValues = []QianchuanEstimateEffectV10DeepExternalAction{
	"AD_CONVERT_TYPE_LIVE_PAY_ROI",
}

// NewQianchuanEstimateEffectV10DeepExternalActionFromValue returns a pointer to a valid QianchuanEstimateEffectV10DeepExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanEstimateEffectV10DeepExternalActionFromValue(v string) (*QianchuanEstimateEffectV10DeepExternalAction, error) {
	ev := QianchuanEstimateEffectV10DeepExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanEstimateEffectV10DeepExternalAction: valid values are %v", v, AllowedQianchuanEstimateEffectV10DeepExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanEstimateEffectV10DeepExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanEstimateEffectV10DeepExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_estimate_effect_v1.0_deep_external_action value
func (v QianchuanEstimateEffectV10DeepExternalAction) Ptr() *QianchuanEstimateEffectV10DeepExternalAction {
	return &v
}
