/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueSmartphoneGetV2Response struct for ClueSmartphoneGetV2Response
type ClueSmartphoneGetV2Response struct {
	//
	Code *int64                           `json:"code,omitempty"`
	Data *ClueSmartphoneGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
