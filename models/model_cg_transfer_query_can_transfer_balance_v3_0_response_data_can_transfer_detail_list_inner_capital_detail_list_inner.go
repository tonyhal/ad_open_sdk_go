/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferQueryCanTransferBalanceV30ResponseDataCanTransferDetailListInnerCapitalDetailListInner struct for CgTransferQueryCanTransferBalanceV30ResponseDataCanTransferDetailListInnerCapitalDetailListInner
type CgTransferQueryCanTransferBalanceV30ResponseDataCanTransferDetailListInnerCapitalDetailListInner struct {
	CapitalType CgTransferQueryCanTransferBalanceV30DataCanTransferDetailListCapitalDetailListCapitalType `json:"capital_type"`
	// 减款方可转资金余额(单位：分)
	TransferBalance int64 `json:"transfer_balance"`
}
