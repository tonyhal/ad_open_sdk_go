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

// QianchuanAwemeReportOrderGetV10FilteringExternalAction
type QianchuanAwemeReportOrderGetV10FilteringExternalAction string

// List of qianchuan_aweme_report_order_get_v1.0_filtering_external_action
const (
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_QianchuanAwemeReportOrderGetV10FilteringExternalAction     QianchuanAwemeReportOrderGetV10FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_QianchuanAwemeReportOrderGetV10FilteringExternalAction           QianchuanAwemeReportOrderGetV10FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_QianchuanAwemeReportOrderGetV10FilteringExternalAction             QianchuanAwemeReportOrderGetV10FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_ROI_QianchuanAwemeReportOrderGetV10FilteringExternalAction                      QianchuanAwemeReportOrderGetV10FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_ROI"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_QianchuanAwemeReportOrderGetV10FilteringExternalAction      QianchuanAwemeReportOrderGetV10FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_QianchuanAwemeReportOrderGetV10FilteringExternalAction         QianchuanAwemeReportOrderGetV10FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_QianchuanAwemeReportOrderGetV10FilteringExternalAction             QianchuanAwemeReportOrderGetV10FilteringExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_QianchuanAwemeReportOrderGetV10FilteringExternalAction              QianchuanAwemeReportOrderGetV10FilteringExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_QianchuanAwemeReportOrderGetV10FilteringExternalAction                   QianchuanAwemeReportOrderGetV10FilteringExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_SHOPPING_QianchuanAwemeReportOrderGetV10FilteringExternalAction                      QianchuanAwemeReportOrderGetV10FilteringExternalAction = "AD_CONVERT_TYPE_SHOPPING"
	AD_CONVERT_TYPE_VIDEO_ROI_QianchuanAwemeReportOrderGetV10FilteringExternalAction                     QianchuanAwemeReportOrderGetV10FilteringExternalAction = "AD_CONVERT_TYPE_VIDEO_ROI"
	AD_CONVERT_TYPE_LIVE_EFFECTIVE_EFFECTIVE_VIEW_QianchuanAwemeReportOrderGetV10FilteringExternalAction QianchuanAwemeReportOrderGetV10FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_EFFECTIVE_EFFECTIVE_VIEW"
)

// All allowed values of QianchuanAwemeReportOrderGetV10FilteringExternalAction enum
var AllowedQianchuanAwemeReportOrderGetV10FilteringExternalActionEnumValues = []QianchuanAwemeReportOrderGetV10FilteringExternalAction{
	"AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_LIVE_ENTER_ACTION",
	"AD_CONVERT_TYPE_LIVE_ROI",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY",
	"AD_CONVERT_TYPE_NEW_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_MUST_BUY",
	"AD_CONVERT_TYPE_SHOPPING",
	"AD_CONVERT_TYPE_VIDEO_ROI",
	"AD_CONVERT_TYPE_LIVE_EFFECTIVE_EFFECTIVE_VIEW",
}

// NewQianchuanAwemeReportOrderGetV10FilteringExternalActionFromValue returns a pointer to a valid QianchuanAwemeReportOrderGetV10FilteringExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeReportOrderGetV10FilteringExternalActionFromValue(v string) (*QianchuanAwemeReportOrderGetV10FilteringExternalAction, error) {
	ev := QianchuanAwemeReportOrderGetV10FilteringExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeReportOrderGetV10FilteringExternalAction: valid values are %v", v, AllowedQianchuanAwemeReportOrderGetV10FilteringExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeReportOrderGetV10FilteringExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeReportOrderGetV10FilteringExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_report_order_get_v1.0_filtering_external_action value
func (v QianchuanAwemeReportOrderGetV10FilteringExternalAction) Ptr() *QianchuanAwemeReportOrderGetV10FilteringExternalAction {
	return &v
}
