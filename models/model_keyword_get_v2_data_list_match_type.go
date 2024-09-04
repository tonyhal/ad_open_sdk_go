/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordGetV2DataListMatchType
type KeywordGetV2DataListMatchType string

// List of keyword_get_v2_data_list_match_type
const (
	PHRASE_KeywordGetV2DataListMatchType    KeywordGetV2DataListMatchType = "PHRASE"
	EXTENSIVE_KeywordGetV2DataListMatchType KeywordGetV2DataListMatchType = "EXTENSIVE"
	PRECISION_KeywordGetV2DataListMatchType KeywordGetV2DataListMatchType = "PRECISION"
)

// Ptr returns reference to keyword_get_v2_data_list_match_type value
func (v KeywordGetV2DataListMatchType) Ptr() *KeywordGetV2DataListMatchType {
	return &v
}
