/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SubscribeAccountsRemoveV30Request struct for SubscribeAccountsRemoveV30Request
type SubscribeAccountsRemoveV30Request struct {
	//
	AdvertiserIds []int64 `json:"advertiser_ids"`
	//
	AppId int64 `json:"app_id"`
	//
	CoreUserId int64 `json:"core_user_id"`
	//
	Events []string `json:"events,omitempty"`
	//
	SubscribeTaskId int64 `json:"subscribe_task_id"`
}
