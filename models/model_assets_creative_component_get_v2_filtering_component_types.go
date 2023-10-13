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

// AssetsCreativeComponentGetV2FilteringComponentTypes
type AssetsCreativeComponentGetV2FilteringComponentTypes string

// List of assets_creative_component_get_v2_filtering_component_types
const (
	CHOICE_MAGNET_AssetsCreativeComponentGetV2FilteringComponentTypes           AssetsCreativeComponentGetV2FilteringComponentTypes = "CHOICE_MAGNET"
	UNION_LIGHT_INTERACTIVE_AssetsCreativeComponentGetV2FilteringComponentTypes AssetsCreativeComponentGetV2FilteringComponentTypes = "UNION_LIGHT_INTERACTIVE"
	COMMERCE_MAGNET_AssetsCreativeComponentGetV2FilteringComponentTypes         AssetsCreativeComponentGetV2FilteringComponentTypes = "COMMERCE_MAGNET"
	COUPON_MAGNET_AssetsCreativeComponentGetV2FilteringComponentTypes           AssetsCreativeComponentGetV2FilteringComponentTypes = "COUPON_MAGNET"
	VOTE_MAGNET_AssetsCreativeComponentGetV2FilteringComponentTypes             AssetsCreativeComponentGetV2FilteringComponentTypes = "VOTE_MAGNET"
	PROMOTION_CARD_AssetsCreativeComponentGetV2FilteringComponentTypes          AssetsCreativeComponentGetV2FilteringComponentTypes = "PROMOTION_CARD"
	LUCKY_BOX_AssetsCreativeComponentGetV2FilteringComponentTypes               AssetsCreativeComponentGetV2FilteringComponentTypes = "LUCKY_BOX"
	RESERVATION_BUTTON_AssetsCreativeComponentGetV2FilteringComponentTypes      AssetsCreativeComponentGetV2FilteringComponentTypes = "RESERVATION_BUTTON"
	LIGHT_INTER_ACTIVE_AssetsCreativeComponentGetV2FilteringComponentTypes      AssetsCreativeComponentGetV2FilteringComponentTypes = "LIGHT_INTER_ACTIVE"
	GAME_PACK_AssetsCreativeComponentGetV2FilteringComponentTypes               AssetsCreativeComponentGetV2FilteringComponentTypes = "GAME_PACK"
	IMAGE_MAGNET_AssetsCreativeComponentGetV2FilteringComponentTypes            AssetsCreativeComponentGetV2FilteringComponentTypes = "IMAGE_MAGNET"
)

// All allowed values of AssetsCreativeComponentGetV2FilteringComponentTypes enum
var AllowedAssetsCreativeComponentGetV2FilteringComponentTypesEnumValues = []AssetsCreativeComponentGetV2FilteringComponentTypes{
	"CHOICE_MAGNET",
	"UNION_LIGHT_INTERACTIVE",
	"COMMERCE_MAGNET",
	"COUPON_MAGNET",
	"VOTE_MAGNET",
	"PROMOTION_CARD",
	"LUCKY_BOX",
	"RESERVATION_BUTTON",
	"LIGHT_INTER_ACTIVE",
	"GAME_PACK",
	"IMAGE_MAGNET",
}

// NewAssetsCreativeComponentGetV2FilteringComponentTypesFromValue returns a pointer to a valid AssetsCreativeComponentGetV2FilteringComponentTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetsCreativeComponentGetV2FilteringComponentTypesFromValue(v string) (*AssetsCreativeComponentGetV2FilteringComponentTypes, error) {
	ev := AssetsCreativeComponentGetV2FilteringComponentTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssetsCreativeComponentGetV2FilteringComponentTypes: valid values are %v", v, AllowedAssetsCreativeComponentGetV2FilteringComponentTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssetsCreativeComponentGetV2FilteringComponentTypes) IsValid() bool {
	for _, existing := range AllowedAssetsCreativeComponentGetV2FilteringComponentTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to assets_creative_component_get_v2_filtering_component_types value
func (v AssetsCreativeComponentGetV2FilteringComponentTypes) Ptr() *AssetsCreativeComponentGetV2FilteringComponentTypes {
	return &v
}
