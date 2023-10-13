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

// ReportLiveRoomAudiencePortraitGetV2OrderField
type ReportLiveRoomAudiencePortraitGetV2OrderField string

// List of report_live_room_audience_portrait_get_v2_order_field
const (
	LIVE_WATCH_COUNT_ReportLiveRoomAudiencePortraitGetV2OrderField         ReportLiveRoomAudiencePortraitGetV2OrderField = "live_watch_count"
	LIVE_FOLLOW_COUNT_ReportLiveRoomAudiencePortraitGetV2OrderField        ReportLiveRoomAudiencePortraitGetV2OrderField = "live_follow_count"
	LIVE_FANS_COUNT_ReportLiveRoomAudiencePortraitGetV2OrderField          ReportLiveRoomAudiencePortraitGetV2OrderField = "live_fans_count"
	LIVE_SHARE_COUNT_ReportLiveRoomAudiencePortraitGetV2OrderField         ReportLiveRoomAudiencePortraitGetV2OrderField = "live_share_count"
	LIVE_COMMENT_COUNT_ReportLiveRoomAudiencePortraitGetV2OrderField       ReportLiveRoomAudiencePortraitGetV2OrderField = "live_comment_count"
	LIVE_GIFT_COUNT_ReportLiveRoomAudiencePortraitGetV2OrderField          ReportLiveRoomAudiencePortraitGetV2OrderField = "live_gift_count"
	LIVE_GIFT_MONEY_ReportLiveRoomAudiencePortraitGetV2OrderField          ReportLiveRoomAudiencePortraitGetV2OrderField = "live_gift_money"
	CLICK_CART_COUNT_ReportLiveRoomAudiencePortraitGetV2OrderField         ReportLiveRoomAudiencePortraitGetV2OrderField = "click_cart_count"
	CLICK_PRODUCT_COUNT_ReportLiveRoomAudiencePortraitGetV2OrderField      ReportLiveRoomAudiencePortraitGetV2OrderField = "click_product_count"
	LIVE_ORDERS_COUNT_ReportLiveRoomAudiencePortraitGetV2OrderField        ReportLiveRoomAudiencePortraitGetV2OrderField = "live_orders_count"
	PAY_ORDER_COUNT_ReportLiveRoomAudiencePortraitGetV2OrderField          ReportLiveRoomAudiencePortraitGetV2OrderField = "pay_order_count"
	PAY_ORDER_GMV_ReportLiveRoomAudiencePortraitGetV2OrderField            ReportLiveRoomAudiencePortraitGetV2OrderField = "pay_order_gmv"
	PER_CUSTOMER_TRANSACTION_ReportLiveRoomAudiencePortraitGetV2OrderField ReportLiveRoomAudiencePortraitGetV2OrderField = "per_customer_transaction"
)

// All allowed values of ReportLiveRoomAudiencePortraitGetV2OrderField enum
var AllowedReportLiveRoomAudiencePortraitGetV2OrderFieldEnumValues = []ReportLiveRoomAudiencePortraitGetV2OrderField{
	"live_watch_count",
	"live_follow_count",
	"live_fans_count",
	"live_share_count",
	"live_comment_count",
	"live_gift_count",
	"live_gift_money",
	"click_cart_count",
	"click_product_count",
	"live_orders_count",
	"pay_order_count",
	"pay_order_gmv",
	"per_customer_transaction",
}

// NewReportLiveRoomAudiencePortraitGetV2OrderFieldFromValue returns a pointer to a valid ReportLiveRoomAudiencePortraitGetV2OrderField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportLiveRoomAudiencePortraitGetV2OrderFieldFromValue(v string) (*ReportLiveRoomAudiencePortraitGetV2OrderField, error) {
	ev := ReportLiveRoomAudiencePortraitGetV2OrderField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportLiveRoomAudiencePortraitGetV2OrderField: valid values are %v", v, AllowedReportLiveRoomAudiencePortraitGetV2OrderFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportLiveRoomAudiencePortraitGetV2OrderField) IsValid() bool {
	for _, existing := range AllowedReportLiveRoomAudiencePortraitGetV2OrderFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_live_room_audience_portrait_get_v2_order_field value
func (v ReportLiveRoomAudiencePortraitGetV2OrderField) Ptr() *ReportLiveRoomAudiencePortraitGetV2OrderField {
	return &v
}
