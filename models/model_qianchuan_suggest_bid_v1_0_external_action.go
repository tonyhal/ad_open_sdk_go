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

// QianchuanSuggestBidV10ExternalAction
type QianchuanSuggestBidV10ExternalAction string

// List of qianchuan_suggest_bid_v1.0_external_action
const (
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_QianchuanSuggestBidV10ExternalAction QianchuanSuggestBidV10ExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_QianchuanSuggestBidV10ExternalAction       QianchuanSuggestBidV10ExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_QianchuanSuggestBidV10ExternalAction         QianchuanSuggestBidV10ExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_QianchuanSuggestBidV10ExternalAction  QianchuanSuggestBidV10ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_QianchuanSuggestBidV10ExternalAction     QianchuanSuggestBidV10ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_QianchuanSuggestBidV10ExternalAction         QianchuanSuggestBidV10ExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_QianchuanSuggestBidV10ExternalAction          QianchuanSuggestBidV10ExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_QianchuanSuggestBidV10ExternalAction               QianchuanSuggestBidV10ExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_SHOPPING_QianchuanSuggestBidV10ExternalAction                  QianchuanSuggestBidV10ExternalAction = "AD_CONVERT_TYPE_SHOPPING"
)

// All allowed values of QianchuanSuggestBidV10ExternalAction enum
var AllowedQianchuanSuggestBidV10ExternalActionEnumValues = []QianchuanSuggestBidV10ExternalAction{
	"AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_LIVE_ENTER_ACTION",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY",
	"AD_CONVERT_TYPE_NEW_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_MUST_BUY",
	"AD_CONVERT_TYPE_SHOPPING",
}

// NewQianchuanSuggestBidV10ExternalActionFromValue returns a pointer to a valid QianchuanSuggestBidV10ExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanSuggestBidV10ExternalActionFromValue(v string) (*QianchuanSuggestBidV10ExternalAction, error) {
	ev := QianchuanSuggestBidV10ExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanSuggestBidV10ExternalAction: valid values are %v", v, AllowedQianchuanSuggestBidV10ExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanSuggestBidV10ExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanSuggestBidV10ExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_suggest_bid_v1.0_external_action value
func (v QianchuanSuggestBidV10ExternalAction) Ptr() *QianchuanSuggestBidV10ExternalAction {
	return &v
}
