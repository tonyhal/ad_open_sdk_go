/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportLiveRoomProductGetV2OrderField
type ReportLiveRoomProductGetV2OrderField string

// List of report_live_room_product_get_v2_order_field
const (
	CLICK_PRODUCT_COUNT_ReportLiveRoomProductGetV2OrderField  ReportLiveRoomProductGetV2OrderField = "click_product_count"
	CLICK_PRODUCT_UCOUNT_ReportLiveRoomProductGetV2OrderField ReportLiveRoomProductGetV2OrderField = "click_product_ucount"
	LIVE_ORDERS_COUNT_ReportLiveRoomProductGetV2OrderField    ReportLiveRoomProductGetV2OrderField = "live_orders_count"
	ORDER_UCOUNT_ReportLiveRoomProductGetV2OrderField         ReportLiveRoomProductGetV2OrderField = "order_ucount"
	PAY_ORDER_GMV_ReportLiveRoomProductGetV2OrderField        ReportLiveRoomProductGetV2OrderField = "pay_order_gmv"
	PAY_ORDER_COUNT_ReportLiveRoomProductGetV2OrderField      ReportLiveRoomProductGetV2OrderField = "pay_order_count"
	PAY_ORDER_UCOUNT_ReportLiveRoomProductGetV2OrderField     ReportLiveRoomProductGetV2OrderField = "pay_order_ucount"
	ORDER_CONVERT_RATE_ReportLiveRoomProductGetV2OrderField   ReportLiveRoomProductGetV2OrderField = "order_convert_rate"
)

// All allowed values of ReportLiveRoomProductGetV2OrderField enum
var AllowedReportLiveRoomProductGetV2OrderFieldEnumValues = []ReportLiveRoomProductGetV2OrderField{
	"click_product_count",
	"click_product_ucount",
	"live_orders_count",
	"order_ucount",
	"pay_order_gmv",
	"pay_order_count",
	"pay_order_ucount",
	"order_convert_rate",
}

// NewReportLiveRoomProductGetV2OrderFieldFromValue returns a pointer to a valid ReportLiveRoomProductGetV2OrderField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportLiveRoomProductGetV2OrderFieldFromValue(v string) (*ReportLiveRoomProductGetV2OrderField, error) {
	ev := ReportLiveRoomProductGetV2OrderField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportLiveRoomProductGetV2OrderField: valid values are %v", v, AllowedReportLiveRoomProductGetV2OrderFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportLiveRoomProductGetV2OrderField) IsValid() bool {
	for _, existing := range AllowedReportLiveRoomProductGetV2OrderFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_live_room_product_get_v2_order_field value
func (v ReportLiveRoomProductGetV2OrderField) Ptr() *ReportLiveRoomProductGetV2OrderField {
	return &v
}
