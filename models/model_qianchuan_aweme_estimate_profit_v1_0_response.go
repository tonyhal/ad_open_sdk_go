/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeEstimateProfitV10Response struct for QianchuanAwemeEstimateProfitV10Response
type QianchuanAwemeEstimateProfitV10Response struct {
	//
	Code *int64                                       `json:"code,omitempty"`
	Data *QianchuanAwemeEstimateProfitV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
