/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordBatchGetV30ResponseDataProjectsPrivativeInner struct for ToolsPrivativeWordBatchGetV30ResponseDataProjectsPrivativeInner
type ToolsPrivativeWordBatchGetV30ResponseDataProjectsPrivativeInner struct {
	//
	PhraseWords []string `json:"phrase_words,omitempty"`
	//
	PreciseWords []string `json:"precise_words,omitempty"`
	//
	ProjectId *int64 `json:"project_id,omitempty"`
}
