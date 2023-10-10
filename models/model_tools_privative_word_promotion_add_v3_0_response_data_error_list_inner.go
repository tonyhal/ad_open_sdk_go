/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordPromotionAddV30ResponseDataErrorListInner struct for ToolsPrivativeWordPromotionAddV30ResponseDataErrorListInner
type ToolsPrivativeWordPromotionAddV30ResponseDataErrorListInner struct {
	//
	ErrorMessage *string `json:"error_message,omitempty"`
	//
	FailPhraseWords []*ToolsPrivativeWordProjectAddV30ResponseDataErrorListInnerFailPhraseWordsInner `json:"fail_phrase_words,omitempty"`
	//
	FailPreciseWords []*ToolsPrivativeWordProjectAddV30ResponseDataErrorListInnerFailPhraseWordsInner `json:"fail_precise_words,omitempty"`
	//
	PromotionId *int64 `json:"promotion_id,omitempty"`
}
