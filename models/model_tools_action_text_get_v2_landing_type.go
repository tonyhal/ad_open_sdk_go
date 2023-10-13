/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsActionTextGetV2LandingType
type ToolsActionTextGetV2LandingType string

// List of tools_action_text_get_v2_landing_type
const (
	STORE_ToolsActionTextGetV2LandingType          ToolsActionTextGetV2LandingType = "STORE"
	QUICK_APP_ToolsActionTextGetV2LandingType      ToolsActionTextGetV2LandingType = "QUICK_APP"
	BRAND_EXTERNAL_ToolsActionTextGetV2LandingType ToolsActionTextGetV2LandingType = "BRAND_EXTERNAL"
	AWEME_ToolsActionTextGetV2LandingType          ToolsActionTextGetV2LandingType = "AWEME"
	LINK_ToolsActionTextGetV2LandingType           ToolsActionTextGetV2LandingType = "LINK"
	ARTICLE_ToolsActionTextGetV2LandingType        ToolsActionTextGetV2LandingType = "ARTICLE"
	LIVE_ToolsActionTextGetV2LandingType           ToolsActionTextGetV2LandingType = "LIVE"
	DPA_ToolsActionTextGetV2LandingType            ToolsActionTextGetV2LandingType = "DPA"
	SHOP_ToolsActionTextGetV2LandingType           ToolsActionTextGetV2LandingType = "SHOP"
	GOODS_ToolsActionTextGetV2LandingType          ToolsActionTextGetV2LandingType = "GOODS"
	APP_ToolsActionTextGetV2LandingType            ToolsActionTextGetV2LandingType = "APP"
)

// All allowed values of ToolsActionTextGetV2LandingType enum
var AllowedToolsActionTextGetV2LandingTypeEnumValues = []ToolsActionTextGetV2LandingType{
	"STORE",
	"QUICK_APP",
	"BRAND_EXTERNAL",
	"AWEME",
	"LINK",
	"ARTICLE",
	"LIVE",
	"DPA",
	"SHOP",
	"GOODS",
	"APP",
}

// NewToolsActionTextGetV2LandingTypeFromValue returns a pointer to a valid ToolsActionTextGetV2LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsActionTextGetV2LandingTypeFromValue(v string) (*ToolsActionTextGetV2LandingType, error) {
	ev := ToolsActionTextGetV2LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsActionTextGetV2LandingType: valid values are %v", v, AllowedToolsActionTextGetV2LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsActionTextGetV2LandingType) IsValid() bool {
	for _, existing := range AllowedToolsActionTextGetV2LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_action_text_get_v2_landing_type value
func (v ToolsActionTextGetV2LandingType) Ptr() *ToolsActionTextGetV2LandingType {
	return &v
}
