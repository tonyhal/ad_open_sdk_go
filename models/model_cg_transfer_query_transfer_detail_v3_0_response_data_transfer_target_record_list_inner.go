/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferQueryTransferDetailV30ResponseDataTransferTargetRecordListInner struct for CgTransferQueryTransferDetailV30ResponseDataTransferTargetRecordListInner
type CgTransferQueryTransferDetailV30ResponseDataTransferTargetRecordListInner struct {
	// 锚定账户id，1:N的1
	AccountId *int64 `json:"account_id,omitempty"`
	// 目标账户id，1:N的N
	TargetAccountId *int64 `json:"target_account_id,omitempty"`
	// 转账金额(单位：分)
	TransferAmount *int64 `json:"transfer_amount,omitempty"`
	// 转账资金类型列表
	TransferCapitalRecordList []*CgTransferQueryTransferDetailV30ResponseDataTransferTargetRecordListInnerTransferCapitalRecordListInner `json:"transfer_capital_record_list,omitempty"`
	TransferStatus            *CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferStatus                                `json:"transfer_status,omitempty"`
}
