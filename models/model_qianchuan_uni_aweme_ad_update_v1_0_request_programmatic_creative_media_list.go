/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniAwemeAdUpdateV10RequestProgrammaticCreativeMediaList
type QianchuanUniAwemeAdUpdateV10RequestProgrammaticCreativeMediaList struct {
	//
	BlockVideoMaterial []*QianchuanUniAwemeAdUpdateV10RequestProgrammaticCreativeMediaListBlockVideoMaterialInner `json:"block_video_material,omitempty"`
	//
	TitleMaterial []*QianchuanUniAwemeAdUpdateV10RequestProgrammaticCreativeMediaListTitleMaterialInner `json:"title_material,omitempty"`
	//
	VideoMaterial []*QianchuanUniAwemeAdUpdateV10RequestProgrammaticCreativeMediaListVideoMaterialInner `json:"video_material,omitempty"`
}
