/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeDetailGetV30ResponseDataAdData 基础创意信息
type CreativeDetailGetV30ResponseDataAdData struct {
	// 是否允许客户端下载视频
	AdDownloadStatus *int64 `json:"ad_download_status,omitempty"`
	// 创意标签
	AdKeywords []string `json:"ad_keywords,omitempty"`
	//
	AnchorId          *string                                          `json:"anchor_id,omitempty"`
	AnchorRelatedType *CreativeDetailGetV30DataAdDataAnchorRelatedType `json:"anchor_related_type,omitempty"`
	AnchorType        *CreativeDetailGetV30DataAdDataAnchorType        `json:"anchor_type,omitempty"`
	// 应用名
	AppName *string `json:"app_name,omitempty"`
	// 是否开启自动派生创意
	CreativeAutoGenerateSwitch *int64 `json:"creative_auto_generate_switch,omitempty"`
	// 落地页链接字段选择
	DpaExternalUrlField *string `json:"dpa_external_url_field,omitempty"`
	// 启用动态创意类型
	DynamicCreativeSwitch []*CreativeDetailGetV30DataAdDataDynamicCreativeSwitch `json:"dynamic_creative_switch,omitempty"`
	EnableSmartSource     *CreativeDetailGetV30DataAdDataEnableSmartSource       `json:"enable_smart_source,omitempty"`
	// 落地页链接
	ExternalUrl *string `json:"external_url,omitempty"`
	// 落地页检测参数
	ExternalUrlParams *string `json:"external_url_params,omitempty"`
	// 绑定抖音号
	IesCoreUserId *string `json:"ies_core_user_id,omitempty"`
	// 是否关闭评论
	IsCommentDisable *int64 `json:"is_comment_disable,omitempty"`
	// 主页作品列表隐藏广告内容
	IsFeedAndFavSee *int64 `json:"is_feed_and_fav_see,omitempty"`
	// 自动生成视频素材
	IsPresentedVideo *int64                                                 `json:"is_presented_video,omitempty"`
	MiniProgramInfo  *CreativeDetailGetV30ResponseDataAdDataMiniProgramInfo `json:"mini_program_info,omitempty"`
	// 直达链接
	OpenUrl    *string                                   `json:"open_url,omitempty"`
	ParamsType *CreativeDetailGetV30DataAdDataParamsType `json:"params_type,omitempty"`
	// 搭配试玩素材URL
	PlayableUrl   *string                                      `json:"playable_url,omitempty"`
	PriorityTrial *CreativeDetailGetV30DataAdDataPriorityTrial `json:"priority_trial,omitempty"`
	// 广告来源
	Source *string `json:"source,omitempty"`
	//
	Supplements []*CreativeDetailGetV30ResponseDataAdDataSupplementsInner `json:"supplements,omitempty"`
	// 三级行业ID
	ThirdIndustryId *int64 `json:"third_industry_id,omitempty"`
	// Android应用下载详情页
	WebUrl *string `json:"web_url,omitempty"`
}
