/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoMaterialClearTaskCreateV2Request struct for FileVideoMaterialClearTaskCreateV2Request
type FileVideoMaterialClearTaskCreateV2Request struct {
	// 广告主id
	AdvertiserId    int64                                                    `json:"advertiser_id"`
	ClearTaskParams FileVideoMaterialClearTaskCreateV2RequestClearTaskParams `json:"clear_task_params"`
}
