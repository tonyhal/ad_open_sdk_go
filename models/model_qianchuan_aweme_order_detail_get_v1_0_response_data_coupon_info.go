/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderDetailGetV10ResponseDataCouponInfo
type QianchuanAwemeOrderDetailGetV10ResponseDataCouponInfo struct {
	// 券的总金额
	CouponAmount *float64 `json:"coupon_amount,omitempty"`
	// 券类型。1表示满减券；4表示满赠券
	CouponType *int64 `json:"coupon_type,omitempty"`
}
