/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeAuthAuthShareAdShareV2ResponseData
type ToolsAwemeAuthAuthShareAdShareV2ResponseData struct {
	// 共享失败的账户及原因
	ErrorList []*ToolsAwemeAuthAuthShareAdShareV2ResponseDataErrorListInner `json:"error_list,omitempty"`
	// 共享成功的账户列表
	SuccessList []int64 `json:"success_list,omitempty"`
}
