/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeGetV2FilteringLandingType
type CreativeGetV2FilteringLandingType string

// List of creative_get_v2_filtering_landing_type
const (
	APP_CreativeGetV2FilteringLandingType            CreativeGetV2FilteringLandingType = "APP"
	ARTICLE_CreativeGetV2FilteringLandingType        CreativeGetV2FilteringLandingType = "ARTICLE"
	AWEME_CreativeGetV2FilteringLandingType          CreativeGetV2FilteringLandingType = "AWEME"
	BRAND_EXTERNAL_CreativeGetV2FilteringLandingType CreativeGetV2FilteringLandingType = "BRAND_EXTERNAL"
	DPA_CreativeGetV2FilteringLandingType            CreativeGetV2FilteringLandingType = "DPA"
	GOODS_CreativeGetV2FilteringLandingType          CreativeGetV2FilteringLandingType = "GOODS"
	LINK_CreativeGetV2FilteringLandingType           CreativeGetV2FilteringLandingType = "LINK"
	LIVE_CreativeGetV2FilteringLandingType           CreativeGetV2FilteringLandingType = "LIVE"
	MICRO_GAME_CreativeGetV2FilteringLandingType     CreativeGetV2FilteringLandingType = "MICRO_GAME"
	QUICK_APP_CreativeGetV2FilteringLandingType      CreativeGetV2FilteringLandingType = "QUICK_APP"
	SHOP_CreativeGetV2FilteringLandingType           CreativeGetV2FilteringLandingType = "SHOP"
	STORE_CreativeGetV2FilteringLandingType          CreativeGetV2FilteringLandingType = "STORE"
)

// All allowed values of CreativeGetV2FilteringLandingType enum
var AllowedCreativeGetV2FilteringLandingTypeEnumValues = []CreativeGetV2FilteringLandingType{
	"APP",
	"ARTICLE",
	"AWEME",
	"BRAND_EXTERNAL",
	"DPA",
	"GOODS",
	"LINK",
	"LIVE",
	"MICRO_GAME",
	"QUICK_APP",
	"SHOP",
	"STORE",
}

// NewCreativeGetV2FilteringLandingTypeFromValue returns a pointer to a valid CreativeGetV2FilteringLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeGetV2FilteringLandingTypeFromValue(v string) (*CreativeGetV2FilteringLandingType, error) {
	ev := CreativeGetV2FilteringLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeGetV2FilteringLandingType: valid values are %v", v, AllowedCreativeGetV2FilteringLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeGetV2FilteringLandingType) IsValid() bool {
	for _, existing := range AllowedCreativeGetV2FilteringLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_get_v2_filtering_landing_type value
func (v CreativeGetV2FilteringLandingType) Ptr() *CreativeGetV2FilteringLandingType {
	return &v
}
