/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselUpdateV2Request struct for CarouselUpdateV2Request
type CarouselUpdateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Carousels []*CarouselUpdateV2RequestCarouselsInner `json:"carousels"`
}
