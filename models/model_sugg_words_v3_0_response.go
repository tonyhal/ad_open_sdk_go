/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SuggWordsV30Response struct for SuggWordsV30Response
type SuggWordsV30Response struct {
	//
	Code *int64                    `json:"code,omitempty"`
	Data *SuggWordsV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
