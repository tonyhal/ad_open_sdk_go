/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferQueryTransferBalanceV30ResponseData
type CgTransferQueryTransferBalanceV30ResponseData struct {
	// 账户金额列表
	AccontAmountDetailList []*CgTransferQueryTransferBalanceV30ResponseDataAccontAmountDetailListInner `json:"accont_amount_detail_list,omitempty"`
}
