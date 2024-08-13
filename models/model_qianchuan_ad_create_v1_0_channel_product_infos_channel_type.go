/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10ChannelProductInfosChannelType
type QianchuanAdCreateV10ChannelProductInfosChannelType string

// List of qianchuan_ad_create_v1.0_channel_product_infos_channel_type
const (
	SHOP_SELL_QianchuanAdCreateV10ChannelProductInfosChannelType QianchuanAdCreateV10ChannelProductInfosChannelType = "SHOP_SELL"
	STAR_SELL_QianchuanAdCreateV10ChannelProductInfosChannelType QianchuanAdCreateV10ChannelProductInfosChannelType = "STAR_SELL"
)

// All allowed values of QianchuanAdCreateV10ChannelProductInfosChannelType enum
var AllowedQianchuanAdCreateV10ChannelProductInfosChannelTypeEnumValues = []QianchuanAdCreateV10ChannelProductInfosChannelType{
	"SHOP_SELL",
	"STAR_SELL",
}

// NewQianchuanAdCreateV10ChannelProductInfosChannelTypeFromValue returns a pointer to a valid QianchuanAdCreateV10ChannelProductInfosChannelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10ChannelProductInfosChannelTypeFromValue(v string) (*QianchuanAdCreateV10ChannelProductInfosChannelType, error) {
	ev := QianchuanAdCreateV10ChannelProductInfosChannelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10ChannelProductInfosChannelType: valid values are %v", v, AllowedQianchuanAdCreateV10ChannelProductInfosChannelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10ChannelProductInfosChannelType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10ChannelProductInfosChannelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_channel_product_infos_channel_type value
func (v QianchuanAdCreateV10ChannelProductInfosChannelType) Ptr() *QianchuanAdCreateV10ChannelProductInfosChannelType {
	return &v
}
