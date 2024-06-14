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

// CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus
type CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus string

// List of cg_transfer_query_transfer_detail_v3.0_data_transfer_target_record_list_transfer_status
const (
	NO_TRANSFER_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus      CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus = "NO_TRANSFER"
	TRANSFER_FAILED_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus  CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus = "TRANSFER_FAILED"
	TRANSFER_ING_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus     CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus = "TRANSFER_ING"
	TRANSFER_PART_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus    CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus = "TRANSFER_PART"
	TRANSFER_SUCCESS_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus = "TRANSFER_SUCCESS"
)

// All allowed values of CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus enum
var AllowedCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatusEnumValues = []CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus{
	"NO_TRANSFER",
	"TRANSFER_FAILED",
	"TRANSFER_ING",
	"TRANSFER_PART",
	"TRANSFER_SUCCESS",
}

// NewCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatusFromValue returns a pointer to a valid CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatusFromValue(v string) (*CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus, error) {
	ev := CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus: valid values are %v", v, AllowedCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus) IsValid() bool {
	for _, existing := range AllowedCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_query_transfer_detail_v3.0_data_transfer_target_record_list_transfer_status value
func (v CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus) Ptr() *CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus {
	return &v
}
