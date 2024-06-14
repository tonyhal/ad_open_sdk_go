/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCampaignGetV2ResponseDataListInner struct for ReportCampaignGetV2ResponseDataListInner
type ReportCampaignGetV2ResponseDataListInner struct {
	Ac *ReportCampaignGetV2DataListAc `json:"ac,omitempty"`
	//
	Active *int64 `json:"active,omitempty"`
	//
	ActiveCost *float64 `json:"active_cost,omitempty"`
	//
	ActivePayAmount *int64 `json:"active_pay_amount,omitempty"`
	//
	ActivePayCost *float64 `json:"active_pay_cost,omitempty"`
	//
	ActivePayRate *float64 `json:"active_pay_rate,omitempty"`
	//
	ActiveRate *float64 `json:"active_rate,omitempty"`
	//
	ActiveRegisterCost *float64 `json:"active_register_cost,omitempty"`
	//
	ActiveRegisterRate *float64 `json:"active_register_rate,omitempty"`
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AdName *string `json:"ad_name,omitempty"`
	//
	AdTag *string `json:"ad_tag,omitempty"`
	//
	AdvancedCreativeCounselClick *int64 `json:"advanced_creative_counsel_click,omitempty"`
	//
	AdvancedCreativeCouponAddition *int64 `json:"advanced_creative_coupon_addition,omitempty"`
	//
	AdvancedCreativeFormClick *int64 `json:"advanced_creative_form_click,omitempty"`
	//
	AdvancedCreativeFormSubmit *int64 `json:"advanced_creative_form_submit,omitempty"`
	//
	AdvancedCreativePhoneClick *int64 `json:"advanced_creative_phone_click,omitempty"`
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	Age *string `json:"age,omitempty"`
	//
	ApprovalCount *float64 `json:"approval_count,omitempty"`
	//
	AttributionActivePay7dCost *float64 `json:"attribution_active_pay_7d_cost,omitempty"`
	//
	AttributionActivePay7dCount *int64 `json:"attribution_active_pay_7d_count,omitempty"`
	//
	AttributionActivePay7dPerCount *int64 `json:"attribution_active_pay_7d_per_count,omitempty"`
	//
	AttributionActivePay7dRate *float64 `json:"attribution_active_pay_7d_rate,omitempty"`
	//
	AttributionActivePayIntraOneDayAmount *float64 `json:"attribution_active_pay_intra_one_day_amount,omitempty"`
	//
	AttributionActivePayIntraOneDayCost *float64 `json:"attribution_active_pay_intra_one_day_cost,omitempty"`
	//
	AttributionActivePayIntraOneDayCount *int64 `json:"attribution_active_pay_intra_one_day_count,omitempty"`
	//
	AttributionActivePayIntraOneDayRate *float64 `json:"attribution_active_pay_intra_one_day_rate,omitempty"`
	//
	AttributionActivePayIntraOneDayRoi *float64 `json:"attribution_active_pay_intra_one_day_roi,omitempty"`
	//
	AttributionConvert *int64 `json:"attribution_convert,omitempty"`
	//
	AttributionConvertCost *float64 `json:"attribution_convert_cost,omitempty"`
	//
	AttributionDayAcitvePayCount *int64 `json:"attribution_day_acitve_pay_count,omitempty"`
	//
	AttributionDayActivePayCount *int64 `json:"attribution_day_active_pay_count,omitempty"`
	//
	AttributionDeepConvert *int64 `json:"attribution_deep_convert,omitempty"`
	//
	AttributionDeepConvertCost *float64 `json:"attribution_deep_convert_cost,omitempty"`
	//
	AttributionGameInAppLtv1day *float64 `json:"attribution_game_in_app_ltv_1day,omitempty"`
	//
	AttributionGameInAppLtv2days *float64 `json:"attribution_game_in_app_ltv_2days,omitempty"`
	//
	AttributionGameInAppLtv3days *float64 `json:"attribution_game_in_app_ltv_3days,omitempty"`
	//
	AttributionGameInAppLtv4days *float64 `json:"attribution_game_in_app_ltv_4days,omitempty"`
	//
	AttributionGameInAppLtv5days *float64 `json:"attribution_game_in_app_ltv_5days,omitempty"`
	//
	AttributionGameInAppLtv6days *float64 `json:"attribution_game_in_app_ltv_6days,omitempty"`
	//
	AttributionGameInAppLtv7days *float64 `json:"attribution_game_in_app_ltv_7days,omitempty"`
	//
	AttributionGameInAppLtv8days *float64 `json:"attribution_game_in_app_ltv_8days,omitempty"`
	//
	AttributionGameInAppRoi1day *float64 `json:"attribution_game_in_app_roi_1day,omitempty"`
	//
	AttributionGameInAppRoi2days *float64 `json:"attribution_game_in_app_roi_2days,omitempty"`
	//
	AttributionGameInAppRoi3days *float64 `json:"attribution_game_in_app_roi_3days,omitempty"`
	//
	AttributionGameInAppRoi4days *float64 `json:"attribution_game_in_app_roi_4days,omitempty"`
	//
	AttributionGameInAppRoi5days *float64 `json:"attribution_game_in_app_roi_5days,omitempty"`
	//
	AttributionGameInAppRoi6days *float64 `json:"attribution_game_in_app_roi_6days,omitempty"`
	//
	AttributionGameInAppRoi7days *float64 `json:"attribution_game_in_app_roi_7days,omitempty"`
	//
	AttributionGameInAppRoi8days *float64 `json:"attribution_game_in_app_roi_8days,omitempty"`
	//
	AttributionGamePay7dCost *float64 `json:"attribution_game_pay_7d_cost,omitempty"`
	//
	AttributionGamePay7dCount *int64 `json:"attribution_game_pay_7d_count,omitempty"`
	//
	AttributionMicroGame0dLtv *float64 `json:"attribution_micro_game_0d_ltv,omitempty"`
	//
	AttributionMicroGame0dRoi *float64 `json:"attribution_micro_game_0d_roi,omitempty"`
	//
	AttributionMicroGame3dLtv *float64 `json:"attribution_micro_game_3d_ltv,omitempty"`
	//
	AttributionMicroGame3dRoi *float64 `json:"attribution_micro_game_3d_roi,omitempty"`
	//
	AttributionMicroGame7dLtv *float64 `json:"attribution_micro_game_7d_ltv,omitempty"`
	//
	AttributionMicroGame7dRoi *float64 `json:"attribution_micro_game_7d_roi,omitempty"`
	//
	AttributionNextDayOpenCnt *int64 `json:"attribution_next_day_open_cnt,omitempty"`
	//
	AttributionNextDayOpenCost *float64 `json:"attribution_next_day_open_cost,omitempty"`
	//
	AttributionNextDayOpenRate *float64 `json:"attribution_next_day_open_rate,omitempty"`
	//
	AttributionRetention2dCnt *int64 `json:"attribution_retention_2d_cnt,omitempty"`
	//
	AttributionRetention2dCost *float64 `json:"attribution_retention_2d_cost,omitempty"`
	//
	AttributionRetention2dRate *float64 `json:"attribution_retention_2d_rate,omitempty"`
	//
	AttributionRetention3dCnt *int64 `json:"attribution_retention_3d_cnt,omitempty"`
	//
	AttributionRetention3dCost *float64 `json:"attribution_retention_3d_cost,omitempty"`
	//
	AttributionRetention3dRate *float64 `json:"attribution_retention_3d_rate,omitempty"`
	//
	AttributionRetention4dCnt *int64 `json:"attribution_retention_4d_cnt,omitempty"`
	//
	AttributionRetention4dCost *float64 `json:"attribution_retention_4d_cost,omitempty"`
	//
	AttributionRetention4dRate *float64 `json:"attribution_retention_4d_rate,omitempty"`
	//
	AttributionRetention5dCnt *int64 `json:"attribution_retention_5d_cnt,omitempty"`
	//
	AttributionRetention5dCost *float64 `json:"attribution_retention_5d_cost,omitempty"`
	//
	AttributionRetention5dRate *float64 `json:"attribution_retention_5d_rate,omitempty"`
	//
	AttributionRetention6dCnt *int64 `json:"attribution_retention_6d_cnt,omitempty"`
	//
	AttributionRetention6dCost *float64 `json:"attribution_retention_6d_cost,omitempty"`
	//
	AttributionRetention6dRate *float64 `json:"attribution_retention_6d_rate,omitempty"`
	//
	AttributionRetention7dCnt *int64 `json:"attribution_retention_7d_cnt,omitempty"`
	//
	AttributionRetention7dCost *float64 `json:"attribution_retention_7d_cost,omitempty"`
	//
	AttributionRetention7dRate *float64 `json:"attribution_retention_7d_rate,omitempty"`
	//
	AttributionRetention7dSumCnt *int64 `json:"attribution_retention_7d_sum_cnt,omitempty"`
	//
	AttributionRetention7dTotalCost *float64 `json:"attribution_retention_7d_total_cost,omitempty"`
	//
	AttributionWechatFirstPay30dCost *float64 `json:"attribution_wechat_first_pay_30d_cost,omitempty"`
	//
	AttributionWechatFirstPay30dCount *int64 `json:"attribution_wechat_first_pay_30d_count,omitempty"`
	//
	AttributionWechatFirstPay30dRate *float64 `json:"attribution_wechat_first_pay_30d_rate,omitempty"`
	//
	AttributionWechatLogin30dCost *float64 `json:"attribution_wechat_login_30d_cost,omitempty"`
	//
	AttributionWechatLogin30dCount *int64 `json:"attribution_wechat_login_30d_count,omitempty"`
	//
	AttributionWechatPay30dAmount *float64 `json:"attribution_wechat_pay_30d_amount,omitempty"`
	//
	AttributionWechatPay30dRoi *float64 `json:"attribution_wechat_pay_30d_roi,omitempty"`
	//
	AveragePlayProgress *float64 `json:"average_play_progress,omitempty"`
	//
	AveragePlayTimePerPlay *float64 `json:"average_play_time_per_play,omitempty"`
	//
	AverageVideoPlay *float64 `json:"average_video_play,omitempty"`
	//
	AvgClickCost *float64 `json:"avg_click_cost,omitempty"`
	//
	AvgRank *float64 `json:"avg_rank,omitempty"`
	//
	AvgShowCost *float64 `json:"avg_show_cost,omitempty"`
	//
	Bidword *string `json:"bidword,omitempty"`
	//
	BidwordId *int64 `json:"bidword_id,omitempty"`
	//
	Button *int64 `json:"button,omitempty"`
	//
	CampaignId *int64 `json:"campaign_id,omitempty"`
	//
	CampaignName *string                                  `json:"campaign_name,omitempty"`
	CampaignType *ReportCampaignGetV2DataListCampaignType `json:"campaign_type,omitempty"`
	//
	CardShow *int64 `json:"card_show,omitempty"`
	//
	CityName *string `json:"city_name,omitempty"`
	//
	Click *int64 `json:"click,omitempty"`
	//
	ClickCallDy *int64 `json:"click_call_dy,omitempty"`
	//
	ClickDownload *int64 `json:"click_download,omitempty"`
	//
	ClickInstall *int64 `json:"click_install,omitempty"`
	//
	ClickLandingPage *int64 `json:"click_landing_page,omitempty"`
	//
	ClickShopwindow *int64 `json:"click_shopwindow,omitempty"`
	//
	ClickWebsite *int64 `json:"click_website,omitempty"`
	//
	Comment *int64 `json:"comment,omitempty"`
	//
	CommuteFirstPayCount *float64 `json:"commute_first_pay_count,omitempty"`
	//
	Consult *int64 `json:"consult,omitempty"`
	//
	ConsultEffective *int64 `json:"consult_effective,omitempty"`
	//
	Convert *int64 `json:"convert,omitempty"`
	//
	ConvertCost *float64 `json:"convert_cost,omitempty"`
	//
	ConvertRate *float64 `json:"convert_rate,omitempty"`
	//
	ConvertShowRate *float64 `json:"convert_show_rate,omitempty"`
	//
	Cost *float64 `json:"cost,omitempty"`
	//
	Coupon *int64 `json:"coupon,omitempty"`
	//
	CouponSinglePage *int64 `json:"coupon_single_page,omitempty"`
	//
	Cpa *float64 `json:"cpa,omitempty"`
	//
	Cpc *float64 `json:"cpc,omitempty"`
	//
	Cpm *float64 `json:"cpm,omitempty"`
	//
	CreativeId           *int64                                           `json:"creative_id,omitempty"`
	CreativeMaterialMode *ReportCampaignGetV2DataListCreativeMaterialMode `json:"creative_material_mode,omitempty"`
	//
	Ctr *float64 `json:"ctr,omitempty"`
	//
	CustomerEffective *int64 `json:"customer_effective,omitempty"`
	//
	DeepConvert *int64 `json:"deep_convert,omitempty"`
	//
	DeepConvertCost *float64 `json:"deep_convert_cost,omitempty"`
	//
	DeepConvertRate *float64 `json:"deep_convert_rate,omitempty"`
	//
	Dislike *int64 `json:"dislike,omitempty"`
	//
	Download *int64 `json:"download,omitempty"`
	//
	DownloadFinish *int64 `json:"download_finish,omitempty"`
	//
	DownloadFinishCost *float64 `json:"download_finish_cost,omitempty"`
	//
	DownloadFinishRate *float64 `json:"download_finish_rate,omitempty"`
	//
	DownloadStart *int64 `json:"download_start,omitempty"`
	//
	DownloadStartCost *float64 `json:"download_start_cost,omitempty"`
	//
	DownloadStartRate *float64                                   `json:"download_start_rate,omitempty"`
	ExternalAction    *ReportCampaignGetV2DataListExternalAction `json:"external_action,omitempty"`
	//
	FirstOrderCount *float64 `json:"first_order_count,omitempty"`
	//
	FirstRentalOrderCount *float64 `json:"first_rental_order_count,omitempty"`
	//
	Follow *int64 `json:"follow,omitempty"`
	//
	Form *int64 `json:"form,omitempty"`
	//
	GameAddiction *int64 `json:"game_addiction,omitempty"`
	//
	GameAddictionCost *float64 `json:"game_addiction_cost,omitempty"`
	//
	GameAddictionRate *float64 `json:"game_addiction_rate,omitempty"`
	//
	GamePayCost *int64 `json:"game_pay_cost,omitempty"`
	//
	GamePayCount *int64                             `json:"game_pay_count,omitempty"`
	Gender       *ReportCampaignGetV2DataListGender `json:"gender,omitempty"`
	//
	HomeVisited *int64 `json:"home_visited,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	//
	IesChallengeClick *int64 `json:"ies_challenge_click,omitempty"`
	//
	IesMusicClick *int64                                `json:"ies_music_click,omitempty"`
	ImageMode     *ReportCampaignGetV2DataListImageMode `json:"image_mode,omitempty"`
	//
	InAppCart *int64 `json:"in_app_cart,omitempty"`
	//
	InAppDetailUv *int64 `json:"in_app_detail_uv,omitempty"`
	//
	InAppOrder *int64 `json:"in_app_order,omitempty"`
	//
	InAppPay *int64 `json:"in_app_pay,omitempty"`
	//
	InAppUv *int64 `json:"in_app_uv,omitempty"`
	//
	InstallFinish *int64 `json:"install_finish,omitempty"`
	//
	InstallFinishCost *float64 `json:"install_finish_cost,omitempty"`
	//
	InstallFinishRate *float64 `json:"install_finish_rate,omitempty"`
	//
	InteractPerCost *float64 `json:"interact_per_cost,omitempty"`
	//
	InterestTag *string                                 `json:"interest_tag,omitempty"`
	Inventory   *ReportCampaignGetV2DataListInventory   `json:"inventory,omitempty"`
	LandingType *ReportCampaignGetV2DataListLandingType `json:"landing_type,omitempty"`
	//
	Like *int64 `json:"like,omitempty"`
	//
	LiveComponentClickCost *float64 `json:"live_component_click_cost,omitempty"`
	//
	LiveComponentClickCount *int64 `json:"live_component_click_count,omitempty"`
	//
	LiveComponentClickRate *float64 `json:"live_component_click_rate,omitempty"`
	//
	LiveFansClubJoinCnt *int64 `json:"live_fans_club_join_cnt,omitempty"`
	//
	LiveWatchOneMinuteCount *int64 `json:"live_watch_one_minute_count,omitempty"`
	//
	LoanCompletion *int64 `json:"loan_completion,omitempty"`
	//
	LoanCompletionCost *float64 `json:"loan_completion_cost,omitempty"`
	//
	LoanCompletionRate *float64 `json:"loan_completion_rate,omitempty"`
	//
	LoanCredit *int64 `json:"loan_credit,omitempty"`
	//
	LoanCreditCost *float64 `json:"loan_credit_cost,omitempty"`
	//
	LoanCreditRate *float64 `json:"loan_credit_rate,omitempty"`
	//
	LocationClick *int64 `json:"location_click,omitempty"`
	//
	Lottery *int64 `json:"lottery,omitempty"`
	//
	LubanLiveClickProductCnt *int64 `json:"luban_live_click_product_cnt,omitempty"`
	//
	LubanLiveCommentCnt *int64 `json:"luban_live_comment_cnt,omitempty"`
	//
	LubanLiveEnterCnt *int64 `json:"luban_live_enter_cnt,omitempty"`
	//
	LubanLiveFollowCnt *int64 `json:"luban_live_follow_cnt,omitempty"`
	//
	LubanLiveGiftAmount *int64 `json:"luban_live_gift_amount,omitempty"`
	//
	LubanLiveGiftCnt *int64 `json:"luban_live_gift_cnt,omitempty"`
	//
	LubanLivePayOrderCount *int64 `json:"luban_live_pay_order_count,omitempty"`
	//
	LubanLivePayOrderStatCost *float64 `json:"luban_live_pay_order_stat_cost,omitempty"`
	//
	LubanLiveShareCnt *int64 `json:"luban_live_share_cnt,omitempty"`
	//
	LubanLiveSlidecartClickCnt *int64 `json:"luban_live_slidecart_click_cnt,omitempty"`
	//
	LubanOrderCnt *int64 `json:"luban_order_cnt,omitempty"`
	//
	LubanOrderRoi *float64 `json:"luban_order_roi,omitempty"`
	//
	LubanOrderStatAmount *float64 `json:"luban_order_stat_amount,omitempty"`
	//
	MapSearch *int64 `json:"map_search,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	Message *int64 `json:"message,omitempty"`
	//
	MessageAction *int64 `json:"message_action,omitempty"`
	//
	NextDayOpen *int64 `json:"next_day_open,omitempty"`
	//
	NextDayOpenCost *float64 `json:"next_day_open_cost,omitempty"`
	//
	NextDayOpenRate *float64 `json:"next_day_open_rate,omitempty"`
	//
	PayAmountRoi *float64 `json:"pay_amount_roi,omitempty"`
	//
	PayCount *int64 `json:"pay_count,omitempty"`
	//
	Phone *int64 `json:"phone,omitempty"`
	//
	PhoneConfirm *int64 `json:"phone_confirm,omitempty"`
	//
	PhoneConnect *int64 `json:"phone_connect,omitempty"`
	//
	PhoneEffective *int64 `json:"phone_effective,omitempty"`
	//
	Platform *string `json:"platform,omitempty"`
	//
	Play100FeedBreak *int64 `json:"play_100_feed_break,omitempty"`
	//
	Play100FeedBreakRate *float64 `json:"play_100_feed_break_rate,omitempty"`
	//
	Play25FeedBreak *int64 `json:"play_25_feed_break,omitempty"`
	//
	Play25FeedBreakRate *float64 `json:"play_25_feed_break_rate,omitempty"`
	//
	Play50FeedBreak *int64 `json:"play_50_feed_break,omitempty"`
	//
	Play50FeedBreakRate *float64 `json:"play_50_feed_break_rate,omitempty"`
	//
	Play75FeedBreak *int64 `json:"play_75_feed_break,omitempty"`
	//
	Play75FeedBreakRate *float64 `json:"play_75_feed_break_rate,omitempty"`
	//
	PlayDuration *int64 `json:"play_duration,omitempty"`
	//
	PlayDuration10s *int64 `json:"play_duration_10s,omitempty"`
	//
	PlayDuration10sRate *float64 `json:"play_duration_10s_rate,omitempty"`
	//
	PlayDuration2s *int64 `json:"play_duration_2s,omitempty"`
	//
	PlayDuration2sRate *float64 `json:"play_duration_2s_rate,omitempty"`
	//
	PlayDuration3s *int64 `json:"play_duration_3s,omitempty"`
	//
	PlayDuration3sRate *float64 `json:"play_duration_3s_rate,omitempty"`
	//
	PlayDuration5sRate *float64 `json:"play_duration_5s_rate,omitempty"`
	//
	PlayDurationSum *int64 `json:"play_duration_sum,omitempty"`
	//
	PlayOver *int64 `json:"play_over,omitempty"`
	//
	PlayOverRate *float64 `json:"play_over_rate,omitempty"`
	//
	PlayableId *int64 `json:"playable_id,omitempty"`
	//
	PlayableName *string `json:"playable_name,omitempty"`
	//
	PlayableOrientation *string `json:"playable_orientation,omitempty"`
	//
	PlayablePreviewUrl *string `json:"playable_preview_url,omitempty"`
	//
	PlayableUrl *string `json:"playable_url,omitempty"`
	//
	PoiAddressClick *int64 `json:"poi_address_click,omitempty"`
	//
	PoiCollect *int64 `json:"poi_collect,omitempty"`
	//
	PreConvertCost *int64 `json:"pre_convert_cost,omitempty"`
	//
	PreConvertCount *int64 `json:"pre_convert_count,omitempty"`
	//
	PreConvertRate *float64 `json:"pre_convert_rate,omitempty"`
	//
	PreLoanCredit *int64 `json:"pre_loan_credit,omitempty"`
	//
	PreLoanCreditCost *float64                                    `json:"pre_loan_credit_cost,omitempty"`
	Pricing           *ReportCampaignGetV2DataListPricing         `json:"pricing,omitempty"`
	PricingCategory   *ReportCampaignGetV2DataListPricingCategory `json:"pricing_category,omitempty"`
	//
	ProvinceName *string `json:"province_name,omitempty"`
	//
	Qq *int64 `json:"qq,omitempty"`
	//
	Query               *string                                         `json:"query,omitempty"`
	RealRecallMatchType *ReportCampaignGetV2DataListRealRecallMatchType `json:"real_recall_match_type,omitempty"`
	//
	Redirect *int64 `json:"redirect,omitempty"`
	//
	RedirectToShop *int64 `json:"redirect_to_shop,omitempty"`
	//
	Register *int64 `json:"register,omitempty"`
	//
	Report *int64 `json:"report,omitempty"`
	//
	Share *int64 `json:"share,omitempty"`
	//
	Shopping *int64 `json:"shopping,omitempty"`
	//
	Show *int64 `json:"show,omitempty"`
	//
	StatDatetime *string `json:"stat_datetime,omitempty"`
	//
	StatPayAmount *float64 `json:"stat_pay_amount,omitempty"`
	//
	StatUnionLtv0 *float64 `json:"stat_union_ltv_0,omitempty"`
	//
	StatUnionLtv3 *float64 `json:"stat_union_ltv_3,omitempty"`
	//
	StatUnionLtv7 *float64 `json:"stat_union_ltv_7,omitempty"`
	//
	SubmitCertificationCount *float64 `json:"submit_certification_count,omitempty"`
	//
	TotalPlay *int64 `json:"total_play,omitempty"`
	//
	UnionRoi0 *float64 `json:"union_roi_0,omitempty"`
	//
	UnionRoi3 *float64 `json:"union_roi_3,omitempty"`
	//
	UnionRoi7 *float64 `json:"union_roi_7,omitempty"`
	//
	ValidPlay *int64 `json:"valid_play,omitempty"`
	//
	ValidPlayCost *float64 `json:"valid_play_cost,omitempty"`
	//
	ValidPlayRate *float64 `json:"valid_play_rate,omitempty"`
	//
	View *int64 `json:"view,omitempty"`
	//
	Vote *int64 `json:"vote,omitempty"`
	//
	Wechat *int64 `json:"wechat,omitempty"`
	//
	WechatFirstPayCost *float64 `json:"wechat_first_pay_cost,omitempty"`
	//
	WechatFirstPayCount *int64 `json:"wechat_first_pay_count,omitempty"`
	//
	WechatFirstPayRate *float64 `json:"wechat_first_pay_rate,omitempty"`
	//
	WechatLoginCost *float64 `json:"wechat_login_cost,omitempty"`
	//
	WechatLoginCount *int64 `json:"wechat_login_count,omitempty"`
	//
	WechatPayAmount *float64 `json:"wechat_pay_amount,omitempty"`
	//
	WifiPlay *int64 `json:"wifi_play,omitempty"`
	//
	WifiPlayRate *float64 `json:"wifi_play_rate,omitempty"`
}
