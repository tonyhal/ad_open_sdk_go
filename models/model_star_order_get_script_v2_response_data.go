/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetScriptV2ResponseData
type StarOrderGetScriptV2ResponseData struct {
	// 任务脚本信息
	OrderScriptList []*StarOrderGetScriptV2ResponseDataOrderScriptListInner `json:"order_script_list,omitempty"`
}
