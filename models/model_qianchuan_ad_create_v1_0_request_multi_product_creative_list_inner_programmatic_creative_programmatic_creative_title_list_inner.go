/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10RequestMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeTitleListInner struct for QianchuanAdCreateV10RequestMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeTitleListInner
type QianchuanAdCreateV10RequestMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeTitleListInner struct {
	//
	AwemeCarouselId *int64 `json:"aweme_carousel_id,omitempty"`
	//
	DynamicWords []*QianchuanAdCreateV10RequestMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeTitleListInnerDynamicWordsInner `json:"dynamic_words,omitempty"`
	//
	Title     string                                                                                                  `json:"title"`
	TitleType *QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType `json:"title_type,omitempty"`
}
