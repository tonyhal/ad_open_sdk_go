/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletWalletInfoGetV30ResponseDataWalletInfoInner struct for SharedWalletWalletInfoGetV30ResponseDataWalletInfoInner
type SharedWalletWalletInfoGetV30ResponseDataWalletInfoInner struct {
	CommonWalletInfo *SharedWalletWalletInfoGetV30ResponseDataWalletInfoInnerCommonWalletInfo `json:"common_wallet_info,omitempty"`
	MainWalletInfo   *SharedWalletWalletInfoGetV30ResponseDataWalletInfoInnerMainWalletInfo   `json:"main_wallet_info,omitempty"`
	SubWalletInfo    *SharedWalletWalletInfoGetV30ResponseDataWalletInfoInnerSubWalletInfo    `json:"sub_wallet_info,omitempty"`
	// 共享钱包ID
	WalletId   *int64                                                `json:"wallet_id,omitempty"`
	WalletType *SharedWalletWalletInfoGetV30DataWalletInfoWalletType `json:"wallet_type,omitempty"`
}
