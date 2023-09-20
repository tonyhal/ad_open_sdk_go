/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordProjectAddV30ResponseDataErrorListInner struct for ToolsPrivativeWordProjectAddV30ResponseDataErrorListInner
type ToolsPrivativeWordProjectAddV30ResponseDataErrorListInner struct {
	//
	ErrorMessage *string `json:"error_message,omitempty"`
	//
	FailPhraseWords []*ToolsPrivativeWordProjectAddV30ResponseDataErrorListInnerFailPhraseWordsInner `json:"fail_phrase_words,omitempty"`
	//
	FailPreciseWords []*ToolsPrivativeWordProjectAddV30ResponseDataErrorListInnerFailPhraseWordsInner `json:"fail_precise_words,omitempty"`
	//
	ProjectId *int64 `json:"project_id,omitempty"`
}
