/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselCreateV2ResponseDataCarousel
type CarouselCreateV2ResponseDataCarousel struct {
	Audio *CarouselCreateV2ResponseDataCarouselAudio `json:"audio,omitempty"`
	//
	CarouselId *int64 `json:"carousel_id,omitempty"`
	//
	CarouselType *int64 `json:"carousel_type,omitempty"`
	//
	FileName *string `json:"file_name,omitempty"`
	//
	Images []*CarouselCreateV2ResponseDataCarouselImagesInner `json:"images,omitempty"`
	//
	Uri *string `json:"uri,omitempty"`
}
