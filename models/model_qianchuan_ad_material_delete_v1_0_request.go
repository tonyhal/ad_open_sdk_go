/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdMaterialDeleteV10Request struct for QianchuanAdMaterialDeleteV10Request
type QianchuanAdMaterialDeleteV10Request struct {
	// 计划ID
	AdId int64 `json:"ad_id"`
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 待删除素材ID
	MaterialIds []int64 `json:"material_ids"`
}
