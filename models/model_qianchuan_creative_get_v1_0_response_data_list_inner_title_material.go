/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCreativeGetV10ResponseDataListInnerTitleMaterial
type QianchuanCreativeGetV10ResponseDataListInnerTitleMaterial struct {
	//
	DynamicWords []*QianchuanCreativeGetV10ResponseDataListInnerTitleMaterialDynamicWordsInner `json:"dynamic_words,omitempty"`
	//
	Title *string `json:"title,omitempty"`
}
