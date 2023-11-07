/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeDetailGetV30ResponseDataCreativeListInnerTitleMaterial 创意标题素材
type CreativeDetailGetV30ResponseDataCreativeListInnerTitleMaterial struct {
	//
	BidwordList []*CreativeDetailGetV30ResponseDataCreativeListInnerTitleMaterialBidwordListInner `json:"bidword_list,omitempty"`
	// DPA词包ID列表
	DpaWordList []*CreativeDetailGetV30ResponseDataCreativeTitleMaterialsInnerDpaWordListInner `json:"dpa_word_list,omitempty"`
	// 创意标题
	Title *string `json:"title,omitempty"`
	// 动态词包ID列表
	WordList []*CreativeDetailGetV30ResponseDataCreativeAbstractMaterialsInnerTextAbstractInfoWordListInner `json:"word_list,omitempty"`
}
