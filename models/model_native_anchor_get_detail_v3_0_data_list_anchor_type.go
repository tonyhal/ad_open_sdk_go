/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetDetailV30DataListAnchorType
type NativeAnchorGetDetailV30DataListAnchorType string

// List of native_anchor_get_detail_v3.0_data_list_anchor_type
const (
	APP_GAME_NativeAnchorGetDetailV30DataListAnchorType             NativeAnchorGetDetailV30DataListAnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_NativeAnchorGetDetailV30DataListAnchorType NativeAnchorGetDetailV30DataListAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_NativeAnchorGetDetailV30DataListAnchorType             NativeAnchorGetDetailV30DataListAnchorType = "APP_SHOP"
	INSURANCE_NativeAnchorGetDetailV30DataListAnchorType            NativeAnchorGetDetailV30DataListAnchorType = "INSURANCE"
	ONLINE_SUBSCRIBE_NativeAnchorGetDetailV30DataListAnchorType     NativeAnchorGetDetailV30DataListAnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_NativeAnchorGetDetailV30DataListAnchorType         NativeAnchorGetDetailV30DataListAnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_NativeAnchorGetDetailV30DataListAnchorType        NativeAnchorGetDetailV30DataListAnchorType = "SHOPPING_CART"
)

// Ptr returns reference to native_anchor_get_detail_v3.0_data_list_anchor_type value
func (v NativeAnchorGetDetailV30DataListAnchorType) Ptr() *NativeAnchorGetDetailV30DataListAnchorType {
	return &v
}
