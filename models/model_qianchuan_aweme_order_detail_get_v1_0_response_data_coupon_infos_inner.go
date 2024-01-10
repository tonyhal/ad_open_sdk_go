/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderDetailGetV10ResponseDataCouponInfosInner struct for QianchuanAwemeOrderDetailGetV10ResponseDataCouponInfosInner
type QianchuanAwemeOrderDetailGetV10ResponseDataCouponInfosInner struct {
	// 券的总金额
	CouponTotalAmount *float64 `json:"coupon_total_amount,omitempty"`
	// 券类型。1表示满减券；4表示满赠券
	CouponType *int64 `json:"coupon_type,omitempty"`
}
