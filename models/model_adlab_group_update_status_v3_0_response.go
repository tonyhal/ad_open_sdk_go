/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateStatusV30Response struct for AdlabGroupUpdateStatusV30Response
type AdlabGroupUpdateStatusV30Response struct {
	//
	Code *int64                                 `json:"code,omitempty"`
	Data *AdlabGroupUpdateStatusV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
