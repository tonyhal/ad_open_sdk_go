/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletMainWalletGetV30ResponseDataMainWalletInfo 资金共享大钱包信息
type SharedWalletMainWalletGetV30ResponseDataMainWalletInfo struct {
	AllocatedBalance *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalance `json:"allocated_balance,omitempty"`
	// 大钱包ID
	MainWalletId *int64 `json:"main_wallet_id,omitempty"`
	// 钱包总余额(单位元)
	TotalBalance       *float64                                                                  `json:"total_balance,omitempty"`
	UnallocatedBalance *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalance `json:"unallocated_balance,omitempty"`
	// 钱包描述
	WalletDescription *string `json:"wallet_description,omitempty"`
	// 钱包标签
	WalletLabel []string `json:"wallet_label,omitempty"`
	// 钱包名称
	WalletName *string `json:"wallet_name,omitempty"`
}
