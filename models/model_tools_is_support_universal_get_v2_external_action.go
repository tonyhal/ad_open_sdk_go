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

// ToolsIsSupportUniversalGetV2ExternalAction
type ToolsIsSupportUniversalGetV2ExternalAction string

// List of tools_is_support_universal_get_v2_external_action
const (
	AD_CONVERT_PAGE_VIEW_ToolsIsSupportUniversalGetV2ExternalAction                        ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_PAGE_VIEW"
	AD_CONVERT_TYPE_ACTIVE_ToolsIsSupportUniversalGetV2ExternalAction                      ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_ACTIVE"
	AD_CONVERT_TYPE_ACTIVE_REGISTER_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	AD_CONVERT_TYPE_ANCHOR_CLICK_ToolsIsSupportUniversalGetV2ExternalAction                ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_ANCHOR_CLICK"
	AD_CONVERT_TYPE_APPLET_CLICK_ToolsIsSupportUniversalGetV2ExternalAction                ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_APPLET_CLICK"
	AD_CONVERT_TYPE_APP_CART_ToolsIsSupportUniversalGetV2ExternalAction                    ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_APP_CART"
	AD_CONVERT_TYPE_APP_DETAIL_UV_ToolsIsSupportUniversalGetV2ExternalAction               ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_APP_DETAIL_UV"
	AD_CONVERT_TYPE_APP_ORDER_ToolsIsSupportUniversalGetV2ExternalAction                   ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_APP_ORDER"
	AD_CONVERT_TYPE_APP_PAY_ToolsIsSupportUniversalGetV2ExternalAction                     ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_APP_PAY"
	AD_CONVERT_TYPE_APP_UV_ToolsIsSupportUniversalGetV2ExternalAction                      ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_APP_UV"
	AD_CONVERT_TYPE_AUTHORIZATION_ToolsIsSupportUniversalGetV2ExternalAction               ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_AUTHORIZATION"
	AD_CONVERT_TYPE_BANKCARD_INFORMATION_ToolsIsSupportUniversalGetV2ExternalAction        ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_BANKCARD_INFORMATION"
	AD_CONVERT_TYPE_BOOST_ToolsIsSupportUniversalGetV2ExternalAction                       ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_BOOST"
	AD_CONVERT_TYPE_BUTTON_ToolsIsSupportUniversalGetV2ExternalAction                      ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_BUTTON"
	AD_CONVERT_TYPE_CERTIFICATION_INFORMATION_ToolsIsSupportUniversalGetV2ExternalAction   ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CERTIFICATION_INFORMATION"
	AD_CONVERT_TYPE_CLICK_CALL_DY_ToolsIsSupportUniversalGetV2ExternalAction               ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CLICK_CALL_DY"
	AD_CONVERT_TYPE_CLICK_DOWNLOAD_ToolsIsSupportUniversalGetV2ExternalAction              ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CLICK_DOWNLOAD"
	AD_CONVERT_TYPE_CLICK_LANDING_PAGE_ToolsIsSupportUniversalGetV2ExternalAction          ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CLICK_LANDING_PAGE"
	AD_CONVERT_TYPE_CLICK_NUM_ToolsIsSupportUniversalGetV2ExternalAction                   ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CLICK_NUM"
	AD_CONVERT_TYPE_CLICK_SHOPWINDOW_ToolsIsSupportUniversalGetV2ExternalAction            ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CLICK_SHOPWINDOW"
	AD_CONVERT_TYPE_CLICK_WEBSITE_ToolsIsSupportUniversalGetV2ExternalAction               ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CLICK_WEBSITE"
	AD_CONVERT_TYPE_CLUE_CONFIRM_ToolsIsSupportUniversalGetV2ExternalAction                ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CLUE_CONFIRM"
	AD_CONVERT_TYPE_CLUE_HIGH_INTENTION_ToolsIsSupportUniversalGetV2ExternalAction         ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CLUE_HIGH_INTENTION"
	AD_CONVERT_TYPE_CLUE_INTERFLOW_ToolsIsSupportUniversalGetV2ExternalAction              ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CLUE_INTERFLOW"
	AD_CONVERT_TYPE_CLUE_PAY_SUCCEED_ToolsIsSupportUniversalGetV2ExternalAction            ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CLUE_PAY_SUCCEED"
	AD_CONVERT_TYPE_COMMENT_ACTION_ToolsIsSupportUniversalGetV2ExternalAction              ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_COMMENT_ACTION"
	AD_CONVERT_TYPE_COMMODITY_CLICK_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_COMMODITY_CLICK"
	AD_CONVERT_TYPE_CONSULT_ToolsIsSupportUniversalGetV2ExternalAction                     ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CONSULT"
	AD_CONVERT_TYPE_CONSULT_CLUE_ToolsIsSupportUniversalGetV2ExternalAction                ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CONSULT_CLUE"
	AD_CONVERT_TYPE_CONSULT_EFFECTIVE_ToolsIsSupportUniversalGetV2ExternalAction           ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CONSULT_EFFECTIVE"
	AD_CONVERT_TYPE_COUPON_ToolsIsSupportUniversalGetV2ExternalAction                      ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_COUPON"
	AD_CONVERT_TYPE_CREATE_GAMEROLE_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CREATE_GAMEROLE"
	AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE_ToolsIsSupportUniversalGetV2ExternalAction          ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE"
	AD_CONVERT_TYPE_DEEP_PURCHASE_ToolsIsSupportUniversalGetV2ExternalAction               ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_DEEP_PURCHASE"
	AD_CONVERT_TYPE_DIALBACK_ToolsIsSupportUniversalGetV2ExternalAction                    ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_DIALBACK"
	AD_CONVERT_TYPE_DIALBACK_CONFIRM_ToolsIsSupportUniversalGetV2ExternalAction            ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONFIRM"
	AD_CONVERT_TYPE_DIALBACK_CONNECT_ToolsIsSupportUniversalGetV2ExternalAction            ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONNECT"
	AD_CONVERT_TYPE_DOWNLOAD_FINISH_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_FINISH"
	AD_CONVERT_TYPE_DOWNLOAD_START_ToolsIsSupportUniversalGetV2ExternalAction              ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_START"
	AD_CONVERT_TYPE_EFFECTIVE_COPY_ToolsIsSupportUniversalGetV2ExternalAction              ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_COPY"
	AD_CONVERT_TYPE_EFFECTIVE_PLAY_ToolsIsSupportUniversalGetV2ExternalAction              ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_PLAY"
	AD_CONVERT_TYPE_ENTER_HOMEPAGE_ToolsIsSupportUniversalGetV2ExternalAction              ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_ENTER_HOMEPAGE"
	AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE_ToolsIsSupportUniversalGetV2ExternalAction          ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_FIRST_RENTAL_ORDER_ToolsIsSupportUniversalGetV2ExternalAction          ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_FIRST_RENTAL_ORDER"
	AD_CONVERT_TYPE_FOLLOW_ACTION_ToolsIsSupportUniversalGetV2ExternalAction               ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_FOLLOW_ACTION"
	AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT_ToolsIsSupportUniversalGetV2ExternalAction        ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT"
	AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER_ToolsIsSupportUniversalGetV2ExternalAction           ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER"
	AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH_ToolsIsSupportUniversalGetV2ExternalAction    ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH"
	AD_CONVERT_TYPE_FORM_ToolsIsSupportUniversalGetV2ExternalAction                        ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_FORM"
	AD_CONVERT_TYPE_FORM_ANSWER_ToolsIsSupportUniversalGetV2ExternalAction                 ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_FORM_ANSWER"
	AD_CONVERT_TYPE_FORM_CONNECT_ToolsIsSupportUniversalGetV2ExternalAction                ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_FORM_CONNECT"
	AD_CONVERT_TYPE_FORM_DEEP_ToolsIsSupportUniversalGetV2ExternalAction                   ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_FORM_DEEP"
	AD_CONVERT_TYPE_GAME_ADDICTION_ToolsIsSupportUniversalGetV2ExternalAction              ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_GAME_ADDICTION"
	AD_CONVERT_TYPE_HIGH_VALUE_CLUE_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_HIGH_VALUE_CLUE"
	AD_CONVERT_TYPE_IDCARD_INFORMATION_ToolsIsSupportUniversalGetV2ExternalAction          ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_IDCARD_INFORMATION"
	AD_CONVERT_TYPE_INSTALL_FINISH_ToolsIsSupportUniversalGetV2ExternalAction              ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_INSTALL_FINISH"
	AD_CONVERT_TYPE_INTENTION_CLUE_ToolsIsSupportUniversalGetV2ExternalAction              ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_INTENTION_CLUE"
	AD_CONVERT_TYPE_INTERACTION_ToolsIsSupportUniversalGetV2ExternalAction                 ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_INTERACTION"
	AD_CONVERT_TYPE_INVALID_CLUE_ToolsIsSupportUniversalGetV2ExternalAction                ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_INVALID_CLUE"
	AD_CONVERT_TYPE_IPU_QUALIFY_ToolsIsSupportUniversalGetV2ExternalAction                 ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_IPU_QUALIFY"
	AD_CONVERT_TYPE_LIKE_ACTION_ToolsIsSupportUniversalGetV2ExternalAction                 ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIKE_ACTION"
	AD_CONVERT_TYPE_LINK_ACTION_ToolsIsSupportUniversalGetV2ExternalAction                 ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LINK_ACTION"
	AD_CONVERT_TYPE_LIVE_APPOINTMENT_ToolsIsSupportUniversalGetV2ExternalAction            ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_APPOINTMENT"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_ToolsIsSupportUniversalGetV2ExternalAction   ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_ToolsIsSupportUniversalGetV2ExternalAction         ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK_ToolsIsSupportUniversalGetV2ExternalAction        ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_ToolsIsSupportUniversalGetV2ExternalAction           ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_FANS_ACTION_ToolsIsSupportUniversalGetV2ExternalAction            ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_FANS_ACTION"
	AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON_ToolsIsSupportUniversalGetV2ExternalAction          ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON"
	AD_CONVERT_TYPE_LIVE_GIFT_ACTION_ToolsIsSupportUniversalGetV2ExternalAction            ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_GIFT_ACTION"
	AD_CONVERT_TYPE_LIVE_HOMEPAGE_ToolsIsSupportUniversalGetV2ExternalAction               ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_LIVE_JOIN_GROUP_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_JOIN_GROUP"
	AD_CONVERT_TYPE_LIVE_NATIVE_ACITON_ToolsIsSupportUniversalGetV2ExternalAction          ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_NATIVE_ACITON"
	AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION_ToolsIsSupportUniversalGetV2ExternalAction ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION"
	AD_CONVERT_TYPE_LIVE_STAY_TIME_ToolsIsSupportUniversalGetV2ExternalAction              ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_STAY_TIME"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_ToolsIsSupportUniversalGetV2ExternalAction    ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_ToolsIsSupportUniversalGetV2ExternalAction       ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_LOAN_ToolsIsSupportUniversalGetV2ExternalAction                        ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LOAN"
	AD_CONVERT_TYPE_LOAN_COMPLETION_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LOAN_COMPLETION"
	AD_CONVERT_TYPE_LOAN_CREDIT_ToolsIsSupportUniversalGetV2ExternalAction                 ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LOAN_CREDIT"
	AD_CONVERT_TYPE_LOCATION_ACTION_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LOCATION_ACTION"
	AD_CONVERT_TYPE_LOTTERY_ToolsIsSupportUniversalGetV2ExternalAction                     ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LOTTERY"
	AD_CONVERT_TYPE_LT_ROI_ToolsIsSupportUniversalGetV2ExternalAction                      ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_LT_ROI"
	AD_CONVERT_TYPE_MAP_SEARCH_ToolsIsSupportUniversalGetV2ExternalAction                  ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_MAP_SEARCH"
	AD_CONVERT_TYPE_MESSAGE_ToolsIsSupportUniversalGetV2ExternalAction                     ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_MESSAGE"
	AD_CONVERT_TYPE_MESSAGE_ACTION_ToolsIsSupportUniversalGetV2ExternalAction              ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_MESSAGE_ACTION"
	AD_CONVERT_TYPE_MESSAGE_CLICK_ToolsIsSupportUniversalGetV2ExternalAction               ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLICK"
	AD_CONVERT_TYPE_MESSAGE_CLUE_ToolsIsSupportUniversalGetV2ExternalAction                ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLUE"
	AD_CONVERT_TYPE_MESSAGE_INTERACTION_ToolsIsSupportUniversalGetV2ExternalAction         ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_MESSAGE_INTERACTION"
	AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP_ToolsIsSupportUniversalGetV2ExternalAction          ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP"
	AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS_ToolsIsSupportUniversalGetV2ExternalAction       ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS"
	AD_CONVERT_TYPE_MESSAGE_SERVICE_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_MESSAGE_SERVICE"
	AD_CONVERT_TYPE_MULTIPLE_ToolsIsSupportUniversalGetV2ExternalAction                    ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_MULTIPLE"
	AD_CONVERT_TYPE_NATIVE_ACTION_ToolsIsSupportUniversalGetV2ExternalAction               ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_NATIVE_ACTION"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_ToolsIsSupportUniversalGetV2ExternalAction           ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_NEXT_DAY_OPEN_ToolsIsSupportUniversalGetV2ExternalAction               ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_NOTIFY_DOWNLOAD_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_NOTIFY_DOWNLOAD"
	AD_CONVERT_TYPE_NOTIFY_FORM_ToolsIsSupportUniversalGetV2ExternalAction                 ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_NOTIFY_FORM"
	AD_CONVERT_TYPE_OTHER_ToolsIsSupportUniversalGetV2ExternalAction                       ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_OTHER"
	AD_CONVERT_TYPE_OTO_PAY_ToolsIsSupportUniversalGetV2ExternalAction                     ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_OTO_PAY"
	AD_CONVERT_TYPE_OTO_STAY_TIME_ToolsIsSupportUniversalGetV2ExternalAction               ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_OTO_STAY_TIME"
	AD_CONVERT_TYPE_PAID_CLUE_ToolsIsSupportUniversalGetV2ExternalAction                   ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_PAID_CLUE"
	AD_CONVERT_TYPE_PAY_ToolsIsSupportUniversalGetV2ExternalAction                         ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_PAY"
	AD_CONVERT_TYPE_PERSONAL_INFORMATION_ToolsIsSupportUniversalGetV2ExternalAction        ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_PERSONAL_INFORMATION"
	AD_CONVERT_TYPE_PHONE_ToolsIsSupportUniversalGetV2ExternalAction                       ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_PHONE"
	AD_CONVERT_TYPE_PHONE_CONFIRM_ToolsIsSupportUniversalGetV2ExternalAction               ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_PHONE_CONFIRM"
	AD_CONVERT_TYPE_PHONE_CONNECT_ToolsIsSupportUniversalGetV2ExternalAction               ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_PHONE_CONNECT"
	AD_CONVERT_TYPE_PHONE_EFFECTIVE_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_PHONE_EFFECTIVE"
	AD_CONVERT_TYPE_POI_ADDRESS_CLICK_ToolsIsSupportUniversalGetV2ExternalAction           ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_POI_ADDRESS_CLICK"
	AD_CONVERT_TYPE_POI_COLLECT_ToolsIsSupportUniversalGetV2ExternalAction                 ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_POI_COLLECT"
	AD_CONVERT_TYPE_POI_MULTIPLE_ToolsIsSupportUniversalGetV2ExternalAction                ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_POI_MULTIPLE"
	AD_CONVERT_TYPE_PREMIUM_PAYMENT_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_PREMIUM_PAYMENT"
	AD_CONVERT_TYPE_PREMIUM_ROI_ToolsIsSupportUniversalGetV2ExternalAction                 ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_PREMIUM_ROI"
	AD_CONVERT_TYPE_PRE_LOAN_CREDIT_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_PRE_LOAN_CREDIT"
	AD_CONVERT_TYPE_PURCHASE_OF_GOODS_ToolsIsSupportUniversalGetV2ExternalAction           ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_PURCHASE_OF_GOODS"
	AD_CONVERT_TYPE_PURCHASE_ROI_ToolsIsSupportUniversalGetV2ExternalAction                ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_ToolsIsSupportUniversalGetV2ExternalAction            ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_ToolsIsSupportUniversalGetV2ExternalAction                 ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_QQ_ToolsIsSupportUniversalGetV2ExternalAction                          ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_QQ"
	AD_CONVERT_TYPE_REDIRECT_ToolsIsSupportUniversalGetV2ExternalAction                    ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_REDIRECT"
	AD_CONVERT_TYPE_REDIRECT_TO_SHOP_ToolsIsSupportUniversalGetV2ExternalAction            ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_SHOP"
	AD_CONVERT_TYPE_REDIRECT_TO_STORE_ToolsIsSupportUniversalGetV2ExternalAction           ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_STORE"
	AD_CONVERT_TYPE_RESERVATION_ToolsIsSupportUniversalGetV2ExternalAction                 ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_RESERVATION"
	AD_CONVERT_TYPE_RSS_ToolsIsSupportUniversalGetV2ExternalAction                         ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_RSS"
	AD_CONVERT_TYPE_SALES_LEAD_ToolsIsSupportUniversalGetV2ExternalAction                  ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_SALES_LEAD"
	AD_CONVERT_TYPE_SHARE_ACTION_ToolsIsSupportUniversalGetV2ExternalAction                ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_SHARE_ACTION"
	AD_CONVERT_TYPE_SHOPPING_ToolsIsSupportUniversalGetV2ExternalAction                    ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_SHOPPING"
	AD_CONVERT_TYPE_SHOPPING_ACTION_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_SHOPPING_ACTION"
	AD_CONVERT_TYPE_SHOPPING_CART_ToolsIsSupportUniversalGetV2ExternalAction               ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_SHOPPING_CART"
	AD_CONVERT_TYPE_SHOW_OFF_NUM_ToolsIsSupportUniversalGetV2ExternalAction                ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_SHOW_OFF_NUM"
	AD_CONVERT_TYPE_STAY_TIME_ToolsIsSupportUniversalGetV2ExternalAction                   ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_STAY_TIME"
	AD_CONVERT_TYPE_SUBMIT_CERTIFICATION_ToolsIsSupportUniversalGetV2ExternalAction        ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_SUBMIT_CERTIFICATION"
	AD_CONVERT_TYPE_SUCCESSORDER_ACTION_ToolsIsSupportUniversalGetV2ExternalAction         ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_UG_ROI_ToolsIsSupportUniversalGetV2ExternalAction                      ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_UG_ROI"
	AD_CONVERT_TYPE_UPDATE_LEVEL_ToolsIsSupportUniversalGetV2ExternalAction                ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_UPDATE_LEVEL"
	AD_CONVERT_TYPE_VIEW_ToolsIsSupportUniversalGetV2ExternalAction                        ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_VIEW"
	AD_CONVERT_TYPE_VOTE_ToolsIsSupportUniversalGetV2ExternalAction                        ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_VOTE"
	AD_CONVERT_TYPE_WECHAT_ToolsIsSupportUniversalGetV2ExternalAction                      ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_WECHAT"
	AD_CONVERT_TYPE_WECHAT_PAY_ToolsIsSupportUniversalGetV2ExternalAction                  ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_WECHAT_PAY"
	AD_CONVERT_TYPE_WECHAT_REGISTER_ToolsIsSupportUniversalGetV2ExternalAction             ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_WECHAT_REGISTER"
	AD_CONVERT_TYPE_XPATH_ToolsIsSupportUniversalGetV2ExternalAction                       ToolsIsSupportUniversalGetV2ExternalAction = "AD_CONVERT_TYPE_XPATH"
)

// All allowed values of ToolsIsSupportUniversalGetV2ExternalAction enum
var AllowedToolsIsSupportUniversalGetV2ExternalActionEnumValues = []ToolsIsSupportUniversalGetV2ExternalAction{
	"AD_CONVERT_PAGE_VIEW",
	"AD_CONVERT_TYPE_ACTIVE",
	"AD_CONVERT_TYPE_ACTIVE_REGISTER",
	"AD_CONVERT_TYPE_ANCHOR_CLICK",
	"AD_CONVERT_TYPE_APPLET_CLICK",
	"AD_CONVERT_TYPE_APP_CART",
	"AD_CONVERT_TYPE_APP_DETAIL_UV",
	"AD_CONVERT_TYPE_APP_ORDER",
	"AD_CONVERT_TYPE_APP_PAY",
	"AD_CONVERT_TYPE_APP_UV",
	"AD_CONVERT_TYPE_AUTHORIZATION",
	"AD_CONVERT_TYPE_BANKCARD_INFORMATION",
	"AD_CONVERT_TYPE_BOOST",
	"AD_CONVERT_TYPE_BUTTON",
	"AD_CONVERT_TYPE_CERTIFICATION_INFORMATION",
	"AD_CONVERT_TYPE_CLICK_CALL_DY",
	"AD_CONVERT_TYPE_CLICK_DOWNLOAD",
	"AD_CONVERT_TYPE_CLICK_LANDING_PAGE",
	"AD_CONVERT_TYPE_CLICK_NUM",
	"AD_CONVERT_TYPE_CLICK_SHOPWINDOW",
	"AD_CONVERT_TYPE_CLICK_WEBSITE",
	"AD_CONVERT_TYPE_CLUE_CONFIRM",
	"AD_CONVERT_TYPE_CLUE_HIGH_INTENTION",
	"AD_CONVERT_TYPE_CLUE_INTERFLOW",
	"AD_CONVERT_TYPE_CLUE_PAY_SUCCEED",
	"AD_CONVERT_TYPE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_COMMODITY_CLICK",
	"AD_CONVERT_TYPE_CONSULT",
	"AD_CONVERT_TYPE_CONSULT_CLUE",
	"AD_CONVERT_TYPE_CONSULT_EFFECTIVE",
	"AD_CONVERT_TYPE_COUPON",
	"AD_CONVERT_TYPE_CREATE_GAMEROLE",
	"AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE",
	"AD_CONVERT_TYPE_DEEP_PURCHASE",
	"AD_CONVERT_TYPE_DIALBACK",
	"AD_CONVERT_TYPE_DIALBACK_CONFIRM",
	"AD_CONVERT_TYPE_DIALBACK_CONNECT",
	"AD_CONVERT_TYPE_DOWNLOAD_FINISH",
	"AD_CONVERT_TYPE_DOWNLOAD_START",
	"AD_CONVERT_TYPE_EFFECTIVE_COPY",
	"AD_CONVERT_TYPE_EFFECTIVE_PLAY",
	"AD_CONVERT_TYPE_ENTER_HOMEPAGE",
	"AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE",
	"AD_CONVERT_TYPE_FIRST_RENTAL_ORDER",
	"AD_CONVERT_TYPE_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT",
	"AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER",
	"AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH",
	"AD_CONVERT_TYPE_FORM",
	"AD_CONVERT_TYPE_FORM_ANSWER",
	"AD_CONVERT_TYPE_FORM_CONNECT",
	"AD_CONVERT_TYPE_FORM_DEEP",
	"AD_CONVERT_TYPE_GAME_ADDICTION",
	"AD_CONVERT_TYPE_HIGH_VALUE_CLUE",
	"AD_CONVERT_TYPE_IDCARD_INFORMATION",
	"AD_CONVERT_TYPE_INSTALL_FINISH",
	"AD_CONVERT_TYPE_INTENTION_CLUE",
	"AD_CONVERT_TYPE_INTERACTION",
	"AD_CONVERT_TYPE_INVALID_CLUE",
	"AD_CONVERT_TYPE_IPU_QUALIFY",
	"AD_CONVERT_TYPE_LIKE_ACTION",
	"AD_CONVERT_TYPE_LINK_ACTION",
	"AD_CONVERT_TYPE_LIVE_APPOINTMENT",
	"AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK",
	"AD_CONVERT_TYPE_LIVE_ENTER_ACTION",
	"AD_CONVERT_TYPE_LIVE_FANS_ACTION",
	"AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON",
	"AD_CONVERT_TYPE_LIVE_GIFT_ACTION",
	"AD_CONVERT_TYPE_LIVE_HOMEPAGE",
	"AD_CONVERT_TYPE_LIVE_JOIN_GROUP",
	"AD_CONVERT_TYPE_LIVE_NATIVE_ACITON",
	"AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION",
	"AD_CONVERT_TYPE_LIVE_STAY_TIME",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY",
	"AD_CONVERT_TYPE_LOAN",
	"AD_CONVERT_TYPE_LOAN_COMPLETION",
	"AD_CONVERT_TYPE_LOAN_CREDIT",
	"AD_CONVERT_TYPE_LOCATION_ACTION",
	"AD_CONVERT_TYPE_LOTTERY",
	"AD_CONVERT_TYPE_LT_ROI",
	"AD_CONVERT_TYPE_MAP_SEARCH",
	"AD_CONVERT_TYPE_MESSAGE",
	"AD_CONVERT_TYPE_MESSAGE_ACTION",
	"AD_CONVERT_TYPE_MESSAGE_CLICK",
	"AD_CONVERT_TYPE_MESSAGE_CLUE",
	"AD_CONVERT_TYPE_MESSAGE_INTERACTION",
	"AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP",
	"AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS",
	"AD_CONVERT_TYPE_MESSAGE_SERVICE",
	"AD_CONVERT_TYPE_MULTIPLE",
	"AD_CONVERT_TYPE_NATIVE_ACTION",
	"AD_CONVERT_TYPE_NEW_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_NEXT_DAY_OPEN",
	"AD_CONVERT_TYPE_NOTIFY_DOWNLOAD",
	"AD_CONVERT_TYPE_NOTIFY_FORM",
	"AD_CONVERT_TYPE_OTHER",
	"AD_CONVERT_TYPE_OTO_PAY",
	"AD_CONVERT_TYPE_OTO_STAY_TIME",
	"AD_CONVERT_TYPE_PAID_CLUE",
	"AD_CONVERT_TYPE_PAY",
	"AD_CONVERT_TYPE_PERSONAL_INFORMATION",
	"AD_CONVERT_TYPE_PHONE",
	"AD_CONVERT_TYPE_PHONE_CONFIRM",
	"AD_CONVERT_TYPE_PHONE_CONNECT",
	"AD_CONVERT_TYPE_PHONE_EFFECTIVE",
	"AD_CONVERT_TYPE_POI_ADDRESS_CLICK",
	"AD_CONVERT_TYPE_POI_COLLECT",
	"AD_CONVERT_TYPE_POI_MULTIPLE",
	"AD_CONVERT_TYPE_PREMIUM_PAYMENT",
	"AD_CONVERT_TYPE_PREMIUM_ROI",
	"AD_CONVERT_TYPE_PRE_LOAN_CREDIT",
	"AD_CONVERT_TYPE_PURCHASE_OF_GOODS",
	"AD_CONVERT_TYPE_PURCHASE_ROI",
	"AD_CONVERT_TYPE_QC_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_MUST_BUY",
	"AD_CONVERT_TYPE_QQ",
	"AD_CONVERT_TYPE_REDIRECT",
	"AD_CONVERT_TYPE_REDIRECT_TO_SHOP",
	"AD_CONVERT_TYPE_REDIRECT_TO_STORE",
	"AD_CONVERT_TYPE_RESERVATION",
	"AD_CONVERT_TYPE_RSS",
	"AD_CONVERT_TYPE_SALES_LEAD",
	"AD_CONVERT_TYPE_SHARE_ACTION",
	"AD_CONVERT_TYPE_SHOPPING",
	"AD_CONVERT_TYPE_SHOPPING_ACTION",
	"AD_CONVERT_TYPE_SHOPPING_CART",
	"AD_CONVERT_TYPE_SHOW_OFF_NUM",
	"AD_CONVERT_TYPE_STAY_TIME",
	"AD_CONVERT_TYPE_SUBMIT_CERTIFICATION",
	"AD_CONVERT_TYPE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_UG_ROI",
	"AD_CONVERT_TYPE_UPDATE_LEVEL",
	"AD_CONVERT_TYPE_VIEW",
	"AD_CONVERT_TYPE_VOTE",
	"AD_CONVERT_TYPE_WECHAT",
	"AD_CONVERT_TYPE_WECHAT_PAY",
	"AD_CONVERT_TYPE_WECHAT_REGISTER",
	"AD_CONVERT_TYPE_XPATH",
}

// NewToolsIsSupportUniversalGetV2ExternalActionFromValue returns a pointer to a valid ToolsIsSupportUniversalGetV2ExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsIsSupportUniversalGetV2ExternalActionFromValue(v string) (*ToolsIsSupportUniversalGetV2ExternalAction, error) {
	ev := ToolsIsSupportUniversalGetV2ExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsIsSupportUniversalGetV2ExternalAction: valid values are %v", v, AllowedToolsIsSupportUniversalGetV2ExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsIsSupportUniversalGetV2ExternalAction) IsValid() bool {
	for _, existing := range AllowedToolsIsSupportUniversalGetV2ExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_is_support_universal_get_v2_external_action value
func (v ToolsIsSupportUniversalGetV2ExternalAction) Ptr() *ToolsIsSupportUniversalGetV2ExternalAction {
	return &v
}
