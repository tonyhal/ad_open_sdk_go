/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordDeleteV2V2Request struct for KeywordDeleteV2V2Request
type KeywordDeleteV2V2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	KeywordIds []int64 `json:"keyword_ids"`
}