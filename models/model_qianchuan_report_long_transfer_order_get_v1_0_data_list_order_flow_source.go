/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource
type QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource string

// List of qianchuan_report_long_transfer_order_get_v1.0_data_list_order_flow_source
const (
	ACTIVITY_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource                    QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "ACTIVITY"
	DOUYIN_SHOPPING_CENTER_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource      QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "DOUYIN_SHOPPING_CENTER"
	GENERAL_SEARCH_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource              QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "GENERAL_SEARCH"
	GUESS_YOU_LIKE_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource              QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "GUESS_YOU_LIKE"
	HOMEPAGE_FOLLOW_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource             QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "HOMEPAGE_FOLLOW"
	LIVE_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource                        QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "LIVE"
	LIVE_OTHER_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource                  QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "LIVE_OTHER"
	OTHERS_HOMEPAGE_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource             QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "OTHERS_HOMEPAGE"
	OTHER_PROFILE_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource               QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "OTHER_PROFILE"
	PRODUCT_CARD_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource                QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "PRODUCT_CARD"
	PRODUCT_CARD_GENERAL_SEARCH_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "PRODUCT_CARD_GENERAL_SEARCH"
	PRODUCT_CARD_OTHER_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource          QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "PRODUCT_CARD_OTHER"
	QIANCHUAN_PROMOTE_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource           QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "QIANCHUAN_PROMOTE"
	RECOMMEND_LIVE_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource              QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "RECOMMEND_LIVE"
	RECOMMEND_PRODUCT_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource           QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "RECOMMEND_PRODUCT"
	RECOMMEND_VIDEO_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource             QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "RECOMMEND_VIDEO"
	SHOP_WINDOW_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource                 QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "SHOP_WINDOW"
	UNKNOWN_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource                     QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "UNKNOWN"
	VIDEO_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource                       QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "VIDEO"
	VIDEO_ACTIVITY_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource              QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "VIDEO_ACTIVITY"
	VIDEO_GENERAL_SEARCH_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource        QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "VIDEO_GENERAL_SEARCH"
	VIDEO_HOMEPAGE_FOLLOW_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource       QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "VIDEO_HOMEPAGE_FOLLOW"
	VIDEO_OTHER_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource                 QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "VIDEO_OTHER"
	VIDEO_TO_LIVE_QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource               QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource = "VIDEO_TO_LIVE"
)

// All allowed values of QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource enum
var AllowedQianchuanReportLongTransferOrderGetV10DataListOrderFlowSourceEnumValues = []QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource{
	"ACTIVITY",
	"DOUYIN_SHOPPING_CENTER",
	"GENERAL_SEARCH",
	"GUESS_YOU_LIKE",
	"HOMEPAGE_FOLLOW",
	"LIVE",
	"LIVE_OTHER",
	"OTHERS_HOMEPAGE",
	"OTHER_PROFILE",
	"PRODUCT_CARD",
	"PRODUCT_CARD_GENERAL_SEARCH",
	"PRODUCT_CARD_OTHER",
	"QIANCHUAN_PROMOTE",
	"RECOMMEND_LIVE",
	"RECOMMEND_PRODUCT",
	"RECOMMEND_VIDEO",
	"SHOP_WINDOW",
	"UNKNOWN",
	"VIDEO",
	"VIDEO_ACTIVITY",
	"VIDEO_GENERAL_SEARCH",
	"VIDEO_HOMEPAGE_FOLLOW",
	"VIDEO_OTHER",
	"VIDEO_TO_LIVE",
}

// NewQianchuanReportLongTransferOrderGetV10DataListOrderFlowSourceFromValue returns a pointer to a valid QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportLongTransferOrderGetV10DataListOrderFlowSourceFromValue(v string) (*QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource, error) {
	ev := QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource: valid values are %v", v, AllowedQianchuanReportLongTransferOrderGetV10DataListOrderFlowSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource) IsValid() bool {
	for _, existing := range AllowedQianchuanReportLongTransferOrderGetV10DataListOrderFlowSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_long_transfer_order_get_v1.0_data_list_order_flow_source value
func (v QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource) Ptr() *QianchuanReportLongTransferOrderGetV10DataListOrderFlowSource {
	return &v
}
