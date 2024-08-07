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

// CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus
type CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus string

// List of cg_transfer_wallet_transfer_detail_v3.0_data_transfer_wallet_record_list_transfer_status
const (
	NO_TRANSFER_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus      CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus = "NO_TRANSFER"
	TRANSFER_FAILED_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus  CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus = "TRANSFER_FAILED"
	TRANSFER_ING_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus     CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus = "TRANSFER_ING"
	TRANSFER_PART_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus    CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus = "TRANSFER_PART"
	TRANSFER_SUCCESS_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus = "TRANSFER_SUCCESS"
)

// All allowed values of CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus enum
var AllowedCgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatusEnumValues = []CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus{
	"NO_TRANSFER",
	"TRANSFER_FAILED",
	"TRANSFER_ING",
	"TRANSFER_PART",
	"TRANSFER_SUCCESS",
}

// NewCgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatusFromValue returns a pointer to a valid CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatusFromValue(v string) (*CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus, error) {
	ev := CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus: valid values are %v", v, AllowedCgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus) IsValid() bool {
	for _, existing := range AllowedCgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_wallet_transfer_detail_v3.0_data_transfer_wallet_record_list_transfer_status value
func (v CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus) Ptr() *CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus {
	return &v
}
