/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskUnshareV30Request struct for StardeliveryTaskUnshareV30Request
type StardeliveryTaskUnshareV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	// 取消共享账号范围，一次最多操作50个
	AdvertiserIds []int64 `json:"advertiser_ids"`
	// 任务ID
	StarTaskId int64 `json:"star_task_id"`
}
