/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetDetailV30ResponseDataListInnerShoppingCartAnchor
type NativeAnchorGetDetailV30ResponseDataListInnerShoppingCartAnchor struct {
	//
	ExternalUrl *string                                                     `json:"external_url,omitempty"`
	LinkType    *NativeAnchorGetDetailV30DataListShoppingCartAnchorLinkType `json:"link_type,omitempty"`
	//
	OpenUrl *string `json:"open_url,omitempty"`
	//
	ProductImages []*NativeAnchorGetDetailV30ResponseDataListInnerShoppingCartAnchorProductImagesInner `json:"product_images,omitempty"`
	//
	ProductPrice *float64 `json:"product_price,omitempty"`
	//
	ProductSource *string `json:"product_source,omitempty"`
	//
	ProductTitle *string `json:"product_title,omitempty"`
	//
	Title *string `json:"title,omitempty"`
}
