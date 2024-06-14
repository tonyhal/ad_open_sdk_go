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

// QianchuanAwemeEstimateProfitV10DeliverySettingBidType
type QianchuanAwemeEstimateProfitV10DeliverySettingBidType string

// List of qianchuan_aweme_estimate_profit_v1.0_delivery_setting_bid_type
const (
	AUTO_BID_QianchuanAwemeEstimateProfitV10DeliverySettingBidType   QianchuanAwemeEstimateProfitV10DeliverySettingBidType = "AUTO_BID"
	MANUAL_BID_QianchuanAwemeEstimateProfitV10DeliverySettingBidType QianchuanAwemeEstimateProfitV10DeliverySettingBidType = "MANUAL_BID"
)

// All allowed values of QianchuanAwemeEstimateProfitV10DeliverySettingBidType enum
var AllowedQianchuanAwemeEstimateProfitV10DeliverySettingBidTypeEnumValues = []QianchuanAwemeEstimateProfitV10DeliverySettingBidType{
	"AUTO_BID",
	"MANUAL_BID",
}

// NewQianchuanAwemeEstimateProfitV10DeliverySettingBidTypeFromValue returns a pointer to a valid QianchuanAwemeEstimateProfitV10DeliverySettingBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeEstimateProfitV10DeliverySettingBidTypeFromValue(v string) (*QianchuanAwemeEstimateProfitV10DeliverySettingBidType, error) {
	ev := QianchuanAwemeEstimateProfitV10DeliverySettingBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeEstimateProfitV10DeliverySettingBidType: valid values are %v", v, AllowedQianchuanAwemeEstimateProfitV10DeliverySettingBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeEstimateProfitV10DeliverySettingBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeEstimateProfitV10DeliverySettingBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_estimate_profit_v1.0_delivery_setting_bid_type value
func (v QianchuanAwemeEstimateProfitV10DeliverySettingBidType) Ptr() *QianchuanAwemeEstimateProfitV10DeliverySettingBidType {
	return &v
}
