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

// ReportCampaignGetV2OrderField
type ReportCampaignGetV2OrderField string

// List of report_campaign_get_v2_order_field
const (
	ACTIVE_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "active"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "advanced_creative_phone_click"
	QQ_ReportCampaignGetV2OrderField                                          ReportCampaignGetV2OrderField = "qq"
	POI_COLLECT_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "poi_collect"
	VALID_PLAY_RATE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "valid_play_rate"
	PLAY_50_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_50_feed_break"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_7days"
	PLAY_50_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_50_feed_break_rate"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_7d_roi"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_day_active_pay_count"
	REDIRECT_TO_SHOP_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "redirect_to_shop"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_7days"
	AVERAGE_PLAY_PROGRESS_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "average_play_progress"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportCampaignGetV2OrderField                 ReportCampaignGetV2OrderField = "live_watch_one_minute_count"
	ACTIVE_RATE_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "active_rate"
	CLICK_WEBSITE_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_website"
	LOAN_CREDIT_RATE_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "loan_credit_rate"
	SHOW_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "show"
	DOWNLOAD_FINISH_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "download_finish_rate"
	LUBAN_LIVE_SHARE_CNT_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "luban_live_share_cnt"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_0d_roi"
	COST_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "cost"
	WIFI_PLAY_RATE_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "wifi_play_rate"
	DOWNLOAD_FINISH_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "download_finish"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_5days"
	PLAY_75_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_75_feed_break"
	TOTAL_PLAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "total_play"
	ATTRIBUTION_RETENTION_6D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_6d_cnt"
	CLICK_DOWNLOAD_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "click_download"
	AVG_RANK_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "avg_rank"
	PLAY_DURATION_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "play_duration"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "average_play_time_per_play"
	IN_APP_CART_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "in_app_cart"
	PRE_CONVERT_RATE_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "pre_convert_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_1day"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_7d_ltv"
	ATTRIBUTION_RETENTION_4D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_4d_cnt"
	CLICK_INSTALL_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_install"
	ATTRIBUTION_RETENTION_7D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_7d_rate"
	ATTRIBUTION_RETENTION_4D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_4d_cost"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_game_pay_7d_cost"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportCampaignGetV2OrderField       ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	PLAY_DURATION_2S_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "play_duration_2s"
	ACTIVE_COST_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "active_cost"
	ACTIVE_PAY_COST_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "active_pay_cost"
	LUBAN_LIVE_GIFT_CNT_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "luban_live_gift_cnt"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "advanced_creative_coupon_addition"
	LIVE_COMPONENT_CLICK_RATE_ReportCampaignGetV2OrderField                   ReportCampaignGetV2OrderField = "live_component_click_rate"
	ACTIVE_PAY_AMOUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "active_pay_amount"
	ATTRIBUTION_CONVERT_COST_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "attribution_convert_cost"
	PLAY_25_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_25_feed_break"
	COMMUTE_FIRST_PAY_COUNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "commute_first_pay_count"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "luban_live_slidecart_click_cnt"
	IN_APP_DETAIL_UV_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "in_app_detail_uv"
	LUBAN_LIVE_FOLLOW_CNT_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "luban_live_follow_cnt"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportCampaignGetV2OrderField   ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	ATTRIBUTION_RETENTION_3D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_3d_cnt"
	DEEP_CONVERT_RATE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "deep_convert_rate"
	PRE_LOAN_CREDIT_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "pre_loan_credit_cost"
	LIVE_COMPONENT_CLICK_COST_ReportCampaignGetV2OrderField                   ReportCampaignGetV2OrderField = "live_component_click_cost"
	LOAN_CREDIT_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "loan_credit"
	PLAY_OVER_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "play_over"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportCampaignGetV2OrderField             ReportCampaignGetV2OrderField = "attribution_active_pay_7d_count"
	CPC_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpc"
	ATTRIBUTION_RETENTION_7D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_7d_cost"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportCampaignGetV2OrderField         ReportCampaignGetV2OrderField = "attribution_active_pay_7d_per_count"
	IN_APP_PAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "in_app_pay"
	STAT_UNION_LTV_0_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_0"
	AVG_SHOW_COST_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "avg_show_cost"
	STAT_UNION_LTV_7_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_7"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_3days"
	LUBAN_ORDER_STAT_AMOUNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "luban_order_stat_amount"
	POI_ADDRESS_CLICK_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "poi_address_click"
	UNION_ROI_3_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_3"
	GAME_PAY_COUNT_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "game_pay_count"
	NEXT_DAY_OPEN_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "next_day_open"
	ATTRIBUTION_RETENTION_4D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_4d_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_3days"
	MESSAGE_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "message"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_wechat_pay_30d_amount"
	SHARE_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "share"
	MAP_SEARCH_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "map_search"
	CLICK_SHOPWINDOW_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "click_shopwindow"
	ATTRIBUTION_RETENTION_2D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_2d_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_4days"
	VIEW_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "view"
	CPA_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpa"
	PLAY_DURATION_2S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_2s_rate"
	COMMENT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "comment"
	ATTRIBUTION_RETENTION_2D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_2d_cnt"
	ATTRIBUTION_CONVERT_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "attribution_convert"
	PLAY_DURATION_SUM_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "play_duration_sum"
	LUBAN_ORDER_CNT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "luban_order_cnt"
	ATTRIBUTION_RETENTION_3D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_3d_rate"
	NEXT_DAY_OPEN_RATE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "next_day_open_rate"
	PAY_AMOUNT_ROI_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "pay_amount_roi"
	VALID_PLAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "valid_play"
	FORM_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "form"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_next_day_open_rate"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_wechat_login_30d_cost"
	ATTRIBUTION_RETENTION_7D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_7d_cnt"
	PLAY_DURATION_3S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_3s_rate"
	DEEP_CONVERT_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "deep_convert"
	ATTRIBUTION_RETENTION_5D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_5d_rate"
	LOAN_COMPLETION_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "loan_completion_rate"
	LIVE_FANS_CLUB_JOIN_CNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "live_fans_club_join_cnt"
	IN_APP_ORDER_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "in_app_order"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportCampaignGetV2OrderField ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportCampaignGetV2OrderField  ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	UNION_ROI_0_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_0"
	IES_CHALLENGE_CLICK_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "ies_challenge_click"
	LUBAN_ORDER_ROI_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "luban_order_roi"
	PLAY_75_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_75_feed_break_rate"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_day_acitve_pay_count"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_3d_ltv"
	CONVERT_RATE_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "convert_rate"
	WECHAT_LOGIN_COUNT_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "wechat_login_count"
	GAME_ADDICTION_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "game_addiction_cost"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_next_day_open_cost"
	CLICK_LANDING_PAGE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "click_landing_page"
	GAME_PAY_COST_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "game_pay_cost"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportCampaignGetV2OrderField          ReportCampaignGetV2OrderField = "attribution_wechat_login_30d_count"
	DOWNLOAD_START_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "download_start_rate"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_wechat_pay_30d_roi"
	HOME_VISITED_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "home_visited"
	WECHAT_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "wechat"
	ACTIVE_PAY_RATE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "active_pay_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_2days"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_5days"
	VOTE_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "vote"
	CUSTOMER_EFFECTIVE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "customer_effective"
	LUBAN_LIVE_ENTER_CNT_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "luban_live_enter_cnt"
	PAY_COUNT_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "pay_count"
	FIRST_RENTAL_ORDER_COUNT_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "first_rental_order_count"
	GAME_ADDICTION_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "game_addiction_rate"
	DISLIKE_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "dislike"
	LIVE_COMPONENT_CLICK_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "live_component_click_count"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_6days"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_retention_7d_sum_cnt"
	NEXT_DAY_OPEN_COST_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "next_day_open_cost"
	AVERAGE_VIDEO_PLAY_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "average_video_play"
	LUBAN_LIVE_GIFT_AMOUNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "luban_live_gift_amount"
	ATTRIBUTION_RETENTION_5D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_5d_cnt"
	PHONE_CONNECT_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "phone_connect"
	PRE_CONVERT_COST_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "pre_convert_cost"
	UNION_ROI_7_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_7"
	CLICK_CALL_DY_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_call_dy"
	PLAY_25_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_25_feed_break_rate"
	APPROVAL_COUNT_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "approval_count"
	CTR_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "ctr"
	AVG_CLICK_COST_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "avg_click_cost"
	PLAY_OVER_RATE_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "play_over_rate"
	WECHAT_LOGIN_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "wechat_login_cost"
	CONSULT_EFFECTIVE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "consult_effective"
	PLAY_100_FEED_BREAK_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "play_100_feed_break"
	ATTRIBUTION_RETENTION_6D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_6d_cost"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "luban_live_click_product_cnt"
	PHONE_CONFIRM_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "phone_confirm"
	DOWNLOAD_FINISH_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "download_finish_cost"
	INSTALL_FINISH_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "install_finish_rate"
	PLAY_DURATION_3S_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "play_duration_3s"
	WECHAT_FIRST_PAY_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "wechat_first_pay_rate"
	STAT_PAY_AMOUNT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "stat_pay_amount"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_next_day_open_cnt"
	PRE_LOAN_CREDIT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "pre_loan_credit"
	PHONE_EFFECTIVE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "phone_effective"
	REDIRECT_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "redirect"
	FOLLOW_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "follow"
	ATTRIBUTION_RETENTION_3D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_3d_cost"
	PRE_CONVERT_COUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "pre_convert_count"
	BUTTON_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "button"
	LIKE_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "like"
	PLAY_DURATION_10S_RATE_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "play_duration_10s_rate"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportCampaignGetV2OrderField   ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	COUPON_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "coupon"
	ACTIVE_REGISTER_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "active_register_cost"
	INTERACT_PER_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "interact_per_cost"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "advanced_creative_form_submit"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_2days"
	PLAY_DURATION_5S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_5s_rate"
	GAME_ADDICTION_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "game_addiction"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_4days"
	SHOPPING_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "shopping"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_8days"
	ADVANCED_CREATIVE_FORM_CLICK_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "advanced_creative_form_click"
	PHONE_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "phone"
	CARD_SHOW_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "card_show"
	DEEP_CONVERT_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "deep_convert_cost"
	REGISTER_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "register"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_game_pay_7d_count"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "luban_live_pay_order_stat_cost"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportCampaignGetV2OrderField      ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	LOCATION_CLICK_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "location_click"
	MESSAGE_ACTION_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "message_action"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportCampaignGetV2OrderField    ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	STAT_UNION_LTV_3_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_3"
	IES_MUSIC_CLICK_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "ies_music_click"
	FIRST_ORDER_COUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "first_order_count"
	PLAY_DURATION_10S_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "play_duration_10s"
	LOAN_CREDIT_COST_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "loan_credit_cost"
	LOAN_COMPLETION_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "loan_completion"
	ATTRIBUTION_RETENTION_5D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_5d_cost"
	INSTALL_FINISH_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "install_finish_cost"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportCampaignGetV2OrderField       ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_0d_ltv"
	ACTIVE_REGISTER_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "active_register_rate"
	CONVERT_COST_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "convert_cost"
	ATTRIBUTION_RETENTION_6D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_6d_rate"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_deep_convert_cost"
	WECHAT_FIRST_PAY_COUNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "wechat_first_pay_count"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_1day"
	DOWNLOAD_START_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "download_start"
	CONVERT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "convert"
	INSTALL_FINISH_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "install_finish"
	LOTTERY_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "lottery"
	IN_APP_UV_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "in_app_uv"
	CONSULT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "consult"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_6days"
	COUPON_SINGLE_PAGE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "coupon_single_page"
	DOWNLOAD_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "download"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportCampaignGetV2OrderField         ReportCampaignGetV2OrderField = "attribution_retention_7d_total_cost"
	CLICK_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "click"
	WIFI_PLAY_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "wifi_play"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportCampaignGetV2OrderField             ReportCampaignGetV2OrderField = "advanced_creative_counsel_click"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_active_pay_7d_cost"
	LOAN_COMPLETION_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "loan_completion_cost"
	LUBAN_LIVE_COMMENT_CNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "luban_live_comment_cnt"
	REPORT_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "report"
	VALID_PLAY_COST_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "valid_play_cost"
	CONVERT_SHOW_RATE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "convert_show_rate"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "luban_live_pay_order_count"
	WECHAT_PAY_AMOUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "wechat_pay_amount"
	PLAY_100_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "play_100_feed_break_rate"
	DOWNLOAD_START_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "download_start_cost"
	CPM_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpm"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_active_pay_7d_rate"
	WECHAT_FIRST_PAY_COST_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "wechat_first_pay_cost"
	ATTRIBUTION_DEEP_CONVERT_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "attribution_deep_convert"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_8days"
	ATTRIBUTION_RETENTION_2D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_2d_rate"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_3d_roi"
	SUBMIT_CERTIFICATION_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "submit_certification_count"
)

// All allowed values of ReportCampaignGetV2OrderField enum
var AllowedReportCampaignGetV2OrderFieldEnumValues = []ReportCampaignGetV2OrderField{
	"active",
	"advanced_creative_phone_click",
	"qq",
	"poi_collect",
	"valid_play_rate",
	"play_50_feed_break",
	"attribution_game_in_app_roi_7days",
	"play_50_feed_break_rate",
	"attribution_micro_game_7d_roi",
	"attribution_day_active_pay_count",
	"redirect_to_shop",
	"attribution_game_in_app_ltv_7days",
	"average_play_progress",
	"live_watch_one_minute_count",
	"active_rate",
	"click_website",
	"loan_credit_rate",
	"show",
	"download_finish_rate",
	"luban_live_share_cnt",
	"attribution_micro_game_0d_roi",
	"cost",
	"wifi_play_rate",
	"download_finish",
	"attribution_game_in_app_ltv_5days",
	"play_75_feed_break",
	"total_play",
	"attribution_retention_6d_cnt",
	"click_download",
	"avg_rank",
	"play_duration",
	"average_play_time_per_play",
	"in_app_cart",
	"pre_convert_rate",
	"attribution_game_in_app_ltv_1day",
	"attribution_micro_game_7d_ltv",
	"attribution_retention_4d_cnt",
	"click_install",
	"attribution_retention_7d_rate",
	"attribution_retention_4d_cost",
	"attribution_game_pay_7d_cost",
	"attribution_wechat_first_pay_30d_rate",
	"play_duration_2s",
	"active_cost",
	"active_pay_cost",
	"luban_live_gift_cnt",
	"advanced_creative_coupon_addition",
	"live_component_click_rate",
	"active_pay_amount",
	"attribution_convert_cost",
	"play_25_feed_break",
	"commute_first_pay_count",
	"luban_live_slidecart_click_cnt",
	"in_app_detail_uv",
	"luban_live_follow_cnt",
	"attribution_active_pay_intra_one_day_cost",
	"attribution_retention_3d_cnt",
	"deep_convert_rate",
	"pre_loan_credit_cost",
	"live_component_click_cost",
	"loan_credit",
	"play_over",
	"attribution_active_pay_7d_count",
	"cpc",
	"attribution_retention_7d_cost",
	"attribution_active_pay_7d_per_count",
	"in_app_pay",
	"stat_union_ltv_0",
	"avg_show_cost",
	"stat_union_ltv_7",
	"attribution_game_in_app_roi_3days",
	"luban_order_stat_amount",
	"poi_address_click",
	"union_roi_3",
	"game_pay_count",
	"next_day_open",
	"attribution_retention_4d_rate",
	"attribution_game_in_app_ltv_3days",
	"message",
	"attribution_wechat_pay_30d_amount",
	"share",
	"map_search",
	"click_shopwindow",
	"attribution_retention_2d_cost",
	"attribution_game_in_app_roi_4days",
	"view",
	"cpa",
	"play_duration_2s_rate",
	"comment",
	"attribution_retention_2d_cnt",
	"attribution_convert",
	"play_duration_sum",
	"luban_order_cnt",
	"attribution_retention_3d_rate",
	"next_day_open_rate",
	"pay_amount_roi",
	"valid_play",
	"form",
	"attribution_next_day_open_rate",
	"attribution_wechat_login_30d_cost",
	"attribution_retention_7d_cnt",
	"play_duration_3s_rate",
	"deep_convert",
	"attribution_retention_5d_rate",
	"loan_completion_rate",
	"live_fans_club_join_cnt",
	"in_app_order",
	"attribution_active_pay_intra_one_day_amount",
	"attribution_active_pay_intra_one_day_count",
	"union_roi_0",
	"ies_challenge_click",
	"luban_order_roi",
	"play_75_feed_break_rate",
	"attribution_day_acitve_pay_count",
	"attribution_micro_game_3d_ltv",
	"convert_rate",
	"wechat_login_count",
	"game_addiction_cost",
	"attribution_next_day_open_cost",
	"click_landing_page",
	"game_pay_cost",
	"attribution_wechat_login_30d_count",
	"download_start_rate",
	"attribution_wechat_pay_30d_roi",
	"home_visited",
	"wechat",
	"active_pay_rate",
	"attribution_game_in_app_roi_2days",
	"attribution_game_in_app_roi_5days",
	"vote",
	"customer_effective",
	"luban_live_enter_cnt",
	"pay_count",
	"first_rental_order_count",
	"game_addiction_rate",
	"dislike",
	"live_component_click_count",
	"attribution_game_in_app_ltv_6days",
	"attribution_retention_7d_sum_cnt",
	"next_day_open_cost",
	"average_video_play",
	"luban_live_gift_amount",
	"attribution_retention_5d_cnt",
	"phone_connect",
	"pre_convert_cost",
	"union_roi_7",
	"click_call_dy",
	"play_25_feed_break_rate",
	"approval_count",
	"ctr",
	"avg_click_cost",
	"play_over_rate",
	"wechat_login_cost",
	"consult_effective",
	"play_100_feed_break",
	"attribution_retention_6d_cost",
	"luban_live_click_product_cnt",
	"phone_confirm",
	"download_finish_cost",
	"install_finish_rate",
	"play_duration_3s",
	"wechat_first_pay_rate",
	"stat_pay_amount",
	"attribution_next_day_open_cnt",
	"pre_loan_credit",
	"phone_effective",
	"redirect",
	"follow",
	"attribution_retention_3d_cost",
	"pre_convert_count",
	"button",
	"like",
	"play_duration_10s_rate",
	"attribution_active_pay_intra_one_day_rate",
	"coupon",
	"active_register_cost",
	"interact_per_cost",
	"advanced_creative_form_submit",
	"attribution_game_in_app_ltv_2days",
	"play_duration_5s_rate",
	"game_addiction",
	"attribution_game_in_app_ltv_4days",
	"shopping",
	"attribution_game_in_app_roi_8days",
	"advanced_creative_form_click",
	"phone",
	"card_show",
	"deep_convert_cost",
	"register",
	"attribution_game_pay_7d_count",
	"luban_live_pay_order_stat_cost",
	"attribution_wechat_first_pay_30d_count",
	"location_click",
	"message_action",
	"attribution_active_pay_intra_one_day_roi",
	"stat_union_ltv_3",
	"ies_music_click",
	"first_order_count",
	"play_duration_10s",
	"loan_credit_cost",
	"loan_completion",
	"attribution_retention_5d_cost",
	"install_finish_cost",
	"attribution_wechat_first_pay_30d_cost",
	"attribution_micro_game_0d_ltv",
	"active_register_rate",
	"convert_cost",
	"attribution_retention_6d_rate",
	"attribution_deep_convert_cost",
	"wechat_first_pay_count",
	"attribution_game_in_app_roi_1day",
	"download_start",
	"convert",
	"install_finish",
	"lottery",
	"in_app_uv",
	"consult",
	"attribution_game_in_app_roi_6days",
	"coupon_single_page",
	"download",
	"attribution_retention_7d_total_cost",
	"click",
	"wifi_play",
	"advanced_creative_counsel_click",
	"attribution_active_pay_7d_cost",
	"loan_completion_cost",
	"luban_live_comment_cnt",
	"report",
	"valid_play_cost",
	"convert_show_rate",
	"luban_live_pay_order_count",
	"wechat_pay_amount",
	"play_100_feed_break_rate",
	"download_start_cost",
	"cpm",
	"attribution_active_pay_7d_rate",
	"wechat_first_pay_cost",
	"attribution_deep_convert",
	"attribution_game_in_app_ltv_8days",
	"attribution_retention_2d_rate",
	"attribution_micro_game_3d_roi",
	"submit_certification_count",
}

// NewReportCampaignGetV2OrderFieldFromValue returns a pointer to a valid ReportCampaignGetV2OrderField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2OrderFieldFromValue(v string) (*ReportCampaignGetV2OrderField, error) {
	ev := ReportCampaignGetV2OrderField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2OrderField: valid values are %v", v, AllowedReportCampaignGetV2OrderFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2OrderField) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2OrderFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_order_field value
func (v ReportCampaignGetV2OrderField) Ptr() *ReportCampaignGetV2OrderField {
	return &v
}
