/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanKeywordCheckV10ResponseData
type QianchuanKeywordCheckV10ResponseData struct {
	//
	FailList []*QianchuanKeywordCheckV10ResponseDataFailListInner `json:"fail_list"`
	//
	SucList []string `json:"suc_list"`
}