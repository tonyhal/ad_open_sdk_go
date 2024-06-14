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

// CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform
type CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform string

// List of cg_transfer_wallet_transfer_can_transfer_balance_v3.0_data_can_transfer_detail_list_remitter_capital_detail_list_platform
const (
	AD_CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform        CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform = "AD"
	AD_ALL_CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform    CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform = "AD_ALL"
	BENDITUI_CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform  CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform = "BENDITUI"
	QIANCHUAN_CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform = "QIANCHUAN"
	STAR_CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform      CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform = "STAR"
)

// All allowed values of CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform enum
var AllowedCgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatformEnumValues = []CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform{
	"AD",
	"AD_ALL",
	"BENDITUI",
	"QIANCHUAN",
	"STAR",
}

// NewCgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatformFromValue returns a pointer to a valid CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatformFromValue(v string) (*CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform, error) {
	ev := CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform: valid values are %v", v, AllowedCgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform) IsValid() bool {
	for _, existing := range AllowedCgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_wallet_transfer_can_transfer_balance_v3.0_data_can_transfer_detail_list_remitter_capital_detail_list_platform value
func (v CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform) Ptr() *CgTransferWalletTransferCanTransferBalanceV30DataCanTransferDetailListRemitterCapitalDetailListPlatform {
	return &v
}
