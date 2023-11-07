/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponGetV2DataListActivityType
type ClueCouponGetV2DataListActivityType string

// List of clue_coupon_get_v2_data_list_activity_type
const (
	DIRECT_NEED_PHONE_ClueCouponGetV2DataListActivityType     ClueCouponGetV2DataListActivityType = "DIRECT_NEED_PHONE"
	DIRECT_NOT_NEED_PHONE_ClueCouponGetV2DataListActivityType ClueCouponGetV2DataListActivityType = "DIRECT_NOT_NEED_PHONE"
)

// All allowed values of ClueCouponGetV2DataListActivityType enum
var AllowedClueCouponGetV2DataListActivityTypeEnumValues = []ClueCouponGetV2DataListActivityType{
	"DIRECT_NEED_PHONE",
	"DIRECT_NOT_NEED_PHONE",
}

// NewClueCouponGetV2DataListActivityTypeFromValue returns a pointer to a valid ClueCouponGetV2DataListActivityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponGetV2DataListActivityTypeFromValue(v string) (*ClueCouponGetV2DataListActivityType, error) {
	ev := ClueCouponGetV2DataListActivityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponGetV2DataListActivityType: valid values are %v", v, AllowedClueCouponGetV2DataListActivityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponGetV2DataListActivityType) IsValid() bool {
	for _, existing := range AllowedClueCouponGetV2DataListActivityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_get_v2_data_list_activity_type value
func (v ClueCouponGetV2DataListActivityType) Ptr() *ClueCouponGetV2DataListActivityType {
	return &v
}
