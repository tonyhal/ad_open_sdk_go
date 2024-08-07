/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30RequestNativeSetting
type PromotionCreateV30RequestNativeSetting struct {
	AnchorRelatedType *PromotionCreateV30NativeSettingAnchorRelatedType `json:"anchor_related_type,omitempty"`
	//
	AwemeId *string `json:"aweme_id,omitempty"`
	//
	AwemeIds         []string                                         `json:"aweme_ids,omitempty"`
	AwemeSettingType *PromotionCreateV30NativeSettingAwemeSettingType `json:"aweme_setting_type,omitempty"`
	//
	ExcludeAwemeIds []string                                        `json:"exclude_aweme_ids,omitempty"`
	IsFeedAndFavSee *PromotionCreateV30NativeSettingIsFeedAndFavSee `json:"is_feed_and_fav_see,omitempty"`
}
