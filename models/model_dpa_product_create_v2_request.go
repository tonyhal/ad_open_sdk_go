/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaProductCreateV2Request struct for DpaProductCreateV2Request
type DpaProductCreateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	PlatformId  int64                                `json:"platform_id"`
	ProductInfo DpaProductCreateV2RequestProductInfo `json:"product_info"`
}
