/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdGetV10DataListDeliverySettingSmartBidType
type QianchuanAdGetV10DataListDeliverySettingSmartBidType string

// List of qianchuan_ad_get_v1.0_data_list_delivery_setting_smart_bid_type
const (
	SMART_BID_CONSERVATIVE_QianchuanAdGetV10DataListDeliverySettingSmartBidType QianchuanAdGetV10DataListDeliverySettingSmartBidType = "SMART_BID_CONSERVATIVE"
	SMART_BID_CUSTOM_QianchuanAdGetV10DataListDeliverySettingSmartBidType       QianchuanAdGetV10DataListDeliverySettingSmartBidType = "SMART_BID_CUSTOM"
)

// All allowed values of QianchuanAdGetV10DataListDeliverySettingSmartBidType enum
var AllowedQianchuanAdGetV10DataListDeliverySettingSmartBidTypeEnumValues = []QianchuanAdGetV10DataListDeliverySettingSmartBidType{
	"SMART_BID_CONSERVATIVE",
	"SMART_BID_CUSTOM",
}

// NewQianchuanAdGetV10DataListDeliverySettingSmartBidTypeFromValue returns a pointer to a valid QianchuanAdGetV10DataListDeliverySettingSmartBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10DataListDeliverySettingSmartBidTypeFromValue(v string) (*QianchuanAdGetV10DataListDeliverySettingSmartBidType, error) {
	ev := QianchuanAdGetV10DataListDeliverySettingSmartBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10DataListDeliverySettingSmartBidType: valid values are %v", v, AllowedQianchuanAdGetV10DataListDeliverySettingSmartBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10DataListDeliverySettingSmartBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10DataListDeliverySettingSmartBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_data_list_delivery_setting_smart_bid_type value
func (v QianchuanAdGetV10DataListDeliverySettingSmartBidType) Ptr() *QianchuanAdGetV10DataListDeliverySettingSmartBidType {
	return &v
}
