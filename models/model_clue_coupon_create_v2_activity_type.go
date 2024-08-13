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

// ClueCouponCreateV2ActivityType
type ClueCouponCreateV2ActivityType string

// List of clue_coupon_create_v2_activity_type
const (
	DIRECT_NOT_NEED_PHONE_ClueCouponCreateV2ActivityType ClueCouponCreateV2ActivityType = "DIRECT_NOT_NEED_PHONE"
	DIRECT_NEED_PHONE_ClueCouponCreateV2ActivityType     ClueCouponCreateV2ActivityType = "DIRECT_NEED_PHONE"
)

// All allowed values of ClueCouponCreateV2ActivityType enum
var AllowedClueCouponCreateV2ActivityTypeEnumValues = []ClueCouponCreateV2ActivityType{
	"DIRECT_NOT_NEED_PHONE",
	"DIRECT_NEED_PHONE",
}

// NewClueCouponCreateV2ActivityTypeFromValue returns a pointer to a valid ClueCouponCreateV2ActivityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponCreateV2ActivityTypeFromValue(v string) (*ClueCouponCreateV2ActivityType, error) {
	ev := ClueCouponCreateV2ActivityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponCreateV2ActivityType: valid values are %v", v, AllowedClueCouponCreateV2ActivityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponCreateV2ActivityType) IsValid() bool {
	for _, existing := range AllowedClueCouponCreateV2ActivityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_create_v2_activity_type value
func (v ClueCouponCreateV2ActivityType) Ptr() *ClueCouponCreateV2ActivityType {
	return &v
}
