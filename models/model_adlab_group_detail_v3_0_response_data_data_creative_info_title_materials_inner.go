/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupDetailV30ResponseDataDataCreativeInfoTitleMaterialsInner struct for AdlabGroupDetailV30ResponseDataDataCreativeInfoTitleMaterialsInner
type AdlabGroupDetailV30ResponseDataDataCreativeInfoTitleMaterialsInner struct {
	// 标题信息
	Title *string `json:"title,omitempty"`
	// 词包信息
	WordList []*AdlabGroupDetailV30ResponseDataDataCreativeInfoTitleMaterialsInnerWordListInner `json:"word_list,omitempty"`
}