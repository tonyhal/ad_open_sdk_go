/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataAudienceSmartInterestAction
type QianchuanAdDetailGetV10DataAudienceSmartInterestAction string

// List of qianchuan_ad_detail_get_v1.0_data_audience_smart_interest_action
const (
	CUSTOM_QianchuanAdDetailGetV10DataAudienceSmartInterestAction    QianchuanAdDetailGetV10DataAudienceSmartInterestAction = "CUSTOM"
	RECOMMEND_QianchuanAdDetailGetV10DataAudienceSmartInterestAction QianchuanAdDetailGetV10DataAudienceSmartInterestAction = "RECOMMEND"
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceSmartInterestAction enum
var AllowedQianchuanAdDetailGetV10DataAudienceSmartInterestActionEnumValues = []QianchuanAdDetailGetV10DataAudienceSmartInterestAction{
	"CUSTOM",
	"RECOMMEND",
}

// NewQianchuanAdDetailGetV10DataAudienceSmartInterestActionFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceSmartInterestAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceSmartInterestActionFromValue(v string) (*QianchuanAdDetailGetV10DataAudienceSmartInterestAction, error) {
	ev := QianchuanAdDetailGetV10DataAudienceSmartInterestAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceSmartInterestAction: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceSmartInterestActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceSmartInterestAction) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceSmartInterestActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_smart_interest_action value
func (v QianchuanAdDetailGetV10DataAudienceSmartInterestAction) Ptr() *QianchuanAdDetailGetV10DataAudienceSmartInterestAction {
	return &v
}
