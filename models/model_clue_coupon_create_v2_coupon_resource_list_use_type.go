/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponCreateV2CouponResourceListUseType
type ClueCouponCreateV2CouponResourceListUseType string

// List of clue_coupon_create_v2_coupon_resource_list_use_type
const (
	OFFLINE_ClueCouponCreateV2CouponResourceListUseType ClueCouponCreateV2CouponResourceListUseType = "OFFLINE"
	LINK_ClueCouponCreateV2CouponResourceListUseType    ClueCouponCreateV2CouponResourceListUseType = "LINK"
)

// All allowed values of ClueCouponCreateV2CouponResourceListUseType enum
var AllowedClueCouponCreateV2CouponResourceListUseTypeEnumValues = []ClueCouponCreateV2CouponResourceListUseType{
	"OFFLINE",
	"LINK",
}

// NewClueCouponCreateV2CouponResourceListUseTypeFromValue returns a pointer to a valid ClueCouponCreateV2CouponResourceListUseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponCreateV2CouponResourceListUseTypeFromValue(v string) (*ClueCouponCreateV2CouponResourceListUseType, error) {
	ev := ClueCouponCreateV2CouponResourceListUseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponCreateV2CouponResourceListUseType: valid values are %v", v, AllowedClueCouponCreateV2CouponResourceListUseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponCreateV2CouponResourceListUseType) IsValid() bool {
	for _, existing := range AllowedClueCouponCreateV2CouponResourceListUseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_create_v2_coupon_resource_list_use_type value
func (v ClueCouponCreateV2CouponResourceListUseType) Ptr() *ClueCouponCreateV2CouponResourceListUseType {
	return &v
}
