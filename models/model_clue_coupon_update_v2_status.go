/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponUpdateV2Status
type ClueCouponUpdateV2Status string

// List of clue_coupon_update_v2_status
const (
	AUDIT_DOING_ClueCouponUpdateV2Status ClueCouponUpdateV2Status = "AUDIT_DOING"
	UNAUDITED_ClueCouponUpdateV2Status   ClueCouponUpdateV2Status = "UNAUDITED"
	OFFLINE_ClueCouponUpdateV2Status     ClueCouponUpdateV2Status = "OFFLINE"
	PAUSE_ClueCouponUpdateV2Status       ClueCouponUpdateV2Status = "PAUSE"
	NORMAL_ClueCouponUpdateV2Status      ClueCouponUpdateV2Status = "NORMAL"
	DELETED_ClueCouponUpdateV2Status     ClueCouponUpdateV2Status = "DELETED"
	AUDIT_FAIL_ClueCouponUpdateV2Status  ClueCouponUpdateV2Status = "AUDIT_FAIL"
)

// Ptr returns reference to clue_coupon_update_v2_status value
func (v ClueCouponUpdateV2Status) Ptr() *ClueCouponUpdateV2Status {
	return &v
}
