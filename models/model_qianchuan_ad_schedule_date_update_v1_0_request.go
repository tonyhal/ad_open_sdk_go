/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdScheduleDateUpdateV10Request struct for QianchuanAdScheduleDateUpdateV10Request
type QianchuanAdScheduleDateUpdateV10Request struct {
	//
	AdIds []int64 `json:"ad_ids"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	EndTime      *int64                                       `json:"end_time,omitempty"`
	ScheduleType QianchuanAdScheduleDateUpdateV10ScheduleType `json:"schedule_type"`
}
