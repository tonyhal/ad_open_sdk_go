/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceAdOnlyAllocatedBalance 巨量广告业务线已分配余额(单位元)
type SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceAdOnlyAllocatedBalance struct {
	AvailableBalance   *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceAdOnlyAllocatedBalanceAvailableBalance   `json:"available_balance,omitempty"`
	UnavailableBalance *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceAdOnlyAllocatedBalanceUnavailableBalance `json:"unavailable_balance,omitempty"`
}
