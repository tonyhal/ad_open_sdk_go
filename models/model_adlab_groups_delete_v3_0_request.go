/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupsDeleteV30Request struct for AdlabGroupsDeleteV30Request
type AdlabGroupsDeleteV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 项目id列表
	Ids []int64 `json:"ids"`
}
