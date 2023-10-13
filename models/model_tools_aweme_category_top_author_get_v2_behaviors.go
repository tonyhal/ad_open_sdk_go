/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAwemeCategoryTopAuthorGetV2Behaviors
type ToolsAwemeCategoryTopAuthorGetV2Behaviors string

// List of tools_aweme_category_top_author_get_v2_behaviors
const (
	FOLLOWED_USER_ToolsAwemeCategoryTopAuthorGetV2Behaviors        ToolsAwemeCategoryTopAuthorGetV2Behaviors = "FOLLOWED_USER"
	GOODS_CARTS_ORDER_ToolsAwemeCategoryTopAuthorGetV2Behaviors    ToolsAwemeCategoryTopAuthorGetV2Behaviors = "GOODS_CARTS_ORDER"
	LIVE_EFFECTIVE_WATCH_ToolsAwemeCategoryTopAuthorGetV2Behaviors ToolsAwemeCategoryTopAuthorGetV2Behaviors = "LIVE_EFFECTIVE_WATCH"
	COMMENTED_USER_ToolsAwemeCategoryTopAuthorGetV2Behaviors       ToolsAwemeCategoryTopAuthorGetV2Behaviors = "COMMENTED_USER"
	LIVE_COMMENT_ToolsAwemeCategoryTopAuthorGetV2Behaviors         ToolsAwemeCategoryTopAuthorGetV2Behaviors = "LIVE_COMMENT"
	LIVE_WATCH_ToolsAwemeCategoryTopAuthorGetV2Behaviors           ToolsAwemeCategoryTopAuthorGetV2Behaviors = "LIVE_WATCH"
	LIVE_EXCEPTIONAL_ToolsAwemeCategoryTopAuthorGetV2Behaviors     ToolsAwemeCategoryTopAuthorGetV2Behaviors = "LIVE_EXCEPTIONAL"
	SHARED_USER_ToolsAwemeCategoryTopAuthorGetV2Behaviors          ToolsAwemeCategoryTopAuthorGetV2Behaviors = "SHARED_USER"
	LIVE_GOODS_CLICK_ToolsAwemeCategoryTopAuthorGetV2Behaviors     ToolsAwemeCategoryTopAuthorGetV2Behaviors = "LIVE_GOODS_CLICK"
	LIVE_GOODS_ORDER_ToolsAwemeCategoryTopAuthorGetV2Behaviors     ToolsAwemeCategoryTopAuthorGetV2Behaviors = "LIVE_GOODS_ORDER"
	GOODS_CARTS_CLICK_ToolsAwemeCategoryTopAuthorGetV2Behaviors    ToolsAwemeCategoryTopAuthorGetV2Behaviors = "GOODS_CARTS_CLICK"
	LIKED_USER_ToolsAwemeCategoryTopAuthorGetV2Behaviors           ToolsAwemeCategoryTopAuthorGetV2Behaviors = "LIKED_USER"
)

// All allowed values of ToolsAwemeCategoryTopAuthorGetV2Behaviors enum
var AllowedToolsAwemeCategoryTopAuthorGetV2BehaviorsEnumValues = []ToolsAwemeCategoryTopAuthorGetV2Behaviors{
	"FOLLOWED_USER",
	"GOODS_CARTS_ORDER",
	"LIVE_EFFECTIVE_WATCH",
	"COMMENTED_USER",
	"LIVE_COMMENT",
	"LIVE_WATCH",
	"LIVE_EXCEPTIONAL",
	"SHARED_USER",
	"LIVE_GOODS_CLICK",
	"LIVE_GOODS_ORDER",
	"GOODS_CARTS_CLICK",
	"LIKED_USER",
}

// NewToolsAwemeCategoryTopAuthorGetV2BehaviorsFromValue returns a pointer to a valid ToolsAwemeCategoryTopAuthorGetV2Behaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeCategoryTopAuthorGetV2BehaviorsFromValue(v string) (*ToolsAwemeCategoryTopAuthorGetV2Behaviors, error) {
	ev := ToolsAwemeCategoryTopAuthorGetV2Behaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeCategoryTopAuthorGetV2Behaviors: valid values are %v", v, AllowedToolsAwemeCategoryTopAuthorGetV2BehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeCategoryTopAuthorGetV2Behaviors) IsValid() bool {
	for _, existing := range AllowedToolsAwemeCategoryTopAuthorGetV2BehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_category_top_author_get_v2_behaviors value
func (v ToolsAwemeCategoryTopAuthorGetV2Behaviors) Ptr() *ToolsAwemeCategoryTopAuthorGetV2Behaviors {
	return &v
}
