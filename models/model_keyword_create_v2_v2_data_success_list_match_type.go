/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordCreateV2V2DataSuccessListMatchType
type KeywordCreateV2V2DataSuccessListMatchType string

// List of keyword_create_v2_v2_data_success_list_match_type
const (
	PRECISION_KeywordCreateV2V2DataSuccessListMatchType KeywordCreateV2V2DataSuccessListMatchType = "PRECISION"
	EXTENSIVE_KeywordCreateV2V2DataSuccessListMatchType KeywordCreateV2V2DataSuccessListMatchType = "EXTENSIVE"
	PHRASE_KeywordCreateV2V2DataSuccessListMatchType    KeywordCreateV2V2DataSuccessListMatchType = "PHRASE"
)

// Ptr returns reference to keyword_create_v2_v2_data_success_list_match_type value
func (v KeywordCreateV2V2DataSuccessListMatchType) Ptr() *KeywordCreateV2V2DataSuccessListMatchType {
	return &v
}
