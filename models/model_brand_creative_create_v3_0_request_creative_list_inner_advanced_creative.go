/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreative 附加创意
type BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreative struct {
	// 图片弹出时间
	Duration *int64                                                               `json:"duration,omitempty"`
	Image    *BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreativeImage `json:"image,omitempty"`
}
