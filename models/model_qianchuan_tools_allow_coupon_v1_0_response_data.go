/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanToolsAllowCouponV10ResponseData
type QianchuanToolsAllowCouponV10ResponseData struct {
	//
	AdvAllowCoupon *bool `json:"adv_allow_coupon,omitempty"`
	//
	AwemeAllowCoupon []*QianchuanToolsAllowCouponV10ResponseDataAwemeAllowCouponInner `json:"aweme_allow_coupon,omitempty"`
	//
	ProductAllowCoupon []*QianchuanToolsAllowCouponV10ResponseDataProductAllowCouponInner `json:"product_allow_coupon,omitempty"`
}
