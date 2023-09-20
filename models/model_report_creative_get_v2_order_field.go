/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCreativeGetV2OrderField
type ReportCreativeGetV2OrderField string

// List of report_creative_get_v2_order_field
const (
	PLAY_DURATION_3S_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "play_duration_3s_rate"
	PLAY_DURATION_5S_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "play_duration_5s_rate"
	INSTALL_FINISH_COST_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "install_finish_cost"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_wechat_pay_30d_amount"
	AVERAGE_VIDEO_PLAY_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "average_video_play"
	ATTRIBUTION_RETENTION_7D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_7d_cnt"
	ACTIVE_PAY_COST_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "active_pay_cost"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_7d_roi"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_0d_roi"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_0d_ltv"
	CONVERT_COST_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "convert_cost"
	GAME_ADDICTION_COST_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "game_addiction_cost"
	DEEP_CONVERT_COST_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "deep_convert_cost"
	DOWNLOAD_START_COST_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "download_start_cost"
	ATTRIBUTION_RETENTION_3D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_3d_cost"
	STAT_UNION_LTV_3_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "stat_union_ltv_3"
	GAME_PAY_COST_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "game_pay_cost"
	PLAY_100_FEED_BREAK_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "play_100_feed_break"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_deep_convert_cost"
	LUBAN_LIVE_ENTER_CNT_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "luban_live_enter_cnt"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportCreativeGetV2OrderField   ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	DOWNLOAD_START_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "download_start"
	IN_APP_DETAIL_UV_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "in_app_detail_uv"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportCreativeGetV2OrderField         ReportCreativeGetV2OrderField = "attribution_active_pay_7d_per_count"
	LIVE_COMPONENT_CLICK_COUNT_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "live_component_click_count"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "luban_live_pay_order_stat_cost"
	REGISTER_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "register"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_6days"
	LUBAN_LIVE_GIFT_CNT_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "luban_live_gift_cnt"
	FIRST_RENTAL_ORDER_COUNT_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "first_rental_order_count"
	WECHAT_LOGIN_COST_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "wechat_login_cost"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "luban_live_click_product_cnt"
	FIRST_ORDER_COUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "first_order_count"
	PLAY_75_FEED_BREAK_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "play_75_feed_break"
	VIEW_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "view"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "advanced_creative_coupon_addition"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportCreativeGetV2OrderField ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_1day"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportCreativeGetV2OrderField    ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	GAME_ADDICTION_RATE_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "game_addiction_rate"
	LUBAN_LIVE_COMMENT_CNT_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "luban_live_comment_cnt"
	PRE_LOAN_CREDIT_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "pre_loan_credit_cost"
	CTR_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "ctr"
	COMMENT_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "comment"
	ACTIVE_RATE_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "active_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_3days"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_game_pay_7d_cost"
	LUBAN_LIVE_SHARE_CNT_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "luban_live_share_cnt"
	WECHAT_FIRST_PAY_COST_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "wechat_first_pay_cost"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_retention_7d_sum_cnt"
	CONVERT_SHOW_RATE_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "convert_show_rate"
	ATTRIBUTION_RETENTION_2D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_2d_cost"
	PHONE_CONNECT_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "phone_connect"
	PLAY_50_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "play_50_feed_break_rate"
	CONVERT_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "convert"
	PLAY_DURATION_2S_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "play_duration_2s"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportCreativeGetV2OrderField   ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	ATTRIBUTION_RETENTION_7D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_7d_rate"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_3d_roi"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_3d_ltv"
	LIVE_COMPONENT_CLICK_COST_ReportCreativeGetV2OrderField                   ReportCreativeGetV2OrderField = "live_component_click_cost"
	ACTIVE_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "active"
	DOWNLOAD_FINISH_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "download_finish"
	FORM_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "form"
	CONSULT_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "consult"
	LUBAN_ORDER_ROI_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "luban_order_roi"
	COST_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "cost"
	ATTRIBUTION_RETENTION_5D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_5d_rate"
	PLAY_75_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "play_75_feed_break_rate"
	IN_APP_PAY_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "in_app_pay"
	PRE_CONVERT_COST_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "pre_convert_cost"
	STAT_PAY_AMOUNT_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "stat_pay_amount"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_2days"
	LOAN_CREDIT_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "loan_credit"
	DOWNLOAD_START_RATE_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "download_start_rate"
	VALID_PLAY_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "valid_play"
	POI_COLLECT_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "poi_collect"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_4days"
	PLAY_DURATION_10S_RATE_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "play_duration_10s_rate"
	ATTRIBUTION_RETENTION_2D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_2d_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_5days"
	LUBAN_LIVE_GIFT_AMOUNT_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "luban_live_gift_amount"
	LOTTERY_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "lottery"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_active_pay_7d_rate"
	AVG_RANK_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "avg_rank"
	AVG_CLICK_COST_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "avg_click_cost"
	BUTTON_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "button"
	SUBMIT_CERTIFICATION_COUNT_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "submit_certification_count"
	CLICK_DOWNLOAD_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "click_download"
	WECHAT_LOGIN_COUNT_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "wechat_login_count"
	CONVERT_RATE_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "convert_rate"
	PAY_COUNT_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "pay_count"
	IN_APP_ORDER_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "in_app_order"
	LOAN_COMPLETION_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "loan_completion_cost"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_game_pay_7d_count"
	CPC_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "cpc"
	PRE_CONVERT_COUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "pre_convert_count"
	UNION_ROI_0_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "union_roi_0"
	DEEP_CONVERT_RATE_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "deep_convert_rate"
	GAME_ADDICTION_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "game_addiction"
	ATTRIBUTION_RETENTION_4D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_4d_cost"
	DOWNLOAD_FINISH_RATE_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "download_finish_rate"
	ATTRIBUTION_RETENTION_2D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_2d_cnt"
	ATTRIBUTION_RETENTION_6D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_6d_rate"
	AVG_SHOW_COST_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "avg_show_cost"
	NEXT_DAY_OPEN_RATE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "next_day_open_rate"
	REDIRECT_TO_SHOP_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "redirect_to_shop"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_7days"
	INTERACT_PER_COST_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "interact_per_cost"
	LOAN_CREDIT_COST_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "loan_credit_cost"
	DISLIKE_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "dislike"
	PRE_CONVERT_RATE_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "pre_convert_rate"
	PAY_AMOUNT_ROI_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "pay_amount_roi"
	WECHAT_FIRST_PAY_COUNT_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "wechat_first_pay_count"
	ATTRIBUTION_RETENTION_3D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_3d_rate"
	ATTRIBUTION_DEEP_CONVERT_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "attribution_deep_convert"
	COMMUTE_FIRST_PAY_COUNT_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "commute_first_pay_count"
	STAT_UNION_LTV_0_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "stat_union_ltv_0"
	ATTRIBUTION_RETENTION_6D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_6d_cost"
	PHONE_ReportCreativeGetV2OrderField                                       ReportCreativeGetV2OrderField = "phone"
	CLICK_SHOPWINDOW_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "click_shopwindow"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_next_day_open_cost"
	STAT_UNION_LTV_7_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "stat_union_ltv_7"
	CUSTOMER_EFFECTIVE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "customer_effective"
	TOTAL_PLAY_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "total_play"
	GAME_PAY_COUNT_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "game_pay_count"
	ATTRIBUTION_RETENTION_6D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_6d_cnt"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_6days"
	MESSAGE_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "message"
	UNION_ROI_3_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "union_roi_3"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportCreativeGetV2OrderField                 ReportCreativeGetV2OrderField = "live_watch_one_minute_count"
	ATTRIBUTION_RETENTION_4D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_4d_cnt"
	PLAY_25_FEED_BREAK_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "play_25_feed_break"
	LOAN_CREDIT_RATE_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "loan_credit_rate"
	PLAY_DURATION_SUM_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "play_duration_sum"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportCreativeGetV2OrderField             ReportCreativeGetV2OrderField = "advanced_creative_counsel_click"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "luban_live_slidecart_click_cnt"
	CPA_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "cpa"
	ACTIVE_PAY_AMOUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "active_pay_amount"
	PLAY_DURATION_10S_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "play_duration_10s"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_next_day_open_cnt"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_5days"
	UNION_ROI_7_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "union_roi_7"
	IES_CHALLENGE_CLICK_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "ies_challenge_click"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportCreativeGetV2OrderField       ReportCreativeGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	ACTIVE_COST_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "active_cost"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_wechat_pay_30d_roi"
	PLAY_OVER_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "play_over"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_2days"
	ATTRIBUTION_RETENTION_4D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_4d_rate"
	IN_APP_CART_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "in_app_cart"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_next_day_open_rate"
	PHONE_EFFECTIVE_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "phone_effective"
	VALID_PLAY_COST_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "valid_play_cost"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportCreativeGetV2OrderField          ReportCreativeGetV2OrderField = "attribution_wechat_login_30d_count"
	PLAY_100_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "play_100_feed_break_rate"
	HOME_VISITED_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "home_visited"
	LIVE_COMPONENT_CLICK_RATE_ReportCreativeGetV2OrderField                   ReportCreativeGetV2OrderField = "live_component_click_rate"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_active_pay_7d_cost"
	ATTRIBUTION_RETENTION_7D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_7d_cost"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_wechat_login_30d_cost"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportCreativeGetV2OrderField      ReportCreativeGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_7days"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportCreativeGetV2OrderField  ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	ATTRIBUTION_CONVERT_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "attribution_convert"
	LIKE_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "like"
	CLICK_CALL_DY_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "click_call_dy"
	MESSAGE_ACTION_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "message_action"
	DOWNLOAD_FINISH_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "download_finish_cost"
	INSTALL_FINISH_RATE_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "install_finish_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_8days"
	REDIRECT_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "redirect"
	WECHAT_FIRST_PAY_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "wechat_first_pay_rate"
	ATTRIBUTION_RETENTION_3D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_3d_cnt"
	WIFI_PLAY_RATE_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "wifi_play_rate"
	ADVANCED_CREATIVE_FORM_CLICK_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "advanced_creative_form_click"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_7d_ltv"
	SHARE_ReportCreativeGetV2OrderField                                       ReportCreativeGetV2OrderField = "share"
	INSTALL_FINISH_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "install_finish"
	WECHAT_PAY_AMOUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "wechat_pay_amount"
	APPROVAL_COUNT_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "approval_count"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportCreativeGetV2OrderField         ReportCreativeGetV2OrderField = "attribution_retention_7d_total_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_3days"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_day_acitve_pay_count"
	ATTRIBUTION_RETENTION_5D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_5d_cost"
	ACTIVE_REGISTER_RATE_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "active_register_rate"
	ACTIVE_REGISTER_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "active_register_cost"
	LUBAN_ORDER_STAT_AMOUNT_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "luban_order_stat_amount"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "advanced_creative_phone_click"
	PHONE_CONFIRM_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "phone_confirm"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_1day"
	DOWNLOAD_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "download"
	PLAY_DURATION_3S_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "play_duration_3s"
	ATTRIBUTION_RETENTION_5D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_5d_cnt"
	LUBAN_ORDER_CNT_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "luban_order_cnt"
	NEXT_DAY_OPEN_COST_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "next_day_open_cost"
	LOAN_COMPLETION_RATE_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "loan_completion_rate"
	LOCATION_CLICK_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "location_click"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_day_active_pay_count"
	POI_ADDRESS_CLICK_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "poi_address_click"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_8days"
	CPM_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "cpm"
	PLAY_DURATION_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "play_duration"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_4days"
	PRE_LOAN_CREDIT_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "pre_loan_credit"
	DEEP_CONVERT_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "deep_convert"
	CLICK_WEBSITE_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "click_website"
	COUPON_SINGLE_PAGE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "coupon_single_page"
	PLAY_50_FEED_BREAK_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "play_50_feed_break"
	SHOW_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "show"
	CLICK_LANDING_PAGE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "click_landing_page"
	AVERAGE_PLAY_PROGRESS_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "average_play_progress"
	CARD_SHOW_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "card_show"
	LUBAN_LIVE_FOLLOW_CNT_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "luban_live_follow_cnt"
	IN_APP_UV_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "in_app_uv"
	NEXT_DAY_OPEN_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "next_day_open"
	PLAY_OVER_RATE_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "play_over_rate"
	WECHAT_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "wechat"
	CLICK_ReportCreativeGetV2OrderField                                       ReportCreativeGetV2OrderField = "click"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportCreativeGetV2OrderField       ReportCreativeGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	LIVE_FANS_CLUB_JOIN_CNT_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "live_fans_club_join_cnt"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "average_play_time_per_play"
	VOTE_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "vote"
	PLAY_25_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "play_25_feed_break_rate"
	LOAN_COMPLETION_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "loan_completion"
	ACTIVE_PAY_RATE_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "active_pay_rate"
	MAP_SEARCH_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "map_search"
	SHOPPING_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "shopping"
	COUPON_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "coupon"
	CONSULT_EFFECTIVE_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "consult_effective"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "luban_live_pay_order_count"
	WIFI_PLAY_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "wifi_play"
	FOLLOW_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "follow"
	IES_MUSIC_CLICK_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "ies_music_click"
	VALID_PLAY_RATE_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "valid_play_rate"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "advanced_creative_form_submit"
	REPORT_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "report"
	CLICK_INSTALL_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "click_install"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportCreativeGetV2OrderField             ReportCreativeGetV2OrderField = "attribution_active_pay_7d_count"
	QQ_ReportCreativeGetV2OrderField                                          ReportCreativeGetV2OrderField = "qq"
	PLAY_DURATION_2S_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "play_duration_2s_rate"
	ATTRIBUTION_CONVERT_COST_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "attribution_convert_cost"
)

// All allowed values of ReportCreativeGetV2OrderField enum
var AllowedReportCreativeGetV2OrderFieldEnumValues = []ReportCreativeGetV2OrderField{
	"play_duration_3s_rate",
	"play_duration_5s_rate",
	"install_finish_cost",
	"attribution_wechat_pay_30d_amount",
	"average_video_play",
	"attribution_retention_7d_cnt",
	"active_pay_cost",
	"attribution_micro_game_7d_roi",
	"attribution_micro_game_0d_roi",
	"attribution_micro_game_0d_ltv",
	"convert_cost",
	"game_addiction_cost",
	"deep_convert_cost",
	"download_start_cost",
	"attribution_retention_3d_cost",
	"stat_union_ltv_3",
	"game_pay_cost",
	"play_100_feed_break",
	"attribution_deep_convert_cost",
	"luban_live_enter_cnt",
	"attribution_active_pay_intra_one_day_cost",
	"download_start",
	"in_app_detail_uv",
	"attribution_active_pay_7d_per_count",
	"live_component_click_count",
	"luban_live_pay_order_stat_cost",
	"register",
	"attribution_game_in_app_ltv_6days",
	"luban_live_gift_cnt",
	"first_rental_order_count",
	"wechat_login_cost",
	"luban_live_click_product_cnt",
	"first_order_count",
	"play_75_feed_break",
	"view",
	"advanced_creative_coupon_addition",
	"attribution_active_pay_intra_one_day_amount",
	"attribution_game_in_app_ltv_1day",
	"attribution_active_pay_intra_one_day_roi",
	"game_addiction_rate",
	"luban_live_comment_cnt",
	"pre_loan_credit_cost",
	"ctr",
	"comment",
	"active_rate",
	"attribution_game_in_app_roi_3days",
	"attribution_game_pay_7d_cost",
	"luban_live_share_cnt",
	"wechat_first_pay_cost",
	"attribution_retention_7d_sum_cnt",
	"convert_show_rate",
	"attribution_retention_2d_cost",
	"phone_connect",
	"play_50_feed_break_rate",
	"convert",
	"play_duration_2s",
	"attribution_active_pay_intra_one_day_rate",
	"attribution_retention_7d_rate",
	"attribution_micro_game_3d_roi",
	"attribution_micro_game_3d_ltv",
	"live_component_click_cost",
	"active",
	"download_finish",
	"form",
	"consult",
	"luban_order_roi",
	"cost",
	"attribution_retention_5d_rate",
	"play_75_feed_break_rate",
	"in_app_pay",
	"pre_convert_cost",
	"stat_pay_amount",
	"attribution_game_in_app_ltv_2days",
	"loan_credit",
	"download_start_rate",
	"valid_play",
	"poi_collect",
	"attribution_game_in_app_roi_4days",
	"play_duration_10s_rate",
	"attribution_retention_2d_rate",
	"attribution_game_in_app_ltv_5days",
	"luban_live_gift_amount",
	"lottery",
	"attribution_active_pay_7d_rate",
	"avg_rank",
	"avg_click_cost",
	"button",
	"submit_certification_count",
	"click_download",
	"wechat_login_count",
	"convert_rate",
	"pay_count",
	"in_app_order",
	"loan_completion_cost",
	"attribution_game_pay_7d_count",
	"cpc",
	"pre_convert_count",
	"union_roi_0",
	"deep_convert_rate",
	"game_addiction",
	"attribution_retention_4d_cost",
	"download_finish_rate",
	"attribution_retention_2d_cnt",
	"attribution_retention_6d_rate",
	"avg_show_cost",
	"next_day_open_rate",
	"redirect_to_shop",
	"attribution_game_in_app_roi_7days",
	"interact_per_cost",
	"loan_credit_cost",
	"dislike",
	"pre_convert_rate",
	"pay_amount_roi",
	"wechat_first_pay_count",
	"attribution_retention_3d_rate",
	"attribution_deep_convert",
	"commute_first_pay_count",
	"stat_union_ltv_0",
	"attribution_retention_6d_cost",
	"phone",
	"click_shopwindow",
	"attribution_next_day_open_cost",
	"stat_union_ltv_7",
	"customer_effective",
	"total_play",
	"game_pay_count",
	"attribution_retention_6d_cnt",
	"attribution_game_in_app_roi_6days",
	"message",
	"union_roi_3",
	"live_watch_one_minute_count",
	"attribution_retention_4d_cnt",
	"play_25_feed_break",
	"loan_credit_rate",
	"play_duration_sum",
	"advanced_creative_counsel_click",
	"luban_live_slidecart_click_cnt",
	"cpa",
	"active_pay_amount",
	"play_duration_10s",
	"attribution_next_day_open_cnt",
	"attribution_game_in_app_roi_5days",
	"union_roi_7",
	"ies_challenge_click",
	"attribution_wechat_first_pay_30d_rate",
	"active_cost",
	"attribution_wechat_pay_30d_roi",
	"play_over",
	"attribution_game_in_app_roi_2days",
	"attribution_retention_4d_rate",
	"in_app_cart",
	"attribution_next_day_open_rate",
	"phone_effective",
	"valid_play_cost",
	"attribution_wechat_login_30d_count",
	"play_100_feed_break_rate",
	"home_visited",
	"live_component_click_rate",
	"attribution_active_pay_7d_cost",
	"attribution_retention_7d_cost",
	"attribution_wechat_login_30d_cost",
	"attribution_wechat_first_pay_30d_count",
	"attribution_game_in_app_ltv_7days",
	"attribution_active_pay_intra_one_day_count",
	"attribution_convert",
	"like",
	"click_call_dy",
	"message_action",
	"download_finish_cost",
	"install_finish_rate",
	"attribution_game_in_app_roi_8days",
	"redirect",
	"wechat_first_pay_rate",
	"attribution_retention_3d_cnt",
	"wifi_play_rate",
	"advanced_creative_form_click",
	"attribution_micro_game_7d_ltv",
	"share",
	"install_finish",
	"wechat_pay_amount",
	"approval_count",
	"attribution_retention_7d_total_cost",
	"attribution_game_in_app_ltv_3days",
	"attribution_day_acitve_pay_count",
	"attribution_retention_5d_cost",
	"active_register_rate",
	"active_register_cost",
	"luban_order_stat_amount",
	"advanced_creative_phone_click",
	"phone_confirm",
	"attribution_game_in_app_roi_1day",
	"download",
	"play_duration_3s",
	"attribution_retention_5d_cnt",
	"luban_order_cnt",
	"next_day_open_cost",
	"loan_completion_rate",
	"location_click",
	"attribution_day_active_pay_count",
	"poi_address_click",
	"attribution_game_in_app_ltv_8days",
	"cpm",
	"play_duration",
	"attribution_game_in_app_ltv_4days",
	"pre_loan_credit",
	"deep_convert",
	"click_website",
	"coupon_single_page",
	"play_50_feed_break",
	"show",
	"click_landing_page",
	"average_play_progress",
	"card_show",
	"luban_live_follow_cnt",
	"in_app_uv",
	"next_day_open",
	"play_over_rate",
	"wechat",
	"click",
	"attribution_wechat_first_pay_30d_cost",
	"live_fans_club_join_cnt",
	"average_play_time_per_play",
	"vote",
	"play_25_feed_break_rate",
	"loan_completion",
	"active_pay_rate",
	"map_search",
	"shopping",
	"coupon",
	"consult_effective",
	"luban_live_pay_order_count",
	"wifi_play",
	"follow",
	"ies_music_click",
	"valid_play_rate",
	"advanced_creative_form_submit",
	"report",
	"click_install",
	"attribution_active_pay_7d_count",
	"qq",
	"play_duration_2s_rate",
	"attribution_convert_cost",
}

// NewReportCreativeGetV2OrderFieldFromValue returns a pointer to a valid ReportCreativeGetV2OrderField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2OrderFieldFromValue(v string) (*ReportCreativeGetV2OrderField, error) {
	ev := ReportCreativeGetV2OrderField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2OrderField: valid values are %v", v, AllowedReportCreativeGetV2OrderFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2OrderField) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2OrderFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_order_field value
func (v ReportCreativeGetV2OrderField) Ptr() *ReportCreativeGetV2OrderField {
	return &v
}
