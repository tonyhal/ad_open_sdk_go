/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponDetailV2DataCouponResourceListResourceType
type ClueCouponDetailV2DataCouponResourceListResourceType string

// List of clue_coupon_detail_v2_data_coupon_resource_list_resource_type
const (
	PHYSICAL_ClueCouponDetailV2DataCouponResourceListResourceType ClueCouponDetailV2DataCouponResourceListResourceType = "PHYSICAL"
	COMMON_ClueCouponDetailV2DataCouponResourceListResourceType   ClueCouponDetailV2DataCouponResourceListResourceType = "COMMON"
	DISCOUNT_ClueCouponDetailV2DataCouponResourceListResourceType ClueCouponDetailV2DataCouponResourceListResourceType = "DISCOUNT"
	GAME_ClueCouponDetailV2DataCouponResourceListResourceType     ClueCouponDetailV2DataCouponResourceListResourceType = "GAME"
	FULL_ClueCouponDetailV2DataCouponResourceListResourceType     ClueCouponDetailV2DataCouponResourceListResourceType = "FULL"
)

// Ptr returns reference to clue_coupon_detail_v2_data_coupon_resource_list_resource_type value
func (v ClueCouponDetailV2DataCouponResourceListResourceType) Ptr() *ClueCouponDetailV2DataCouponResourceListResourceType {
	return &v
}
