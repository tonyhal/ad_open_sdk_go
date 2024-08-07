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

// QueryRebateAccountingInfoV2FilteringUseType
type QueryRebateAccountingInfoV2FilteringUseType string

// List of query_rebate_accounting_info_v2_filtering_use_type
const (
	CASH_QueryRebateAccountingInfoV2FilteringUseType       QueryRebateAccountingInfoV2FilteringUseType = "CASH"
	CHARGE_QueryRebateAccountingInfoV2FilteringUseType     QueryRebateAccountingInfoV2FilteringUseType = "CHARGE"
	HEDGING_QueryRebateAccountingInfoV2FilteringUseType    QueryRebateAccountingInfoV2FilteringUseType = "HEDGING"
	NONPAYMENT_QueryRebateAccountingInfoV2FilteringUseType QueryRebateAccountingInfoV2FilteringUseType = "NONPAYMENT"
	REVERT_QueryRebateAccountingInfoV2FilteringUseType     QueryRebateAccountingInfoV2FilteringUseType = "REVERT"
)

// All allowed values of QueryRebateAccountingInfoV2FilteringUseType enum
var AllowedQueryRebateAccountingInfoV2FilteringUseTypeEnumValues = []QueryRebateAccountingInfoV2FilteringUseType{
	"CASH",
	"CHARGE",
	"HEDGING",
	"NONPAYMENT",
	"REVERT",
}

// NewQueryRebateAccountingInfoV2FilteringUseTypeFromValue returns a pointer to a valid QueryRebateAccountingInfoV2FilteringUseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryRebateAccountingInfoV2FilteringUseTypeFromValue(v string) (*QueryRebateAccountingInfoV2FilteringUseType, error) {
	ev := QueryRebateAccountingInfoV2FilteringUseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryRebateAccountingInfoV2FilteringUseType: valid values are %v", v, AllowedQueryRebateAccountingInfoV2FilteringUseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryRebateAccountingInfoV2FilteringUseType) IsValid() bool {
	for _, existing := range AllowedQueryRebateAccountingInfoV2FilteringUseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to query_rebate_accounting_info_v2_filtering_use_type value
func (v QueryRebateAccountingInfoV2FilteringUseType) Ptr() *QueryRebateAccountingInfoV2FilteringUseType {
	return &v
}
