/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10RequestMultiProductCreativeListInnerProgrammaticCreative
type QianchuanAdCreateV10RequestMultiProductCreativeListInnerProgrammaticCreative struct {
	ProgrammaticCreativeCard *QianchuanAdCreateV10RequestMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeCard `json:"programmatic_creative_card,omitempty"`
	//
	ProgrammaticCreativeMediaList []*QianchuanAdCreateV10RequestMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeMediaListInner `json:"programmatic_creative_media_list,omitempty"`
	//
	ProgrammaticCreativeTitleList []*QianchuanAdCreateV10RequestMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeTitleListInner `json:"programmatic_creative_title_list,omitempty"`
}
