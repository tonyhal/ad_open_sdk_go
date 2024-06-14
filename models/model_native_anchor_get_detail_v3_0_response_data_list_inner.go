/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetDetailV30ResponseDataListInner struct for NativeAnchorGetDetailV30ResponseDataListInner
type NativeAnchorGetDetailV30ResponseDataListInner struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AnchorId           *string                                                          `json:"anchor_id,omitempty"`
	AnchorType         *NativeAnchorGetDetailV30DataListAnchorType                      `json:"anchor_type,omitempty"`
	AppEcommerceAnchor *NativeAnchorGetDetailV30ResponseDataListInnerAppEcommerceAnchor `json:"app_ecommerce_anchor,omitempty"`
	//
	CreateTime                *string                                                                 `json:"create_time,omitempty"`
	GameAnchor                *NativeAnchorGetDetailV30ResponseDataListInnerGameAnchor                `json:"game_anchor,omitempty"`
	InsuranceEnterpriseAnchor *NativeAnchorGetDetailV30ResponseDataListInnerInsuranceEnterpriseAnchor `json:"insurance_enterprise_anchor,omitempty"`
	//
	ModifyTime         *string                                                          `json:"modify_time,omitempty"`
	NetServiceAnchor   *NativeAnchorGetDetailV30ResponseDataListInnerNetServiceAnchor   `json:"net_service_anchor,omitempty"`
	PrivateChatAnchor  *NativeAnchorGetDetailV30ResponseDataListInnerPrivateChatAnchor  `json:"private_chat_anchor,omitempty"`
	ShoppingCartAnchor *NativeAnchorGetDetailV30ResponseDataListInnerShoppingCartAnchor `json:"shopping_cart_anchor,omitempty"`
	//
	ToolTitle *string `json:"tool_title,omitempty"`
}
