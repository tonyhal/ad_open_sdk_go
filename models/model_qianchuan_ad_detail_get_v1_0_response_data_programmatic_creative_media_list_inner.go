/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataProgrammaticCreativeMediaListInner struct for QianchuanAdDetailGetV10ResponseDataProgrammaticCreativeMediaListInner
type QianchuanAdDetailGetV10ResponseDataProgrammaticCreativeMediaListInner struct {
	//
	AwemeItemId *int64 `json:"aweme_item_id,omitempty"`
	//
	ImageIds  []string                                                           `json:"image_ids,omitempty"`
	ImageMode *QianchuanAdDetailGetV10DataProgrammaticCreativeMediaListImageMode `json:"image_mode,omitempty"`
	//
	IsAutoGenerate *int64 `json:"is_auto_generate,omitempty"`
	//
	Title *string `json:"title,omitempty"`
	//
	VideoCoverId *string `json:"video_cover_id,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
}
