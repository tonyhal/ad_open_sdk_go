/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30RequestPromotionMaterialsTitleMaterialListInner struct for PromotionUpdateV30RequestPromotionMaterialsTitleMaterialListInner
type PromotionUpdateV30RequestPromotionMaterialsTitleMaterialListInner struct {
	//
	BidwordList []*PromotionCreateV30RequestPromotionMaterialsTextAbstractListInnerBidwordListInner `json:"bidword_list,omitempty"`
	//
	DpaWordList []int64 `json:"dpa_word_list,omitempty"`
	//
	Title string `json:"title"`
	//
	WordList []int64 `json:"word_list,omitempty"`
}
