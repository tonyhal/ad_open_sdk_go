/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletTransactionDetailGetV30ResponseData
type SharedWalletTransactionDetailGetV30ResponseData struct {
	PageInfo *SharedWalletTransactionDetailGetV30ResponseDataPageInfo `json:"page_info,omitempty"`
	// 流水明细结果
	Results []*SharedWalletTransactionDetailGetV30ResponseDataResultsInner `json:"results,omitempty"`
}
