/*
API Title

巨量引擎开放平台

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CommonResponse struct for CommonResponse
type CommonResponse struct {
	Code      *int64                 `json:"code,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
	Message   *string                `json:"message,omitempty"`
	RequestId *string                `json:"request_id,omitempty"`
}
