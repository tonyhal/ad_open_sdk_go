/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCompensateStatusGetV10Response struct for QianchuanAdCompensateStatusGetV10Response
type QianchuanAdCompensateStatusGetV10Response struct {
	//
	Code *int64                                         `json:"code,omitempty"`
	Data *QianchuanAdCompensateStatusGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
