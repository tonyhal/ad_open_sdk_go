/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2OrderField
type ReportAdGetV2OrderField string

// List of report_ad_get_v2_order_field
const (
	CONSULT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "consult"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_active_pay_7d_rate"
	ACTIVE_PAY_AMOUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "active_pay_amount"
	LUBAN_LIVE_SHARE_CNT_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "luban_live_share_cnt"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_6days"
	PLAY_25_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_25_feed_break_rate"
	PLAY_100_FEED_BREAK_RATE_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "play_100_feed_break_rate"
	ATTRIBUTION_RETENTION_7D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_7d_cnt"
	POI_COLLECT_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "poi_collect"
	POI_ADDRESS_CLICK_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "poi_address_click"
	DOWNLOAD_START_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "download_start_cost"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_day_active_pay_count"
	APPROVAL_COUNT_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "approval_count"
	CLICK_LANDING_PAGE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "click_landing_page"
	IN_APP_CART_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "in_app_cart"
	LOAN_CREDIT_RATE_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "loan_credit_rate"
	WECHAT_FIRST_PAY_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "wechat_first_pay_rate"
	LOCATION_CLICK_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "location_click"
	ATTRIBUTION_RETENTION_2D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_2d_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_3days"
	PLAY_OVER_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "play_over"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_2days"
	PLAY_DURATION_10S_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "play_duration_10s"
	LOAN_COMPLETION_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "loan_completion_cost"
	PHONE_CONNECT_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "phone_connect"
	CONVERT_SHOW_RATE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "convert_show_rate"
	STAT_PAY_AMOUNT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "stat_pay_amount"
	IN_APP_UV_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "in_app_uv"
	WECHAT_LOGIN_COUNT_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "wechat_login_count"
	FIRST_RENTAL_ORDER_COUNT_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "first_rental_order_count"
	INSTALL_FINISH_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "install_finish"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_7d_roi"
	ACTIVE_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "active"
	LUBAN_LIVE_COMMENT_CNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "luban_live_comment_cnt"
	PRE_CONVERT_COUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "pre_convert_count"
	DOWNLOAD_FINISH_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "download_finish_rate"
	GAME_PAY_COUNT_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "game_pay_count"
	VALID_PLAY_RATE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "valid_play_rate"
	WECHAT_FIRST_PAY_COST_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "wechat_first_pay_cost"
	MESSAGE_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "message"
	ACTIVE_REGISTER_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "active_register_rate"
	CPM_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpm"
	ATTRIBUTION_RETENTION_2D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_2d_cnt"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "luban_live_pay_order_stat_cost"
	INTERACT_PER_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "interact_per_cost"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportAdGetV2OrderField      ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	REPORT_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "report"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_7days"
	ACTIVE_PAY_RATE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "active_pay_rate"
	ATTRIBUTION_RETENTION_5D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_5d_cnt"
	LIVE_COMPONENT_CLICK_COST_ReportAdGetV2OrderField                   ReportAdGetV2OrderField = "live_component_click_cost"
	VOTE_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "vote"
	PHONE_EFFECTIVE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "phone_effective"
	DOWNLOAD_FINISH_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "download_finish_cost"
	GAME_ADDICTION_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "game_addiction"
	VALID_PLAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "valid_play"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportAdGetV2OrderField                 ReportAdGetV2OrderField = "live_watch_one_minute_count"
	WIFI_PLAY_RATE_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "wifi_play_rate"
	PLAY_75_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_75_feed_break_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_2days"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "luban_live_slidecart_click_cnt"
	FORM_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "form"
	ATTRIBUTION_RETENTION_5D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_5d_rate"
	AVERAGE_VIDEO_PLAY_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "average_video_play"
	PAY_COUNT_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "pay_count"
	PLAY_DURATION_3S_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "play_duration_3s"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportAdGetV2OrderField   ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	IES_MUSIC_CLICK_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "ies_music_click"
	UNION_ROI_0_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_0"
	CLICK_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "click"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_next_day_open_cost"
	PLAY_DURATION_10S_RATE_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "play_duration_10s_rate"
	ATTRIBUTION_RETENTION_3D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_3d_cost"
	COUPON_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "coupon"
	PLAY_DURATION_2S_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "play_duration_2s"
	STAT_UNION_LTV_3_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_3"
	ACTIVE_REGISTER_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "active_register_cost"
	ATTRIBUTION_RETENTION_6D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_6d_rate"
	STAT_UNION_LTV_7_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_7"
	INSTALL_FINISH_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "install_finish_rate"
	LUBAN_ORDER_ROI_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "luban_order_roi"
	COMMENT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "comment"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "luban_live_click_product_cnt"
	PLAY_75_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_75_feed_break"
	REGISTER_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "register"
	WECHAT_PAY_AMOUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "wechat_pay_amount"
	ATTRIBUTION_RETENTION_2D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_2d_cost"
	CPA_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpa"
	LUBAN_LIVE_GIFT_CNT_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "luban_live_gift_cnt"
	UNION_ROI_7_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_7"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_wechat_pay_30d_roi"
	CLICK_CALL_DY_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_call_dy"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_0d_ltv"
	QQ_ReportAdGetV2OrderField                                          ReportAdGetV2OrderField = "qq"
	ACTIVE_RATE_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "active_rate"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportAdGetV2OrderField    ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	PHONE_CONFIRM_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "phone_confirm"
	ACTIVE_COST_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "active_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_4days"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportAdGetV2OrderField         ReportAdGetV2OrderField = "attribution_active_pay_7d_per_count"
	IES_CHALLENGE_CLICK_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "ies_challenge_click"
	LOTTERY_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "lottery"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "average_play_time_per_play"
	GAME_ADDICTION_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "game_addiction_rate"
	CLICK_SHOPWINDOW_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "click_shopwindow"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_wechat_pay_30d_amount"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "luban_live_pay_order_count"
	DOWNLOAD_START_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "download_start"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_active_pay_7d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_3days"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_8days"
	CLICK_DOWNLOAD_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "click_download"
	PRE_LOAN_CREDIT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "pre_loan_credit"
	PHONE_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "phone"
	ATTRIBUTION_CONVERT_COST_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "attribution_convert_cost"
	CONSULT_EFFECTIVE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "consult_effective"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportAdGetV2OrderField          ReportAdGetV2OrderField = "attribution_wechat_login_30d_count"
	BUTTON_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "button"
	VIEW_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "view"
	PAY_AMOUNT_ROI_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "pay_amount_roi"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_7d_ltv"
	ADVANCED_CREATIVE_FORM_CLICK_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "advanced_creative_form_click"
	CUSTOMER_EFFECTIVE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "customer_effective"
	WECHAT_FIRST_PAY_COUNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "wechat_first_pay_count"
	FIRST_ORDER_COUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "first_order_count"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_game_in_app_roi_1day"
	CARD_SHOW_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "card_show"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportAdGetV2OrderField             ReportAdGetV2OrderField = "attribution_active_pay_7d_count"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_day_acitve_pay_count"
	CTR_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "ctr"
	NEXT_DAY_OPEN_RATE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "next_day_open_rate"
	ATTRIBUTION_RETENTION_5D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_5d_cost"
	INSTALL_FINISH_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "install_finish_cost"
	REDIRECT_TO_SHOP_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "redirect_to_shop"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_game_pay_7d_count"
	FOLLOW_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "follow"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportAdGetV2OrderField  ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportAdGetV2OrderField ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_retention_7d_sum_cnt"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_7days"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_8days"
	IN_APP_DETAIL_UV_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "in_app_detail_uv"
	CLICK_WEBSITE_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_website"
	LUBAN_LIVE_FOLLOW_CNT_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "luban_live_follow_cnt"
	LOAN_CREDIT_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "loan_credit"
	LIVE_FANS_CLUB_JOIN_CNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "live_fans_club_join_cnt"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportAdGetV2OrderField             ReportAdGetV2OrderField = "advanced_creative_counsel_click"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportAdGetV2OrderField       ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	SUBMIT_CERTIFICATION_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "submit_certification_count"
	PLAY_OVER_RATE_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "play_over_rate"
	ATTRIBUTION_RETENTION_4D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_4d_cost"
	LOAN_COMPLETION_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "loan_completion"
	WECHAT_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "wechat"
	ATTRIBUTION_RETENTION_6D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_6d_cnt"
	CLICK_INSTALL_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_install"
	MAP_SEARCH_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "map_search"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportAdGetV2OrderField   ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "advanced_creative_phone_click"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_6days"
	ATTRIBUTION_RETENTION_7D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_7d_cost"
	REDIRECT_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "redirect"
	SHOPPING_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "shopping"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "advanced_creative_coupon_addition"
	IN_APP_PAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "in_app_pay"
	DISLIKE_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "dislike"
	PLAY_50_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_50_feed_break_rate"
	LIVE_COMPONENT_CLICK_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "live_component_click_count"
	WECHAT_LOGIN_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "wechat_login_cost"
	LOAN_CREDIT_COST_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "loan_credit_cost"
	COST_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "cost"
	TOTAL_PLAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "total_play"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_4days"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_deep_convert_cost"
	PLAY_50_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_50_feed_break"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_game_in_app_ltv_1day"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_5days"
	DEEP_CONVERT_RATE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "deep_convert_rate"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_game_pay_7d_cost"
	AVERAGE_PLAY_PROGRESS_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "average_play_progress"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_5days"
	CONVERT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "convert"
	GAME_PAY_COST_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "game_pay_cost"
	AVG_SHOW_COST_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "avg_show_cost"
	VALID_PLAY_COST_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "valid_play_cost"
	CPC_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpc"
	SHOW_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "show"
	ACTIVE_PAY_COST_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "active_pay_cost"
	STAT_UNION_LTV_0_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_0"
	ATTRIBUTION_RETENTION_6D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_6d_cost"
	CONVERT_RATE_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "convert_rate"
	ATTRIBUTION_RETENTION_4D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_4d_rate"
	PRE_LOAN_CREDIT_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "pre_loan_credit_cost"
	HOME_VISITED_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "home_visited"
	DOWNLOAD_START_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "download_start_rate"
	PRE_CONVERT_RATE_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "pre_convert_rate"
	LUBAN_LIVE_GIFT_AMOUNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "luban_live_gift_amount"
	AVG_RANK_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "avg_rank"
	ATTRIBUTION_DEEP_CONVERT_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "attribution_deep_convert"
	ATTRIBUTION_RETENTION_7D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_7d_rate"
	PLAY_DURATION_3S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_3s_rate"
	PRE_CONVERT_COST_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "pre_convert_cost"
	DOWNLOAD_FINISH_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "download_finish"
	LIVE_COMPONENT_CLICK_RATE_ReportAdGetV2OrderField                   ReportAdGetV2OrderField = "live_component_click_rate"
	LOAN_COMPLETION_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "loan_completion_rate"
	PLAY_100_FEED_BREAK_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "play_100_feed_break"
	PLAY_DURATION_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "play_duration"
	DOWNLOAD_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "download"
	PLAY_DURATION_SUM_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "play_duration_sum"
	LUBAN_ORDER_STAT_AMOUNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "luban_order_stat_amount"
	ATTRIBUTION_CONVERT_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "attribution_convert"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "advanced_creative_form_submit"
	IN_APP_ORDER_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "in_app_order"
	CONVERT_COST_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "convert_cost"
	ATTRIBUTION_RETENTION_3D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_3d_cnt"
	NEXT_DAY_OPEN_COST_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "next_day_open_cost"
	NEXT_DAY_OPEN_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "next_day_open"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportAdGetV2OrderField       ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	MESSAGE_ACTION_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "message_action"
	PLAY_DURATION_5S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_5s_rate"
	ATTRIBUTION_RETENTION_4D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_4d_cnt"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_3d_ltv"
	PLAY_DURATION_2S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_2s_rate"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_next_day_open_rate"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_next_day_open_cnt"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_wechat_login_30d_cost"
	SHARE_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "share"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportAdGetV2OrderField         ReportAdGetV2OrderField = "attribution_retention_7d_total_cost"
	UNION_ROI_3_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_3"
	LUBAN_ORDER_CNT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "luban_order_cnt"
	ATTRIBUTION_RETENTION_3D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_3d_rate"
	COMMUTE_FIRST_PAY_COUNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "commute_first_pay_count"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_0d_roi"
	PLAY_25_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_25_feed_break"
	DEEP_CONVERT_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "deep_convert_cost"
	GAME_ADDICTION_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "game_addiction_cost"
	COUPON_SINGLE_PAGE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "coupon_single_page"
	AVG_CLICK_COST_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "avg_click_cost"
	LUBAN_LIVE_ENTER_CNT_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "luban_live_enter_cnt"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_3d_roi"
	WIFI_PLAY_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "wifi_play"
	LIKE_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "like"
	DEEP_CONVERT_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "deep_convert"
)

// Ptr returns reference to report_ad_get_v2_order_field value
func (v ReportAdGetV2OrderField) Ptr() *ReportAdGetV2OrderField {
	return &v
}
