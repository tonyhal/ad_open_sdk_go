/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroGameUpdateV30RequestGameLinkInner struct for ToolsMicroGameUpdateV30RequestGameLinkInner
type ToolsMicroGameUpdateV30RequestGameLinkInner struct {
	//
	Id *int64 `json:"id,omitempty"`
	//
	Link        string                                     `json:"link"`
	OperateType ToolsMicroGameUpdateV30GameLinkOperateType `json:"operate_type"`
	//
	Remark string `json:"remark"`
	//
	StartParam *string `json:"start_param,omitempty"`
}
