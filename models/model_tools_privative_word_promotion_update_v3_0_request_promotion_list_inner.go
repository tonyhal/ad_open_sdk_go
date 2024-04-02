/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordPromotionUpdateV30RequestPromotionListInner struct for ToolsPrivativeWordPromotionUpdateV30RequestPromotionListInner
type ToolsPrivativeWordPromotionUpdateV30RequestPromotionListInner struct {
	//
	PhraseWords []string `json:"phrase_words,omitempty"`
	//
	PreciseWords []string `json:"precise_words,omitempty"`
	//
	PromotionId *int64 `json:"promotion_id,omitempty"`
}
