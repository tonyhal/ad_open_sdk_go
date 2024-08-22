/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmGetChallengeItemsDataV2ResponseDataDataListInner struct for StarDemandOmGetChallengeItemsDataV2ResponseDataDataListInner
type StarDemandOmGetChallengeItemsDataV2ResponseDataDataListInner struct {
	//
	AndroidActivateCnt *int64 `json:"android_activate_cnt,omitempty"`
	// 星图达人ID
	AuthorId *int64 `json:"author_id,omitempty"`
	// 达人名称
	AuthorNickname *string `json:"author_nickname,omitempty"`
	//
	CommentCnt *int64 `json:"comment_cnt,omitempty"`
	//
	ComponentClickCnt *int64 `json:"component_click_cnt,omitempty"`
	// 预估广告消耗金额 预估广告消耗金额、数据更新秒级时间戳，示例：{\"value\": 1, \"time\":1703606399}，其中value为视频播放量，time为数据更新秒级时间戳
	EstAdCost *string `json:"est_ad_cost,omitempty"`
	// 预估付费流水金额 预估付费流水、数据更新秒级时间戳，示例：{\"value\": 1, \"time\":1703606399}，其中value为视频播放量，time为数据更新秒级时间戳
	EstSales *string `json:"est_sales,omitempty"`
	//
	IosActivateCnt *int64 `json:"ios_activate_cnt,omitempty"`
	// 视频ID/图文ID
	ItemId *int64 `json:"item_id,omitempty"`
	//
	LikeCnt *int64 `json:"like_cnt,omitempty"`
	//
	Link *string `json:"link,omitempty"`
	// 播放量 视频播放量、数据更新秒级时间戳，示例：{\"value\": 1, \"time\":1703606399}，其中value为视频播放量，time为数据更新秒级时间戳
	Play *string `json:"play,omitempty"`
	//
	PlayVv *int64 `json:"play_vv,omitempty"`
	//
	PromoteCnt *int64 `json:"promote_cnt,omitempty"`
	// 发布时间 秒级时间戳，1658073600 则表示 2022-07-18 00:00:00
	ReleaseTime *int64 `json:"release_time,omitempty"`
	//
	RelevanceAuditResult *int64 `json:"relevance_audit_result,omitempty"`
	//
	RewardAmount *int64 `json:"reward_amount,omitempty"`
	//
	RewardLevel *int64 `json:"reward_level,omitempty"`
	// 已发放达人分成金额 已发放达人分成金额、数据更新秒级时间戳，示例：{\"value\": 1, \"time\":1703606399}，其中value为视频播放量，time为数据更新秒级时间戳
	SettleAdShare *string `json:"settle_ad_share,omitempty"`
	// 已发放达人分佣金额 已发放达人分佣金额、数据更新秒级时间戳，示例：{\"value\": 1, \"time\":1703606399}，其中value为视频播放量，time为数据更新秒级时间戳
	SettleCps *string `json:"settle_cps,omitempty"`
	//
	ShareCnt *int64 `json:"share_cnt,omitempty"`
	// 视频标题
	Title *string `json:"title,omitempty"`
	//
	ValidLikeCnt *int64 `json:"valid_like_cnt,omitempty"`
	//
	ValidPlayVv *int64 `json:"valid_play_vv,omitempty"`
}
