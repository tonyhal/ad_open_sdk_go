/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserFundTransactionGetV2DataListTransactionType
type AdvertiserFundTransactionGetV2DataListTransactionType string

// List of advertiser_fund_transaction_get_v2_data_list_transaction_type
const (
	RECHARGE_AdvertiserFundTransactionGetV2DataListTransactionType AdvertiserFundTransactionGetV2DataListTransactionType = "RECHARGE"
	TRANSFER_AdvertiserFundTransactionGetV2DataListTransactionType AdvertiserFundTransactionGetV2DataListTransactionType = "TRANSFER"
)

// All allowed values of AdvertiserFundTransactionGetV2DataListTransactionType enum
var AllowedAdvertiserFundTransactionGetV2DataListTransactionTypeEnumValues = []AdvertiserFundTransactionGetV2DataListTransactionType{
	"RECHARGE",
	"TRANSFER",
}

// NewAdvertiserFundTransactionGetV2DataListTransactionTypeFromValue returns a pointer to a valid AdvertiserFundTransactionGetV2DataListTransactionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserFundTransactionGetV2DataListTransactionTypeFromValue(v string) (*AdvertiserFundTransactionGetV2DataListTransactionType, error) {
	ev := AdvertiserFundTransactionGetV2DataListTransactionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserFundTransactionGetV2DataListTransactionType: valid values are %v", v, AllowedAdvertiserFundTransactionGetV2DataListTransactionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserFundTransactionGetV2DataListTransactionType) IsValid() bool {
	for _, existing := range AllowedAdvertiserFundTransactionGetV2DataListTransactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_fund_transaction_get_v2_data_list_transaction_type value
func (v AdvertiserFundTransactionGetV2DataListTransactionType) Ptr() *AdvertiserFundTransactionGetV2DataListTransactionType {
	return &v
}
