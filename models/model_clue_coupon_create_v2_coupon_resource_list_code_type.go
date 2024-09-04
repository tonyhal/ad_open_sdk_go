/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponCreateV2CouponResourceListCodeType
type ClueCouponCreateV2CouponResourceListCodeType string

// List of clue_coupon_create_v2_coupon_resource_list_code_type
const (
	COMMON_ClueCouponCreateV2CouponResourceListCodeType   ClueCouponCreateV2CouponResourceListCodeType = "COMMON"
	PLATFORM_ClueCouponCreateV2CouponResourceListCodeType ClueCouponCreateV2CouponResourceListCodeType = "PLATFORM"
	MERCHANT_ClueCouponCreateV2CouponResourceListCodeType ClueCouponCreateV2CouponResourceListCodeType = "MERCHANT"
	API_ClueCouponCreateV2CouponResourceListCodeType      ClueCouponCreateV2CouponResourceListCodeType = "API"
)

// Ptr returns reference to clue_coupon_create_v2_coupon_resource_list_code_type value
func (v ClueCouponCreateV2CouponResourceListCodeType) Ptr() *ClueCouponCreateV2CouponResourceListCodeType {
	return &v
}
