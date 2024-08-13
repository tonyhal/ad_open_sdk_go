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

// ClueCouponDetailV2DataCouponStatus
type ClueCouponDetailV2DataCouponStatus string

// List of clue_coupon_detail_v2_data_coupon_status
const (
	AUDIT_DOING_ClueCouponDetailV2DataCouponStatus ClueCouponDetailV2DataCouponStatus = "AUDIT_DOING"
	OFFLINE_ClueCouponDetailV2DataCouponStatus     ClueCouponDetailV2DataCouponStatus = "OFFLINE"
	UNAUDITED_ClueCouponDetailV2DataCouponStatus   ClueCouponDetailV2DataCouponStatus = "UNAUDITED"
	PAUSE_ClueCouponDetailV2DataCouponStatus       ClueCouponDetailV2DataCouponStatus = "PAUSE"
	NORMAL_ClueCouponDetailV2DataCouponStatus      ClueCouponDetailV2DataCouponStatus = "NORMAL"
	AUDIT_FAIL_ClueCouponDetailV2DataCouponStatus  ClueCouponDetailV2DataCouponStatus = "AUDIT_FAIL"
	DELETED_ClueCouponDetailV2DataCouponStatus     ClueCouponDetailV2DataCouponStatus = "DELETED"
)

// All allowed values of ClueCouponDetailV2DataCouponStatus enum
var AllowedClueCouponDetailV2DataCouponStatusEnumValues = []ClueCouponDetailV2DataCouponStatus{
	"AUDIT_DOING",
	"OFFLINE",
	"UNAUDITED",
	"PAUSE",
	"NORMAL",
	"AUDIT_FAIL",
	"DELETED",
}

// NewClueCouponDetailV2DataCouponStatusFromValue returns a pointer to a valid ClueCouponDetailV2DataCouponStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataCouponStatusFromValue(v string) (*ClueCouponDetailV2DataCouponStatus, error) {
	ev := ClueCouponDetailV2DataCouponStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataCouponStatus: valid values are %v", v, AllowedClueCouponDetailV2DataCouponStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataCouponStatus) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataCouponStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_coupon_status value
func (v ClueCouponDetailV2DataCouponStatus) Ptr() *ClueCouponDetailV2DataCouponStatus {
	return &v
}
