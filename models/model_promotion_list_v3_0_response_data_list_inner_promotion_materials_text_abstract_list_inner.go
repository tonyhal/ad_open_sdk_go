/*
API Title

巨量引擎开放平台

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerPromotionMaterialsTextAbstractListInner struct for PromotionListV30ResponseDataListInnerPromotionMaterialsTextAbstractListInner
type PromotionListV30ResponseDataListInnerPromotionMaterialsTextAbstractListInner struct {
	//
	AbstractText *string `json:"abstract_text,omitempty"`
	//
	BidwordList []*PromotionListV30ResponseDataListInnerPromotionMaterialsTextAbstractListInnerBidwordListInner `json:"bidword_list,omitempty"`
	//
	WordList []int64 `json:"word_list,omitempty"`
}
