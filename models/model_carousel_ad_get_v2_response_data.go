/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselAdGetV2ResponseData
type CarouselAdGetV2ResponseData struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Carousels []*CarouselAdGetV2ResponseDataCarouselsInner `json:"carousels,omitempty"`
}
