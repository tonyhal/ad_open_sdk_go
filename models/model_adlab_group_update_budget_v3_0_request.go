/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateBudgetV30Request struct for AdlabGroupUpdateBudgetV30Request
type AdlabGroupUpdateBudgetV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 项目信息
	Data []*AdlabGroupUpdateBudgetV30RequestDataInner `json:"data"`
}