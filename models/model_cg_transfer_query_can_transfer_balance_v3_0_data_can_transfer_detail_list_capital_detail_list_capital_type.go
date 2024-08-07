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

// CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType
type CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType string

// List of cg_transfer_query_can_transfer_balance_v3.0_data_can_transfer_detail_list_capital_detail_list_capital_type
const (
	CREDIT_BIDDING_CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType = "CREDIT_BIDDING"
	CREDIT_BRAND_CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType   CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType = "CREDIT_BRAND"
	CREDIT_GENERAL_CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType = "CREDIT_GENERAL"
	PREPAY_BIDDING_CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType = "PREPAY_BIDDING"
	PREPAY_BRAND_CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType   CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType = "PREPAY_BRAND"
	PREPAY_GENERAL_CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType = "PREPAY_GENERAL"
)

// All allowed values of CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType enum
var AllowedCgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalTypeEnumValues = []CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType{
	"CREDIT_BIDDING",
	"CREDIT_BRAND",
	"CREDIT_GENERAL",
	"PREPAY_BIDDING",
	"PREPAY_BRAND",
	"PREPAY_GENERAL",
}

// NewCgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalTypeFromValue returns a pointer to a valid CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalTypeFromValue(v string) (*CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType, error) {
	ev := CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType: valid values are %v", v, AllowedCgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType) IsValid() bool {
	for _, existing := range AllowedCgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_query_can_transfer_balance_v3.0_data_can_transfer_detail_list_capital_detail_list_capital_type value
func (v CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType) Ptr() *CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType {
	return &v
}
