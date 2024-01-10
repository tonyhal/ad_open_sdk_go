/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeCreateV2RequestAdData
type CreativeProceduralCreativeCreateV2RequestAdData struct {
	//
	AdCategory       *int64                                                    `json:"ad_category,omitempty"`
	AdDownloadStatus *CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus `json:"ad_download_status,omitempty"`
	//
	AdKeywords []string `json:"ad_keywords,omitempty"`
	//
	AnchorId          *string                                                    `json:"anchor_id,omitempty"`
	AnchorRelatedType *CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType `json:"anchor_related_type,omitempty"`
	AnchorType        *CreativeProceduralCreativeCreateV2AdDataAnchorType        `json:"anchor_type,omitempty"`
	//
	AppName                    *string                                                             `json:"app_name,omitempty"`
	CreativeAutoGenerateSwitch *CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch `json:"creative_auto_generate_switch,omitempty"`
	//
	DpaExternalUrlField *string `json:"dpa_external_url_field,omitempty"`
	//
	DynamicCreativeSwitch []*CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch `json:"dynamic_creative_switch,omitempty"`
	EnableSmartSource     *CreativeProceduralCreativeCreateV2AdDataEnableSmartSource       `json:"enable_smart_source,omitempty"`
	//
	ExternalUrl *string `json:"external_url,omitempty"`
	//
	ExternalUrlParams *string `json:"external_url_params,omitempty"`
	//
	IesCoreUserId    *string                                                         `json:"ies_core_user_id,omitempty"`
	IsCommentDisable *CreativeProceduralCreativeCreateV2AdDataIsCommentDisable       `json:"is_comment_disable,omitempty"`
	IsFeedAndFavSee  *CreativeProceduralCreativeCreateV2AdDataIsFeedAndFavSee        `json:"is_feed_and_fav_see,omitempty"`
	IsPresentedVideo *CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo       `json:"is_presented_video,omitempty"`
	MiniProgramInfo  *CreativeProceduralCreativeCreateV2RequestAdDataMiniProgramInfo `json:"mini_program_info,omitempty"`
	//
	OpenUrl    *string                                             `json:"open_url,omitempty"`
	ParamsType *CreativeProceduralCreativeCreateV2AdDataParamsType `json:"params_type,omitempty"`
	//
	PlayableUrl   *string                                                `json:"playable_url,omitempty"`
	PriorityTrial *CreativeProceduralCreativeCreateV2AdDataPriorityTrial `json:"priority_trial,omitempty"`
	//
	Source *string `json:"source,omitempty"`
	//
	SubLinkIdList []int64 `json:"sub_link_id_list,omitempty"`
	//
	Supplements []*CreativeProceduralCreativeCreateV2RequestAdDataSupplementsInner `json:"supplements,omitempty"`
	//
	ThirdIndustryId *int64 `json:"third_industry_id,omitempty"`
	//
	WebUrl *string `json:"web_url,omitempty"`
}
