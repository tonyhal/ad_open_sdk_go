/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementUploadTaskCreateV2Response struct for ToolsAppManagementUploadTaskCreateV2Response
type ToolsAppManagementUploadTaskCreateV2Response struct {
	//
	Code *int64                                            `json:"code,omitempty"`
	Data *ToolsAppManagementUploadTaskCreateV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}