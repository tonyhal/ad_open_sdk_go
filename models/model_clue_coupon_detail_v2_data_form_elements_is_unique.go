/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponDetailV2DataFormElementsIsUnique
type ClueCouponDetailV2DataFormElementsIsUnique string

// List of clue_coupon_detail_v2_data_form_elements_is_unique
const (
	Enum_0_ClueCouponDetailV2DataFormElementsIsUnique ClueCouponDetailV2DataFormElementsIsUnique = "0"
	Enum_1_ClueCouponDetailV2DataFormElementsIsUnique ClueCouponDetailV2DataFormElementsIsUnique = "1"
)

// All allowed values of ClueCouponDetailV2DataFormElementsIsUnique enum
var AllowedClueCouponDetailV2DataFormElementsIsUniqueEnumValues = []ClueCouponDetailV2DataFormElementsIsUnique{
	"0",
	"1",
}

// NewClueCouponDetailV2DataFormElementsIsUniqueFromValue returns a pointer to a valid ClueCouponDetailV2DataFormElementsIsUnique
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataFormElementsIsUniqueFromValue(v string) (*ClueCouponDetailV2DataFormElementsIsUnique, error) {
	ev := ClueCouponDetailV2DataFormElementsIsUnique(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataFormElementsIsUnique: valid values are %v", v, AllowedClueCouponDetailV2DataFormElementsIsUniqueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataFormElementsIsUnique) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataFormElementsIsUniqueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_form_elements_is_unique value
func (v ClueCouponDetailV2DataFormElementsIsUnique) Ptr() *ClueCouponDetailV2DataFormElementsIsUnique {
	return &v
}
