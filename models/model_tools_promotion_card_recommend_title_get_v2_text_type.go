/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPromotionCardRecommendTitleGetV2TextType
type ToolsPromotionCardRecommendTitleGetV2TextType string

// List of tools_promotion_card_recommend_title_get_v2_text_type
const (
	CALL_TO_ACTION_ToolsPromotionCardRecommendTitleGetV2TextType ToolsPromotionCardRecommendTitleGetV2TextType = "CALL_TO_ACTION"
	CARD_TITLE_ToolsPromotionCardRecommendTitleGetV2TextType     ToolsPromotionCardRecommendTitleGetV2TextType = "CARD_TITLE"
	PROMOTION_ToolsPromotionCardRecommendTitleGetV2TextType      ToolsPromotionCardRecommendTitleGetV2TextType = "PROMOTION"
)

// All allowed values of ToolsPromotionCardRecommendTitleGetV2TextType enum
var AllowedToolsPromotionCardRecommendTitleGetV2TextTypeEnumValues = []ToolsPromotionCardRecommendTitleGetV2TextType{
	"CALL_TO_ACTION",
	"CARD_TITLE",
	"PROMOTION",
}

// NewToolsPromotionCardRecommendTitleGetV2TextTypeFromValue returns a pointer to a valid ToolsPromotionCardRecommendTitleGetV2TextType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPromotionCardRecommendTitleGetV2TextTypeFromValue(v string) (*ToolsPromotionCardRecommendTitleGetV2TextType, error) {
	ev := ToolsPromotionCardRecommendTitleGetV2TextType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPromotionCardRecommendTitleGetV2TextType: valid values are %v", v, AllowedToolsPromotionCardRecommendTitleGetV2TextTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPromotionCardRecommendTitleGetV2TextType) IsValid() bool {
	for _, existing := range AllowedToolsPromotionCardRecommendTitleGetV2TextTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_promotion_card_recommend_title_get_v2_text_type value
func (v ToolsPromotionCardRecommendTitleGetV2TextType) Ptr() *ToolsPromotionCardRecommendTitleGetV2TextType {
	return &v
}
