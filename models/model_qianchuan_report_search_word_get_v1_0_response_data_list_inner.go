/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportSearchWordGetV10ResponseDataListInner struct for QianchuanReportSearchWordGetV10ResponseDataListInner
type QianchuanReportSearchWordGetV10ResponseDataListInner struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AdName *string `json:"ad_name,omitempty"`
	//
	CampaignId *int64 `json:"campaign_id,omitempty"`
	//
	CampaignName *string `json:"campaign_name,omitempty"`
	//
	ClickCnt *float64 `json:"click_cnt,omitempty"`
	// todo
	CpmPlatform *float64 `json:"cpm_platform,omitempty"`
	//
	CreateOrderAmount *float64 `json:"create_order_amount,omitempty"`
	//
	CreateOrderCount *float64 `json:"create_order_count,omitempty"`
	//
	CreateOrderRoi *float64 `json:"create_order_roi,omitempty"`
	//
	Ctr *float64 `json:"ctr,omitempty"`
	//
	DyComment *float64 `json:"dy_comment,omitempty"`
	//
	DyFollow *float64 `json:"dy_follow,omitempty"`
	//
	DyLike *float64 `json:"dy_like,omitempty"`
	//
	DyShare *float64 `json:"dy_share,omitempty"`
	//
	EcpConvertCnt *float64 `json:"ecp_convert_cnt,omitempty"`
	//
	EcpConvertRate *float64 `json:"ecp_convert_rate,omitempty"`
	// todo
	EcpCpaPlatform *float64 `json:"ecp_cpa_platform,omitempty"`
	//
	EndDate *string `json:"end_date,omitempty"`
	//
	KeyWord          *string                                                  `json:"key_word,omitempty"`
	KeywordMatchType *QianchuanReportSearchWordGetV10DataListKeywordMatchType `json:"keyword_match_type,omitempty"`
	//
	LiveFansClubJoinCnt *float64 `json:"live_fans_club_join_cnt,omitempty"`
	//
	LiveWatchOneMinuteCount *float64 `json:"live_watch_one_minute_count,omitempty"`
	//
	LubanLiveClickProductCnt *float64 `json:"luban_live_click_product_cnt,omitempty"`
	//
	LubanLiveCommentCnt *float64 `json:"luban_live_comment_cnt,omitempty"`
	//
	LubanLiveEnterCnt *float64 `json:"luban_live_enter_cnt,omitempty"`
	//
	LubanLiveGiftAmount *float64 `json:"luban_live_gift_amount,omitempty"`
	//
	LubanLiveGiftCnt *float64 `json:"luban_live_gift_cnt,omitempty"`
	//
	LubanLiveShareCnt *float64 `json:"luban_live_share_cnt,omitempty"`
	//
	LubanLiveSlidecartClickCnt *float64 `json:"luban_live_slidecart_click_cnt,omitempty"`
	//
	PayOrderAmount *float64 `json:"pay_order_amount,omitempty"`
	//
	PayOrderCostPerOrder *float64 `json:"pay_order_cost_per_order,omitempty"`
	//
	PayOrderCount *float64 `json:"pay_order_count,omitempty"`
	//
	Play25FeedBreak *float64 `json:"play_25_feed_break,omitempty"`
	//
	Play50FeedBreak *float64 `json:"play_50_feed_break,omitempty"`
	//
	Play75FeedBreak *float64 `json:"play_75_feed_break,omitempty"`
	//
	PlayDuration3s *float64 `json:"play_duration_3s,omitempty"`
	//
	PlayOver *float64 `json:"play_over,omitempty"`
	//
	PlayOverRate *float64 `json:"play_over_rate,omitempty"`
	//
	PrepayAndPayOrderRoi *float64 `json:"prepay_and_pay_order_roi,omitempty"`
	//
	PrepayOrderAmount *float64 `json:"prepay_order_amount,omitempty"`
	//
	PrepayOrderCount *float64 `json:"prepay_order_count,omitempty"`
	//
	SearchWord *string `json:"search_word,omitempty"`
	//
	ShowCnt *float64 `json:"show_cnt,omitempty"`
	//
	StartDate *string `json:"start_date,omitempty"`
	//
	StatCost *float64 `json:"stat_cost,omitempty"`
	//
	TotalPlay *float64 `json:"total_play,omitempty"`
}
