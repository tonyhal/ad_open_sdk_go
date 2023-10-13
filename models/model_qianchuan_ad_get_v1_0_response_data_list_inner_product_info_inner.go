/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdGetV10ResponseDataListInnerProductInfoInner struct for QianchuanAdGetV10ResponseDataListInnerProductInfoInner
type QianchuanAdGetV10ResponseDataListInnerProductInfoInner struct {
	//
	ChannelId   *int64                                           `json:"channel_id,omitempty"`
	ChannelType *QianchuanAdGetV10DataListProductInfoChannelType `json:"channel_type,omitempty"`
	//
	DiscountHigherPrice *float64 `json:"discount_higher_price,omitempty"`
	//
	DiscountLowerPrice *float64 `json:"discount_lower_price,omitempty"`
	//
	DiscountPrice *float64 `json:"discount_price,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	//
	Img *string `json:"img,omitempty"`
	//
	MarketPrice *float64 `json:"market_price,omitempty"`
	//
	Name *string `json:"name,omitempty"`
}
