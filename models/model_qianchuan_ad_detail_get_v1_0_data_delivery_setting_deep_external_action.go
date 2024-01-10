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

// QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction
type QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction string

// List of qianchuan_ad_detail_get_v1.0_data_delivery_setting_deep_external_action
const (
	AD_CONVERT_TYPE_CARD_ACTIVE_QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction                  QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_CARD_ACTIVE"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction    QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction          QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction            QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_PAY_ROI_QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction                 QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_PAY_ROI"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction     QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction        QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7_DAYS_QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction            QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction             QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction                  QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_SHOPPING_QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction                     QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_SHOPPING"
)

// All allowed values of QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction enum
var AllowedQianchuanAdDetailGetV10DataDeliverySettingDeepExternalActionEnumValues = []QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction{
	"AD_CONVERT_TYPE_CARD_ACTIVE",
	"AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_LIVE_ENTER_ACTION",
	"AD_CONVERT_TYPE_LIVE_PAY_ROI",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS",
	"AD_CONVERT_TYPE_NEW_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_MUST_BUY",
	"AD_CONVERT_TYPE_SHOPPING",
}

// NewQianchuanAdDetailGetV10DataDeliverySettingDeepExternalActionFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataDeliverySettingDeepExternalActionFromValue(v string) (*QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction, error) {
	ev := QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataDeliverySettingDeepExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataDeliverySettingDeepExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_delivery_setting_deep_external_action value
func (v QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction) Ptr() *QianchuanAdDetailGetV10DataDeliverySettingDeepExternalAction {
	return &v
}
