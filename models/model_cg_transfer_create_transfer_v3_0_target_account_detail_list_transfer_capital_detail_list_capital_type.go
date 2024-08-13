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

// CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType
type CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType string

// List of cg_transfer_create_transfer_v3.0_target_account_detail_list_transfer_capital_detail_list_capital_type
const (
	CREDIT_BIDDING_CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType = "CREDIT_BIDDING"
	CREDIT_BRAND_CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType   CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType = "CREDIT_BRAND"
	CREDIT_GENERAL_CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType = "CREDIT_GENERAL"
	PREPAY_BIDDING_CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType = "PREPAY_BIDDING"
	PREPAY_BRAND_CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType   CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType = "PREPAY_BRAND"
	PREPAY_GENERAL_CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType = "PREPAY_GENERAL"
)

// All allowed values of CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType enum
var AllowedCgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalTypeEnumValues = []CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType{
	"CREDIT_BIDDING",
	"CREDIT_BRAND",
	"CREDIT_GENERAL",
	"PREPAY_BIDDING",
	"PREPAY_BRAND",
	"PREPAY_GENERAL",
}

// NewCgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalTypeFromValue returns a pointer to a valid CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalTypeFromValue(v string) (*CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType, error) {
	ev := CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType: valid values are %v", v, AllowedCgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType) IsValid() bool {
	for _, existing := range AllowedCgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_create_transfer_v3.0_target_account_detail_list_transfer_capital_detail_list_capital_type value
func (v CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType) Ptr() *CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType {
	return &v
}
