/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoHorizontalImageVideo 横版带图视频
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoHorizontalImageVideo struct {
	Image *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoHorizontalImageVideoImage `json:"image,omitempty"`
	// 标题
	Title *string                                                                                     `json:"title,omitempty"`
	Video *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoHorizontalImageVideoVideo `json:"video,omitempty"`
}
