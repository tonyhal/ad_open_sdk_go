/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdUpdateBidV2Request struct for AdUpdateBidV2Request
type AdUpdateBidV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Data []*AdUpdateBidV2RequestDataInner `json:"data"`
}
