/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalReportMaterialGetV30Response struct for LocalReportMaterialGetV30Response
type LocalReportMaterialGetV30Response struct {
	//
	Code *int64                                 `json:"code,omitempty"`
	Data *LocalReportMaterialGetV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
