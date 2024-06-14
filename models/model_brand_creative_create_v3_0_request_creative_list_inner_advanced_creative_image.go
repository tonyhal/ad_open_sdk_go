/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreativeImage 图片信息
type BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreativeImage struct {
	// 格式
	Format *string `json:"format,omitempty"`
	//
	Hash *string `json:"hash,omitempty"`
	// 高
	Height *int64 `json:"height,omitempty"`
	// 图片URL
	ShowUrl *string `json:"show_url,omitempty"`
	// 图片URI
	Uri *string `json:"uri,omitempty"`
	// 宽
	Width *int64 `json:"width,omitempty"`
}
