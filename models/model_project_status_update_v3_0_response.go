/*
API Title

巨量引擎开放平台

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectStatusUpdateV30Response struct for ProjectStatusUpdateV30Response
type ProjectStatusUpdateV30Response struct {
	//
	Code *int64                              `json:"code,omitempty"`
	Data *ProjectStatusUpdateV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
