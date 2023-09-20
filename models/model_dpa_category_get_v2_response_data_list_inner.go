/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaCategoryGetV2ResponseDataListInner struct for DpaCategoryGetV2ResponseDataListInner
type DpaCategoryGetV2ResponseDataListInner struct {
	// 分类id
	Id *string `json:"id,omitempty"`
	// 分类名称
	Name *string `json:"name,omitempty"`
	// 父级分类id，没有父级则为-1
	Parent *string `json:"parent,omitempty"`
	// 子级分类，嵌套递归
	Subs []map[string]interface{} `json:"subs,omitempty"`
}
