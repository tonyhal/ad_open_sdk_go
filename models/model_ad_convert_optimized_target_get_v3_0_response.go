/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdConvertOptimizedTargetGetV30Response struct for AdConvertOptimizedTargetGetV30Response
type AdConvertOptimizedTargetGetV30Response struct {
	//
	Code *int64                                      `json:"code,omitempty"`
	Data *AdConvertOptimizedTargetGetV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}