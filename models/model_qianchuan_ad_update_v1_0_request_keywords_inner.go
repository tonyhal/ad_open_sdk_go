/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10RequestKeywordsInner struct for QianchuanAdUpdateV10RequestKeywordsInner
type QianchuanAdUpdateV10RequestKeywordsInner struct {
	MatchType QianchuanAdUpdateV10KeywordsMatchType `json:"match_type"`
	//
	Word string `json:"word"`
	//
	WordId *int64 `json:"word_id,omitempty"`
	//
	WordPackageId *int64 `json:"word_package_id,omitempty"`
	//
	WordPackageName *string `json:"word_package_name,omitempty"`
}
