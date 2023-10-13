/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AsyncTaskCreateV2Request struct for AsyncTaskCreateV2Request
type AsyncTaskCreateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Force *bool `json:"force,omitempty"`
	//
	TaskName   string                             `json:"task_name"`
	TaskParams AsyncTaskCreateV2RequestTaskParams `json:"task_params"`
	TaskType   AsyncTaskCreateV2TaskType          `json:"task_type"`
}
