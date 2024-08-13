/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselListV2Filtering
type CarouselListV2Filtering struct {
	//
	AudioId *string `json:"audio_id,omitempty"`
	//
	CarouselIds []int64 `json:"carousel_ids,omitempty"`
	//
	EndTime *string `json:"end_time,omitempty"`
	//
	FileName *string `json:"file_name,omitempty"`
	//
	ImageIds []string `json:"image_ids,omitempty"`
	//
	StartTime *string `json:"start_time,omitempty"`
}
