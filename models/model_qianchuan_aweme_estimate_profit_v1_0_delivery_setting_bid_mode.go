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

// QianchuanAwemeEstimateProfitV10DeliverySettingBidMode
type QianchuanAwemeEstimateProfitV10DeliverySettingBidMode string

// List of qianchuan_aweme_estimate_profit_v1.0_delivery_setting_bid_mode
const (
	PRICING_ACTION_QianchuanAwemeEstimateProfitV10DeliverySettingBidMode QianchuanAwemeEstimateProfitV10DeliverySettingBidMode = "PRICING_ACTION"
	PRICING_PLAY_QianchuanAwemeEstimateProfitV10DeliverySettingBidMode   QianchuanAwemeEstimateProfitV10DeliverySettingBidMode = "PRICING_PLAY"
)

// All allowed values of QianchuanAwemeEstimateProfitV10DeliverySettingBidMode enum
var AllowedQianchuanAwemeEstimateProfitV10DeliverySettingBidModeEnumValues = []QianchuanAwemeEstimateProfitV10DeliverySettingBidMode{
	"PRICING_ACTION",
	"PRICING_PLAY",
}

// NewQianchuanAwemeEstimateProfitV10DeliverySettingBidModeFromValue returns a pointer to a valid QianchuanAwemeEstimateProfitV10DeliverySettingBidMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeEstimateProfitV10DeliverySettingBidModeFromValue(v string) (*QianchuanAwemeEstimateProfitV10DeliverySettingBidMode, error) {
	ev := QianchuanAwemeEstimateProfitV10DeliverySettingBidMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeEstimateProfitV10DeliverySettingBidMode: valid values are %v", v, AllowedQianchuanAwemeEstimateProfitV10DeliverySettingBidModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeEstimateProfitV10DeliverySettingBidMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeEstimateProfitV10DeliverySettingBidModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_estimate_profit_v1.0_delivery_setting_bid_mode value
func (v QianchuanAwemeEstimateProfitV10DeliverySettingBidMode) Ptr() *QianchuanAwemeEstimateProfitV10DeliverySettingBidMode {
	return &v
}
