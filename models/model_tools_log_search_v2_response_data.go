/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLogSearchV2ResponseData
type ToolsLogSearchV2ResponseData struct {
	//
	Logs     []*ToolsLogSearchV2ResponseDataLogsInner     `json:"logs,omitempty"`
	PageInfo *AgentAdvertiserSelectV2ResponseDataPageInfo `json:"page_info,omitempty"`
}
