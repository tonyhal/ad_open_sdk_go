/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanKeywordPackageGetV10ResponseDataWordPackageInfosInner struct for QianchuanKeywordPackageGetV10ResponseDataWordPackageInfosInner
type QianchuanKeywordPackageGetV10ResponseDataWordPackageInfosInner struct {
	//
	Degree *int64 `json:"degree,omitempty"`
	//
	Id int64 `json:"id"`
	//
	KeywordInfos []*QianchuanKeywordPackageGetV10ResponseDataWordPackageInfosInnerKeywordInfosInner `json:"keyword_infos,omitempty"`
	//
	Name string `json:"name"`
	//
	SearchSum *int64 `json:"search_sum,omitempty"`
}