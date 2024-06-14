/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdConvertOptimizedTargetGetV2PromotionContent
type ToolsAdConvertOptimizedTargetGetV2PromotionContent string

// List of tools_ad_convert_optimized_target_get_v2_promotion_content
const (
	DOUYIN_ToolsAdConvertOptimizedTargetGetV2PromotionContent          ToolsAdConvertOptimizedTargetGetV2PromotionContent = "DOUYIN"
	NORMAL_ToolsAdConvertOptimizedTargetGetV2PromotionContent          ToolsAdConvertOptimizedTargetGetV2PromotionContent = "NORMAL"
	AWEME_HOME_PAGE_ToolsAdConvertOptimizedTargetGetV2PromotionContent ToolsAdConvertOptimizedTargetGetV2PromotionContent = "AWEME_HOME_PAGE"
	QUICK_APP_URL_ToolsAdConvertOptimizedTargetGetV2PromotionContent   ToolsAdConvertOptimizedTargetGetV2PromotionContent = "QUICK_APP_URL"
	DOWNLOAD_URL_ToolsAdConvertOptimizedTargetGetV2PromotionContent    ToolsAdConvertOptimizedTargetGetV2PromotionContent = "DOWNLOAD_URL"
	EXTERNAL_URL_ToolsAdConvertOptimizedTargetGetV2PromotionContent    ToolsAdConvertOptimizedTargetGetV2PromotionContent = "EXTERNAL_URL"
	SHOP_ToolsAdConvertOptimizedTargetGetV2PromotionContent            ToolsAdConvertOptimizedTargetGetV2PromotionContent = "SHOP"
	GOODS_LINK_ToolsAdConvertOptimizedTargetGetV2PromotionContent      ToolsAdConvertOptimizedTargetGetV2PromotionContent = "GOODS_LINK"
	LIVE_ROOM_ToolsAdConvertOptimizedTargetGetV2PromotionContent       ToolsAdConvertOptimizedTargetGetV2PromotionContent = "LIVE_ROOM"
	THIRD_PARTY_ToolsAdConvertOptimizedTargetGetV2PromotionContent     ToolsAdConvertOptimizedTargetGetV2PromotionContent = "THIRD_PARTY"
	MICRO_APP_ToolsAdConvertOptimizedTargetGetV2PromotionContent       ToolsAdConvertOptimizedTargetGetV2PromotionContent = "MICRO_APP"
)

// All allowed values of ToolsAdConvertOptimizedTargetGetV2PromotionContent enum
var AllowedToolsAdConvertOptimizedTargetGetV2PromotionContentEnumValues = []ToolsAdConvertOptimizedTargetGetV2PromotionContent{
	"DOUYIN",
	"NORMAL",
	"AWEME_HOME_PAGE",
	"QUICK_APP_URL",
	"DOWNLOAD_URL",
	"EXTERNAL_URL",
	"SHOP",
	"GOODS_LINK",
	"LIVE_ROOM",
	"THIRD_PARTY",
	"MICRO_APP",
}

// NewToolsAdConvertOptimizedTargetGetV2PromotionContentFromValue returns a pointer to a valid ToolsAdConvertOptimizedTargetGetV2PromotionContent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertOptimizedTargetGetV2PromotionContentFromValue(v string) (*ToolsAdConvertOptimizedTargetGetV2PromotionContent, error) {
	ev := ToolsAdConvertOptimizedTargetGetV2PromotionContent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertOptimizedTargetGetV2PromotionContent: valid values are %v", v, AllowedToolsAdConvertOptimizedTargetGetV2PromotionContentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertOptimizedTargetGetV2PromotionContent) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertOptimizedTargetGetV2PromotionContentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_optimized_target_get_v2_promotion_content value
func (v ToolsAdConvertOptimizedTargetGetV2PromotionContent) Ptr() *ToolsAdConvertOptimizedTargetGetV2PromotionContent {
	return &v
}
