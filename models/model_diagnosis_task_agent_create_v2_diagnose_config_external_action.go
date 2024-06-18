/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction
type DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction string

// List of diagnosis_task_agent_create_v2_diagnose_config_external_action
const (
	AD_APP_ACTIVATE_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction         DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_ACTIVATE"
	AD_APP_AUTH_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction             DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_AUTH"
	AD_APP_BOOK_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction             DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_BOOK"
	AD_APP_BUY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction              DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_BUY"
	AD_APP_CLICKS_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction           DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_CLICKS"
	AD_APP_DETAIL_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction           DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_DETAIL"
	AD_APP_DOWNLOADED_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction       DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_DOWNLOADED"
	AD_APP_INSTALLED_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction        DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_INSTALLED"
	AD_APP_KEY_BEHAVIOR_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction     DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_KEY_BEHAVIOR"
	AD_APP_ORDER_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction            DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_ORDER"
	AD_APP_PAY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction              DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_PAY"
	AD_APP_PRE_AUTH_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction         DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_PRE_AUTH"
	AD_APP_PRE_DOWNLOAD_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction     DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_PRE_DOWNLOAD"
	AD_APP_PUSH_ORDER_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction       DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_PUSH_ORDER"
	AD_APP_REGISTER_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction         DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_REGISTER"
	AD_APP_SHOW_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction             DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_SHOW"
	AD_APP_SUBMIT_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction           DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_SUBMIT"
	AD_APP_VIEW_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction             DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_APP_VIEW"
	AD_CLUE_AUTH_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction            DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_AUTH"
	AD_CLUE_BUTTON_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction          DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_BUTTON"
	AD_CLUE_CLICKS_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction          DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_CLICKS"
	AD_CLUE_CONFIRM_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction         DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_CONFIRM"
	AD_CLUE_CONSULT_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction         DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_CONSULT"
	AD_CLUE_CONSULT_MSG_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction     DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_CONSULT_MSG"
	AD_CLUE_COUPON_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction          DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_COUPON"
	AD_CLUE_CUSTOMER_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction        DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_CUSTOMER"
	AD_CLUE_CVT_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction             DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_CVT"
	AD_CLUE_DONE_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction            DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_DONE"
	AD_CLUE_FORM_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction            DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_FORM"
	AD_CLUE_FRIENDS_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction         DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_FRIENDS"
	AD_CLUE_INSURANCE_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction       DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_INSURANCE"
	AD_CLUE_INTENTION_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction       DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_INTENTION"
	AD_CLUE_INTENTION_FORM_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction  DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_INTENTION_FORM"
	AD_CLUE_INTENTION_TEL_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction   DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_INTENTION_TEL"
	AD_CLUE_MESSAGE_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction         DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_MESSAGE"
	AD_CLUE_MONEY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction           DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_MONEY"
	AD_CLUE_MSG_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction             DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_MSG"
	AD_CLUE_PAGE_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction            DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_PAGE"
	AD_CLUE_PAY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction             DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_PAY"
	AD_CLUE_PRE_AUTH_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction        DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_PRE_AUTH"
	AD_CLUE_PROTENTIAL_DEAL_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_PROTENTIAL_DEAL"
	AD_CLUE_PUSH_ORDER_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction      DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_PUSH_ORDER"
	AD_CLUE_REGISTER_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction        DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_REGISTER"
	AD_CLUE_SHOW_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction            DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_SHOW"
	AD_CLUE_SUBMIT_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction          DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_SUBMIT"
	AD_CLUE_TEL_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction             DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_TEL"
	AD_CLUE_TEL_CALL_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction        DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_TEL_CALL"
	AD_CLUE_WX_ADD_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction          DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_WX_ADD"
	AD_CLUE_WX_COPY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction         DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_WX_COPY"
	AD_CLUE_WX_MSG_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction          DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_CLUE_WX_MSG"
	AD_ECP_APP_BUY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction          DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_ECP_APP_BUY"
	AD_ECP_APP_DETAIL_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction       DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_ECP_APP_DETAIL"
	AD_ECP_APP_VIEW_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction         DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_ECP_APP_VIEW"
	AD_ECP_BUTTON_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction           DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_ECP_BUTTON"
	AD_ECP_INTEREST_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction         DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_ECP_INTEREST"
	AD_ECP_SHOP_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction             DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_ECP_SHOP"
	AD_ECP_SHOP_STAY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction        DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_ECP_SHOP_STAY"
	AD_MINIAPP_ACTIVATE_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction     DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_MINIAPP_ACTIVATE"
	AD_MINIAPP_KEY_BEHAVIOR_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_MINIAPP_KEY_BEHAVIOR"
	AD_MINIAPP_PAY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction          DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_MINIAPP_PAY"
	AD_MINIAPP_REGISTER_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction     DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_MINIAPP_REGISTER"
	AD_NATIVE_ACTIVATE_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction      DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_NATIVE_ACTIVATE"
	AD_NATIVE_CLICKS_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction        DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_NATIVE_CLICKS"
	AD_NATIVE_FANS_GROUP_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction    DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_NATIVE_FANS_GROUP"
	AD_NATIVE_FOLLOW_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction        DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_NATIVE_FOLLOW"
	AD_NATIVE_INTERACTIVE_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction   DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_NATIVE_INTERACTIVE"
	AD_NATIVE_LIVE_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction          DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_NATIVE_LIVE"
	AD_NATIVE_LIVE_DONATE_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction   DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_NATIVE_LIVE_DONATE"
	AD_NATIVE_LIVE_PAY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction      DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_NATIVE_LIVE_PAY"
	AD_NATIVE_LIVE_STAY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction     DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_NATIVE_LIVE_STAY"
	AD_NATIVE_LIVE_VIEW_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction     DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_NATIVE_LIVE_VIEW"
	AD_NATIVE_PAY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction           DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_NATIVE_PAY"
	AD_PRODUCT_ACTIVATE_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction     DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_PRODUCT_ACTIVATE"
	AD_PRODUCT_APP_BUY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction      DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_PRODUCT_APP_BUY"
	AD_PRODUCT_APP_DETAIL_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction   DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_PRODUCT_APP_DETAIL"
	AD_PRODUCT_APP_PAY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction      DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_PRODUCT_APP_PAY"
	AD_PRODUCT_APP_VIEW_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction     DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_PRODUCT_APP_VIEW"
	AD_PRODUCT_FORM_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction         DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_PRODUCT_FORM"
	AD_PRODUCT_KEY_BEHAVIOR_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_PRODUCT_KEY_BEHAVIOR"
	AD_PRODUCT_PAY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction          DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_PRODUCT_PAY"
	AD_TINYAPP_ACTIVATE_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction     DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_TINYAPP_ACTIVATE"
	AD_TINYAPP_KEY_BEHAVIOR_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_TINYAPP_KEY_BEHAVIOR"
	AD_TINYAPP_PAY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction          DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "AD_TINYAPP_PAY"
	QC_LIVE_BUY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction             DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_LIVE_BUY"
	QC_LIVE_CHECK_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction           DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_LIVE_CHECK"
	QC_LIVE_COMMENTS_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction        DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_LIVE_COMMENTS"
	QC_LIVE_DEAL_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction            DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_LIVE_DEAL"
	QC_LIVE_ENTRY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction           DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_LIVE_ENTRY"
	QC_LIVE_FANS_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction            DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_LIVE_FANS"
	QC_LIVE_HIT_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction             DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_LIVE_HIT"
	QC_LIVE_PRODUCT_CLICKS_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction  DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_LIVE_PRODUCT_CLICKS"
	QC_LIVE_ROI_CHECK_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction       DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_LIVE_ROI_CHECK"
	QC_LIVE_ROI_DEAL_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction        DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_LIVE_ROI_DEAL"
	QC_LIVE_ROI_QC_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction          DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_LIVE_ROI_QC"
	QC_PRODUCT_BUY_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction          DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_PRODUCT_BUY"
	QC_PRODUCT_COMMENTS_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction     DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_PRODUCT_COMMENTS"
	QC_PRODUCT_FANS_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction         DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_PRODUCT_FANS"
	QC_PRODUCT_INTEREST_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction     DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_PRODUCT_INTEREST"
	QC_PRODUCT_QC_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction           DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_PRODUCT_QC"
	QC_PRODUCT_ROI_DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction          DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction = "QC_PRODUCT_ROI"
)

// All allowed values of DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction enum
var AllowedDiagnosisTaskAgentCreateV2DiagnoseConfigExternalActionEnumValues = []DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction{
	"AD_APP_ACTIVATE",
	"AD_APP_AUTH",
	"AD_APP_BOOK",
	"AD_APP_BUY",
	"AD_APP_CLICKS",
	"AD_APP_DETAIL",
	"AD_APP_DOWNLOADED",
	"AD_APP_INSTALLED",
	"AD_APP_KEY_BEHAVIOR",
	"AD_APP_ORDER",
	"AD_APP_PAY",
	"AD_APP_PRE_AUTH",
	"AD_APP_PRE_DOWNLOAD",
	"AD_APP_PUSH_ORDER",
	"AD_APP_REGISTER",
	"AD_APP_SHOW",
	"AD_APP_SUBMIT",
	"AD_APP_VIEW",
	"AD_CLUE_AUTH",
	"AD_CLUE_BUTTON",
	"AD_CLUE_CLICKS",
	"AD_CLUE_CONFIRM",
	"AD_CLUE_CONSULT",
	"AD_CLUE_CONSULT_MSG",
	"AD_CLUE_COUPON",
	"AD_CLUE_CUSTOMER",
	"AD_CLUE_CVT",
	"AD_CLUE_DONE",
	"AD_CLUE_FORM",
	"AD_CLUE_FRIENDS",
	"AD_CLUE_INSURANCE",
	"AD_CLUE_INTENTION",
	"AD_CLUE_INTENTION_FORM",
	"AD_CLUE_INTENTION_TEL",
	"AD_CLUE_MESSAGE",
	"AD_CLUE_MONEY",
	"AD_CLUE_MSG",
	"AD_CLUE_PAGE",
	"AD_CLUE_PAY",
	"AD_CLUE_PRE_AUTH",
	"AD_CLUE_PROTENTIAL_DEAL",
	"AD_CLUE_PUSH_ORDER",
	"AD_CLUE_REGISTER",
	"AD_CLUE_SHOW",
	"AD_CLUE_SUBMIT",
	"AD_CLUE_TEL",
	"AD_CLUE_TEL_CALL",
	"AD_CLUE_WX_ADD",
	"AD_CLUE_WX_COPY",
	"AD_CLUE_WX_MSG",
	"AD_ECP_APP_BUY",
	"AD_ECP_APP_DETAIL",
	"AD_ECP_APP_VIEW",
	"AD_ECP_BUTTON",
	"AD_ECP_INTEREST",
	"AD_ECP_SHOP",
	"AD_ECP_SHOP_STAY",
	"AD_MINIAPP_ACTIVATE",
	"AD_MINIAPP_KEY_BEHAVIOR",
	"AD_MINIAPP_PAY",
	"AD_MINIAPP_REGISTER",
	"AD_NATIVE_ACTIVATE",
	"AD_NATIVE_CLICKS",
	"AD_NATIVE_FANS_GROUP",
	"AD_NATIVE_FOLLOW",
	"AD_NATIVE_INTERACTIVE",
	"AD_NATIVE_LIVE",
	"AD_NATIVE_LIVE_DONATE",
	"AD_NATIVE_LIVE_PAY",
	"AD_NATIVE_LIVE_STAY",
	"AD_NATIVE_LIVE_VIEW",
	"AD_NATIVE_PAY",
	"AD_PRODUCT_ACTIVATE",
	"AD_PRODUCT_APP_BUY",
	"AD_PRODUCT_APP_DETAIL",
	"AD_PRODUCT_APP_PAY",
	"AD_PRODUCT_APP_VIEW",
	"AD_PRODUCT_FORM",
	"AD_PRODUCT_KEY_BEHAVIOR",
	"AD_PRODUCT_PAY",
	"AD_TINYAPP_ACTIVATE",
	"AD_TINYAPP_KEY_BEHAVIOR",
	"AD_TINYAPP_PAY",
	"QC_LIVE_BUY",
	"QC_LIVE_CHECK",
	"QC_LIVE_COMMENTS",
	"QC_LIVE_DEAL",
	"QC_LIVE_ENTRY",
	"QC_LIVE_FANS",
	"QC_LIVE_HIT",
	"QC_LIVE_PRODUCT_CLICKS",
	"QC_LIVE_ROI_CHECK",
	"QC_LIVE_ROI_DEAL",
	"QC_LIVE_ROI_QC",
	"QC_PRODUCT_BUY",
	"QC_PRODUCT_COMMENTS",
	"QC_PRODUCT_FANS",
	"QC_PRODUCT_INTEREST",
	"QC_PRODUCT_QC",
	"QC_PRODUCT_ROI",
}

// NewDiagnosisTaskAgentCreateV2DiagnoseConfigExternalActionFromValue returns a pointer to a valid DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiagnosisTaskAgentCreateV2DiagnoseConfigExternalActionFromValue(v string) (*DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction, error) {
	ev := DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction: valid values are %v", v, AllowedDiagnosisTaskAgentCreateV2DiagnoseConfigExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction) IsValid() bool {
	for _, existing := range AllowedDiagnosisTaskAgentCreateV2DiagnoseConfigExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to diagnosis_task_agent_create_v2_diagnose_config_external_action value
func (v DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction) Ptr() *DiagnosisTaskAgentCreateV2DiagnoseConfigExternalAction {
	return &v
}