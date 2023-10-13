/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeReportOrderGetV10ResponseDataListInner struct for QianchuanAwemeReportOrderGetV10ResponseDataListInner
type QianchuanAwemeReportOrderGetV10ResponseDataListInner struct {
	// 广告计划id
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 点击次数
	ClickCnt *int64 `json:"click_cnt,omitempty"`
	// 直播间商品点击次数
	ClickProduct *int64 `json:"click_product,omitempty"`
	// 转化数
	ConvertCnt *int64 `json:"convert_cnt,omitempty"`
	// 转化成本
	ConvertCost *float64 `json:"convert_cost,omitempty"`
	// 点击率
	Ctr *float64 `json:"ctr,omitempty"`
	// 评论次数
	DyComment *int64 `json:"dy_comment,omitempty"`
	// 新增粉丝数
	DyFollow *int64 `json:"dy_follow,omitempty"`
	// 主页访问量
	DyHomeVisited *int64 `json:"dy_home_visited,omitempty"`
	// 点赞次数
	DyLike *int64 `json:"dy_like,omitempty"`
	// 分享次数
	DyShare *int64 `json:"dy_share,omitempty"`
	// 转换数，替换convert_cnt
	EcpConvertCnt *int64 `json:"ecp_convert_cnt,omitempty"`
	// 转换成本.替换convert_cost。单位元
	EcpCpaPlatform *float64 `json:"ecp_cpa_platform,omitempty"`
	// 点赞率（视频订单才会返回）
	LikeRate *float64 `json:"like_rate,omitempty"`
	// 直播间新加团人次
	LiveFansClubJoinCnt *int64 `json:"live_fans_club_join_cnt,omitempty"`
	// 直播间查看购物车次数
	LiveSlidecartClickCnt *int64 `json:"live_slidecart_click_cnt,omitempty"`
	// 商品点击次数
	LubanLiveClickProductCnt *int64 `json:"luban_live_click_product_cnt,omitempty"`
	// 直播间评论次数
	LubanLiveCommentCnt *int64 `json:"luban_live_comment_cnt,omitempty"`
	// 直播间观看人次
	LubanLiveEnterCnt *int64 `json:"luban_live_enter_cnt,omitempty"`
	// 直播间新增粉丝次数
	LubanLiveFollowCnt *int64 `json:"luban_live_follow_cnt,omitempty"`
	// 直播间音浪收入
	LubanLiveGiftAmount *float64 `json:"luban_live_gift_amount,omitempty"`
	// 直播间打赏次数
	LubanLiveGiftCnt *int64 `json:"luban_live_gift_cnt,omitempty"`
	// 直播间分享次数
	LubanLiveShareCnt *int64 `json:"luban_live_share_cnt,omitempty"`
	// 查看购物车次数
	LubanLiveSlidecartClickCnt *int64                                                `json:"luban_live_slidecart_click_cnt,omitempty"`
	MarketingGoal              *QianchuanAwemeReportOrderGetV10DataListMarketingGoal `json:"marketing_goal,omitempty"`
	// 成交金额
	PayOrderAmount *float64 `json:"pay_order_amount,omitempty"`
	// 成交订单数
	PayOrderCount *int64 `json:"pay_order_count,omitempty"`
	//
	PlayDuration5sRate *float64 `json:"play_duration_5s_rate,omitempty"`
	// 支付roi
	PrepayAndPayOrderRoi *float64 `json:"prepay_and_pay_order_roi,omitempty"`
	// 广告预售订单金额
	PrepayOrderAmount *float64 `json:"prepay_order_amount,omitempty"`
	// 广告预售订单数
	PrepayOrderCount *int64 `json:"prepay_order_count,omitempty"`
	// 有效看播数
	QianchuanEffectiveViewConvertCount *int64 `json:"qianchuan_effective_view_convert_count,omitempty"`
	// 展示次数
	ShowCnt *int64 `json:"show_cnt,omitempty"`
	// 消耗
	StatCost *float64 `json:"stat_cost,omitempty"`
	// 播放次数
	TotalPlay *int64 `json:"total_play,omitempty"`
}
