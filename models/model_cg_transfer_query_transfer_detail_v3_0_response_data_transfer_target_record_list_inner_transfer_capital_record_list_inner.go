/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferQueryTransferDetailV30ResponseDataTransferTargetRecordListInnerTransferCapitalRecordListInner struct for CgTransferQueryTransferDetailV30ResponseDataTransferTargetRecordListInnerTransferCapitalRecordListInner
type CgTransferQueryTransferDetailV30ResponseDataTransferTargetRecordListInnerTransferCapitalRecordListInner struct {
	CapitalType *CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListCapitalType `json:"capital_type,omitempty"`
	// 失败原因
	FailReason *string `json:"fail_reason,omitempty"`
	// 转账资金金额(单位：分)
	TransferAmount *int64                                                                                               `json:"transfer_amount,omitempty"`
	TransferStatus *CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus `json:"transfer_status,omitempty"`
}
