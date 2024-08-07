/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponCodeGetV2Status
type ClueCouponCodeGetV2Status string

// List of clue_coupon_code_get_v2_status
const (
	USED_ClueCouponCodeGetV2Status      ClueCouponCodeGetV2Status = "USED"
	EXPIRED_ClueCouponCodeGetV2Status   ClueCouponCodeGetV2Status = "EXPIRED"
	ABANDONED_ClueCouponCodeGetV2Status ClueCouponCodeGetV2Status = "ABANDONED"
	INVALID_ClueCouponCodeGetV2Status   ClueCouponCodeGetV2Status = "INVALID"
	VALID_ClueCouponCodeGetV2Status     ClueCouponCodeGetV2Status = "VALID"
)

// All allowed values of ClueCouponCodeGetV2Status enum
var AllowedClueCouponCodeGetV2StatusEnumValues = []ClueCouponCodeGetV2Status{
	"USED",
	"EXPIRED",
	"ABANDONED",
	"INVALID",
	"VALID",
}

// NewClueCouponCodeGetV2StatusFromValue returns a pointer to a valid ClueCouponCodeGetV2Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponCodeGetV2StatusFromValue(v string) (*ClueCouponCodeGetV2Status, error) {
	ev := ClueCouponCodeGetV2Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponCodeGetV2Status: valid values are %v", v, AllowedClueCouponCodeGetV2StatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponCodeGetV2Status) IsValid() bool {
	for _, existing := range AllowedClueCouponCodeGetV2StatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_code_get_v2_status value
func (v ClueCouponCodeGetV2Status) Ptr() *ClueCouponCodeGetV2Status {
	return &v
}
