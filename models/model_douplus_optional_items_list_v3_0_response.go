/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOptionalItemsListV30Response struct for DouplusOptionalItemsListV30Response
type DouplusOptionalItemsListV30Response struct {
	//
	Code *int64                                   `json:"code,omitempty"`
	Data *DouplusOptionalItemsListV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}