/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniAwemeAdUpdateV10ResponseData
type QianchuanUniAwemeAdUpdateV10ResponseData struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	ErrorList []*QianchuanUniAwemeAdUpdateV10ResponseDataErrorListInner `json:"error_list,omitempty"`
}
