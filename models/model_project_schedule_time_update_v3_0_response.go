/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectScheduleTimeUpdateV30Response struct for ProjectScheduleTimeUpdateV30Response
type ProjectScheduleTimeUpdateV30Response struct {
	//
	Code *int64                                    `json:"code,omitempty"`
	Data *ProjectScheduleTimeUpdateV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
