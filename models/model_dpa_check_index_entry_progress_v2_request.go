/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaCheckIndexEntryProgressV2Request struct for DpaCheckIndexEntryProgressV2Request
type DpaCheckIndexEntryProgressV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 文件拉取日期，case:\"20230331\"
	IndexPullDate string `json:"index_pull_date"`
	// 当天文件拉取次数，从1起
	IndexPullTime int64 `json:"index_pull_time"`
	// 商品库ID ，可通过【获取商品库信息】获取
	ProductPlatformId int64 `json:"product_platform_id"`
}
