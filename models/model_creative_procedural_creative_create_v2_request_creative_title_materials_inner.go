/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeCreateV2RequestCreativeTitleMaterialsInner struct for CreativeProceduralCreativeCreateV2RequestCreativeTitleMaterialsInner
type CreativeProceduralCreativeCreateV2RequestCreativeTitleMaterialsInner struct {
	//
	BidwordList []*CreativeProceduralCreativeCreateV2RequestCreativeTitleMaterialsInnerBidwordListInner `json:"bidword_list,omitempty"`
	//
	DpaWordList []*CreativeProceduralCreativeCreateV2RequestCreativeTitleMaterialsInnerDpaWordListInner `json:"dpa_word_list,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	Title string `json:"title"`
	//
	WordList []*CreativeProceduralCreativeCreateV2RequestCreativeTitleMaterialsInnerWordListInner `json:"word_list,omitempty"`
}
