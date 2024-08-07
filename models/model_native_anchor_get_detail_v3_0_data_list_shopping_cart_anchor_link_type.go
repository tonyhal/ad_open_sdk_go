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

// NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType
type NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType string

// List of native_anchor_get_detail_v3.0_data_list_shopping_cart_anchor_link_type
const (
	ONE_JUMP_NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType = "ONE_JUMP"
	TWO_JUMP_NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType = "TWO_JUMP"
)

// All allowed values of NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType enum
var AllowedNativeAnchorGetDetailV30DataListShoppingCartAnchorLinkTypeEnumValues = []NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType{
	"ONE_JUMP",
	"TWO_JUMP",
}

// NewNativeAnchorGetDetailV30DataListShoppingCartAnchorLinkTypeFromValue returns a pointer to a valid NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetDetailV30DataListShoppingCartAnchorLinkTypeFromValue(v string) (*NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType, error) {
	ev := NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType: valid values are %v", v, AllowedNativeAnchorGetDetailV30DataListShoppingCartAnchorLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetDetailV30DataListShoppingCartAnchorLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_detail_v3.0_data_list_shopping_cart_anchor_link_type value
func (v NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType) Ptr() *NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType {
	return &v
}
