/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanKeywordCheckV10Response struct for QianchuanKeywordCheckV10Response
type QianchuanKeywordCheckV10Response struct {
	//
	Code *int64                                `json:"code,omitempty"`
	Data *QianchuanKeywordCheckV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
