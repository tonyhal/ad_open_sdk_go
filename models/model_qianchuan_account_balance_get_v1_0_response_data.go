/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAccountBalanceGetV10ResponseData
type QianchuanAccountBalanceGetV10ResponseData struct {
	// 竞价不可用总余额
	AccountBiddingFrozen *int64 `json:"account_bidding_frozen,omitempty"`
	// 竞价总余额
	AccountBiddingTotal *int64 `json:"account_bidding_total,omitempty"`
	// 竞价可用总余额
	AccountBiddingValid *int64 `json:"account_bidding_valid,omitempty"`
	// 账户品牌冻结余额
	AccountBrandFrozen *int64 `json:"account_brand_frozen,omitempty"`
	// 账户品牌总余额
	AccountBrandTotal *int64 `json:"account_brand_total,omitempty"`
	// 账户品牌可用余额
	AccountBrandValid *int64 `json:"account_brand_valid,omitempty"`
	// 账户不可用总余额
	AccountFrozen *int64 `json:"account_frozen,omitempty"`
	// 通用不可用总余额
	AccountGeneralFrozen *int64 `json:"account_general_frozen,omitempty"`
	// 通用总余额
	AccountGeneralTotal *int64 `json:"account_general_total,omitempty"`
	// 通用可用总余额
	AccountGeneralValid *int64 `json:"account_general_valid,omitempty"`
	// 账户总余额
	AccountTotal *int64 `json:"account_total,omitempty"`
	// 账户可用总余额
	AccountValid *int64 `json:"account_valid,omitempty"`
	// 广告主ID或代理商ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 总共享赠款
	ShareGrantTotal *int64 `json:"share_grant_total,omitempty"`
	// 共享钱包竞价可用余额
	ShareWalletBiddingValid *int64 `json:"share_wallet_bidding_valid,omitempty"`
	// 共享钱包品牌可用余额
	ShareWalletBrandValid *int64 `json:"share_wallet_brand_valid,omitempty"`
	// 共享钱包通用可用余额
	ShareWalletGeneralValid *int64 `json:"share_wallet_general_valid,omitempty"`
	// 共享钱包id
	ShareWalletId *string `json:"share_wallet_id,omitempty"`
	// 共享钱包名称
	ShareWalletName *string `json:"share_wallet_name,omitempty"`
	//
	ShareWalletTotal *int64 `json:"share_wallet_total,omitempty"`
}
