/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPromotionCardRecommendGetV2AdvancedCreativeType
type ToolsPromotionCardRecommendGetV2AdvancedCreativeType string

// List of tools_promotion_card_recommend_get_v2_advanced_creative_type
const (
	ATTACHED_CREATIVE_VOTE_INTERACT_ToolsPromotionCardRecommendGetV2AdvancedCreativeType   ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_VOTE_INTERACT"
	ATTACHED_CREATIVE_CARD_ToolsPromotionCardRecommendGetV2AdvancedCreativeType            ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_CARD"
	ATTACHED_CREATIVE_GAME_SUBSCRIBE_ToolsPromotionCardRecommendGetV2AdvancedCreativeType  ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_GAME_SUBSCRIBE"
	ATTACHED_CREATIVE_PHONE_ToolsPromotionCardRecommendGetV2AdvancedCreativeType           ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_PHONE"
	ATTACHED_CREATIVE_COUPON_ToolsPromotionCardRecommendGetV2AdvancedCreativeType          ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_COUPON"
	ATTACHED_CREATIVE_COMMERCE_CARD_ToolsPromotionCardRecommendGetV2AdvancedCreativeType   ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_COMMERCE_CARD"
	ATTACHED_CREATIVE_FORM_ToolsPromotionCardRecommendGetV2AdvancedCreativeType            ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_FORM"
	ATTACHED_CREATIVE_INTERACT_ToolsPromotionCardRecommendGetV2AdvancedCreativeType        ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_INTERACT"
	ATTACHED_CREATIVE_GAME_FORM_ToolsPromotionCardRecommendGetV2AdvancedCreativeType       ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_GAME_FORM"
	ATTACHED_CREATIVE_SMART_PHONE_ToolsPromotionCardRecommendGetV2AdvancedCreativeType     ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_SMART_PHONE"
	ATTACHED_CREATIVE_DOWNLOAD_CARD_ToolsPromotionCardRecommendGetV2AdvancedCreativeType   ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_DOWNLOAD_CARD"
	ATTACHED_CREATIVE_GAME_PACKAGE_ToolsPromotionCardRecommendGetV2AdvancedCreativeType    ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_GAME_PACKAGE"
	ATTACHED_CREATIVE_LIVE_CARD_ToolsPromotionCardRecommendGetV2AdvancedCreativeType       ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_LIVE_CARD"
	ATTACHED_CREATIVE_CONSULTANT_ToolsPromotionCardRecommendGetV2AdvancedCreativeType      ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_CONSULTANT"
	ATTACHED_CREATIVE_NONE_ToolsPromotionCardRecommendGetV2AdvancedCreativeType            ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_NONE"
	ATTACHED_CREATIVE_APP_ToolsPromotionCardRecommendGetV2AdvancedCreativeType             ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_APP"
	ATTACHED_CREATIVE_COUPON_INTERACT_ToolsPromotionCardRecommendGetV2AdvancedCreativeType ToolsPromotionCardRecommendGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_COUPON_INTERACT"
)

// All allowed values of ToolsPromotionCardRecommendGetV2AdvancedCreativeType enum
var AllowedToolsPromotionCardRecommendGetV2AdvancedCreativeTypeEnumValues = []ToolsPromotionCardRecommendGetV2AdvancedCreativeType{
	"ATTACHED_CREATIVE_VOTE_INTERACT",
	"ATTACHED_CREATIVE_CARD",
	"ATTACHED_CREATIVE_GAME_SUBSCRIBE",
	"ATTACHED_CREATIVE_PHONE",
	"ATTACHED_CREATIVE_COUPON",
	"ATTACHED_CREATIVE_COMMERCE_CARD",
	"ATTACHED_CREATIVE_FORM",
	"ATTACHED_CREATIVE_INTERACT",
	"ATTACHED_CREATIVE_GAME_FORM",
	"ATTACHED_CREATIVE_SMART_PHONE",
	"ATTACHED_CREATIVE_DOWNLOAD_CARD",
	"ATTACHED_CREATIVE_GAME_PACKAGE",
	"ATTACHED_CREATIVE_LIVE_CARD",
	"ATTACHED_CREATIVE_CONSULTANT",
	"ATTACHED_CREATIVE_NONE",
	"ATTACHED_CREATIVE_APP",
	"ATTACHED_CREATIVE_COUPON_INTERACT",
}

// NewToolsPromotionCardRecommendGetV2AdvancedCreativeTypeFromValue returns a pointer to a valid ToolsPromotionCardRecommendGetV2AdvancedCreativeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPromotionCardRecommendGetV2AdvancedCreativeTypeFromValue(v string) (*ToolsPromotionCardRecommendGetV2AdvancedCreativeType, error) {
	ev := ToolsPromotionCardRecommendGetV2AdvancedCreativeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPromotionCardRecommendGetV2AdvancedCreativeType: valid values are %v", v, AllowedToolsPromotionCardRecommendGetV2AdvancedCreativeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPromotionCardRecommendGetV2AdvancedCreativeType) IsValid() bool {
	for _, existing := range AllowedToolsPromotionCardRecommendGetV2AdvancedCreativeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_promotion_card_recommend_get_v2_advanced_creative_type value
func (v ToolsPromotionCardRecommendGetV2AdvancedCreativeType) Ptr() *ToolsPromotionCardRecommendGetV2AdvancedCreativeType {
	return &v
}
