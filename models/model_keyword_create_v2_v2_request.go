/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordCreateV2V2Request struct for KeywordCreateV2V2Request
type KeywordCreateV2V2Request struct {
	//
	AdId int64 `json:"ad_id"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Keywords []*KeywordCreateV2V2RequestKeywordsInner `json:"keywords"`
}
