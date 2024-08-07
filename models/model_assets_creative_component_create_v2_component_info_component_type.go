/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AssetsCreativeComponentCreateV2ComponentInfoComponentType
type AssetsCreativeComponentCreateV2ComponentInfoComponentType string

// List of assets_creative_component_create_v2_component_info_component_type
const (
	RESERVATION_BUTTON_AssetsCreativeComponentCreateV2ComponentInfoComponentType      AssetsCreativeComponentCreateV2ComponentInfoComponentType = "RESERVATION_BUTTON"
	UNION_LIGHT_INTERACTIVE_AssetsCreativeComponentCreateV2ComponentInfoComponentType AssetsCreativeComponentCreateV2ComponentInfoComponentType = "UNION_LIGHT_INTERACTIVE"
	PROMOTION_CARD_AssetsCreativeComponentCreateV2ComponentInfoComponentType          AssetsCreativeComponentCreateV2ComponentInfoComponentType = "PROMOTION_CARD"
	LIGHT_INTER_ACTIVE_AssetsCreativeComponentCreateV2ComponentInfoComponentType      AssetsCreativeComponentCreateV2ComponentInfoComponentType = "LIGHT_INTER_ACTIVE"
	LUCKY_BOX_AssetsCreativeComponentCreateV2ComponentInfoComponentType               AssetsCreativeComponentCreateV2ComponentInfoComponentType = "LUCKY_BOX"
	CHOICE_MAGNET_AssetsCreativeComponentCreateV2ComponentInfoComponentType           AssetsCreativeComponentCreateV2ComponentInfoComponentType = "CHOICE_MAGNET"
	ECOMMERCE_CARD_AssetsCreativeComponentCreateV2ComponentInfoComponentType          AssetsCreativeComponentCreateV2ComponentInfoComponentType = "ECOMMERCE_CARD"
	GAME_PACK_AssetsCreativeComponentCreateV2ComponentInfoComponentType               AssetsCreativeComponentCreateV2ComponentInfoComponentType = "GAME_PACK"
	HALF_INTERACTIVE_AssetsCreativeComponentCreateV2ComponentInfoComponentType        AssetsCreativeComponentCreateV2ComponentInfoComponentType = "HALF_INTERACTIVE"
	COMMERCE_MAGNET_AssetsCreativeComponentCreateV2ComponentInfoComponentType         AssetsCreativeComponentCreateV2ComponentInfoComponentType = "COMMERCE_MAGNET"
	IMAGE_MAGNET_AssetsCreativeComponentCreateV2ComponentInfoComponentType            AssetsCreativeComponentCreateV2ComponentInfoComponentType = "IMAGE_MAGNET"
	COUPON_MAGNET_AssetsCreativeComponentCreateV2ComponentInfoComponentType           AssetsCreativeComponentCreateV2ComponentInfoComponentType = "COUPON_MAGNET"
	VOTE_MAGNET_AssetsCreativeComponentCreateV2ComponentInfoComponentType             AssetsCreativeComponentCreateV2ComponentInfoComponentType = "VOTE_MAGNET"
)

// All allowed values of AssetsCreativeComponentCreateV2ComponentInfoComponentType enum
var AllowedAssetsCreativeComponentCreateV2ComponentInfoComponentTypeEnumValues = []AssetsCreativeComponentCreateV2ComponentInfoComponentType{
	"RESERVATION_BUTTON",
	"UNION_LIGHT_INTERACTIVE",
	"PROMOTION_CARD",
	"LIGHT_INTER_ACTIVE",
	"LUCKY_BOX",
	"CHOICE_MAGNET",
	"ECOMMERCE_CARD",
	"GAME_PACK",
	"HALF_INTERACTIVE",
	"COMMERCE_MAGNET",
	"IMAGE_MAGNET",
	"COUPON_MAGNET",
	"VOTE_MAGNET",
}

// NewAssetsCreativeComponentCreateV2ComponentInfoComponentTypeFromValue returns a pointer to a valid AssetsCreativeComponentCreateV2ComponentInfoComponentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetsCreativeComponentCreateV2ComponentInfoComponentTypeFromValue(v string) (*AssetsCreativeComponentCreateV2ComponentInfoComponentType, error) {
	ev := AssetsCreativeComponentCreateV2ComponentInfoComponentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssetsCreativeComponentCreateV2ComponentInfoComponentType: valid values are %v", v, AllowedAssetsCreativeComponentCreateV2ComponentInfoComponentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssetsCreativeComponentCreateV2ComponentInfoComponentType) IsValid() bool {
	for _, existing := range AllowedAssetsCreativeComponentCreateV2ComponentInfoComponentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to assets_creative_component_create_v2_component_info_component_type value
func (v AssetsCreativeComponentCreateV2ComponentInfoComponentType) Ptr() *AssetsCreativeComponentCreateV2ComponentInfoComponentType {
	return &v
}
