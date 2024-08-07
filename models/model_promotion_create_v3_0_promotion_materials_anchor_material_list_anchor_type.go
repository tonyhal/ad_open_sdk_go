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

// PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType
type PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType string

// List of promotion_create_v3.0_promotion_materials_anchor_material_list_anchor_type
const (
	APP_GAME_PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType             PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType             PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType = "APP_SHOP"
	INSURANCE_PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType            PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType = "INSURANCE"
	ONLINE_SUBSCRIBE_PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType     PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType         PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType        PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType = "SHOPPING_CART"
)

// All allowed values of PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType enum
var AllowedPromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorTypeEnumValues = []PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType{
	"APP_GAME",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
	"INSURANCE",
	"ONLINE_SUBSCRIBE",
	"PRIVATE_CHAT",
	"SHOPPING_CART",
}

// NewPromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorTypeFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorTypeFromValue(v string) (*PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType, error) {
	ev := PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_anchor_material_list_anchor_type value
func (v PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType) Ptr() *PromotionCreateV30PromotionMaterialsAnchorMaterialListAnchorType {
	return &v
}
