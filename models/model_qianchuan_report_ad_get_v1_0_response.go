/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportAdGetV10Response struct for QianchuanReportAdGetV10Response
type QianchuanReportAdGetV10Response struct {
	//
	Code *int64                               `json:"code,omitempty"`
	Data *QianchuanReportAdGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
