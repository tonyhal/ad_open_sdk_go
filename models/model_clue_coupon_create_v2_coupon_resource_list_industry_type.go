/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponCreateV2CouponResourceListIndustryType
type ClueCouponCreateV2CouponResourceListIndustryType string

// List of clue_coupon_create_v2_coupon_resource_list_industry_type
const (
	FOOD_ClueCouponCreateV2CouponResourceListIndustryType          ClueCouponCreateV2CouponResourceListIndustryType = "FOOD"
	GAME_ClueCouponCreateV2CouponResourceListIndustryType          ClueCouponCreateV2CouponResourceListIndustryType = "GAME"
	OTHER_ClueCouponCreateV2CouponResourceListIndustryType         ClueCouponCreateV2CouponResourceListIndustryType = "OTHER"
	FINANCIAL_ClueCouponCreateV2CouponResourceListIndustryType     ClueCouponCreateV2CouponResourceListIndustryType = "FINANCIAL"
	TICKET_ClueCouponCreateV2CouponResourceListIndustryType        ClueCouponCreateV2CouponResourceListIndustryType = "TICKET"
	ENTERTAINMENT_ClueCouponCreateV2CouponResourceListIndustryType ClueCouponCreateV2CouponResourceListIndustryType = "ENTERTAINMENT"
)

// All allowed values of ClueCouponCreateV2CouponResourceListIndustryType enum
var AllowedClueCouponCreateV2CouponResourceListIndustryTypeEnumValues = []ClueCouponCreateV2CouponResourceListIndustryType{
	"FOOD",
	"GAME",
	"OTHER",
	"FINANCIAL",
	"TICKET",
	"ENTERTAINMENT",
}

// NewClueCouponCreateV2CouponResourceListIndustryTypeFromValue returns a pointer to a valid ClueCouponCreateV2CouponResourceListIndustryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponCreateV2CouponResourceListIndustryTypeFromValue(v string) (*ClueCouponCreateV2CouponResourceListIndustryType, error) {
	ev := ClueCouponCreateV2CouponResourceListIndustryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponCreateV2CouponResourceListIndustryType: valid values are %v", v, AllowedClueCouponCreateV2CouponResourceListIndustryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponCreateV2CouponResourceListIndustryType) IsValid() bool {
	for _, existing := range AllowedClueCouponCreateV2CouponResourceListIndustryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_create_v2_coupon_resource_list_industry_type value
func (v ClueCouponCreateV2CouponResourceListIndustryType) Ptr() *ClueCouponCreateV2CouponResourceListIndustryType {
	return &v
}
