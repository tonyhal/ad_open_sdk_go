/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueContactLogListV2Filter
type ToolsClueContactLogListV2Filter struct {
	// 呼叫记录id，若不填则返回线索下的全部话单
	ContactId *string `json:"contact_id,omitempty"`
}
