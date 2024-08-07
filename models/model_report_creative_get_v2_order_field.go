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

// ReportCreativeGetV2OrderField
type ReportCreativeGetV2OrderField string

// List of report_creative_get_v2_order_field
const (
	WECHAT_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "wechat"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_next_day_open_cnt"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportCreativeGetV2OrderField      ReportCreativeGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	LOAN_CREDIT_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "loan_credit"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportCreativeGetV2OrderField  ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	PLAY_DURATION_3S_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "play_duration_3s"
	REDIRECT_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "redirect"
	PLAY_OVER_RATE_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "play_over_rate"
	CTR_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "ctr"
	ACTIVE_RATE_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "active_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_6days"
	LIVE_FANS_CLUB_JOIN_CNT_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "live_fans_club_join_cnt"
	CPM_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "cpm"
	QQ_ReportCreativeGetV2OrderField                                          ReportCreativeGetV2OrderField = "qq"
	LUBAN_ORDER_ROI_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "luban_order_roi"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_3days"
	ATTRIBUTION_RETENTION_4D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_4d_cnt"
	WECHAT_FIRST_PAY_COUNT_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "wechat_first_pay_count"
	PLAY_DURATION_5S_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "play_duration_5s_rate"
	POI_COLLECT_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "poi_collect"
	PLAY_DURATION_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "play_duration"
	ATTRIBUTION_RETENTION_6D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_6d_cnt"
	ATTRIBUTION_RETENTION_3D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_3d_cnt"
	GAME_ADDICTION_RATE_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "game_addiction_rate"
	DISLIKE_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "dislike"
	PLAY_75_FEED_BREAK_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "play_75_feed_break"
	FIRST_RENTAL_ORDER_COUNT_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "first_rental_order_count"
	DEEP_CONVERT_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "deep_convert"
	FORM_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "form"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportCreativeGetV2OrderField   ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	PLAY_25_FEED_BREAK_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "play_25_feed_break"
	PRE_CONVERT_COST_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "pre_convert_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_8days"
	CONSULT_EFFECTIVE_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "consult_effective"
	MESSAGE_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "message"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_wechat_pay_30d_roi"
	LOAN_COMPLETION_RATE_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "loan_completion_rate"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "luban_live_click_product_cnt"
	ATTRIBUTION_DEEP_CONVERT_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "attribution_deep_convert"
	WECHAT_FIRST_PAY_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "wechat_first_pay_rate"
	NEXT_DAY_OPEN_COST_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "next_day_open_cost"
	LUBAN_ORDER_STAT_AMOUNT_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "luban_order_stat_amount"
	PHONE_EFFECTIVE_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "phone_effective"
	PLAY_DURATION_SUM_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "play_duration_sum"
	ACTIVE_COST_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "active_cost"
	PLAY_75_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "play_75_feed_break_rate"
	CUSTOMER_EFFECTIVE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "customer_effective"
	GAME_ADDICTION_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "game_addiction"
	ADVANCED_CREATIVE_FORM_CLICK_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "advanced_creative_form_click"
	COMMUTE_FIRST_PAY_COUNT_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "commute_first_pay_count"
	ATTRIBUTION_RETENTION_2D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_2d_cost"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "luban_live_pay_order_stat_cost"
	SHOW_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "show"
	ATTRIBUTION_RETENTION_5D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_5d_rate"
	NEXT_DAY_OPEN_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "next_day_open"
	ATTRIBUTION_RETENTION_4D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_4d_cost"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportCreativeGetV2OrderField       ReportCreativeGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_2days"
	ATTRIBUTION_RETENTION_5D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_5d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_5days"
	LUBAN_LIVE_SHARE_CNT_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "luban_live_share_cnt"
	ATTRIBUTION_CONVERT_COST_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "attribution_convert_cost"
	ATTRIBUTION_RETENTION_2D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_2d_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_5days"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_7d_ltv"
	WIFI_PLAY_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "wifi_play"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "advanced_creative_coupon_addition"
	ATTRIBUTION_RETENTION_5D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_5d_cnt"
	LUBAN_LIVE_GIFT_AMOUNT_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "luban_live_gift_amount"
	DOWNLOAD_FINISH_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "download_finish"
	CONVERT_SHOW_RATE_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "convert_show_rate"
	AVERAGE_PLAY_PROGRESS_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "average_play_progress"
	LIKE_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "like"
	NEXT_DAY_OPEN_RATE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "next_day_open_rate"
	VALID_PLAY_RATE_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "valid_play_rate"
	PLAY_DURATION_10S_RATE_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "play_duration_10s_rate"
	PLAY_100_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "play_100_feed_break_rate"
	IN_APP_UV_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "in_app_uv"
	LUBAN_LIVE_ENTER_CNT_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "luban_live_enter_cnt"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "advanced_creative_form_submit"
	POI_ADDRESS_CLICK_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "poi_address_click"
	ATTRIBUTION_RETENTION_7D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_7d_cost"
	UNION_ROI_3_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "union_roi_3"
	IES_MUSIC_CLICK_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "ies_music_click"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_7d_roi"
	DEEP_CONVERT_RATE_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "deep_convert_rate"
	LIVE_COMPONENT_CLICK_COST_ReportCreativeGetV2OrderField                   ReportCreativeGetV2OrderField = "live_component_click_cost"
	LUBAN_LIVE_COMMENT_CNT_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "luban_live_comment_cnt"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_active_pay_7d_cost"
	SHOPPING_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "shopping"
	WECHAT_PAY_AMOUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "wechat_pay_amount"
	CONVERT_COST_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "convert_cost"
	ATTRIBUTION_RETENTION_2D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_2d_cnt"
	LUBAN_ORDER_CNT_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "luban_order_cnt"
	PLAY_DURATION_2S_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "play_duration_2s_rate"
	DOWNLOAD_START_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "download_start"
	PLAY_100_FEED_BREAK_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "play_100_feed_break"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_deep_convert_cost"
	IN_APP_DETAIL_UV_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "in_app_detail_uv"
	FOLLOW_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "follow"
	ATTRIBUTION_RETENTION_7D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_7d_cnt"
	PLAY_DURATION_2S_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "play_duration_2s"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_6days"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_day_acitve_pay_count"
	LOCATION_CLICK_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "location_click"
	PRE_LOAN_CREDIT_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "pre_loan_credit_cost"
	INSTALL_FINISH_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "install_finish"
	PHONE_CONFIRM_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "phone_confirm"
	PLAY_DURATION_10S_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "play_duration_10s"
	PLAY_50_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "play_50_feed_break_rate"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_3d_roi"
	WIFI_PLAY_RATE_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "wifi_play_rate"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_active_pay_7d_rate"
	WECHAT_LOGIN_COST_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "wechat_login_cost"
	ATTRIBUTION_RETENTION_3D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_3d_cost"
	PHONE_CONNECT_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "phone_connect"
	PAY_COUNT_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "pay_count"
	PAY_AMOUNT_ROI_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "pay_amount_roi"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_wechat_pay_30d_amount"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_8days"
	COUPON_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "coupon"
	BUTTON_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "button"
	DOWNLOAD_START_RATE_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "download_start_rate"
	ATTRIBUTION_RETENTION_4D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_4d_rate"
	VALID_PLAY_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "valid_play"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_retention_7d_sum_cnt"
	IES_CHALLENGE_CLICK_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "ies_challenge_click"
	APPROVAL_COUNT_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "approval_count"
	GAME_PAY_COUNT_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "game_pay_count"
	STAT_UNION_LTV_3_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "stat_union_ltv_3"
	CLICK_ReportCreativeGetV2OrderField                                       ReportCreativeGetV2OrderField = "click"
	MAP_SEARCH_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "map_search"
	CLICK_CALL_DY_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "click_call_dy"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportCreativeGetV2OrderField         ReportCreativeGetV2OrderField = "attribution_active_pay_7d_per_count"
	WECHAT_LOGIN_COUNT_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "wechat_login_count"
	REGISTER_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "register"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_1day"
	DOWNLOAD_FINISH_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "download_finish_cost"
	AVERAGE_VIDEO_PLAY_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "average_video_play"
	AVG_CLICK_COST_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "avg_click_cost"
	COUPON_SINGLE_PAGE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "coupon_single_page"
	STAT_UNION_LTV_7_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "stat_union_ltv_7"
	INSTALL_FINISH_RATE_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "install_finish_rate"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_next_day_open_rate"
	CONSULT_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "consult"
	DOWNLOAD_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "download"
	ACTIVE_PAY_AMOUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "active_pay_amount"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "luban_live_pay_order_count"
	HOME_VISITED_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "home_visited"
	COST_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "cost"
	LOAN_CREDIT_RATE_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "loan_credit_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_4days"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_day_active_pay_count"
	REDIRECT_TO_SHOP_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "redirect_to_shop"
	INSTALL_FINISH_COST_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "install_finish_cost"
	ATTRIBUTION_RETENTION_7D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_7d_rate"
	PRE_CONVERT_RATE_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "pre_convert_rate"
	VALID_PLAY_COST_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "valid_play_cost"
	ACTIVE_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "active"
	AVG_SHOW_COST_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "avg_show_cost"
	CLICK_LANDING_PAGE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "click_landing_page"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_1day"
	DOWNLOAD_START_COST_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "download_start_cost"
	STAT_PAY_AMOUNT_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "stat_pay_amount"
	LIVE_COMPONENT_CLICK_COUNT_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "live_component_click_count"
	IN_APP_PAY_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "in_app_pay"
	PHONE_ReportCreativeGetV2OrderField                                       ReportCreativeGetV2OrderField = "phone"
	LOAN_COMPLETION_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "loan_completion_cost"
	SHARE_ReportCreativeGetV2OrderField                                       ReportCreativeGetV2OrderField = "share"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportCreativeGetV2OrderField ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "luban_live_slidecart_click_cnt"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportCreativeGetV2OrderField   ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_3days"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_next_day_open_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_2days"
	ACTIVE_PAY_RATE_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "active_pay_rate"
	MESSAGE_ACTION_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "message_action"
	ACTIVE_REGISTER_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "active_register_cost"
	UNION_ROI_0_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "union_roi_0"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportCreativeGetV2OrderField         ReportCreativeGetV2OrderField = "attribution_retention_7d_total_cost"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportCreativeGetV2OrderField             ReportCreativeGetV2OrderField = "attribution_active_pay_7d_count"
	STAT_UNION_LTV_0_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "stat_union_ltv_0"
	CLICK_DOWNLOAD_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "click_download"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_game_pay_7d_count"
	CONVERT_RATE_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "convert_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_7days"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_0d_ltv"
	GAME_ADDICTION_COST_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "game_addiction_cost"
	PLAY_50_FEED_BREAK_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "play_50_feed_break"
	ACTIVE_REGISTER_RATE_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "active_register_rate"
	CARD_SHOW_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "card_show"
	REPORT_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "report"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportCreativeGetV2OrderField    ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	PRE_LOAN_CREDIT_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "pre_loan_credit"
	GAME_PAY_COST_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "game_pay_cost"
	PLAY_DURATION_3S_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "play_duration_3s_rate"
	PLAY_25_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "play_25_feed_break_rate"
	PRE_CONVERT_COUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "pre_convert_count"
	CLICK_INSTALL_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "click_install"
	VIEW_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "view"
	VOTE_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "vote"
	AVG_RANK_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "avg_rank"
	ATTRIBUTION_RETENTION_6D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_6d_rate"
	ATTRIBUTION_RETENTION_3D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_3d_rate"
	INTERACT_PER_COST_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "interact_per_cost"
	WECHAT_FIRST_PAY_COST_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "wechat_first_pay_cost"
	LIVE_COMPONENT_CLICK_RATE_ReportCreativeGetV2OrderField                   ReportCreativeGetV2OrderField = "live_component_click_rate"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "advanced_creative_phone_click"
	IN_APP_CART_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "in_app_cart"
	LOTTERY_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "lottery"
	CPC_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "cpc"
	UNION_ROI_7_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "union_roi_7"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportCreativeGetV2OrderField                 ReportCreativeGetV2OrderField = "live_watch_one_minute_count"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_0d_roi"
	LUBAN_LIVE_FOLLOW_CNT_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "luban_live_follow_cnt"
	PLAY_OVER_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "play_over"
	ATTRIBUTION_RETENTION_6D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_6d_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_4days"
	LOAN_COMPLETION_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "loan_completion"
	LUBAN_LIVE_GIFT_CNT_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "luban_live_gift_cnt"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportCreativeGetV2OrderField             ReportCreativeGetV2OrderField = "advanced_creative_counsel_click"
	ATTRIBUTION_CONVERT_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "attribution_convert"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_wechat_login_30d_cost"
	CLICK_SHOPWINDOW_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "click_shopwindow"
	CONVERT_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "convert"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_game_pay_7d_cost"
	ACTIVE_PAY_COST_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "active_pay_cost"
	LOAN_CREDIT_COST_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "loan_credit_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_7days"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportCreativeGetV2OrderField          ReportCreativeGetV2OrderField = "attribution_wechat_login_30d_count"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportCreativeGetV2OrderField       ReportCreativeGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "average_play_time_per_play"
	COMMENT_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "comment"
	SUBMIT_CERTIFICATION_COUNT_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "submit_certification_count"
	IN_APP_ORDER_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "in_app_order"
	CLICK_WEBSITE_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "click_website"
	CPA_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "cpa"
	DEEP_CONVERT_COST_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "deep_convert_cost"
	TOTAL_PLAY_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "total_play"
	FIRST_ORDER_COUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "first_order_count"
	DOWNLOAD_FINISH_RATE_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "download_finish_rate"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_3d_ltv"
)

// All allowed values of ReportCreativeGetV2OrderField enum
var AllowedReportCreativeGetV2OrderFieldEnumValues = []ReportCreativeGetV2OrderField{
	"wechat",
	"attribution_next_day_open_cnt",
	"attribution_wechat_first_pay_30d_count",
	"loan_credit",
	"attribution_active_pay_intra_one_day_count",
	"play_duration_3s",
	"redirect",
	"play_over_rate",
	"ctr",
	"active_rate",
	"attribution_game_in_app_ltv_6days",
	"live_fans_club_join_cnt",
	"cpm",
	"qq",
	"luban_order_roi",
	"attribution_game_in_app_ltv_3days",
	"attribution_retention_4d_cnt",
	"wechat_first_pay_count",
	"play_duration_5s_rate",
	"poi_collect",
	"play_duration",
	"attribution_retention_6d_cnt",
	"attribution_retention_3d_cnt",
	"game_addiction_rate",
	"dislike",
	"play_75_feed_break",
	"first_rental_order_count",
	"deep_convert",
	"form",
	"attribution_active_pay_intra_one_day_rate",
	"play_25_feed_break",
	"pre_convert_cost",
	"attribution_game_in_app_ltv_8days",
	"consult_effective",
	"message",
	"attribution_wechat_pay_30d_roi",
	"loan_completion_rate",
	"luban_live_click_product_cnt",
	"attribution_deep_convert",
	"wechat_first_pay_rate",
	"next_day_open_cost",
	"luban_order_stat_amount",
	"phone_effective",
	"play_duration_sum",
	"active_cost",
	"play_75_feed_break_rate",
	"customer_effective",
	"game_addiction",
	"advanced_creative_form_click",
	"commute_first_pay_count",
	"attribution_retention_2d_cost",
	"luban_live_pay_order_stat_cost",
	"show",
	"attribution_retention_5d_rate",
	"next_day_open",
	"attribution_retention_4d_cost",
	"attribution_wechat_first_pay_30d_rate",
	"attribution_game_in_app_ltv_2days",
	"attribution_retention_5d_cost",
	"attribution_game_in_app_ltv_5days",
	"luban_live_share_cnt",
	"attribution_convert_cost",
	"attribution_retention_2d_rate",
	"attribution_game_in_app_roi_5days",
	"attribution_micro_game_7d_ltv",
	"wifi_play",
	"advanced_creative_coupon_addition",
	"attribution_retention_5d_cnt",
	"luban_live_gift_amount",
	"download_finish",
	"convert_show_rate",
	"average_play_progress",
	"like",
	"next_day_open_rate",
	"valid_play_rate",
	"play_duration_10s_rate",
	"play_100_feed_break_rate",
	"in_app_uv",
	"luban_live_enter_cnt",
	"advanced_creative_form_submit",
	"poi_address_click",
	"attribution_retention_7d_cost",
	"union_roi_3",
	"ies_music_click",
	"attribution_micro_game_7d_roi",
	"deep_convert_rate",
	"live_component_click_cost",
	"luban_live_comment_cnt",
	"attribution_active_pay_7d_cost",
	"shopping",
	"wechat_pay_amount",
	"convert_cost",
	"attribution_retention_2d_cnt",
	"luban_order_cnt",
	"play_duration_2s_rate",
	"download_start",
	"play_100_feed_break",
	"attribution_deep_convert_cost",
	"in_app_detail_uv",
	"follow",
	"attribution_retention_7d_cnt",
	"play_duration_2s",
	"attribution_game_in_app_roi_6days",
	"attribution_day_acitve_pay_count",
	"location_click",
	"pre_loan_credit_cost",
	"install_finish",
	"phone_confirm",
	"play_duration_10s",
	"play_50_feed_break_rate",
	"attribution_micro_game_3d_roi",
	"wifi_play_rate",
	"attribution_active_pay_7d_rate",
	"wechat_login_cost",
	"attribution_retention_3d_cost",
	"phone_connect",
	"pay_count",
	"pay_amount_roi",
	"attribution_wechat_pay_30d_amount",
	"attribution_game_in_app_roi_8days",
	"coupon",
	"button",
	"download_start_rate",
	"attribution_retention_4d_rate",
	"valid_play",
	"attribution_retention_7d_sum_cnt",
	"ies_challenge_click",
	"approval_count",
	"game_pay_count",
	"stat_union_ltv_3",
	"click",
	"map_search",
	"click_call_dy",
	"attribution_active_pay_7d_per_count",
	"wechat_login_count",
	"register",
	"attribution_game_in_app_roi_1day",
	"download_finish_cost",
	"average_video_play",
	"avg_click_cost",
	"coupon_single_page",
	"stat_union_ltv_7",
	"install_finish_rate",
	"attribution_next_day_open_rate",
	"consult",
	"download",
	"active_pay_amount",
	"luban_live_pay_order_count",
	"home_visited",
	"cost",
	"loan_credit_rate",
	"attribution_game_in_app_ltv_4days",
	"attribution_day_active_pay_count",
	"redirect_to_shop",
	"install_finish_cost",
	"attribution_retention_7d_rate",
	"pre_convert_rate",
	"valid_play_cost",
	"active",
	"avg_show_cost",
	"click_landing_page",
	"attribution_game_in_app_ltv_1day",
	"download_start_cost",
	"stat_pay_amount",
	"live_component_click_count",
	"in_app_pay",
	"phone",
	"loan_completion_cost",
	"share",
	"attribution_active_pay_intra_one_day_amount",
	"luban_live_slidecart_click_cnt",
	"attribution_active_pay_intra_one_day_cost",
	"attribution_game_in_app_roi_3days",
	"attribution_next_day_open_cost",
	"attribution_game_in_app_roi_2days",
	"active_pay_rate",
	"message_action",
	"active_register_cost",
	"union_roi_0",
	"attribution_retention_7d_total_cost",
	"attribution_active_pay_7d_count",
	"stat_union_ltv_0",
	"click_download",
	"attribution_game_pay_7d_count",
	"convert_rate",
	"attribution_game_in_app_ltv_7days",
	"attribution_micro_game_0d_ltv",
	"game_addiction_cost",
	"play_50_feed_break",
	"active_register_rate",
	"card_show",
	"report",
	"attribution_active_pay_intra_one_day_roi",
	"pre_loan_credit",
	"game_pay_cost",
	"play_duration_3s_rate",
	"play_25_feed_break_rate",
	"pre_convert_count",
	"click_install",
	"view",
	"vote",
	"avg_rank",
	"attribution_retention_6d_rate",
	"attribution_retention_3d_rate",
	"interact_per_cost",
	"wechat_first_pay_cost",
	"live_component_click_rate",
	"advanced_creative_phone_click",
	"in_app_cart",
	"lottery",
	"cpc",
	"union_roi_7",
	"live_watch_one_minute_count",
	"attribution_micro_game_0d_roi",
	"luban_live_follow_cnt",
	"play_over",
	"attribution_retention_6d_cost",
	"attribution_game_in_app_roi_4days",
	"loan_completion",
	"luban_live_gift_cnt",
	"advanced_creative_counsel_click",
	"attribution_convert",
	"attribution_wechat_login_30d_cost",
	"click_shopwindow",
	"convert",
	"attribution_game_pay_7d_cost",
	"active_pay_cost",
	"loan_credit_cost",
	"attribution_game_in_app_roi_7days",
	"attribution_wechat_login_30d_count",
	"attribution_wechat_first_pay_30d_cost",
	"average_play_time_per_play",
	"comment",
	"submit_certification_count",
	"in_app_order",
	"click_website",
	"cpa",
	"deep_convert_cost",
	"total_play",
	"first_order_count",
	"download_finish_rate",
	"attribution_micro_game_3d_ltv",
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
