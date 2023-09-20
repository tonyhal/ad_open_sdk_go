/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
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
	ACTIVE_REGISTER_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "active_register_cost"
	COMMUTE_FIRST_PAY_COUNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "commute_first_pay_count"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportCampaignGetV2OrderField  ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	ACTIVE_PAY_COST_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "active_pay_cost"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_game_pay_7d_cost"
	IN_APP_CART_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "in_app_cart"
	INSTALL_FINISH_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "install_finish_rate"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "advanced_creative_coupon_addition"
	PLAY_OVER_RATE_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "play_over_rate"
	ATTRIBUTION_RETENTION_3D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_3d_rate"
	DOWNLOAD_FINISH_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "download_finish"
	STAT_UNION_LTV_0_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_0"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_day_acitve_pay_count"
	ADVANCED_CREATIVE_FORM_CLICK_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "advanced_creative_form_click"
	COST_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "cost"
	NEXT_DAY_OPEN_RATE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "next_day_open_rate"
	COMMENT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "comment"
	CPA_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpa"
	CLICK_SHOPWINDOW_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "click_shopwindow"
	AVG_RANK_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "avg_rank"
	ATTRIBUTION_RETENTION_3D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_3d_cost"
	QQ_ReportCampaignGetV2OrderField                                          ReportCampaignGetV2OrderField = "qq"
	LOTTERY_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "lottery"
	WECHAT_LOGIN_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "wechat_login_cost"
	PRE_CONVERT_RATE_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "pre_convert_rate"
	PHONE_CONFIRM_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "phone_confirm"
	LIVE_COMPONENT_CLICK_COST_ReportCampaignGetV2OrderField                   ReportCampaignGetV2OrderField = "live_component_click_cost"
	PLAY_DURATION_10S_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "play_duration_10s"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_4days"
	DOWNLOAD_START_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "download_start"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_day_active_pay_count"
	CONVERT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "convert"
	CONSULT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "consult"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_wechat_pay_30d_amount"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportCampaignGetV2OrderField             ReportCampaignGetV2OrderField = "attribution_active_pay_7d_count"
	LOAN_COMPLETION_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "loan_completion_rate"
	CONVERT_RATE_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "convert_rate"
	LUBAN_LIVE_FOLLOW_CNT_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "luban_live_follow_cnt"
	INSTALL_FINISH_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "install_finish"
	ATTRIBUTION_RETENTION_7D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_7d_cnt"
	LOAN_COMPLETION_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "loan_completion"
	PLAY_DURATION_SUM_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "play_duration_sum"
	DISLIKE_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "dislike"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_1day"
	CLICK_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "click"
	AVG_CLICK_COST_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "avg_click_cost"
	LUBAN_LIVE_COMMENT_CNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "luban_live_comment_cnt"
	SUBMIT_CERTIFICATION_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "submit_certification_count"
	ACTIVE_PAY_RATE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "active_pay_rate"
	PLAY_DURATION_3S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_3s_rate"
	LIVE_COMPONENT_CLICK_RATE_ReportCampaignGetV2OrderField                   ReportCampaignGetV2OrderField = "live_component_click_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_6days"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_active_pay_7d_cost"
	WECHAT_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "wechat"
	CONVERT_COST_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "convert_cost"
	IES_CHALLENGE_CLICK_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "ies_challenge_click"
	PLAY_DURATION_2S_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "play_duration_2s"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_5days"
	ATTRIBUTION_RETENTION_6D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_6d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_7days"
	UNION_ROI_0_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_0"
	ATTRIBUTION_CONVERT_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "attribution_convert"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportCampaignGetV2OrderField ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportCampaignGetV2OrderField    ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	ACTIVE_REGISTER_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "active_register_rate"
	ATTRIBUTION_RETENTION_4D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_4d_cnt"
	ATTRIBUTION_DEEP_CONVERT_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "attribution_deep_convert"
	CPC_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpc"
	LIVE_COMPONENT_CLICK_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "live_component_click_count"
	CTR_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "ctr"
	REGISTER_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "register"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_wechat_pay_30d_roi"
	LUBAN_LIVE_GIFT_AMOUNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "luban_live_gift_amount"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportCampaignGetV2OrderField       ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	LOAN_CREDIT_COST_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "loan_credit_cost"
	PHONE_CONNECT_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "phone_connect"
	CLICK_DOWNLOAD_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "click_download"
	ATTRIBUTION_RETENTION_6D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_6d_rate"
	AVG_SHOW_COST_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "avg_show_cost"
	APPROVAL_COUNT_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "approval_count"
	PLAY_DURATION_10S_RATE_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "play_duration_10s_rate"
	WECHAT_LOGIN_COUNT_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "wechat_login_count"
	MAP_SEARCH_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "map_search"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_active_pay_7d_rate"
	SHARE_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "share"
	PAY_COUNT_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "pay_count"
	LUBAN_LIVE_GIFT_CNT_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "luban_live_gift_cnt"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_8days"
	IN_APP_ORDER_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "in_app_order"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_3days"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_7d_ltv"
	HOME_VISITED_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "home_visited"
	INTERACT_PER_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "interact_per_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_4days"
	FORM_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "form"
	GAME_ADDICTION_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "game_addiction"
	GAME_PAY_COST_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "game_pay_cost"
	DOWNLOAD_FINISH_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "download_finish_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_3days"
	CONSULT_EFFECTIVE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "consult_effective"
	VALID_PLAY_COST_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "valid_play_cost"
	PAY_AMOUNT_ROI_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "pay_amount_roi"
	LUBAN_LIVE_ENTER_CNT_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "luban_live_enter_cnt"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_5days"
	WIFI_PLAY_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "wifi_play"
	ATTRIBUTION_RETENTION_2D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_2d_cost"
	ATTRIBUTION_RETENTION_4D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_4d_rate"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_3d_roi"
	POI_ADDRESS_CLICK_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "poi_address_click"
	ACTIVE_COST_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "active_cost"
	ATTRIBUTION_RETENTION_5D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_5d_rate"
	CLICK_WEBSITE_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_website"
	PLAY_75_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_75_feed_break"
	LOAN_CREDIT_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "loan_credit"
	VOTE_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "vote"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "luban_live_pay_order_count"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_deep_convert_cost"
	PRE_CONVERT_COUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "pre_convert_count"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_next_day_open_cost"
	ATTRIBUTION_RETENTION_4D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_4d_cost"
	SHOW_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "show"
	ATTRIBUTION_RETENTION_3D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_3d_cnt"
	LOAN_CREDIT_RATE_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "loan_credit_rate"
	PLAY_DURATION_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "play_duration"
	MESSAGE_ACTION_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "message_action"
	REDIRECT_TO_SHOP_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "redirect_to_shop"
	PLAY_50_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_50_feed_break_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_2days"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_7d_roi"
	PLAY_75_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_75_feed_break_rate"
	UNION_ROI_7_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_7"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "luban_live_slidecart_click_cnt"
	ATTRIBUTION_RETENTION_5D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_5d_cnt"
	PLAY_DURATION_5S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_5s_rate"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_0d_ltv"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportCampaignGetV2OrderField                 ReportCampaignGetV2OrderField = "live_watch_one_minute_count"
	LIKE_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "like"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportCampaignGetV2OrderField      ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	PRE_LOAN_CREDIT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "pre_loan_credit"
	IES_MUSIC_CLICK_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "ies_music_click"
	LIVE_FANS_CLUB_JOIN_CNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "live_fans_club_join_cnt"
	REPORT_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "report"
	AVERAGE_VIDEO_PLAY_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "average_video_play"
	DOWNLOAD_FINISH_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "download_finish_rate"
	DEEP_CONVERT_RATE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "deep_convert_rate"
	STAT_UNION_LTV_7_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_7"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "average_play_time_per_play"
	COUPON_SINGLE_PAGE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "coupon_single_page"
	DOWNLOAD_START_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "download_start_rate"
	NEXT_DAY_OPEN_COST_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "next_day_open_cost"
	CLICK_CALL_DY_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_call_dy"
	ATTRIBUTION_RETENTION_6D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_6d_cnt"
	LUBAN_LIVE_SHARE_CNT_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "luban_live_share_cnt"
	CLICK_LANDING_PAGE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "click_landing_page"
	ATTRIBUTION_RETENTION_5D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_5d_cost"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportCampaignGetV2OrderField         ReportCampaignGetV2OrderField = "attribution_retention_7d_total_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_6days"
	IN_APP_UV_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "in_app_uv"
	LUBAN_ORDER_CNT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "luban_order_cnt"
	ACTIVE_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "active"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_next_day_open_cnt"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "luban_live_click_product_cnt"
	PLAY_25_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_25_feed_break"
	GAME_ADDICTION_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "game_addiction_rate"
	DOWNLOAD_START_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "download_start_cost"
	IN_APP_PAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "in_app_pay"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportCampaignGetV2OrderField             ReportCampaignGetV2OrderField = "advanced_creative_counsel_click"
	PRE_CONVERT_COST_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "pre_convert_cost"
	LOAN_COMPLETION_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "loan_completion_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_8days"
	VIEW_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "view"
	REDIRECT_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "redirect"
	FIRST_ORDER_COUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "first_order_count"
	ACTIVE_RATE_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "active_rate"
	IN_APP_DETAIL_UV_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "in_app_detail_uv"
	ATTRIBUTION_RETENTION_7D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_7d_rate"
	SHOPPING_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "shopping"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportCampaignGetV2OrderField   ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	ATTRIBUTION_CONVERT_COST_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "attribution_convert_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_2days"
	PLAY_OVER_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "play_over"
	NEXT_DAY_OPEN_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "next_day_open"
	BUTTON_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "button"
	PLAY_100_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "play_100_feed_break_rate"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_0d_roi"
	PLAY_50_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_50_feed_break"
	VALID_PLAY_RATE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "valid_play_rate"
	PHONE_EFFECTIVE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "phone_effective"
	PHONE_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "phone"
	GAME_ADDICTION_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "game_addiction_cost"
	COUPON_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "coupon"
	PLAY_DURATION_3S_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "play_duration_3s"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_3d_ltv"
	ATTRIBUTION_RETENTION_7D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_7d_cost"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "advanced_creative_form_submit"
	LOCATION_CLICK_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "location_click"
	ACTIVE_PAY_AMOUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "active_pay_amount"
	POI_COLLECT_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "poi_collect"
	CARD_SHOW_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "card_show"
	PLAY_DURATION_2S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_2s_rate"
	CUSTOMER_EFFECTIVE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "customer_effective"
	ATTRIBUTION_RETENTION_2D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_2d_rate"
	WECHAT_FIRST_PAY_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "wechat_first_pay_rate"
	TOTAL_PLAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "total_play"
	WECHAT_FIRST_PAY_COST_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "wechat_first_pay_cost"
	FOLLOW_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "follow"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_game_pay_7d_count"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportCampaignGetV2OrderField   ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	LUBAN_ORDER_STAT_AMOUNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "luban_order_stat_amount"
	GAME_PAY_COUNT_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "game_pay_count"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportCampaignGetV2OrderField       ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	PLAY_25_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_25_feed_break_rate"
	MESSAGE_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "message"
	WECHAT_FIRST_PAY_COUNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "wechat_first_pay_count"
	DEEP_CONVERT_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "deep_convert"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_retention_7d_sum_cnt"
	VALID_PLAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "valid_play"
	CPM_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpm"
	CONVERT_SHOW_RATE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "convert_show_rate"
	WECHAT_PAY_AMOUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "wechat_pay_amount"
	INSTALL_FINISH_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "install_finish_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_1day"
	CLICK_INSTALL_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_install"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_next_day_open_rate"
	LUBAN_ORDER_ROI_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "luban_order_roi"
	ATTRIBUTION_RETENTION_2D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_2d_cnt"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportCampaignGetV2OrderField         ReportCampaignGetV2OrderField = "attribution_active_pay_7d_per_count"
	STAT_UNION_LTV_3_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_3"
	PLAY_100_FEED_BREAK_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "play_100_feed_break"
	DEEP_CONVERT_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "deep_convert_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_7days"
	UNION_ROI_3_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_3"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportCampaignGetV2OrderField          ReportCampaignGetV2OrderField = "attribution_wechat_login_30d_count"
	STAT_PAY_AMOUNT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "stat_pay_amount"
	PRE_LOAN_CREDIT_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "pre_loan_credit_cost"
	WIFI_PLAY_RATE_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "wifi_play_rate"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_wechat_login_30d_cost"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "advanced_creative_phone_click"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "luban_live_pay_order_stat_cost"
	AVERAGE_PLAY_PROGRESS_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "average_play_progress"
	FIRST_RENTAL_ORDER_COUNT_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "first_rental_order_count"
	DOWNLOAD_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "download"
)

// All allowed values of ReportCampaignGetV2OrderField enum
var AllowedReportCampaignGetV2OrderFieldEnumValues = []ReportCampaignGetV2OrderField{
	"active_register_cost",
	"commute_first_pay_count",
	"attribution_active_pay_intra_one_day_count",
	"active_pay_cost",
	"attribution_game_pay_7d_cost",
	"in_app_cart",
	"install_finish_rate",
	"advanced_creative_coupon_addition",
	"play_over_rate",
	"attribution_retention_3d_rate",
	"download_finish",
	"stat_union_ltv_0",
	"attribution_day_acitve_pay_count",
	"advanced_creative_form_click",
	"cost",
	"next_day_open_rate",
	"comment",
	"cpa",
	"click_shopwindow",
	"avg_rank",
	"attribution_retention_3d_cost",
	"qq",
	"lottery",
	"wechat_login_cost",
	"pre_convert_rate",
	"phone_confirm",
	"live_component_click_cost",
	"play_duration_10s",
	"attribution_game_in_app_roi_4days",
	"download_start",
	"attribution_day_active_pay_count",
	"convert",
	"consult",
	"attribution_wechat_pay_30d_amount",
	"attribution_active_pay_7d_count",
	"loan_completion_rate",
	"convert_rate",
	"luban_live_follow_cnt",
	"install_finish",
	"attribution_retention_7d_cnt",
	"loan_completion",
	"play_duration_sum",
	"dislike",
	"attribution_game_in_app_roi_1day",
	"click",
	"avg_click_cost",
	"luban_live_comment_cnt",
	"submit_certification_count",
	"active_pay_rate",
	"play_duration_3s_rate",
	"live_component_click_rate",
	"attribution_game_in_app_ltv_6days",
	"attribution_active_pay_7d_cost",
	"wechat",
	"convert_cost",
	"ies_challenge_click",
	"play_duration_2s",
	"attribution_game_in_app_ltv_5days",
	"attribution_retention_6d_cost",
	"attribution_game_in_app_ltv_7days",
	"union_roi_0",
	"attribution_convert",
	"attribution_active_pay_intra_one_day_amount",
	"attribution_active_pay_intra_one_day_roi",
	"active_register_rate",
	"attribution_retention_4d_cnt",
	"attribution_deep_convert",
	"cpc",
	"live_component_click_count",
	"ctr",
	"register",
	"attribution_wechat_pay_30d_roi",
	"luban_live_gift_amount",
	"attribution_wechat_first_pay_30d_rate",
	"loan_credit_cost",
	"phone_connect",
	"click_download",
	"attribution_retention_6d_rate",
	"avg_show_cost",
	"approval_count",
	"play_duration_10s_rate",
	"wechat_login_count",
	"map_search",
	"attribution_active_pay_7d_rate",
	"share",
	"pay_count",
	"luban_live_gift_cnt",
	"attribution_game_in_app_roi_8days",
	"in_app_order",
	"attribution_game_in_app_roi_3days",
	"attribution_micro_game_7d_ltv",
	"home_visited",
	"interact_per_cost",
	"attribution_game_in_app_ltv_4days",
	"form",
	"game_addiction",
	"game_pay_cost",
	"download_finish_cost",
	"attribution_game_in_app_ltv_3days",
	"consult_effective",
	"valid_play_cost",
	"pay_amount_roi",
	"luban_live_enter_cnt",
	"attribution_game_in_app_roi_5days",
	"wifi_play",
	"attribution_retention_2d_cost",
	"attribution_retention_4d_rate",
	"attribution_micro_game_3d_roi",
	"poi_address_click",
	"active_cost",
	"attribution_retention_5d_rate",
	"click_website",
	"play_75_feed_break",
	"loan_credit",
	"vote",
	"luban_live_pay_order_count",
	"attribution_deep_convert_cost",
	"pre_convert_count",
	"attribution_next_day_open_cost",
	"attribution_retention_4d_cost",
	"show",
	"attribution_retention_3d_cnt",
	"loan_credit_rate",
	"play_duration",
	"message_action",
	"redirect_to_shop",
	"play_50_feed_break_rate",
	"attribution_game_in_app_roi_2days",
	"attribution_micro_game_7d_roi",
	"play_75_feed_break_rate",
	"union_roi_7",
	"luban_live_slidecart_click_cnt",
	"attribution_retention_5d_cnt",
	"play_duration_5s_rate",
	"attribution_micro_game_0d_ltv",
	"live_watch_one_minute_count",
	"like",
	"attribution_wechat_first_pay_30d_count",
	"pre_loan_credit",
	"ies_music_click",
	"live_fans_club_join_cnt",
	"report",
	"average_video_play",
	"download_finish_rate",
	"deep_convert_rate",
	"stat_union_ltv_7",
	"average_play_time_per_play",
	"coupon_single_page",
	"download_start_rate",
	"next_day_open_cost",
	"click_call_dy",
	"attribution_retention_6d_cnt",
	"luban_live_share_cnt",
	"click_landing_page",
	"attribution_retention_5d_cost",
	"attribution_retention_7d_total_cost",
	"attribution_game_in_app_roi_6days",
	"in_app_uv",
	"luban_order_cnt",
	"active",
	"attribution_next_day_open_cnt",
	"luban_live_click_product_cnt",
	"play_25_feed_break",
	"game_addiction_rate",
	"download_start_cost",
	"in_app_pay",
	"advanced_creative_counsel_click",
	"pre_convert_cost",
	"loan_completion_cost",
	"attribution_game_in_app_ltv_8days",
	"view",
	"redirect",
	"first_order_count",
	"active_rate",
	"in_app_detail_uv",
	"attribution_retention_7d_rate",
	"shopping",
	"attribution_active_pay_intra_one_day_cost",
	"attribution_convert_cost",
	"attribution_game_in_app_ltv_2days",
	"play_over",
	"next_day_open",
	"button",
	"play_100_feed_break_rate",
	"attribution_micro_game_0d_roi",
	"play_50_feed_break",
	"valid_play_rate",
	"phone_effective",
	"phone",
	"game_addiction_cost",
	"coupon",
	"play_duration_3s",
	"attribution_micro_game_3d_ltv",
	"attribution_retention_7d_cost",
	"advanced_creative_form_submit",
	"location_click",
	"active_pay_amount",
	"poi_collect",
	"card_show",
	"play_duration_2s_rate",
	"customer_effective",
	"attribution_retention_2d_rate",
	"wechat_first_pay_rate",
	"total_play",
	"wechat_first_pay_cost",
	"follow",
	"attribution_game_pay_7d_count",
	"attribution_active_pay_intra_one_day_rate",
	"luban_order_stat_amount",
	"game_pay_count",
	"attribution_wechat_first_pay_30d_cost",
	"play_25_feed_break_rate",
	"message",
	"wechat_first_pay_count",
	"deep_convert",
	"attribution_retention_7d_sum_cnt",
	"valid_play",
	"cpm",
	"convert_show_rate",
	"wechat_pay_amount",
	"install_finish_cost",
	"attribution_game_in_app_ltv_1day",
	"click_install",
	"attribution_next_day_open_rate",
	"luban_order_roi",
	"attribution_retention_2d_cnt",
	"attribution_active_pay_7d_per_count",
	"stat_union_ltv_3",
	"play_100_feed_break",
	"deep_convert_cost",
	"attribution_game_in_app_roi_7days",
	"union_roi_3",
	"attribution_wechat_login_30d_count",
	"stat_pay_amount",
	"pre_loan_credit_cost",
	"wifi_play_rate",
	"attribution_wechat_login_30d_cost",
	"advanced_creative_phone_click",
	"luban_live_pay_order_stat_cost",
	"average_play_progress",
	"first_rental_order_count",
	"download",
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
