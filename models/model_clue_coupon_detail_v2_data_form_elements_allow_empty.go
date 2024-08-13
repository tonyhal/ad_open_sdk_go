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

// ClueCouponDetailV2DataFormElementsAllowEmpty
type ClueCouponDetailV2DataFormElementsAllowEmpty string

// List of clue_coupon_detail_v2_data_form_elements_allow_empty
const (
	Enum_0_ClueCouponDetailV2DataFormElementsAllowEmpty ClueCouponDetailV2DataFormElementsAllowEmpty = "0"
	Enum_1_ClueCouponDetailV2DataFormElementsAllowEmpty ClueCouponDetailV2DataFormElementsAllowEmpty = "1"
)

// All allowed values of ClueCouponDetailV2DataFormElementsAllowEmpty enum
var AllowedClueCouponDetailV2DataFormElementsAllowEmptyEnumValues = []ClueCouponDetailV2DataFormElementsAllowEmpty{
	"0",
	"1",
}

// NewClueCouponDetailV2DataFormElementsAllowEmptyFromValue returns a pointer to a valid ClueCouponDetailV2DataFormElementsAllowEmpty
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataFormElementsAllowEmptyFromValue(v string) (*ClueCouponDetailV2DataFormElementsAllowEmpty, error) {
	ev := ClueCouponDetailV2DataFormElementsAllowEmpty(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataFormElementsAllowEmpty: valid values are %v", v, AllowedClueCouponDetailV2DataFormElementsAllowEmptyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataFormElementsAllowEmpty) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataFormElementsAllowEmptyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_form_elements_allow_empty value
func (v ClueCouponDetailV2DataFormElementsAllowEmpty) Ptr() *ClueCouponDetailV2DataFormElementsAllowEmpty {
	return &v
}
