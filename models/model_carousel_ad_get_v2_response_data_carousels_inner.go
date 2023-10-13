/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselAdGetV2ResponseDataCarouselsInner struct for CarouselAdGetV2ResponseDataCarouselsInner
type CarouselAdGetV2ResponseDataCarouselsInner struct {
	Audio *CarouselAdGetV2ResponseDataCarouselsInnerAudio `json:"audio,omitempty"`
	//
	CarouselType *int64 `json:"carousel_type,omitempty"`
	//
	FileName *string `json:"file_name,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	//
	Images []*CarouselAdGetV2ResponseDataCarouselsInnerImagesInner `json:"images,omitempty"`
	//
	Uri *string `json:"uri,omitempty"`
}