/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferQueryTransferBalanceV30Response struct for CgTransferQueryTransferBalanceV30Response
type CgTransferQueryTransferBalanceV30Response struct {
	//
	Code *int64                                         `json:"code,omitempty"`
	Data *CgTransferQueryTransferBalanceV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
