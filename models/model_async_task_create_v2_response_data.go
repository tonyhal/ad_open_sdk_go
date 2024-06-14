/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AsyncTaskCreateV2ResponseData
type AsyncTaskCreateV2ResponseData struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	ErrMsg *string `json:"err_msg,omitempty"`
	//
	FileSize *int64 `json:"file_size,omitempty"`
	//
	TaskId *int64 `json:"task_id,omitempty"`
	//
	TaskName *string `json:"task_name,omitempty"`
	//
	TaskParams map[string]interface{} `json:"task_params,omitempty"`
	//
	TaskStatus *string `json:"task_status,omitempty"`
	//
	TaskType *string `json:"task_type,omitempty"`
}
