/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferCreateTransferV30RequestTargetAccountDetailListInnerTransferCapitalDetailListInner struct for CgTransferCreateTransferV30RequestTargetAccountDetailListInnerTransferCapitalDetailListInner
type CgTransferCreateTransferV30RequestTargetAccountDetailListInnerTransferCapitalDetailListInner struct {
	CapitalType CgTransferCreateTransferV30TargetAccountDetailListTransferCapitalDetailListCapitalType `json:"capital_type"`
	// 转账资金金额(单位：分)
	TransferAmount int64 `json:"transfer_amount"`
}
