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

// QianchuanToolsSmartBoostAdBoostSetV10OptType
type QianchuanToolsSmartBoostAdBoostSetV10OptType string

// List of qianchuan_tools_smart_boost_ad_boost_set_v1.0_opt_type
const (
	CANCEL_SUBSCRIBE_QianchuanToolsSmartBoostAdBoostSetV10OptType QianchuanToolsSmartBoostAdBoostSetV10OptType = "CANCEL_SUBSCRIBE"
	START_RAISE_QianchuanToolsSmartBoostAdBoostSetV10OptType      QianchuanToolsSmartBoostAdBoostSetV10OptType = "START_RAISE"
	STOP_RAISE_QianchuanToolsSmartBoostAdBoostSetV10OptType       QianchuanToolsSmartBoostAdBoostSetV10OptType = "STOP_RAISE"
	SUBSCRIBE_RAISE_QianchuanToolsSmartBoostAdBoostSetV10OptType  QianchuanToolsSmartBoostAdBoostSetV10OptType = "SUBSCRIBE_RAISE"
)

// All allowed values of QianchuanToolsSmartBoostAdBoostSetV10OptType enum
var AllowedQianchuanToolsSmartBoostAdBoostSetV10OptTypeEnumValues = []QianchuanToolsSmartBoostAdBoostSetV10OptType{
	"CANCEL_SUBSCRIBE",
	"START_RAISE",
	"STOP_RAISE",
	"SUBSCRIBE_RAISE",
}

// NewQianchuanToolsSmartBoostAdBoostSetV10OptTypeFromValue returns a pointer to a valid QianchuanToolsSmartBoostAdBoostSetV10OptType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsSmartBoostAdBoostSetV10OptTypeFromValue(v string) (*QianchuanToolsSmartBoostAdBoostSetV10OptType, error) {
	ev := QianchuanToolsSmartBoostAdBoostSetV10OptType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsSmartBoostAdBoostSetV10OptType: valid values are %v", v, AllowedQianchuanToolsSmartBoostAdBoostSetV10OptTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsSmartBoostAdBoostSetV10OptType) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsSmartBoostAdBoostSetV10OptTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_smart_boost_ad_boost_set_v1.0_opt_type value
func (v QianchuanToolsSmartBoostAdBoostSetV10OptType) Ptr() *QianchuanToolsSmartBoostAdBoostSetV10OptType {
	return &v
}
