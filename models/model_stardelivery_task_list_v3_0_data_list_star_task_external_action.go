/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StardeliveryTaskListV30DataListStarTaskExternalAction
type StardeliveryTaskListV30DataListStarTaskExternalAction string

// List of stardelivery_task_list_v3.0_data_list_star_task_external_action
const (
	AD_CONVERT_PAGE_VIEW_StardeliveryTaskListV30DataListStarTaskExternalAction                        StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_PAGE_VIEW"
	AD_CONVERT_PHONE_CONNECT_StardeliveryTaskListV30DataListStarTaskExternalAction                    StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_PHONE_CONNECT"
	AD_CONVERT_TYPE_ACTIVE_StardeliveryTaskListV30DataListStarTaskExternalAction                      StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_ACTIVE"
	AD_CONVERT_TYPE_ACTIVE_REGISTER_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	AD_CONVERT_TYPE_ANCHOR_CLICK_StardeliveryTaskListV30DataListStarTaskExternalAction                StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_ANCHOR_CLICK"
	AD_CONVERT_TYPE_APPLET_CLICK_StardeliveryTaskListV30DataListStarTaskExternalAction                StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_APPLET_CLICK"
	AD_CONVERT_TYPE_APP_CART_StardeliveryTaskListV30DataListStarTaskExternalAction                    StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_APP_CART"
	AD_CONVERT_TYPE_APP_DETAIL_UV_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_APP_DETAIL_UV"
	AD_CONVERT_TYPE_APP_ORDER_StardeliveryTaskListV30DataListStarTaskExternalAction                   StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_APP_ORDER"
	AD_CONVERT_TYPE_APP_PAY_StardeliveryTaskListV30DataListStarTaskExternalAction                     StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_APP_PAY"
	AD_CONVERT_TYPE_APP_UV_StardeliveryTaskListV30DataListStarTaskExternalAction                      StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_APP_UV"
	AD_CONVERT_TYPE_ARPU0_StardeliveryTaskListV30DataListStarTaskExternalAction                       StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_ARPU0"
	AD_CONVERT_TYPE_AUTHORIZATION_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_AUTHORIZATION"
	AD_CONVERT_TYPE_BANKCARD_INFORMATION_StardeliveryTaskListV30DataListStarTaskExternalAction        StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_BANKCARD_INFORMATION"
	AD_CONVERT_TYPE_BOOST_StardeliveryTaskListV30DataListStarTaskExternalAction                       StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_BOOST"
	AD_CONVERT_TYPE_BUTTON_StardeliveryTaskListV30DataListStarTaskExternalAction                      StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_BUTTON"
	AD_CONVERT_TYPE_CERTIFICATION_INFORMATION_StardeliveryTaskListV30DataListStarTaskExternalAction   StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CERTIFICATION_INFORMATION"
	AD_CONVERT_TYPE_CLASS_StardeliveryTaskListV30DataListStarTaskExternalAction                       StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CLASS"
	AD_CONVERT_TYPE_CLICK_CALL_DY_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CLICK_CALL_DY"
	AD_CONVERT_TYPE_CLICK_DOWNLOAD_StardeliveryTaskListV30DataListStarTaskExternalAction              StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CLICK_DOWNLOAD"
	AD_CONVERT_TYPE_CLICK_LANDING_PAGE_StardeliveryTaskListV30DataListStarTaskExternalAction          StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CLICK_LANDING_PAGE"
	AD_CONVERT_TYPE_CLICK_NUM_StardeliveryTaskListV30DataListStarTaskExternalAction                   StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CLICK_NUM"
	AD_CONVERT_TYPE_CLICK_SHOPWINDOW_StardeliveryTaskListV30DataListStarTaskExternalAction            StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CLICK_SHOPWINDOW"
	AD_CONVERT_TYPE_CLICK_WEBSITE_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CLICK_WEBSITE"
	AD_CONVERT_TYPE_CLUE_CONFIRM_StardeliveryTaskListV30DataListStarTaskExternalAction                StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CLUE_CONFIRM"
	AD_CONVERT_TYPE_CLUE_HIGH_INTENTION_StardeliveryTaskListV30DataListStarTaskExternalAction         StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CLUE_HIGH_INTENTION"
	AD_CONVERT_TYPE_CLUE_INTERFLOW_StardeliveryTaskListV30DataListStarTaskExternalAction              StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CLUE_INTERFLOW"
	AD_CONVERT_TYPE_CLUE_PAY_SUCCEED_StardeliveryTaskListV30DataListStarTaskExternalAction            StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CLUE_PAY_SUCCEED"
	AD_CONVERT_TYPE_COMMENT_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction              StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_COMMENT_ACTION"
	AD_CONVERT_TYPE_COMMODITY_CLICK_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_COMMODITY_CLICK"
	AD_CONVERT_TYPE_CONSULT_StardeliveryTaskListV30DataListStarTaskExternalAction                     StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CONSULT"
	AD_CONVERT_TYPE_CONSULT_CLUE_StardeliveryTaskListV30DataListStarTaskExternalAction                StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CONSULT_CLUE"
	AD_CONVERT_TYPE_CONSULT_EFFECTIVE_StardeliveryTaskListV30DataListStarTaskExternalAction           StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CONSULT_EFFECTIVE"
	AD_CONVERT_TYPE_COUPON_StardeliveryTaskListV30DataListStarTaskExternalAction                      StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_COUPON"
	AD_CONVERT_TYPE_CREATE_GAMEROLE_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CREATE_GAMEROLE"
	AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE_StardeliveryTaskListV30DataListStarTaskExternalAction          StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE"
	AD_CONVERT_TYPE_DEEP_PURCHASE_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_DEEP_PURCHASE"
	AD_CONVERT_TYPE_DIALBACK_StardeliveryTaskListV30DataListStarTaskExternalAction                    StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_DIALBACK"
	AD_CONVERT_TYPE_DIALBACK_CONFIRM_StardeliveryTaskListV30DataListStarTaskExternalAction            StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONFIRM"
	AD_CONVERT_TYPE_DIALBACK_CONNECT_StardeliveryTaskListV30DataListStarTaskExternalAction            StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONNECT"
	AD_CONVERT_TYPE_DOWNLOAD_FINISH_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_FINISH"
	AD_CONVERT_TYPE_DOWNLOAD_START_StardeliveryTaskListV30DataListStarTaskExternalAction              StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_START"
	AD_CONVERT_TYPE_EFFECTIVE_COPY_StardeliveryTaskListV30DataListStarTaskExternalAction              StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_COPY"
	AD_CONVERT_TYPE_EFFECTIVE_PLAY_StardeliveryTaskListV30DataListStarTaskExternalAction              StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_PLAY"
	AD_CONVERT_TYPE_ENTER_HOMEPAGE_StardeliveryTaskListV30DataListStarTaskExternalAction              StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_ENTER_HOMEPAGE"
	AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE_StardeliveryTaskListV30DataListStarTaskExternalAction          StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_FIRST_CLASS_StardeliveryTaskListV30DataListStarTaskExternalAction                 StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_FIRST_CLASS"
	AD_CONVERT_TYPE_FIRST_RENTAL_ORDER_StardeliveryTaskListV30DataListStarTaskExternalAction          StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_FIRST_RENTAL_ORDER"
	AD_CONVERT_TYPE_FOLLOW_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_FOLLOW_ACTION"
	AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT_StardeliveryTaskListV30DataListStarTaskExternalAction        StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT"
	AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER_StardeliveryTaskListV30DataListStarTaskExternalAction           StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER"
	AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH_StardeliveryTaskListV30DataListStarTaskExternalAction    StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH"
	AD_CONVERT_TYPE_FORM_StardeliveryTaskListV30DataListStarTaskExternalAction                        StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_FORM"
	AD_CONVERT_TYPE_FORM_ANSWER_StardeliveryTaskListV30DataListStarTaskExternalAction                 StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_FORM_ANSWER"
	AD_CONVERT_TYPE_FORM_CONNECT_StardeliveryTaskListV30DataListStarTaskExternalAction                StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_FORM_CONNECT"
	AD_CONVERT_TYPE_FORM_DEEP_StardeliveryTaskListV30DataListStarTaskExternalAction                   StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_FORM_DEEP"
	AD_CONVERT_TYPE_FORTUNE_SMALL_DEAL_StardeliveryTaskListV30DataListStarTaskExternalAction          StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_FORTUNE_SMALL_DEAL"
	AD_CONVERT_TYPE_GAME_ADDICTION_StardeliveryTaskListV30DataListStarTaskExternalAction              StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_GAME_ADDICTION"
	AD_CONVERT_TYPE_HIGH_VALUE_CLUE_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_HIGH_VALUE_CLUE"
	AD_CONVERT_TYPE_IDCARD_INFORMATION_StardeliveryTaskListV30DataListStarTaskExternalAction          StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_IDCARD_INFORMATION"
	AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN_StardeliveryTaskListV30DataListStarTaskExternalAction         StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_INSTALL_FINISH_StardeliveryTaskListV30DataListStarTaskExternalAction              StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_INSTALL_FINISH"
	AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN_StardeliveryTaskListV30DataListStarTaskExternalAction    StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN"
	AD_CONVERT_TYPE_INTENTION_CLUE_StardeliveryTaskListV30DataListStarTaskExternalAction              StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_INTENTION_CLUE"
	AD_CONVERT_TYPE_INTERACTION_StardeliveryTaskListV30DataListStarTaskExternalAction                 StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_INTERACTION"
	AD_CONVERT_TYPE_INVALID_CLUE_StardeliveryTaskListV30DataListStarTaskExternalAction                StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_INVALID_CLUE"
	AD_CONVERT_TYPE_IPU_QUALIFY_StardeliveryTaskListV30DataListStarTaskExternalAction                 StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_IPU_QUALIFY"
	AD_CONVERT_TYPE_LIKE_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction                 StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIKE_ACTION"
	AD_CONVERT_TYPE_LINK_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction                 StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LINK_ACTION"
	AD_CONVERT_TYPE_LIVE_APPOINTMENT_StardeliveryTaskListV30DataListStarTaskExternalAction            StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_APPOINTMENT"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction   StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction         StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK_StardeliveryTaskListV30DataListStarTaskExternalAction        StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction           StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_FANS_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction            StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_FANS_ACTION"
	AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON_StardeliveryTaskListV30DataListStarTaskExternalAction          StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON"
	AD_CONVERT_TYPE_LIVE_GIFT_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction            StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_GIFT_ACTION"
	AD_CONVERT_TYPE_LIVE_HOMEPAGE_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_LIVE_JOIN_GROUP_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_JOIN_GROUP"
	AD_CONVERT_TYPE_LIVE_NATIVE_ACITON_StardeliveryTaskListV30DataListStarTaskExternalAction          StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_NATIVE_ACITON"
	AD_CONVERT_TYPE_LIVE_OTO_CLICK_StardeliveryTaskListV30DataListStarTaskExternalAction              StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_CLICK"
	AD_CONVERT_TYPE_LIVE_OTO_PAY_StardeliveryTaskListV30DataListStarTaskExternalAction                StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_PAY"
	AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK_StardeliveryTaskListV30DataListStarTaskExternalAction          StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK"
	AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION"
	AD_CONVERT_TYPE_LIVE_STAY_TIME_StardeliveryTaskListV30DataListStarTaskExternalAction              StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_STAY_TIME"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction    StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_StardeliveryTaskListV30DataListStarTaskExternalAction       StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_LOAN_StardeliveryTaskListV30DataListStarTaskExternalAction                        StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LOAN"
	AD_CONVERT_TYPE_LOAN_COMPLETION_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LOAN_COMPLETION"
	AD_CONVERT_TYPE_LOAN_CREDIT_StardeliveryTaskListV30DataListStarTaskExternalAction                 StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LOAN_CREDIT"
	AD_CONVERT_TYPE_LOCATION_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LOCATION_ACTION"
	AD_CONVERT_TYPE_LOTTERY_StardeliveryTaskListV30DataListStarTaskExternalAction                     StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LOTTERY"
	AD_CONVERT_TYPE_LT_ROI_StardeliveryTaskListV30DataListStarTaskExternalAction                      StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_LT_ROI"
	AD_CONVERT_TYPE_MAP_SEARCH_StardeliveryTaskListV30DataListStarTaskExternalAction                  StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_MAP_SEARCH"
	AD_CONVERT_TYPE_MESSAGE_StardeliveryTaskListV30DataListStarTaskExternalAction                     StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE"
	AD_CONVERT_TYPE_MESSAGE_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction              StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE_ACTION"
	AD_CONVERT_TYPE_MESSAGE_CLICK_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLICK"
	AD_CONVERT_TYPE_MESSAGE_CLUE_StardeliveryTaskListV30DataListStarTaskExternalAction                StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLUE"
	AD_CONVERT_TYPE_MESSAGE_INTERACTION_StardeliveryTaskListV30DataListStarTaskExternalAction         StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE_INTERACTION"
	AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP_StardeliveryTaskListV30DataListStarTaskExternalAction          StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP"
	AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS_StardeliveryTaskListV30DataListStarTaskExternalAction       StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS"
	AD_CONVERT_TYPE_MESSAGE_SERVICE_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE_SERVICE"
	AD_CONVERT_TYPE_MULTIPLE_StardeliveryTaskListV30DataListStarTaskExternalAction                    StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_MULTIPLE"
	AD_CONVERT_TYPE_NATIVE_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_NATIVE_ACTION"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction           StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_NEXT_DAY_OPEN_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_NOTIFY_DOWNLOAD_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_NOTIFY_DOWNLOAD"
	AD_CONVERT_TYPE_NOTIFY_FORM_StardeliveryTaskListV30DataListStarTaskExternalAction                 StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_NOTIFY_FORM"
	AD_CONVERT_TYPE_OTHER_StardeliveryTaskListV30DataListStarTaskExternalAction                       StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_OTHER"
	AD_CONVERT_TYPE_OTO_PAY_StardeliveryTaskListV30DataListStarTaskExternalAction                     StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_OTO_PAY"
	AD_CONVERT_TYPE_OTO_STAY_TIME_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_OTO_STAY_TIME"
	AD_CONVERT_TYPE_PAID_CLUE_StardeliveryTaskListV30DataListStarTaskExternalAction                   StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PAID_CLUE"
	AD_CONVERT_TYPE_PAY_StardeliveryTaskListV30DataListStarTaskExternalAction                         StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PAY"
	AD_CONVERT_TYPE_PERSONAL_INFORMATION_StardeliveryTaskListV30DataListStarTaskExternalAction        StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PERSONAL_INFORMATION"
	AD_CONVERT_TYPE_PHONE_StardeliveryTaskListV30DataListStarTaskExternalAction                       StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PHONE"
	AD_CONVERT_TYPE_PHONE_CONFIRM_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PHONE_CONFIRM"
	AD_CONVERT_TYPE_PHONE_CONNECT_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PHONE_CONNECT"
	AD_CONVERT_TYPE_PHONE_EFFECTIVE_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PHONE_EFFECTIVE"
	AD_CONVERT_TYPE_POI_ADDRESS_CLICK_StardeliveryTaskListV30DataListStarTaskExternalAction           StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_POI_ADDRESS_CLICK"
	AD_CONVERT_TYPE_POI_COLLECT_StardeliveryTaskListV30DataListStarTaskExternalAction                 StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_POI_COLLECT"
	AD_CONVERT_TYPE_POI_MULTIPLE_StardeliveryTaskListV30DataListStarTaskExternalAction                StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_POI_MULTIPLE"
	AD_CONVERT_TYPE_PREMIUM_PAYMENT_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PREMIUM_PAYMENT"
	AD_CONVERT_TYPE_PREMIUM_ROI_StardeliveryTaskListV30DataListStarTaskExternalAction                 StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PREMIUM_ROI"
	AD_CONVERT_TYPE_PREMIUM_UPGEADE_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PREMIUM_UPGEADE"
	AD_CONVERT_TYPE_PRE_LOAN_CREDIT_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PRE_LOAN_CREDIT"
	AD_CONVERT_TYPE_PURCHASE_OF_GOODS_StardeliveryTaskListV30DataListStarTaskExternalAction           StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PURCHASE_OF_GOODS"
	AD_CONVERT_TYPE_PURCHASE_ROI_StardeliveryTaskListV30DataListStarTaskExternalAction                StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI"
	AD_CONVERT_TYPE_PURCHASE_ROI_2_D_StardeliveryTaskListV30DataListStarTaskExternalAction            StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI_2D"
	AD_CONVERT_TYPE_PURCHASE_ROI_7_D_StardeliveryTaskListV30DataListStarTaskExternalAction            StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI_7D"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction            StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_StardeliveryTaskListV30DataListStarTaskExternalAction                 StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_QQ_StardeliveryTaskListV30DataListStarTaskExternalAction                          StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_QQ"
	AD_CONVERT_TYPE_REDIRECT_StardeliveryTaskListV30DataListStarTaskExternalAction                    StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_REDIRECT"
	AD_CONVERT_TYPE_REDIRECT_TO_SHOP_StardeliveryTaskListV30DataListStarTaskExternalAction            StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_SHOP"
	AD_CONVERT_TYPE_REDIRECT_TO_STORE_StardeliveryTaskListV30DataListStarTaskExternalAction           StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_STORE"
	AD_CONVERT_TYPE_RESERVATION_StardeliveryTaskListV30DataListStarTaskExternalAction                 StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_RESERVATION"
	AD_CONVERT_TYPE_RETENTION_7_D_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_RETENTION_7D"
	AD_CONVERT_TYPE_RETENTION_DAYS_StardeliveryTaskListV30DataListStarTaskExternalAction              StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_RETENTION_DAYS"
	AD_CONVERT_TYPE_RSS_StardeliveryTaskListV30DataListStarTaskExternalAction                         StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_RSS"
	AD_CONVERT_TYPE_SALES_LEAD_StardeliveryTaskListV30DataListStarTaskExternalAction                  StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_SALES_LEAD"
	AD_CONVERT_TYPE_SERVICE_WITHDRAW_StardeliveryTaskListV30DataListStarTaskExternalAction            StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_SERVICE_WITHDRAW"
	AD_CONVERT_TYPE_SHARE_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction                StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_SHARE_ACTION"
	AD_CONVERT_TYPE_SHOPPING_StardeliveryTaskListV30DataListStarTaskExternalAction                    StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_SHOPPING"
	AD_CONVERT_TYPE_SHOPPING_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_SHOPPING_ACTION"
	AD_CONVERT_TYPE_SHOPPING_CART_StardeliveryTaskListV30DataListStarTaskExternalAction               StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_SHOPPING_CART"
	AD_CONVERT_TYPE_SHOW_OFF_NUM_StardeliveryTaskListV30DataListStarTaskExternalAction                StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_SHOW_OFF_NUM"
	AD_CONVERT_TYPE_STAY_TIME_StardeliveryTaskListV30DataListStarTaskExternalAction                   StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_STAY_TIME"
	AD_CONVERT_TYPE_SUBMIT_CERTIFICATION_StardeliveryTaskListV30DataListStarTaskExternalAction        StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_SUBMIT_CERTIFICATION"
	AD_CONVERT_TYPE_SUCCESSORDER_ACTION_StardeliveryTaskListV30DataListStarTaskExternalAction         StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_UG_ROI_StardeliveryTaskListV30DataListStarTaskExternalAction                      StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_UG_ROI"
	AD_CONVERT_TYPE_UPDATE_LEVEL_StardeliveryTaskListV30DataListStarTaskExternalAction                StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_UPDATE_LEVEL"
	AD_CONVERT_TYPE_VIEW_StardeliveryTaskListV30DataListStarTaskExternalAction                        StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_VIEW"
	AD_CONVERT_TYPE_VOTE_StardeliveryTaskListV30DataListStarTaskExternalAction                        StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_VOTE"
	AD_CONVERT_TYPE_WECHAT_StardeliveryTaskListV30DataListStarTaskExternalAction                      StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_WECHAT"
	AD_CONVERT_TYPE_WECHAT_FIRST_MSG_StardeliveryTaskListV30DataListStarTaskExternalAction            StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_WECHAT_FIRST_MSG"
	AD_CONVERT_TYPE_WECHAT_PAY_StardeliveryTaskListV30DataListStarTaskExternalAction                  StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_WECHAT_PAY"
	AD_CONVERT_TYPE_WECHAT_QRCODE_SHOW_StardeliveryTaskListV30DataListStarTaskExternalAction          StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_WECHAT_QRCODE_SHOW"
	AD_CONVERT_TYPE_WECHAT_QRCODE_TRY_StardeliveryTaskListV30DataListStarTaskExternalAction           StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_WECHAT_QRCODE_TRY"
	AD_CONVERT_TYPE_WECHAT_REGISTER_StardeliveryTaskListV30DataListStarTaskExternalAction             StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_WECHAT_REGISTER"
	AD_CONVERT_TYPE_WECHAT_WECOM_ADD_StardeliveryTaskListV30DataListStarTaskExternalAction            StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_WECHAT_WECOM_ADD"
	AD_CONVERT_TYPE_XPATH_StardeliveryTaskListV30DataListStarTaskExternalAction                       StardeliveryTaskListV30DataListStarTaskExternalAction = "AD_CONVERT_TYPE_XPATH"
)

// All allowed values of StardeliveryTaskListV30DataListStarTaskExternalAction enum
var AllowedStardeliveryTaskListV30DataListStarTaskExternalActionEnumValues = []StardeliveryTaskListV30DataListStarTaskExternalAction{
	"AD_CONVERT_PAGE_VIEW",
	"AD_CONVERT_PHONE_CONNECT",
	"AD_CONVERT_TYPE_ACTIVE",
	"AD_CONVERT_TYPE_ACTIVE_REGISTER",
	"AD_CONVERT_TYPE_ANCHOR_CLICK",
	"AD_CONVERT_TYPE_APPLET_CLICK",
	"AD_CONVERT_TYPE_APP_CART",
	"AD_CONVERT_TYPE_APP_DETAIL_UV",
	"AD_CONVERT_TYPE_APP_ORDER",
	"AD_CONVERT_TYPE_APP_PAY",
	"AD_CONVERT_TYPE_APP_UV",
	"AD_CONVERT_TYPE_ARPU0",
	"AD_CONVERT_TYPE_AUTHORIZATION",
	"AD_CONVERT_TYPE_BANKCARD_INFORMATION",
	"AD_CONVERT_TYPE_BOOST",
	"AD_CONVERT_TYPE_BUTTON",
	"AD_CONVERT_TYPE_CERTIFICATION_INFORMATION",
	"AD_CONVERT_TYPE_CLASS",
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
	"AD_CONVERT_TYPE_FIRST_CLASS",
	"AD_CONVERT_TYPE_FIRST_RENTAL_ORDER",
	"AD_CONVERT_TYPE_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT",
	"AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER",
	"AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH",
	"AD_CONVERT_TYPE_FORM",
	"AD_CONVERT_TYPE_FORM_ANSWER",
	"AD_CONVERT_TYPE_FORM_CONNECT",
	"AD_CONVERT_TYPE_FORM_DEEP",
	"AD_CONVERT_TYPE_FORTUNE_SMALL_DEAL",
	"AD_CONVERT_TYPE_GAME_ADDICTION",
	"AD_CONVERT_TYPE_HIGH_VALUE_CLUE",
	"AD_CONVERT_TYPE_IDCARD_INFORMATION",
	"AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN",
	"AD_CONVERT_TYPE_INSTALL_FINISH",
	"AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN",
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
	"AD_CONVERT_TYPE_LIVE_OTO_CLICK",
	"AD_CONVERT_TYPE_LIVE_OTO_PAY",
	"AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK",
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
	"AD_CONVERT_TYPE_PREMIUM_UPGEADE",
	"AD_CONVERT_TYPE_PRE_LOAN_CREDIT",
	"AD_CONVERT_TYPE_PURCHASE_OF_GOODS",
	"AD_CONVERT_TYPE_PURCHASE_ROI",
	"AD_CONVERT_TYPE_PURCHASE_ROI_2D",
	"AD_CONVERT_TYPE_PURCHASE_ROI_7D",
	"AD_CONVERT_TYPE_QC_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_MUST_BUY",
	"AD_CONVERT_TYPE_QQ",
	"AD_CONVERT_TYPE_REDIRECT",
	"AD_CONVERT_TYPE_REDIRECT_TO_SHOP",
	"AD_CONVERT_TYPE_REDIRECT_TO_STORE",
	"AD_CONVERT_TYPE_RESERVATION",
	"AD_CONVERT_TYPE_RETENTION_7D",
	"AD_CONVERT_TYPE_RETENTION_DAYS",
	"AD_CONVERT_TYPE_RSS",
	"AD_CONVERT_TYPE_SALES_LEAD",
	"AD_CONVERT_TYPE_SERVICE_WITHDRAW",
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
	"AD_CONVERT_TYPE_WECHAT_FIRST_MSG",
	"AD_CONVERT_TYPE_WECHAT_PAY",
	"AD_CONVERT_TYPE_WECHAT_QRCODE_SHOW",
	"AD_CONVERT_TYPE_WECHAT_QRCODE_TRY",
	"AD_CONVERT_TYPE_WECHAT_REGISTER",
	"AD_CONVERT_TYPE_WECHAT_WECOM_ADD",
	"AD_CONVERT_TYPE_XPATH",
}

// NewStardeliveryTaskListV30DataListStarTaskExternalActionFromValue returns a pointer to a valid StardeliveryTaskListV30DataListStarTaskExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStardeliveryTaskListV30DataListStarTaskExternalActionFromValue(v string) (*StardeliveryTaskListV30DataListStarTaskExternalAction, error) {
	ev := StardeliveryTaskListV30DataListStarTaskExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StardeliveryTaskListV30DataListStarTaskExternalAction: valid values are %v", v, AllowedStardeliveryTaskListV30DataListStarTaskExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StardeliveryTaskListV30DataListStarTaskExternalAction) IsValid() bool {
	for _, existing := range AllowedStardeliveryTaskListV30DataListStarTaskExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stardelivery_task_list_v3.0_data_list_star_task_external_action value
func (v StardeliveryTaskListV30DataListStarTaskExternalAction) Ptr() *StardeliveryTaskListV30DataListStarTaskExternalAction {
	return &v
}
