/*
API Title

巨量引擎开放平台

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectListV30ResponseDataListInnerKeywordsInner struct for ProjectListV30ResponseDataListInnerKeywordsInner
type ProjectListV30ResponseDataListInnerKeywordsInner struct {
	BidType   *ProjectListV30DataListKeywordsBidType   `json:"bid_type,omitempty"`
	MatchType *ProjectListV30DataListKeywordsMatchType `json:"match_type,omitempty"`
	// 关键名称
	Word *string `json:"word,omitempty"`
}
