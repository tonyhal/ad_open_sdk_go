/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormCreateV2RequestExtendInfoCountConfig
type ClueFormCreateV2RequestExtendInfoCountConfig struct {
	CountType *ClueFormCreateV2ExtendInfoCountConfigCountType `json:"count_type,omitempty"`
	//
	Prefix *string `json:"prefix,omitempty"`
	//
	StartNum *int64 `json:"start_num,omitempty"`
	//
	Suffix *string `json:"suffix,omitempty"`
}
