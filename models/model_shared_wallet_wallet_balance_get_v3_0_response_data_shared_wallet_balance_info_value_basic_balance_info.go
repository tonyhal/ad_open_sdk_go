/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValueBasicBalanceInfo 常用余额信息
type SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValueBasicBalanceInfo struct {
	// 非赠款总余额(单位元)
	NonGrantBalance *float64 `json:"non_grant_balance,omitempty"`
	// 非赠款可转余额(单位元)
	NonGrantTransferableBalance *float64 `json:"non_grant_transferable_balance,omitempty"`
	// 非赠款可用余额(单位元)
	NonGrantValidBalance *float64 `json:"non_grant_valid_balance,omitempty"`
	// 总余额(单位元)
	TotalBalance *float64 `json:"total_balance,omitempty"`
	// 总可转余额(单位元)
	TotalTransferableBalance *float64 `json:"total_transferable_balance,omitempty"`
	// 总可用余额(单位元)
	TotalValidBalance *float64 `json:"total_valid_balance,omitempty"`
}
