/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaProductDeleteV2Request struct for DpaProductDeleteV2Request
type DpaProductDeleteV2Request struct {
	// 广告主id
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 商品库id
	PlatformId int64 `json:"platform_id"`
	// 商品id
	ProductId int64 `json:"product_id"`
}
