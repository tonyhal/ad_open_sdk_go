/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordUpdateV2V2ResponseData
type KeywordUpdateV2V2ResponseData struct {
	//
	ErrorList []*KeywordUpdateV2V2ResponseDataErrorListInner `json:"error_list,omitempty"`
	//
	SuccessList []*KeywordUpdateV2V2ResponseDataSuccessListInner `json:"success_list,omitempty"`
}
