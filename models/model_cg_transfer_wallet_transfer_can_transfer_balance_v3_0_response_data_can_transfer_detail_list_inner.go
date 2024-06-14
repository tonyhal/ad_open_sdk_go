/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferWalletTransferCanTransferBalanceV30ResponseDataCanTransferDetailListInner struct for CgTransferWalletTransferCanTransferBalanceV30ResponseDataCanTransferDetailListInner
type CgTransferWalletTransferCanTransferBalanceV30ResponseDataCanTransferDetailListInner struct {
	// 减款钱包非品牌资金最大可转出金额（单位：分）
	NonBrandMaxTransferBalance *int64 `json:"non_brand_max_transfer_balance,omitempty"`
	// 加款钱包可转余额信息列表
	PayeeTransferAmountDetailList []*CgTransferWalletTransferCanTransferBalanceV30ResponseDataCanTransferDetailListInnerPayeeTransferAmountDetailListInner `json:"payee_transfer_amount_detail_list,omitempty"`
	// 减款钱包可转资金列表
	RemitterCapitalDetailList []*CgTransferWalletTransferCanTransferBalanceV30ResponseDataCanTransferDetailListInnerRemitterCapitalDetailListInner `json:"remitter_capital_detail_list,omitempty"`
	// 减款钱包id
	RemitterWalletId *int64 `json:"remitter_wallet_id,omitempty"`
}
