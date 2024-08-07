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

// SharedWalletDailyStatGetV30DataResultsWalletType
type SharedWalletDailyStatGetV30DataResultsWalletType string

// List of shared_wallet_daily_stat_get_v3.0_data_results_wallet_type
const (
	MAIN_WALLET_SharedWalletDailyStatGetV30DataResultsWalletType        SharedWalletDailyStatGetV30DataResultsWalletType = "MAIN_WALLET"
	SUB_CONSUME_WALLET_SharedWalletDailyStatGetV30DataResultsWalletType SharedWalletDailyStatGetV30DataResultsWalletType = "SUB_CONSUME_WALLET"
	SUB_MANAGE_WALLET_SharedWalletDailyStatGetV30DataResultsWalletType  SharedWalletDailyStatGetV30DataResultsWalletType = "SUB_MANAGE_WALLET"
)

// All allowed values of SharedWalletDailyStatGetV30DataResultsWalletType enum
var AllowedSharedWalletDailyStatGetV30DataResultsWalletTypeEnumValues = []SharedWalletDailyStatGetV30DataResultsWalletType{
	"MAIN_WALLET",
	"SUB_CONSUME_WALLET",
	"SUB_MANAGE_WALLET",
}

// NewSharedWalletDailyStatGetV30DataResultsWalletTypeFromValue returns a pointer to a valid SharedWalletDailyStatGetV30DataResultsWalletType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharedWalletDailyStatGetV30DataResultsWalletTypeFromValue(v string) (*SharedWalletDailyStatGetV30DataResultsWalletType, error) {
	ev := SharedWalletDailyStatGetV30DataResultsWalletType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharedWalletDailyStatGetV30DataResultsWalletType: valid values are %v", v, AllowedSharedWalletDailyStatGetV30DataResultsWalletTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharedWalletDailyStatGetV30DataResultsWalletType) IsValid() bool {
	for _, existing := range AllowedSharedWalletDailyStatGetV30DataResultsWalletTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to shared_wallet_daily_stat_get_v3.0_data_results_wallet_type value
func (v SharedWalletDailyStatGetV30DataResultsWalletType) Ptr() *SharedWalletDailyStatGetV30DataResultsWalletType {
	return &v
}
