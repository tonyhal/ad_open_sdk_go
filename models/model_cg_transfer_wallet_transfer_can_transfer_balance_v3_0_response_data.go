/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferWalletTransferCanTransferBalanceV30ResponseData
type CgTransferWalletTransferCanTransferBalanceV30ResponseData struct {
	// 可转余额信息列表
	CanTransferDetailList []*CgTransferWalletTransferCanTransferBalanceV30ResponseDataCanTransferDetailListInner `json:"can_transfer_detail_list,omitempty"`
}