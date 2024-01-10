/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordCampaignAddV2RequestCampaignListInner struct for ToolsPrivativeWordCampaignAddV2RequestCampaignListInner
type ToolsPrivativeWordCampaignAddV2RequestCampaignListInner struct {
	//
	CampaignId int64 `json:"campaign_id"`
	//
	PhraseWords []string `json:"phrase_words,omitempty"`
	//
	PreciseWords []string `json:"precise_words,omitempty"`
}
