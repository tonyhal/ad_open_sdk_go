/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniAwemeAuthorizedGetV10ResponseDataPageInfo
type QianchuanUniAwemeAuthorizedGetV10ResponseDataPageInfo struct {
	//
	Page int64 `json:"page"`
	//
	PageSize int64 `json:"page_size"`
	//
	TotalPage *int64 `json:"total_page,omitempty"`
	//
	TotolNumer *int64 `json:"totol_numer,omitempty"`
}
