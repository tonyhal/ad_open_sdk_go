/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordCreateV2V2ResponseData
type KeywordCreateV2V2ResponseData struct {
	//
	ErrorList []*KeywordCreateV2V2ResponseDataErrorListInner `json:"error_list,omitempty"`
	//
	SuccessList []*KeywordCreateV2V2ResponseDataSuccessListInner `json:"success_list,omitempty"`
}
