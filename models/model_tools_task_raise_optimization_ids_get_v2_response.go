/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsTaskRaiseOptimizationIdsGetV2Response struct for ToolsTaskRaiseOptimizationIdsGetV2Response
type ToolsTaskRaiseOptimizationIdsGetV2Response struct {
	//
	Code *int64                                          `json:"code,omitempty"`
	Data *ToolsTaskRaiseOptimizationIdsGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
