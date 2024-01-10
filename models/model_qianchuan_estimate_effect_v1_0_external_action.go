/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanEstimateEffectV10ExternalAction
type QianchuanEstimateEffectV10ExternalAction string

// List of qianchuan_estimate_effect_v1.0_external_action
const (
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_QianchuanEstimateEffectV10ExternalAction QianchuanEstimateEffectV10ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
)

// All allowed values of QianchuanEstimateEffectV10ExternalAction enum
var AllowedQianchuanEstimateEffectV10ExternalActionEnumValues = []QianchuanEstimateEffectV10ExternalAction{
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY",
}

// NewQianchuanEstimateEffectV10ExternalActionFromValue returns a pointer to a valid QianchuanEstimateEffectV10ExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanEstimateEffectV10ExternalActionFromValue(v string) (*QianchuanEstimateEffectV10ExternalAction, error) {
	ev := QianchuanEstimateEffectV10ExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanEstimateEffectV10ExternalAction: valid values are %v", v, AllowedQianchuanEstimateEffectV10ExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanEstimateEffectV10ExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanEstimateEffectV10ExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_estimate_effect_v1.0_external_action value
func (v QianchuanEstimateEffectV10ExternalAction) Ptr() *QianchuanEstimateEffectV10ExternalAction {
	return &v
}
