/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeAdvancedCreativeInfoCard 图文磁贴
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeAdvancedCreativeInfoCard struct {
	// 组件弹出时间
	AdvancedDuration *int64                                                                                  `json:"advanced_duration,omitempty"`
	ImageInfo        *BrandCreativeGetV30ResponseDataCreativesInnerCreativeAdvancedCreativeInfoCardImageInfo `json:"image_info,omitempty"`
}
