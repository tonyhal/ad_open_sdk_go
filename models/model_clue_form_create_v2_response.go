/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormCreateV2Response struct for ClueFormCreateV2Response
type ClueFormCreateV2Response struct {
	//
	Code *int64                        `json:"code,omitempty"`
	Data *ClueFormCreateV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
