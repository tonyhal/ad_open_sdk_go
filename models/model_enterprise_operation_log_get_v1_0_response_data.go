/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseOperationLogGetV10ResponseData
type EnterpriseOperationLogGetV10ResponseData struct {
	//
	ActionCount *int64 `json:"action_count,omitempty"`
	//
	List []*EnterpriseOperationLogGetV10ResponseDataListInner `json:"list,omitempty"`
	//
	Offset *int64 `json:"offset,omitempty"`
}