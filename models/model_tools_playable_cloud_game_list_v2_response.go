/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableCloudGameListV2Response struct for ToolsPlayableCloudGameListV2Response
type ToolsPlayableCloudGameListV2Response struct {
	//
	Code *int64                                    `json:"code,omitempty"`
	Data *ToolsPlayableCloudGameListV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
