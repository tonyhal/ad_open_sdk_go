/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskUnshareV30ResponseData
type StardeliveryTaskUnshareV30ResponseData struct {
	//
	ErrorList []*StardeliveryTaskUnshareV30ResponseDataErrorListInner `json:"error_list,omitempty"`
	//
	SuccessList []int64 `json:"success_list,omitempty"`
}
