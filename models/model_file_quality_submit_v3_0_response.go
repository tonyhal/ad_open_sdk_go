/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileQualitySubmitV30Response struct for FileQualitySubmitV30Response
type FileQualitySubmitV30Response struct {
	//
	Code *int64                            `json:"code,omitempty"`
	Data *FileQualitySubmitV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
