/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeUpdateV30RequestAdvancedCreative 附加创意
type BrandCreativeUpdateV30RequestAdvancedCreative struct {
	// 时长
	Duration *int64                                              `json:"duration,omitempty"`
	Image    *BrandCreativeUpdateV30RequestAdvancedCreativeImage `json:"image,omitempty"`
}