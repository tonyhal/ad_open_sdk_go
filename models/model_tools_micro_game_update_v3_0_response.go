/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroGameUpdateV30Response struct for ToolsMicroGameUpdateV30Response
type ToolsMicroGameUpdateV30Response struct {
	//
	Code *int64                               `json:"code,omitempty"`
	Data *ToolsMicroGameUpdateV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
