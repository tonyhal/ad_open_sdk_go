/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordAdAddV2RequestAdListInner struct for ToolsPrivativeWordAdAddV2RequestAdListInner
type ToolsPrivativeWordAdAddV2RequestAdListInner struct {
	//
	AdId int64 `json:"ad_id"`
	//
	PhraseWords []string `json:"phrase_words,omitempty"`
	//
	PreciseWords []string `json:"precise_words,omitempty"`
}
