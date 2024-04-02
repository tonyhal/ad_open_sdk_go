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

// NativeAnchorQrcodePreviewGetV30AnchorType
type NativeAnchorQrcodePreviewGetV30AnchorType string

// List of native_anchor_qrcode_preview_get_v3.0_anchor_type
const (
	APP_GAME_NativeAnchorQrcodePreviewGetV30AnchorType             NativeAnchorQrcodePreviewGetV30AnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_NativeAnchorQrcodePreviewGetV30AnchorType NativeAnchorQrcodePreviewGetV30AnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_NativeAnchorQrcodePreviewGetV30AnchorType             NativeAnchorQrcodePreviewGetV30AnchorType = "APP_SHOP"
	INSURANCE_NativeAnchorQrcodePreviewGetV30AnchorType            NativeAnchorQrcodePreviewGetV30AnchorType = "INSURANCE"
	MICRO_APP_NativeAnchorQrcodePreviewGetV30AnchorType            NativeAnchorQrcodePreviewGetV30AnchorType = "MICRO_APP"
	MICRO_GAME_NativeAnchorQrcodePreviewGetV30AnchorType           NativeAnchorQrcodePreviewGetV30AnchorType = "MICRO_GAME"
	ONLINE_SUBSCRIBE_NativeAnchorQrcodePreviewGetV30AnchorType     NativeAnchorQrcodePreviewGetV30AnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_NativeAnchorQrcodePreviewGetV30AnchorType         NativeAnchorQrcodePreviewGetV30AnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_NativeAnchorQrcodePreviewGetV30AnchorType        NativeAnchorQrcodePreviewGetV30AnchorType = "SHOPPING_CART"
)

// All allowed values of NativeAnchorQrcodePreviewGetV30AnchorType enum
var AllowedNativeAnchorQrcodePreviewGetV30AnchorTypeEnumValues = []NativeAnchorQrcodePreviewGetV30AnchorType{
	"APP_GAME",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
	"INSURANCE",
	"MICRO_APP",
	"MICRO_GAME",
	"ONLINE_SUBSCRIBE",
	"PRIVATE_CHAT",
	"SHOPPING_CART",
}

// NewNativeAnchorQrcodePreviewGetV30AnchorTypeFromValue returns a pointer to a valid NativeAnchorQrcodePreviewGetV30AnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorQrcodePreviewGetV30AnchorTypeFromValue(v string) (*NativeAnchorQrcodePreviewGetV30AnchorType, error) {
	ev := NativeAnchorQrcodePreviewGetV30AnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorQrcodePreviewGetV30AnchorType: valid values are %v", v, AllowedNativeAnchorQrcodePreviewGetV30AnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorQrcodePreviewGetV30AnchorType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorQrcodePreviewGetV30AnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_qrcode_preview_get_v3.0_anchor_type value
func (v NativeAnchorQrcodePreviewGetV30AnchorType) Ptr() *NativeAnchorQrcodePreviewGetV30AnchorType {
	return &v
}
