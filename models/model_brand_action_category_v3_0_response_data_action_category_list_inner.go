/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandActionCategoryV30ResponseDataActionCategoryListInner struct for BrandActionCategoryV30ResponseDataActionCategoryListInner
type BrandActionCategoryV30ResponseDataActionCategoryListInner struct {
	// 二级兴趣
	Children []*BrandActionCategoryV30ResponseDataActionCategoryListInnerChildrenInner `json:"children,omitempty"`
	// 一级兴趣ID
	Id *int64 `json:"id,omitempty"`
	// 一级兴趣名称
	Name *string `json:"name,omitempty"`
}