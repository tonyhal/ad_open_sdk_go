/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordUpdateV30ResponseData
type KeywordUpdateV30ResponseData struct {
	//
	ErrorList []*KeywordUpdateV30ResponseDataErrorListInner `json:"error_list,omitempty"`
	//
	SuccessList []*KeywordUpdateV30ResponseDataSuccessListInner `json:"success_list,omitempty"`
}
