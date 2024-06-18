/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QueryRebateBalanceV2FilteringStampType
type QueryRebateBalanceV2FilteringStampType string

// List of query_rebate_balance_v2_filtering_stamp_type
const (
	ONLINE_QueryRebateBalanceV2FilteringStampType  QueryRebateBalanceV2FilteringStampType = "ONLINE"
	DEFAULT_QueryRebateBalanceV2FilteringStampType QueryRebateBalanceV2FilteringStampType = "DEFAULT"
)

// All allowed values of QueryRebateBalanceV2FilteringStampType enum
var AllowedQueryRebateBalanceV2FilteringStampTypeEnumValues = []QueryRebateBalanceV2FilteringStampType{
	"ONLINE",
	"DEFAULT",
}

// NewQueryRebateBalanceV2FilteringStampTypeFromValue returns a pointer to a valid QueryRebateBalanceV2FilteringStampType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryRebateBalanceV2FilteringStampTypeFromValue(v string) (*QueryRebateBalanceV2FilteringStampType, error) {
	ev := QueryRebateBalanceV2FilteringStampType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryRebateBalanceV2FilteringStampType: valid values are %v", v, AllowedQueryRebateBalanceV2FilteringStampTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryRebateBalanceV2FilteringStampType) IsValid() bool {
	for _, existing := range AllowedQueryRebateBalanceV2FilteringStampTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to query_rebate_balance_v2_filtering_stamp_type value
func (v QueryRebateBalanceV2FilteringStampType) Ptr() *QueryRebateBalanceV2FilteringStampType {
	return &v
}