/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletDailyStatGetV30Response struct for SharedWalletDailyStatGetV30Response
type SharedWalletDailyStatGetV30Response struct {
	//
	Code *int64                                   `json:"code,omitempty"`
	Data *SharedWalletDailyStatGetV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}