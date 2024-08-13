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

// SharedWalletTransactionDetailGetV30DataResultsPayeeType
type SharedWalletTransactionDetailGetV30DataResultsPayeeType string

// List of shared_wallet_transaction_detail_get_v3.0_data_results_payee_type
const (
	AGENT_SharedWalletTransactionDetailGetV30DataResultsPayeeType         SharedWalletTransactionDetailGetV30DataResultsPayeeType = "AGENT"
	SHARED_WALLET_SharedWalletTransactionDetailGetV30DataResultsPayeeType SharedWalletTransactionDetailGetV30DataResultsPayeeType = "SHARED_WALLET"
)

// All allowed values of SharedWalletTransactionDetailGetV30DataResultsPayeeType enum
var AllowedSharedWalletTransactionDetailGetV30DataResultsPayeeTypeEnumValues = []SharedWalletTransactionDetailGetV30DataResultsPayeeType{
	"AGENT",
	"SHARED_WALLET",
}

// NewSharedWalletTransactionDetailGetV30DataResultsPayeeTypeFromValue returns a pointer to a valid SharedWalletTransactionDetailGetV30DataResultsPayeeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharedWalletTransactionDetailGetV30DataResultsPayeeTypeFromValue(v string) (*SharedWalletTransactionDetailGetV30DataResultsPayeeType, error) {
	ev := SharedWalletTransactionDetailGetV30DataResultsPayeeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharedWalletTransactionDetailGetV30DataResultsPayeeType: valid values are %v", v, AllowedSharedWalletTransactionDetailGetV30DataResultsPayeeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharedWalletTransactionDetailGetV30DataResultsPayeeType) IsValid() bool {
	for _, existing := range AllowedSharedWalletTransactionDetailGetV30DataResultsPayeeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to shared_wallet_transaction_detail_get_v3.0_data_results_payee_type value
func (v SharedWalletTransactionDetailGetV30DataResultsPayeeType) Ptr() *SharedWalletTransactionDetailGetV30DataResultsPayeeType {
	return &v
}
