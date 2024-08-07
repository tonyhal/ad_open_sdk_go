/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportProductDailyAsyncTaskCreateV30Request struct for ReportProductDailyAsyncTaskCreateV30Request
type ReportProductDailyAsyncTaskCreateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	EndDate string `json:"end_date"`
	//
	Fields    []string                                              `json:"fields,omitempty"`
	Filtering *ReportProductDailyAsyncTaskCreateV30RequestFiltering `json:"filtering,omitempty"`
	//
	ProductPlatformId int64 `json:"product_platform_id"`
	//
	StartDate string `json:"start_date"`
	//
	TaskName string `json:"task_name"`
}
