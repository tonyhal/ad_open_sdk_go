/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryRebateBalanceV2ResponseData
type QueryRebateBalanceV2ResponseData struct {
	PageInfo *QueryRebateBalanceV2ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	Rebates []*QueryRebateBalanceV2ResponseDataRebatesInner `json:"rebates,omitempty"`
}
