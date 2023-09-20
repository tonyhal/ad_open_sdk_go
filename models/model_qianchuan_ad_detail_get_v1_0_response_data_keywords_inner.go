/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataKeywordsInner struct for QianchuanAdDetailGetV10ResponseDataKeywordsInner
type QianchuanAdDetailGetV10ResponseDataKeywordsInner struct {
	//
	Id        *int64                                        `json:"id,omitempty"`
	MatchType *QianchuanAdDetailGetV10DataKeywordsMatchType `json:"match_type,omitempty"`
	Status    *QianchuanAdDetailGetV10DataKeywordsStatus    `json:"status,omitempty"`
	//
	Word *string `json:"word,omitempty"`
	//
	WordId *int64 `json:"word_id,omitempty"`
	//
	WordPackageId *int64 `json:"word_package_id,omitempty"`
	//
	WordPackageName *string `json:"word_package_name,omitempty"`
}