/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportLiveRoomFlowCategoryGetV2OrderField
type ReportLiveRoomFlowCategoryGetV2OrderField string

// List of report_live_room_flow_category_get_v2_order_field
const (
	LIVE_WATCH_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                      ReportLiveRoomFlowCategoryGetV2OrderField = "live_watch_count"
	LIVE_WATCH_UCOUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                     ReportLiveRoomFlowCategoryGetV2OrderField = "live_watch_ucount"
	LIVE_WATCH_DURATION_ReportLiveRoomFlowCategoryGetV2OrderField                                   ReportLiveRoomFlowCategoryGetV2OrderField = "live_watch_duration"
	PAY_ORDER_GMV_ReportLiveRoomFlowCategoryGetV2OrderField                                         ReportLiveRoomFlowCategoryGetV2OrderField = "pay_order_gmv"
	LIVE_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                            ReportLiveRoomFlowCategoryGetV2OrderField = "live_count"
	LIVE_AVG_WATCH_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                  ReportLiveRoomFlowCategoryGetV2OrderField = "live_avg_watch_count"
	LIVE_AVG_WATCH_DURATION_ReportLiveRoomFlowCategoryGetV2OrderField                               ReportLiveRoomFlowCategoryGetV2OrderField = "live_avg_watch_duration"
	LIVE_PAY_AVG_ORDER_GMV_ReportLiveRoomFlowCategoryGetV2OrderField                                ReportLiveRoomFlowCategoryGetV2OrderField = "live_pay_avg_order_gmv"
	LIVE_DURATION_60S_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                               ReportLiveRoomFlowCategoryGetV2OrderField = "live_duration_60s_count"
	LIVE_DURATION_60S_UCOUNT_ReportLiveRoomFlowCategoryGetV2OrderField                              ReportLiveRoomFlowCategoryGetV2OrderField = "live_duration_60s_ucount"
	LIVE_ONLINE_UCOUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                    ReportLiveRoomFlowCategoryGetV2OrderField = "live_online_ucount"
	CLICK_PRODUCT_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                   ReportLiveRoomFlowCategoryGetV2OrderField = "click_product_count"
	CLICK_PRODUCT_UCOUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                  ReportLiveRoomFlowCategoryGetV2OrderField = "click_product_ucount"
	LIVE_ORDERS_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                     ReportLiveRoomFlowCategoryGetV2OrderField = "live_orders_count"
	ORDER_UCOUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                          ReportLiveRoomFlowCategoryGetV2OrderField = "order_ucount"
	PAY_ORDER_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                       ReportLiveRoomFlowCategoryGetV2OrderField = "pay_order_count"
	PAY_ORDER_UCOUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                      ReportLiveRoomFlowCategoryGetV2OrderField = "pay_order_ucount"
	ORDER_CONVERT_RATE_ReportLiveRoomFlowCategoryGetV2OrderField                                    ReportLiveRoomFlowCategoryGetV2OrderField = "order_convert_rate"
	PER_CUSTOMER_TRANSACTION_ReportLiveRoomFlowCategoryGetV2OrderField                              ReportLiveRoomFlowCategoryGetV2OrderField = "per_customer_transaction"
	LIVE_FOLLOW_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                     ReportLiveRoomFlowCategoryGetV2OrderField = "live_follow_count"
	LIVE_FANS_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                       ReportLiveRoomFlowCategoryGetV2OrderField = "live_fans_count"
	LIVE_COMMENT_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                    ReportLiveRoomFlowCategoryGetV2OrderField = "live_comment_count"
	LIVE_SHARE_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                      ReportLiveRoomFlowCategoryGetV2OrderField = "live_share_count"
	LIVE_GIFT_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                       ReportLiveRoomFlowCategoryGetV2OrderField = "live_gift_count"
	LIVE_GIFT_MONEY_ReportLiveRoomFlowCategoryGetV2OrderField                                       ReportLiveRoomFlowCategoryGetV2OrderField = "live_gift_money"
	CLICK_CART_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                      ReportLiveRoomFlowCategoryGetV2OrderField = "click_cart_count"
	LIVE_DURATION_60S_COUNT_TO_LIVE_WATCH_COUNT_RATE_ReportLiveRoomFlowCategoryGetV2OrderField      ReportLiveRoomFlowCategoryGetV2OrderField = "live_duration_60s_count_to_live_watch_count_rate"
	CLICK_PRODUCT_COUNT_TO_LIVE_DURATION_60S_COUNT_RATE_ReportLiveRoomFlowCategoryGetV2OrderField   ReportLiveRoomFlowCategoryGetV2OrderField = "click_product_count_to_live_duration_60s_count_rate"
	LIVE_ORDERS_COUNT_TO_CLICK_PRODUCT_COUNT_RATE_ReportLiveRoomFlowCategoryGetV2OrderField         ReportLiveRoomFlowCategoryGetV2OrderField = "live_orders_count_to_click_product_count_rate"
	PAY_ORDER_COUNT_TO_LIVE_ORDERS_COUNT_RATE_ReportLiveRoomFlowCategoryGetV2OrderField             ReportLiveRoomFlowCategoryGetV2OrderField = "pay_order_count_to_live_orders_count_rate"
	LIVE_DURATION_60S_UCOUNT_TO_LIVE_WATCH_UCOUNT_RATE_ReportLiveRoomFlowCategoryGetV2OrderField    ReportLiveRoomFlowCategoryGetV2OrderField = "live_duration_60s_ucount_to_live_watch_ucount_rate"
	CLICK_PRODUCT_UCOUNT_TO_LIVE_DURATION_60S_UCOUNT_RATE_ReportLiveRoomFlowCategoryGetV2OrderField ReportLiveRoomFlowCategoryGetV2OrderField = "click_product_ucount_to_live_duration_60s_ucount_rate"
	ORDER_UCOUNT_TO_CLICK_PRODUCT_UCOUNT_RATE_ReportLiveRoomFlowCategoryGetV2OrderField             ReportLiveRoomFlowCategoryGetV2OrderField = "order_ucount_to_click_product_ucount_rate"
	PAY_ORDER_UCOUNT_TO_ORDER_UCOUNT_RATE_ReportLiveRoomFlowCategoryGetV2OrderField                 ReportLiveRoomFlowCategoryGetV2OrderField = "pay_order_ucount_to_order_ucount_rate"
	LIVE_CARD_ICON_COMPONENT_CLICK_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                  ReportLiveRoomFlowCategoryGetV2OrderField = "live_card_icon_component_click_count"
	LIVE_CARD_ICON_COMPONENT_SHOW_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                   ReportLiveRoomFlowCategoryGetV2OrderField = "live_card_icon_component_show_count"
	LIVE_COMPONENT_CTR_ReportLiveRoomFlowCategoryGetV2OrderField                                    ReportLiveRoomFlowCategoryGetV2OrderField = "live_component_ctr"
	LIVE_APP_DOWNLOAD_START_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                         ReportLiveRoomFlowCategoryGetV2OrderField = "live_app_download_start_count"
	LIVE_APP_DOWNLOAD_FINISH_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                        ReportLiveRoomFlowCategoryGetV2OrderField = "live_app_download_finish_count"
	LIVE_APP_INSTALL_FINISH_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                         ReportLiveRoomFlowCategoryGetV2OrderField = "live_app_install_finish_count"
	LIVE_APP_ACTIVE_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                 ReportLiveRoomFlowCategoryGetV2OrderField = "live_app_active_count"
	LIVE_APP_ACTIVE_REGISTER_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                        ReportLiveRoomFlowCategoryGetV2OrderField = "live_app_active_register_count"
	LIVE_APP_GAME_PAY_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                               ReportLiveRoomFlowCategoryGetV2OrderField = "live_app_game_pay_count"
	LIVE_APP_ACTIVE_PAY_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                             ReportLiveRoomFlowCategoryGetV2OrderField = "live_app_active_pay_count"
	LIVE_APP_ATTRIBUTION_NEXT_DAY_OPEN_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField              ReportLiveRoomFlowCategoryGetV2OrderField = "live_app_attribution_next_day_open_count"
	LIVE_APP_ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportLiveRoomFlowCategoryGetV2OrderField               ReportLiveRoomFlowCategoryGetV2OrderField = "live_app_attribution_next_day_open_rate"
	LIVE_GROUPBUY_ORDER_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                             ReportLiveRoomFlowCategoryGetV2OrderField = "live_groupbuy_order_count"
	STAT_LIVE_GROUPBUY_ORDER_GMV_ReportLiveRoomFlowCategoryGetV2OrderField                          ReportLiveRoomFlowCategoryGetV2OrderField = "stat_live_groupbuy_order_gmv"
	LIVE_FORM_SUBMIT_COUNT_ReportLiveRoomFlowCategoryGetV2OrderField                                ReportLiveRoomFlowCategoryGetV2OrderField = "live_form_submit_count"
)

// All allowed values of ReportLiveRoomFlowCategoryGetV2OrderField enum
var AllowedReportLiveRoomFlowCategoryGetV2OrderFieldEnumValues = []ReportLiveRoomFlowCategoryGetV2OrderField{
	"live_watch_count",
	"live_watch_ucount",
	"live_watch_duration",
	"pay_order_gmv",
	"live_count",
	"live_avg_watch_count",
	"live_avg_watch_duration",
	"live_pay_avg_order_gmv",
	"live_duration_60s_count",
	"live_duration_60s_ucount",
	"live_online_ucount",
	"click_product_count",
	"click_product_ucount",
	"live_orders_count",
	"order_ucount",
	"pay_order_count",
	"pay_order_ucount",
	"order_convert_rate",
	"per_customer_transaction",
	"live_follow_count",
	"live_fans_count",
	"live_comment_count",
	"live_share_count",
	"live_gift_count",
	"live_gift_money",
	"click_cart_count",
	"live_duration_60s_count_to_live_watch_count_rate",
	"click_product_count_to_live_duration_60s_count_rate",
	"live_orders_count_to_click_product_count_rate",
	"pay_order_count_to_live_orders_count_rate",
	"live_duration_60s_ucount_to_live_watch_ucount_rate",
	"click_product_ucount_to_live_duration_60s_ucount_rate",
	"order_ucount_to_click_product_ucount_rate",
	"pay_order_ucount_to_order_ucount_rate",
	"live_card_icon_component_click_count",
	"live_card_icon_component_show_count",
	"live_component_ctr",
	"live_app_download_start_count",
	"live_app_download_finish_count",
	"live_app_install_finish_count",
	"live_app_active_count",
	"live_app_active_register_count",
	"live_app_game_pay_count",
	"live_app_active_pay_count",
	"live_app_attribution_next_day_open_count",
	"live_app_attribution_next_day_open_rate",
	"live_groupbuy_order_count",
	"stat_live_groupbuy_order_gmv",
	"live_form_submit_count",
}

// NewReportLiveRoomFlowCategoryGetV2OrderFieldFromValue returns a pointer to a valid ReportLiveRoomFlowCategoryGetV2OrderField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportLiveRoomFlowCategoryGetV2OrderFieldFromValue(v string) (*ReportLiveRoomFlowCategoryGetV2OrderField, error) {
	ev := ReportLiveRoomFlowCategoryGetV2OrderField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportLiveRoomFlowCategoryGetV2OrderField: valid values are %v", v, AllowedReportLiveRoomFlowCategoryGetV2OrderFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportLiveRoomFlowCategoryGetV2OrderField) IsValid() bool {
	for _, existing := range AllowedReportLiveRoomFlowCategoryGetV2OrderFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_live_room_flow_category_get_v2_order_field value
func (v ReportLiveRoomFlowCategoryGetV2OrderField) Ptr() *ReportLiveRoomFlowCategoryGetV2OrderField {
	return &v
}
