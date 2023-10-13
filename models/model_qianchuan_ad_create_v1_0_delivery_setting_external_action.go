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

// QianchuanAdCreateV10DeliverySettingExternalAction
type QianchuanAdCreateV10DeliverySettingExternalAction string

// List of qianchuan_ad_create_v1.0_delivery_setting_external_action
const (
	AD_CONVERT_TYPE_CARD_ACTIVE_QianchuanAdCreateV10DeliverySettingExternalAction                  QianchuanAdCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_CARD_ACTIVE"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_QianchuanAdCreateV10DeliverySettingExternalAction    QianchuanAdCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_QianchuanAdCreateV10DeliverySettingExternalAction          QianchuanAdCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_QianchuanAdCreateV10DeliverySettingExternalAction            QianchuanAdCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_PAY_ROI_QianchuanAdCreateV10DeliverySettingExternalAction                 QianchuanAdCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_PAY_ROI"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_QianchuanAdCreateV10DeliverySettingExternalAction     QianchuanAdCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_QianchuanAdCreateV10DeliverySettingExternalAction        QianchuanAdCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7_DAYS_QianchuanAdCreateV10DeliverySettingExternalAction QianchuanAdCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_SETTLE_QianchuanAdCreateV10DeliverySettingExternalAction     QianchuanAdCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_SETTLE"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_QianchuanAdCreateV10DeliverySettingExternalAction            QianchuanAdCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_QianchuanAdCreateV10DeliverySettingExternalAction             QianchuanAdCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_QianchuanAdCreateV10DeliverySettingExternalAction                  QianchuanAdCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_SHOPPING_QianchuanAdCreateV10DeliverySettingExternalAction                     QianchuanAdCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_SHOPPING"
)

// All allowed values of QianchuanAdCreateV10DeliverySettingExternalAction enum
var AllowedQianchuanAdCreateV10DeliverySettingExternalActionEnumValues = []QianchuanAdCreateV10DeliverySettingExternalAction{
	"AD_CONVERT_TYPE_CARD_ACTIVE",
	"AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_LIVE_ENTER_ACTION",
	"AD_CONVERT_TYPE_LIVE_PAY_ROI",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_SETTLE",
	"AD_CONVERT_TYPE_NEW_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_MUST_BUY",
	"AD_CONVERT_TYPE_SHOPPING",
}

// NewQianchuanAdCreateV10DeliverySettingExternalActionFromValue returns a pointer to a valid QianchuanAdCreateV10DeliverySettingExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10DeliverySettingExternalActionFromValue(v string) (*QianchuanAdCreateV10DeliverySettingExternalAction, error) {
	ev := QianchuanAdCreateV10DeliverySettingExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10DeliverySettingExternalAction: valid values are %v", v, AllowedQianchuanAdCreateV10DeliverySettingExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10DeliverySettingExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10DeliverySettingExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_delivery_setting_external_action value
func (v QianchuanAdCreateV10DeliverySettingExternalAction) Ptr() *QianchuanAdCreateV10DeliverySettingExternalAction {
	return &v
}
