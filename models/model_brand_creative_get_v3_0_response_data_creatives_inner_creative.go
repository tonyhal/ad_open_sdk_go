/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreative 创意详情信息
type BrandCreativeGetV30ResponseDataCreativesInnerCreative struct {
	AdvancedCreativeInfo *BrandCreativeGetV30ResponseDataCreativesInnerCreativeAdvancedCreativeInfo `json:"advanced_creative_info,omitempty"`
	AvatarIcon           *BrandCreativeGetV30ResponseDataCreativesInnerCreativeAvatarIcon           `json:"avatar_icon,omitempty"`
	AwemeItemInfo        *BrandCreativeGetV30ResponseDataCreativesInnerCreativeAwemeItemInfo        `json:"aweme_item_info,omitempty"`
	BusinessInfo         *BrandCreativeGetV30ResponseDataCreativesInnerCreativeBusinessInfo         `json:"business_info,omitempty"`
	// 创意ID
	CreativeId   *string                                                            `json:"creative_id,omitempty"`
	CreativeInfo *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfo `json:"creative_info,omitempty"`
	ExternalInfo *BrandCreativeGetV30ResponseDataCreativesInnerCreativeExternalInfo `json:"external_info,omitempty"`
	// 文章ID
	GroupId *string `json:"group_id,omitempty"`
	// 抖音用户Id
	IesCoreUserId     *string                                                                 `json:"ies_core_user_id,omitempty"`
	InteractiveModule *BrandCreativeGetV30ResponseDataCreativesInnerCreativeInteractiveModule `json:"interactive_module,omitempty"`
	// 来源
	Source *string `json:"source,omitempty"`
	// 开屏按钮文案
	SplashButtonText *string                                                              `json:"splash_button_text,omitempty"`
	SplashCreative   *BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreative `json:"splash_creative,omitempty"`
	// 标题
	Title *string `json:"title,omitempty"`
	// 按钮颜色配置
	VideoThemeColor *string `json:"video_theme_color,omitempty"`
}