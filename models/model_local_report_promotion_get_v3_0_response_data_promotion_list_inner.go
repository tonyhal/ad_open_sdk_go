/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalReportPromotionGetV30ResponseDataPromotionListInner struct for LocalReportPromotionGetV30ResponseDataPromotionListInner
type LocalReportPromotionGetV30ResponseDataPromotionListInner struct {
	// 私信留资数(计费时间)
	AttributionClueMessageCnt *int64 `json:"attribution_clue_message_cnt,omitempty"`
	// 团购线索数(计费时间)
	AttributionCluePayOrderCnt *int64 `json:"attribution_clue_pay_order_cnt,omitempty"`
	// 转化率(计费时间)
	AttributionConversionRate *float64 `json:"attribution_conversion_rate,omitempty"`
	// 转化数(计费时间)
	AttributionConvertCnt *int64 `json:"attribution_convert_cnt,omitempty"`
	// 转化成本(元)(计费时间)
	AttributionConvertCost *float64 `json:"attribution_convert_cost,omitempty"`
	// 表单提交数(计费时间)
	AttributionFormCnt *int64 `json:"attribution_form_cnt,omitempty"`
	// 意向表单数(计费时间)
	AttributionIntentionFormCnt *int64 `json:"attribution_intention_form_cnt,omitempty"`
	// 意向咨询数(计费时间)
	AttributionIntentionMessageClueCnt *int64 `json:"attribution_intention_message_clue_cnt,omitempty"`
	// 意向话单数(计费时间)
	AttributionIntentionPhoneCnt *int64 `json:"attribution_intention_phone_cnt,omitempty"`
	// 私信咨询数(计费时间)
	AttributionMessageActionCnt *int64 `json:"attribution_message_action_cnt,omitempty"`
	// 电话拨打数(计费时间)
	AttributionPhoneConfirmCnt *int64 `json:"attribution_phone_confirm_cnt,omitempty"`
	// 电话接通数(计费时间)
	AttributionPhoneConnectCnt *int64 `json:"attribution_phone_connect_cnt,omitempty"`
	// 点击次数
	ClickCnt *int64 `json:"click_cnt,omitempty"`
	// 私信留资数
	ClueMessageCount *int64 `json:"clue_message_count,omitempty"`
	// 团购线索数
	CluePayOrderCnt *int64 `json:"clue_pay_order_cnt,omitempty"`
	// 转化成本(元)
	ConversionCost *float64 `json:"conversion_cost,omitempty"`
	// 转化率
	ConversionRate *float64 `json:"conversion_rate,omitempty"`
	// 转化数
	ConvertCnt *int64 `json:"convert_cnt,omitempty"`
	// 点击均价(元)
	CpcPlatform *float64 `json:"cpc_platform,omitempty"`
	// 平均千次展示费用(元)
	CpmPlatform *float64 `json:"cpm_platform,omitempty"`
	// 点击率
	Ctr *float64 `json:"ctr,omitempty"`
	// 视频收藏次数
	DyCollect *int64 `json:"dy_collect,omitempty"`
	// 视频评论次数
	DyComment *int64 `json:"dy_comment,omitempty"`
	// 粉丝量
	DyFollow *int64 `json:"dy_follow,omitempty"`
	// 视频点赞次数
	DyLike *int64 `json:"dy_like,omitempty"`
	// 视频点赞率
	DyLikeRate *float64 `json:"dy_like_rate,omitempty"`
	// 视频分享次数
	DyShare *int64 `json:"dy_share,omitempty"`
	// 表单提交数
	FormCnt *int64 `json:"form_cnt,omitempty"`
	// 意向表单数
	IntentionFormCnt *int64 `json:"intention_form_cnt,omitempty"`
	// 意向咨询数
	IntentionMessageClueCnt *int64 `json:"intention_message_clue_cnt,omitempty"`
	// 意向话单数
	IntentionPhoneCnt *int64 `json:"intention_phone_cnt,omitempty"`
	// 直播间超1分钟停留次数
	LiveWatchOneMinuteCount *int64 `json:"live_watch_one_minute_count,omitempty"`
	// 直播间评论次数
	LubanLiveCommentCnt *int64 `json:"luban_live_comment_cnt,omitempty"`
	// 直播间观看次数
	LubanLiveEnterCnt *int64 `json:"luban_live_enter_cnt,omitempty"`
	// 直播间分享次数
	LubanLiveShareCnt *int64 `json:"luban_live_share_cnt,omitempty"`
	// 私信咨询数
	MessageActionCnt *int64 `json:"message_action_cnt,omitempty"`
	// 电话拨打数
	PhoneConfirmCnt *int64 `json:"phone_confirm_cnt,omitempty"`
	// 电话接通数
	PhoneConnectCnt *int64 `json:"phone_connect_cnt,omitempty"`
	// 视频25%进度播放次数
	Play25FeedBreak *int64 `json:"play_25_feed_break,omitempty"`
	// 视频50%进度播放次数
	Play50FeedBreak *int64 `json:"play_50_feed_break,omitempty"`
	// 视频75%进度播放次数
	Play75FeedBreak *int64 `json:"play_75_feed_break,omitempty"`
	// 视频3s播放次数
	PlayDuration3s *int64 `json:"play_duration_3s,omitempty"`
	// 视频5s播放次数
	PlayDuration5s *int64 `json:"play_duration_5s,omitempty"`
	// 视频5s播放率
	PlayDuration5sShowCntRate *float64 `json:"play_duration_5s_show_cnt_rate,omitempty"`
	// 视频播放完成次数
	PlayOver *int64 `json:"play_over,omitempty"`
	// 视频完播率
	PlayOverRate *float64 `json:"play_over_rate,omitempty"`
	// 项目id
	ProjectId *int64 `json:"project_id,omitempty"`
	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`
	// 广告id
	PromotionId *int64 `json:"promotion_id,omitempty"`
	// 广告名称
	PromotionName *string `json:"promotion_name,omitempty"`
	// 展示次数
	ShowCnt *int64 `json:"show_cnt,omitempty"`
	// 消耗(元)
	StatCost *float64 `json:"stat_cost,omitempty"`
	// 时间-天
	StatTimeDay *string `json:"stat_time_day,omitempty"`
	// 时间-小时
	StatTimeHour *string `json:"stat_time_hour,omitempty"`
	// 视频播放次数
	TotalPlay *int64 `json:"total_play,omitempty"`
}