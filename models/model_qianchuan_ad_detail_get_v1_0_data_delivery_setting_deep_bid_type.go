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

// QianchuanAdDetailGetV10DataDeliverySettingDeepBidType
type QianchuanAdDetailGetV10DataDeliverySettingDeepBidType string

// List of qianchuan_ad_detail_get_v1.0_data_delivery_setting_deep_bid_type
const (
	COMM_ROI_QianchuanAdDetailGetV10DataDeliverySettingDeepBidType         QianchuanAdDetailGetV10DataDeliverySettingDeepBidType = "COMM_ROI"
	MIN_QianchuanAdDetailGetV10DataDeliverySettingDeepBidType              QianchuanAdDetailGetV10DataDeliverySettingDeepBidType = "MIN"
	MIN_SECOND_STAGE_QianchuanAdDetailGetV10DataDeliverySettingDeepBidType QianchuanAdDetailGetV10DataDeliverySettingDeepBidType = "MIN_SECOND_STAGE"
	PACING_QianchuanAdDetailGetV10DataDeliverySettingDeepBidType           QianchuanAdDetailGetV10DataDeliverySettingDeepBidType = "PACING"
	PAY_ROI_QianchuanAdDetailGetV10DataDeliverySettingDeepBidType          QianchuanAdDetailGetV10DataDeliverySettingDeepBidType = "PAY_ROI"
)

// All allowed values of QianchuanAdDetailGetV10DataDeliverySettingDeepBidType enum
var AllowedQianchuanAdDetailGetV10DataDeliverySettingDeepBidTypeEnumValues = []QianchuanAdDetailGetV10DataDeliverySettingDeepBidType{
	"COMM_ROI",
	"MIN",
	"MIN_SECOND_STAGE",
	"PACING",
	"PAY_ROI",
}

// NewQianchuanAdDetailGetV10DataDeliverySettingDeepBidTypeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataDeliverySettingDeepBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataDeliverySettingDeepBidTypeFromValue(v string) (*QianchuanAdDetailGetV10DataDeliverySettingDeepBidType, error) {
	ev := QianchuanAdDetailGetV10DataDeliverySettingDeepBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataDeliverySettingDeepBidType: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataDeliverySettingDeepBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataDeliverySettingDeepBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataDeliverySettingDeepBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_delivery_setting_deep_bid_type value
func (v QianchuanAdDetailGetV10DataDeliverySettingDeepBidType) Ptr() *QianchuanAdDetailGetV10DataDeliverySettingDeepBidType {
	return &v
}
