/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupCreateV30Response struct for AdlabGroupCreateV30Response
type AdlabGroupCreateV30Response struct {
	//
	Code *int64                           `json:"code,omitempty"`
	Data *AdlabGroupCreateV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
