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

// CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus
type CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus string

// List of cg_transfer_wallet_transfer_list_v3.0_data_record_list_transfer_capital_record_list_transfer_status
const (
	NO_TRANSFER_CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus      CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus = "NO_TRANSFER"
	TRANSFER_FAILED_CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus  CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus = "TRANSFER_FAILED"
	TRANSFER_ING_CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus     CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus = "TRANSFER_ING"
	TRANSFER_PART_CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus    CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus = "TRANSFER_PART"
	TRANSFER_SUCCESS_CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus = "TRANSFER_SUCCESS"
)

// All allowed values of CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus enum
var AllowedCgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatusEnumValues = []CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus{
	"NO_TRANSFER",
	"TRANSFER_FAILED",
	"TRANSFER_ING",
	"TRANSFER_PART",
	"TRANSFER_SUCCESS",
}

// NewCgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatusFromValue returns a pointer to a valid CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatusFromValue(v string) (*CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus, error) {
	ev := CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus: valid values are %v", v, AllowedCgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus) IsValid() bool {
	for _, existing := range AllowedCgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_wallet_transfer_list_v3.0_data_record_list_transfer_capital_record_list_transfer_status value
func (v CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus) Ptr() *CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus {
	return &v
}
