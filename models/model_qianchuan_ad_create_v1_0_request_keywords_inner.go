/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10RequestKeywordsInner struct for QianchuanAdCreateV10RequestKeywordsInner
type QianchuanAdCreateV10RequestKeywordsInner struct {
	MatchType QianchuanAdCreateV10KeywordsMatchType `json:"match_type"`
	//
	Word string `json:"word"`
	//
	WordId *int64 `json:"word_id,omitempty"`
	//
	WordPackageId *int64 `json:"word_package_id,omitempty"`
	//
	WordPackageName *string `json:"word_package_name,omitempty"`
}
