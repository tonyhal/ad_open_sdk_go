/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SuggWordsV30ResponseData
type SuggWordsV30ResponseData struct {
	// 错误
	Err map[string]interface{} `json:"err,omitempty"`
	//
	List []*SuggWordsV30ResponseDataListInner `json:"list,omitempty"`
}
