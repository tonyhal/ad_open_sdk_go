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

// CreativeProceduralCreativeCreateV2AdDataAnchorType
type CreativeProceduralCreativeCreateV2AdDataAnchorType string

// List of creative_procedural_creative_create_v2_ad_data_anchor_type
const (
	APP_INTERNET_SERVICE_CreativeProceduralCreativeCreateV2AdDataAnchorType CreativeProceduralCreativeCreateV2AdDataAnchorType = "APP_INTERNET_SERVICE"
	APP_GAME_CreativeProceduralCreativeCreateV2AdDataAnchorType             CreativeProceduralCreativeCreateV2AdDataAnchorType = "APP_GAME"
	ONLINE_SUBSCRIBE_CreativeProceduralCreativeCreateV2AdDataAnchorType     CreativeProceduralCreativeCreateV2AdDataAnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_CreativeProceduralCreativeCreateV2AdDataAnchorType         CreativeProceduralCreativeCreateV2AdDataAnchorType = "PRIVATE_CHAT"
	INSURANCE_CreativeProceduralCreativeCreateV2AdDataAnchorType            CreativeProceduralCreativeCreateV2AdDataAnchorType = "INSURANCE"
	APP_SHOP_CreativeProceduralCreativeCreateV2AdDataAnchorType             CreativeProceduralCreativeCreateV2AdDataAnchorType = "APP_SHOP"
	SHOPPING_CART_CreativeProceduralCreativeCreateV2AdDataAnchorType        CreativeProceduralCreativeCreateV2AdDataAnchorType = "SHOPPING_CART"
)

// All allowed values of CreativeProceduralCreativeCreateV2AdDataAnchorType enum
var AllowedCreativeProceduralCreativeCreateV2AdDataAnchorTypeEnumValues = []CreativeProceduralCreativeCreateV2AdDataAnchorType{
	"APP_INTERNET_SERVICE",
	"APP_GAME",
	"ONLINE_SUBSCRIBE",
	"PRIVATE_CHAT",
	"INSURANCE",
	"APP_SHOP",
	"SHOPPING_CART",
}

// NewCreativeProceduralCreativeCreateV2AdDataAnchorTypeFromValue returns a pointer to a valid CreativeProceduralCreativeCreateV2AdDataAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeCreateV2AdDataAnchorTypeFromValue(v string) (*CreativeProceduralCreativeCreateV2AdDataAnchorType, error) {
	ev := CreativeProceduralCreativeCreateV2AdDataAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeCreateV2AdDataAnchorType: valid values are %v", v, AllowedCreativeProceduralCreativeCreateV2AdDataAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeCreateV2AdDataAnchorType) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeCreateV2AdDataAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_create_v2_ad_data_anchor_type value
func (v CreativeProceduralCreativeCreateV2AdDataAnchorType) Ptr() *CreativeProceduralCreativeCreateV2AdDataAnchorType {
	return &v
}
