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

// QianchuanAdCreateV10DeliverySettingDeepExternalAction
type QianchuanAdCreateV10DeliverySettingDeepExternalAction string

// List of qianchuan_ad_create_v1.0_delivery_setting_deep_external_action
const (
	AD_CONVERT_TYPE_CARD_ACTIVE_QianchuanAdCreateV10DeliverySettingDeepExternalAction  QianchuanAdCreateV10DeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_CARD_ACTIVE"
	AD_CONVERT_TYPE_LIVE_PAY_ROI_QianchuanAdCreateV10DeliverySettingDeepExternalAction QianchuanAdCreateV10DeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_PAY_ROI"
)

// All allowed values of QianchuanAdCreateV10DeliverySettingDeepExternalAction enum
var AllowedQianchuanAdCreateV10DeliverySettingDeepExternalActionEnumValues = []QianchuanAdCreateV10DeliverySettingDeepExternalAction{
	"AD_CONVERT_TYPE_CARD_ACTIVE",
	"AD_CONVERT_TYPE_LIVE_PAY_ROI",
}

// NewQianchuanAdCreateV10DeliverySettingDeepExternalActionFromValue returns a pointer to a valid QianchuanAdCreateV10DeliverySettingDeepExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10DeliverySettingDeepExternalActionFromValue(v string) (*QianchuanAdCreateV10DeliverySettingDeepExternalAction, error) {
	ev := QianchuanAdCreateV10DeliverySettingDeepExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10DeliverySettingDeepExternalAction: valid values are %v", v, AllowedQianchuanAdCreateV10DeliverySettingDeepExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10DeliverySettingDeepExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10DeliverySettingDeepExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_delivery_setting_deep_external_action value
func (v QianchuanAdCreateV10DeliverySettingDeepExternalAction) Ptr() *QianchuanAdCreateV10DeliverySettingDeepExternalAction {
	return &v
}
