/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeCustomCreativeUpdateV2RequestCreativeListInnerTitleMaterial
type CreativeCustomCreativeUpdateV2RequestCreativeListInnerTitleMaterial struct {
	//
	BidwordList []*CreativeCustomCreativeUpdateV2RequestCreativeListInnerTitleMaterialBidwordListInner `json:"bidword_list,omitempty"`
	//
	DpaWordList []*CreativeCustomCreativeUpdateV2RequestCreativeListInnerTitleMaterialDpaWordListInner `json:"dpa_word_list,omitempty"`
	//
	Title string `json:"title"`
	//
	WordList []*CreativeCustomCreativeUpdateV2RequestCreativeListInnerTitleMaterialWordListInner `json:"word_list,omitempty"`
}
