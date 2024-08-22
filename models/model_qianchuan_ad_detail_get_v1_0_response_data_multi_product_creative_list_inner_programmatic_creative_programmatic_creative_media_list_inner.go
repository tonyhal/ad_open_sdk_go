/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeMediaListInner struct for QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeMediaListInner
type QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeMediaListInner struct {
	//
	AwemeCarouselId *int64 `json:"aweme_carousel_id,omitempty"`
	//
	AwemeItemId   *int64                                                                                                                               `json:"aweme_item_id,omitempty"`
	CarouselAudio *QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeMediaListInnerCarouselAudio `json:"carousel_audio,omitempty"`
	//
	CarouselDescription *string `json:"carousel_description,omitempty"`
	//
	CarouselId *int64 `json:"carousel_id,omitempty"`
	//
	CarouselImages []*QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeMediaListInnerCarouselImagesInner `json:"carousel_images,omitempty"`
	//
	ImageIds  []string                                                                                                       `json:"image_ids,omitempty"`
	ImageMode *QianchuanAdDetailGetV10DataMultiProductCreativeListProgrammaticCreativeProgrammaticCreativeMediaListImageMode `json:"image_mode,omitempty"`
	//
	IsAutoGenerate *int64 `json:"is_auto_generate,omitempty"`
	//
	Title *string `json:"title,omitempty"`
	//
	VideoCoverId *string `json:"video_cover_id,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
	//
	VideoPosterUrl *string `json:"video_poster_url,omitempty"`
	//
	VideoUrl *string `json:"video_url,omitempty"`
}
