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

// CgTransferWalletTransferCreateV30AccountType
type CgTransferWalletTransferCreateV30AccountType string

// List of cg_transfer_wallet_transfer_create_v3.0_account_type
const (
	AD_CgTransferWalletTransferCreateV30AccountType        CgTransferWalletTransferCreateV30AccountType = "AD"
	AGENT_CgTransferWalletTransferCreateV30AccountType     CgTransferWalletTransferCreateV30AccountType = "AGENT"
	LOCAL_CgTransferWalletTransferCreateV30AccountType     CgTransferWalletTransferCreateV30AccountType = "LOCAL"
	QIANCHUAN_CgTransferWalletTransferCreateV30AccountType CgTransferWalletTransferCreateV30AccountType = "QIANCHUAN"
)

// All allowed values of CgTransferWalletTransferCreateV30AccountType enum
var AllowedCgTransferWalletTransferCreateV30AccountTypeEnumValues = []CgTransferWalletTransferCreateV30AccountType{
	"AD",
	"AGENT",
	"LOCAL",
	"QIANCHUAN",
}

// NewCgTransferWalletTransferCreateV30AccountTypeFromValue returns a pointer to a valid CgTransferWalletTransferCreateV30AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferWalletTransferCreateV30AccountTypeFromValue(v string) (*CgTransferWalletTransferCreateV30AccountType, error) {
	ev := CgTransferWalletTransferCreateV30AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferWalletTransferCreateV30AccountType: valid values are %v", v, AllowedCgTransferWalletTransferCreateV30AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferWalletTransferCreateV30AccountType) IsValid() bool {
	for _, existing := range AllowedCgTransferWalletTransferCreateV30AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_wallet_transfer_create_v3.0_account_type value
func (v CgTransferWalletTransferCreateV30AccountType) Ptr() *CgTransferWalletTransferCreateV30AccountType {
	return &v
}