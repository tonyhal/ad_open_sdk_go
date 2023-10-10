/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPromotionCardRecommendTitleGetV2ContentType
type ToolsPromotionCardRecommendTitleGetV2ContentType string

// List of tools_promotion_card_recommend_title_get_v2_content_type
const (
	PHONE_ToolsPromotionCardRecommendTitleGetV2ContentType          ToolsPromotionCardRecommendTitleGetV2ContentType = "PHONE"
	CARD_ToolsPromotionCardRecommendTitleGetV2ContentType           ToolsPromotionCardRecommendTitleGetV2ContentType = "CARD"
	GAME_PACKAGE_ToolsPromotionCardRecommendTitleGetV2ContentType   ToolsPromotionCardRecommendTitleGetV2ContentType = "GAME_PACKAGE"
	GAME_FORM_ToolsPromotionCardRecommendTitleGetV2ContentType      ToolsPromotionCardRecommendTitleGetV2ContentType = "GAME_FORM"
	TABLE_ToolsPromotionCardRecommendTitleGetV2ContentType          ToolsPromotionCardRecommendTitleGetV2ContentType = "TABLE"
	GAME_SUBSCRIBE_ToolsPromotionCardRecommendTitleGetV2ContentType ToolsPromotionCardRecommendTitleGetV2ContentType = "GAME_SUBSCRIBE"
	CONSULT_ToolsPromotionCardRecommendTitleGetV2ContentType        ToolsPromotionCardRecommendTitleGetV2ContentType = "CONSULT"
	UNKNOWN_ToolsPromotionCardRecommendTitleGetV2ContentType        ToolsPromotionCardRecommendTitleGetV2ContentType = "UNKNOWN"
	DOWNLOAD_ToolsPromotionCardRecommendTitleGetV2ContentType       ToolsPromotionCardRecommendTitleGetV2ContentType = "DOWNLOAD"
	LANDING_ToolsPromotionCardRecommendTitleGetV2ContentType        ToolsPromotionCardRecommendTitleGetV2ContentType = "LANDING"
)

// All allowed values of ToolsPromotionCardRecommendTitleGetV2ContentType enum
var AllowedToolsPromotionCardRecommendTitleGetV2ContentTypeEnumValues = []ToolsPromotionCardRecommendTitleGetV2ContentType{
	"PHONE",
	"CARD",
	"GAME_PACKAGE",
	"GAME_FORM",
	"TABLE",
	"GAME_SUBSCRIBE",
	"CONSULT",
	"UNKNOWN",
	"DOWNLOAD",
	"LANDING",
}

// NewToolsPromotionCardRecommendTitleGetV2ContentTypeFromValue returns a pointer to a valid ToolsPromotionCardRecommendTitleGetV2ContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPromotionCardRecommendTitleGetV2ContentTypeFromValue(v string) (*ToolsPromotionCardRecommendTitleGetV2ContentType, error) {
	ev := ToolsPromotionCardRecommendTitleGetV2ContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPromotionCardRecommendTitleGetV2ContentType: valid values are %v", v, AllowedToolsPromotionCardRecommendTitleGetV2ContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPromotionCardRecommendTitleGetV2ContentType) IsValid() bool {
	for _, existing := range AllowedToolsPromotionCardRecommendTitleGetV2ContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_promotion_card_recommend_title_get_v2_content_type value
func (v ToolsPromotionCardRecommendTitleGetV2ContentType) Ptr() *ToolsPromotionCardRecommendTitleGetV2ContentType {
	return &v
}
